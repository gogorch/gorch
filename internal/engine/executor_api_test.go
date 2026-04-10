package engine

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/gogorch/gorch/recorder"
	"github.com/stretchr/testify/assert"
)

type spyRecorder struct {
	mu sync.Mutex

	useOperatorName bool
	useCalls        []bool

	startCount int
	stopCount  int
	resetCount int
	recorded   []string
}

func (r *spyRecorder) UseOperatorName(on bool) recorder.Recorder {
	r.mu.Lock()
	r.useOperatorName = on
	r.useCalls = append(r.useCalls, on)
	r.mu.Unlock()
	return r
}

func (r *spyRecorder) Start() {
	r.mu.Lock()
	r.startCount++
	r.mu.Unlock()
}

func (r *spyRecorder) Stop() {
	r.mu.Lock()
	r.stopCount++
	r.mu.Unlock()
}

func (r *spyRecorder) Elapsed() time.Duration { return 0 }

func (r *spyRecorder) Record(name, seq string) recorder.Done {
	r.mu.Lock()
	key := seq
	if r.useOperatorName {
		key = name
	}
	r.recorded = append(r.recorded, key)
	r.mu.Unlock()
	return func(status int) {}
}

func (r *spyRecorder) Report() recorder.Report {
	return recorder.Report{}
}

func (r *spyRecorder) Reset() {
	r.mu.Lock()
	r.resetCount++
	r.mu.Unlock()
}

func (r *spyRecorder) snapshot() (useCalls []bool, start, stop, reset int, recorded []string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	useCalls = append(useCalls, r.useCalls...)
	recorded = append(recorded, r.recorded...)
	return useCalls, r.startCount, r.stopCount, r.resetCount, recorded
}

func TestExecutorSettersWithCustomRecorder(t *testing.T) {
	clean()
	defer clean()
	registerTestOperator(t)

	p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() ChangeValueOP(val=2) }`)
	assert.Nil(t, err)

	eng, err := New(p)
	assert.Nil(t, err)
	exe := eng.Start("test")
	assert.NotNil(t, exe)

	val := &BeChangeValue{}
	assert.Nil(t, exe.Inject(&val, &t))

	rec := &spyRecorder{}
	exe.SetRecorder(rec)
	exe.SetLogOperatorName()
	exe.SetLogThreshold(time.Hour)
	exe.SetTimeout(3 * time.Second)

	assert.Nil(t, exe.Execute(context.Background()))
	assert.Equal(t, int64(2), val.Val)

	useCalls, start, stop, reset, recorded := rec.snapshot()
	assert.Equal(t, []bool{true}, useCalls)
	assert.Equal(t, 1, start)
	assert.Equal(t, 1, stop)
	assert.Equal(t, 1, reset)
	assert.Contains(t, recorded, "ChangeValueOP")
}

func TestTypedAPIWrappers(t *testing.T) {
	ctx := NewContext()
	defer releaseContext(ctx)

	v := int64(11)
	assert.Nil(t, RegisterTyped(ctx, &v, false))

	var got int64
	assert.Nil(t, MutableTyped(ctx, &got))
	assert.Equal(t, int64(11), got)

	v2 := int64(12)
	assert.EqualError(t, RegisterTyped(ctx, &v2, false), "register error: duplicate type, error type: int64")
	assert.Nil(t, RegisterTyped(ctx, &v2, true))

	assert.Nil(t, MutableTyped(ctx, &got))
	assert.Equal(t, int64(12), got)

	var miss string
	assert.EqualError(t, MutableTyped(ctx, &miss), "mutable error: ins not found, error type: string")
}
