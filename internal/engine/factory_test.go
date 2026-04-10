package engine

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/stretchr/testify/assert"
)

type generatedInvalidDIFieldOp struct {
	Bad int `inject:""`
}

func (o *generatedInvalidDIFieldOp) Execute(ctx *Context) error { return nil }
func (o *generatedInvalidDIFieldOp) IsGenerateOperatorCode()    {}

type strictModePlainOp struct{}

func (o *strictModePlainOp) Execute(ctx *Context) error { return nil }

type strictModeGeneratedOp struct{}

func (o *strictModeGeneratedOp) Execute(ctx *Context) error { return nil }
func (o *strictModeGeneratedOp) IsGenerateOperatorCode()    {}

func TestGeneratedOperatorSkipAnalyze(t *testing.T) {
	clean()
	defer clean()

	assert.NotPanics(t, func() {
		err := RegisterOperator[generatedInvalidDIFieldOp]("generatedInvalidDIFieldOp", 9001)
		assert.Nil(t, err)
	})
}

func TestStrictModeRejectPlainOperatorOnExecute(t *testing.T) {
	clean()
	defer clean()
	assert.Nil(t, RegisterOperator[strictModePlainOp]("strictModePlainOp", 9004))

	p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() strictModePlainOp }`)
	assert.Nil(t, err)

	eng, err := New(p)
	assert.Nil(t, err)
	eng.SetStrictGeneratedOnly(true)
	exe := eng.Start("test")
	assert.NotNil(t, exe)

	runErr := exe.Execute(context.Background())
	assert.Error(t, runErr)
	assert.Contains(t, runErr.Error(), "must implement IsGenerateOperatorCode in strict mode")
}

func TestStrictModeAllowGeneratedOperatorOnExecute(t *testing.T) {
	clean()
	defer clean()
	assert.Nil(t, RegisterOperator[strictModeGeneratedOp]("strictModeGeneratedOp", 9005))

	p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() strictModeGeneratedOp }`)
	assert.Nil(t, err)

	eng, err := New(p)
	assert.Nil(t, err)
	eng.SetStrictGeneratedOnly(true)

	exe := eng.Start("test")
	assert.NotNil(t, exe)
	assert.Nil(t, exe.Execute(context.Background()))
}

func TestStrictModeIsEngineScoped(t *testing.T) {
	clean()
	defer clean()
	assert.Nil(t, RegisterOperator[strictModePlainOp]("strictModePlainOp", 9006))

	p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() strictModePlainOp }`)
	assert.Nil(t, err)

	engStrict, err := New(p)
	assert.Nil(t, err)
	engLoose, err := New(p)
	assert.Nil(t, err)

	engStrict.SetStrictGeneratedOnly(true)
	engLoose.SetStrictGeneratedOnly(false)

	assert.Error(t, engStrict.Start("test").Execute(context.Background()))
	assert.Nil(t, engLoose.Start("test").Execute(context.Background()))
}

func TestRuntimeStatsFallbackObservation(t *testing.T) {
	clean()
	defer clean()
	assert.Nil(t, RegisterOperator[strictModePlainOp]("strictModePlainOp", 9007))
	assert.Nil(t, RegisterOperator[strictModeGeneratedOp]("strictModeGeneratedOp", 9008))

	p, err := iparser.Parse(`
START("plain"){ NO_CHECK_MISS() strictModePlainOp }
START("gen"){ NO_CHECK_MISS() strictModeGeneratedOp }
`)
	assert.Nil(t, err)

	eng, err := New(p)
	assert.Nil(t, err)
	eng.SetStrictGeneratedOnly(false)

	assert.Nil(t, eng.Start("plain").Execute(context.Background()))
	assert.Nil(t, eng.Start("gen").Execute(context.Background()))

	stats := SnapshotRuntimeStats()
	assert.GreaterOrEqual(t, stats.FallbackDIExecCount, uint64(1))
	assert.GreaterOrEqual(t, stats.GeneratedExecCount, uint64(1))
}

func TestRegisterOperatorConcurrent(t *testing.T) {
	clean()
	defer clean()

	var wg sync.WaitGroup
	errCh := make(chan error, 50)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			name := fmt.Sprintf("strictModeGeneratedOp-%d", idx)
			seq := uint(10000 + idx)
			errCh <- RegisterOperator[strictModeGeneratedOp](name, seq)
		}(i)
	}
	wg.Wait()
	close(errCh)

	for err := range errCh {
		assert.Nil(t, err)
	}
}
