package engine

import (
	"context"
	"testing"
	"time"

	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/stretchr/testify/assert"
)

// runDefectDSL 构造引擎并执行指定 START，返回执行错误与累加值。
func runDefectDSL(t *testing.T, dsl string, startName string, timeout time.Duration) (error, int64) {
	t.Helper()
	registerTestOperator(t)
	defer clean()

	p, err := iparser.Parse(dsl)
	if err != nil {
		return err, 0
	}
	eng, err := New(p)
	if err != nil {
		return err, 0
	}
	val := &BeChangeValue{Val: 0}
	exec := eng.Start(startName)
	if exec == nil {
		return errStartNotFound, 0
	}
	if err := exec.Inject(&val, &t); err != nil {
		return err, 0
	}
	if timeout >= 0 {
		exec.SetTimeout(timeout)
	}
	return exec.Execute(context.Background()), val.Get()
}

// LOOP 内的 BREAK 只应结束本次循环，不应把 skipSerial 传播到外层串行序列。
func TestRegressionBreakInsideLoopShouldNotSkipOuterSerial(t *testing.T) {
	err, val := runDefectDSL(t, `START("test"){
		NO_CHECK_MISS()
		AddValueOp(delta=1)
		-> LOOP(MAX_ITER=3){ AddValueOp(delta=10) -> BREAK() }
		-> AddValueOp(delta=100)
	}`, "test", 0)
	assert.Nil(t, err)
	assert.Equal(t, int64(111), val, "BREAK 之后的串行算子必须继续执行")
}

// RETRY 的循环体是串行序列时，内部算子的错误必须向上传递，否则永远不会重试。
func TestRegressionRetryBodySerialMustPropagateError(t *testing.T) {
	err, val := runDefectDSL(t, `START("test"){
		NO_CHECK_MISS()
		RETRY(MAX_TIMES=3){
			RetryFailThenPassOp(failTimes=2) -> NothingOp()
		}
	}`, "test", 0)
	assert.Nil(t, err)
	assert.Equal(t, int64(3), val, "串行体的错误被吞掉，RETRY 只执行了一次")
}

// 已执行的 Executor 再次 Execute 必须返回错误，而不是 nil 指针 panic。
func TestRegressionExecutorReuseReturnsError(t *testing.T) {
	registerTestOperator(t)
	defer clean()

	p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() NothingOp() }`)
	assert.Nil(t, err)
	eng, err := New(p)
	assert.Nil(t, err)

	val := &BeChangeValue{}
	exec := eng.Start("test")
	assert.NotNil(t, exec)
	assert.Nil(t, exec.Inject(&val))
	assert.Nil(t, exec.Execute(context.Background()))

	var second error
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("复用 Executor 触发了 panic: %v", r)
			}
		}()
		second = exec.Execute(context.Background())
	}()
	assert.NotNil(t, second, "复用 Executor 必须返回明确错误")
}

// Start 传入不存在的名称时不能返回 nil，否则调用方会 nil 解引用。
func TestRegressionStartUnknownNameReturnsErrorExecutor(t *testing.T) {
	registerTestOperator(t)
	defer clean()

	p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() NothingOp() }`)
	assert.Nil(t, err)
	eng, err := New(p)
	assert.Nil(t, err)

	exec := eng.Start("not-exist")
	assert.NotNil(t, exec, "未知 START 名称必须返回可调用的 Executor")
	if exec != nil {
		assert.NotNil(t, exec.Execute(context.Background()))
	}
}

// New(nil) 必须返回错误而不是 panic。
func TestRegressionNewNilPrimaryReturnsError(t *testing.T) {
	registerTestOperator(t)
	defer clean()

	var eng Engine
	var err error
	assert.NotPanics(t, func() {
		eng, err = New(nil)
	})
	assert.Nil(t, eng)
	assert.NotNil(t, err)
}

// SetTimeout(0) 应表示关闭引擎级超时；当前实现直接忽略 0。
func TestRegressionSetTimeoutZeroDisablesTimeout(t *testing.T) {
	err, _ := runDefectDSL(t, `START("test"){ NO_CHECK_MISS() SleepOp(sleep=60ms) }`, "test", 0)
	assert.Nil(t, err, "SetTimeout(0) 应关闭引擎超时，SleepOp 60ms 不应被判定超时")
}

// 同一个 Engine 派生的多个 Executor 并发执行必须互不干扰。
func TestRegressionConcurrentExecutorsSameEngine(t *testing.T) {
	registerTestOperator(t)
	defer clean()

	p, err := iparser.Parse(`START("test"){
		NO_CHECK_MISS()
		AddValueOp(delta=1) -> SleepOp(sleep=5ms) -> AddValueOp(delta=10)
	}`)
	assert.Nil(t, err)
	eng, err := New(p)
	assert.Nil(t, err)

	const n = 8
	errs := make(chan error, n)
	vals := make(chan int64, n)
	for i := 0; i < n; i++ {
		go func() {
			val := &BeChangeValue{}
			exec := eng.Start("test")
			if exec == nil {
				errs <- errStartNotFound
				return
			}
			if err := exec.Inject(&val); err != nil {
				errs <- err
				return
			}
			errs <- exec.Execute(context.Background())
			vals <- val.Get()
		}()
	}
	for i := 0; i < n; i++ {
		assert.Nil(t, <-errs)
		assert.Equal(t, int64(11), <-vals)
	}
}
