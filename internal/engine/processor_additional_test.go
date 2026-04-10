package engine

import (
	"testing"

	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/gogorch/gorch/recorder"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentProcessorCoverage(t *testing.T) {
	t.Run("concurrent_execute", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				[AddValueOp(delta=1), AddValueOp(delta=2)]
				-> ChangeValueOP(val=3, if=3)
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(3), val.Val)
			},
		})
	})

	t.Run("concurrent_prepare_error", func(t *testing.T) {
		clean()
		defer clean()

		p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() [UnknownOP] }`)
		assert.Nil(t, err)

		_, err = New(p)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "operator UnknownOP not register")
		assert.Contains(t, err.Error(), "prepare concurrent processor error")
	})
}

func TestOnFinishSemantics(t *testing.T) {
	t.Run("run_on_finish_when_main_failed", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				ON_FINISH(){ ChangeValueOP(val=9) }
				TestOperatorState(status=1)
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, fatal1, res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("on_finish_error_sets_global_error", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				ON_FINISH(){ TestOperatorState(status=1) }
				ChangeValueOP(val=2)
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, fatal1, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})
}

func TestOnFinishProcessorDirect(t *testing.T) {
	clean()
	defer clean()
	registerTestOperator(t)

	p, err := iparser.Parse(`START("test"){
		NO_CHECK_MISS()
		ON_FINISH(){ ChangeValueOP(val=6) }
		NothingOp
	}`)
	assert.Nil(t, err)

	stmt := p.Starters.Map["test"].OnFinishStmt
	proc := createProcessor(stmt)
	assert.NotNil(t, proc)
	assert.Nil(t, proc.Prepare())

	ctx := NewContext()
	defer releaseContext(ctx)
	ctx.recorder = recorder.New()
	ctx.recorder.Start()
	defer ctx.recorder.Reset()

	val := &BeChangeValue{}
	tt := t
	assert.Nil(t, registerAny(ctx.container, &val, true))
	assert.Nil(t, registerAny(ctx.container, &tt, true))

	assert.Nil(t, proc.Execute(ctx))
	assert.Equal(t, int64(6), val.Val)
}
