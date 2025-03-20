package engine

import (
	"context"
	"sync"

	"github.com/gorch/gorch/mlog"
	"github.com/gorch/gorch/pool"
	"github.com/gorch/gorch/recorder"
)

// Context 框架执行每个算子的上下文，每次执行算子都会获得一个全新的Context对象
// 但Context的属性，比如container容器属性，每次执行整个框架保持不变，所以在单次框架执行期间，每个算子拿到的container对象保持不变，
// 这样可以方便的让数据在算子间传递。
type Context struct {
	// Context 调用者的context.Context
	ctx context.Context

	// insHolder 对象保持工具
	// 对象容器用于存储执行过程中，需要在算子间传递的对象
	// 对象容器在一次执行中，仅仅创建一次对象，然后在Context.clone过程中依次传递给新的Context对象
	*container

	// 算子参数获取接口
	*args

	// interrupt 全局中断器，每次执行engine都只有一个对象
	// 一旦中断，则引擎不会继续执行（正在执行的逻辑会继续执行，直到结束）
	*interrupt

	// Logger 日志输出接口
	// 用户可以实现接口后，用自己的日志handler来输出日志
	mlog.Logger

	// recorder 日志记录器
	recorder recorder.Recorder

	// waiter 等待所有的协程执行完毕，才会最终退出执行框架
	waiter *sync.WaitGroup

	// skipSerial 在一个serial中，设置为true后，当前算子后面的算子跳过执行
	// 注意：仅在serial中生效，其他执行场景无效
	skipSerial bool

	// 在wrap指令场景下，调用当前wrap算子的下一个wrap算子，直至最终被wrap的processor
	nextWrap func() error

	// switchCase 在SWTICH指令场景下，调用指定的case算子
	// 注意，支持调用多个case，在框架内部，会并发的执行多个case，且总是会傻傻的等待所有case执行完成后，才会结束SWITCH指令
	switchCase func(...string) error
}

var (
	contextPool = pool.NewObjectPool(func(ctx *Context) {
		ctx.reset()
	})
)

func NewContext() *Context {
	nctx := contextPool.Get()
	if nctx == nil {
		nctx = new(Context)
	}

	nctx.ctx = context.TODO()
	nctx.container = newContainer()
	nctx.args = nil
	nctx.interrupt = newInterrupt()
	nctx.recorder = nil
	nctx.Logger = printLogger
	nctx.waiter = pool.WaitGroupPool.Get()
	nctx.skipSerial = false
	nctx.nextWrap = nil
	nctx.switchCase = nil

	return nctx
}

// releaseContext 本次调用结束后，释放Context对象，Context对象包含的其他子对象，也会被回收，请谨慎使用
func releaseContext(ctx *Context) {
	releaseContainer(ctx.container)
	releaseInterrupt(ctx.interrupt)
	pool.WaitGroupPool.Put(ctx.waiter)
	contextPool.Put(ctx)
}

// reset 重置Context对象
func (ctx *Context) reset() {
	ctx.ctx = nil
	ctx.container = nil
	ctx.args = nil
	ctx.interrupt = nil
	ctx.recorder = nil
	ctx.Logger = printLogger
	ctx.waiter = nil
	ctx.skipSerial = false
	ctx.nextWrap = nil
	ctx.switchCase = nil
}

// putPool 将当前Context对象放回对象池，请注意，本函数不会将当前Context对象的内容清空
// Context对象的清空操作，在clone或者NewContext时执行
func (ctx *Context) putPool() {
	if ctx != nil {
		contextPool.Put(ctx)
	}
}

// Clone 拷贝一个Context对象
func (ctx *Context) clone() *Context {
	nctx := contextPool.Get()
	if nctx == nil {
		nctx = new(Context)
	}
	nctx.ctx = ctx.ctx
	nctx.container = ctx.container
	nctx.args = nil
	nctx.recorder = ctx.recorder
	nctx.waiter = ctx.waiter
	nctx.interrupt = ctx.interrupt
	nctx.Logger = ctx.Logger
	nctx.skipSerial = false
	nctx.nextWrap = nil
	nctx.switchCase = nil

	return nctx
}

// Next 不是一个Wrap算子，无法调用
// 仅在WRAP指令场景下，用户在WRAP算子内调用本函数时，会执行被包装的逻辑，直至结束
func (ctx *Context) Next() error {
	if ctx.nextWrap == nil {
		return errNotWrapOperator
	}

	return ctx.nextWrap()
}

// Switch 在SWITCH指令执行的算子内，选择图上的其他CASE执行
// 如果算子不在SWITCH指令内，调用时会返回errNotSwitchOperator错误
func (ctx *Context) Switch(cases ...string) error {
	if ctx.switchCase == nil {
		return errNotSwitchOperator
	}

	return ctx.switchCase(cases...)
}

// SkipSerial 跳过当前串行执行，在其他场景内不生效
func (ctx *Context) SkipSerial() error {
	ctx.skipSerial = true
	return nil
}

func (ctx *Context) GetRoutine(name string) Routine {
	return ctx.getRoutine(name)
}

// Interrupt 返回中断器
func (ctx *Context) Interrupt() Interrupt {
	return ctx.interrupt
}
