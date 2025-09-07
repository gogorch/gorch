package engine

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOperator(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})

	t.Run("engine_timeout", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() SleepOp(sleep=100ms) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, errEngineExecuteTimeout, res)
				assert.Equal(t, int64(0), val.Val)
			},
			engineTimeout: 10 * time.Millisecond,
		})
	})

	t.Run("operator_timeout", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() SleepOp(sleep=200ms, timeout=50ms) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, errOperatorExecuteTimeout, res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("operator_timeout_ignore", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() @SleepOp(sleep=10ms, timeout=1ms) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, nil, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})

	t.Run("operator_timeout_success", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() SleepOp(sleep=10ms, timeout=20) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, nil, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})

	t.Run("operator_timeout_context_cancel", func(t *testing.T) {
		ctx := context.TODO()
		ctx, cancel := context.WithCancel(ctx)
		go func() {
			time.Sleep(time.Millisecond * 2)
			cancel()
		}()
		MyTestRun(t, TestValus{
			name: "test",
			ctx:  ctx,
			g:    `START("test"){ NO_CHECK_MISS() SleepOp(sleep=10ms) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, context.Canceled, res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("operator_timeout_context_interrupt", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ NO_CHECK_MISS() 
				ContextInterrupt(sleep=2ms) //睡眠2ms之后中断
					-> SleepOp(sleep=10ms, timeout=10ms) 
					-> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, errors.New("exit by ContextInterrupt"), res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("operator_panic", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() PanicOp }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, errors.New("operator PanicOp execute panic"), res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("operator_status_fatal", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() TestOperatorState(status=1) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, fatal1, res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("operator_status_fatal_ignore", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() @TestOperatorState(status=1) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, nil, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})

	t.Run("operator_status_fatal_with_msg", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() TestOperatorState(status=2) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, fatal2Cause(errors.New("fatal2WithErr")), res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("operator_status_fatal_with_msg_ignore", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() @TestOperatorState(status=2) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, nil, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})

	t.Run("operator_status_info", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() TestOperatorState(status=3) -> ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, nil, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})
}

func TestGoAndWaitOperator(t *testing.T) {
	// 等待routine成功
	t.Run("operator_go_wait_success", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS() 

				GO(GoOperator(sleep=4ms, val=10), "goOperator") 
				-> SleepOp(sleep=2ms) // 等待goroutine启动
				-> NothingOp(WAIT("goOperator")) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, nil, res)
				assert.Equal(t, int64(10), val.Val)
			},
		})
	})

	// 依赖gorutine启动，经常会出现失败
	// t.Run("operator_go_wait_not_start", func(t *testing.T) {
	// 	MyTestRun(t, TestValus{
	// 		name: "test",
	// 		g: `START("test"){
	// 			NO_CHECK_MISS()

	// 			GO(GoOperator(sleep=5ms, val=10), "goOperator")
	// 			-> NothingOp(WAIT("goOperator")) }`,
	// 		assertFunc: func(res error, val *BeChangeValue) {
	// 			assert.Equal(t, errors.New("routine \"goOperator\" not started"), res)
	// 			assert.Equal(t, int64(10), val.Val)
	// 		},
	// 	})
	// })

	// 等待超时，从当前算子执行时间计算
	t.Run("operator_go_wait_timeout", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS() 

				GO(GoOperator(sleep=20ms, val=10), "goOperator")
				-> SleepOp(sleep=2ms) // 等待goroutine启动
				-> NothingOp(WAIT("goOperator", timeout=5ms)) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.NotNil(t, res)
				assert.Contains(t, res.Error(), "wait routine \"goOperator\" timeout")
				assert.Equal(t, int64(10), val.Val)
			},
		})
	})

	// 等待超时，从routine启动时间开始计算
	t.Run("operator_go_wait_execute_timeout", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS() 

				GO(GoOperator(sleep=50ms, val=10), "goOperator")
				-> SleepOp(sleep=10ms) // 其他算子执行耗时
				-> NothingOp(WAIT("goOperator", totalTimeout=5ms)) }`, // 如果到NothingOp算子执行时，距离goOperator算子启动已经超过10ms，就认为超时
			assertFunc: func(res error, val *BeChangeValue) {
				assert.NotNil(t, res)
				assert.Contains(t, res.Error(), "routine \"goOperator\" execute timeout")
				assert.Equal(t, int64(10), val.Val) // 因为在engine中会等待所有的goroutine执行完成之后才推出，所以值会是10
			},
		})
	})

	t.Run("operator_go_wait_execute_success", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS() 

				GO(GoOperator(sleep=50ms, val=10), "goOperator")
				-> SleepOp(sleep=10ms) // 其他算子执行耗时
				-> NothingOp(WAIT("goOperator", totalTimeout=70ms)) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(10), val.Val)
			},
		})
	})
}
