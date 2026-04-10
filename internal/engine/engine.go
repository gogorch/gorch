package engine

import (
	"context"
	"fmt"
	"reflect"
	"sync/atomic"
	"time"

	"github.com/gogorch/gorch/internal/lang/ast"
	"github.com/gogorch/gorch/mlog"
	"github.com/gogorch/gorch/pool"
	"github.com/gogorch/gorch/recorder"
)

type Engine interface {
	Start(string) Executor
	SetStrictGeneratedOnly(bool)
}

type engine struct {
	starters            map[string]PrepareProcessor
	strictGeneratedOnly atomic.Bool
}

func (eng *engine) Start(key string) Executor {
	st, has := eng.starters[key]
	if has {
		return newExecutor(executorOptions{
			ctx:                 NewContext(),
			startName:           key,
			processor:           st,
			timeout:             8 * time.Second,
			strictGeneratedOnly: eng.strictGeneratedOnly.Load(),
			traceHooks:          getTraceHooks(),
		})
	}

	return nil
}

func (eng *engine) SetStrictGeneratedOnly(enabled bool) {
	eng.strictGeneratedOnly.Store(enabled)
}

// New 创建执行引擎对象
func New(p *ast.Primary) (Engine, error) {
	eng := &engine{
		starters: make(map[string]PrepareProcessor, len(p.Starters.Sort)),
	}
	for _, sName := range p.Starters.Sort {
		smt := p.Starters.Map[sName]
		starter := createProcessor(smt)
		if err := starter.Prepare(); err != nil {
			return nil, err
		}

		eng.starters[sName] = starter
	}

	return eng, nil
}

type Executor interface {
	// 注入对象到执行引擎中
	Inject(vals ...any) error

	// 设置执行超时时间
	// 如果不设置，将使用默认的执行超时时间: 8s
	SetTimeout(to time.Duration)

	// SetRecorder 设置一个算子执行记录器，如果不设置，将使用默认的执行记录器
	SetRecorder(r recorder.Recorder)

	// 默认情况下，在日志里记录算子的序号，
	// 可以通过这个方法，将算子的序号替换为算子的名称
	SetLogOperatorName()

	// 设置打印算子耗时的阈值
	// 默认情况下，只会记录算子的序号(或名称到日志)，表示该算子执行过
	// 当算子的耗时大于这个阈值，才会记录算子的耗时到日志
	SetLogThreshold(t time.Duration)

	// 执行算子
	Execute(ctx context.Context) error
}

type executor struct {
	ctx *Context
	// startName 当前执行入口名（START 名称）
	startName string

	// 执行超时时间
	timeout time.Duration
	// 记录算子的执行结果
	recorder recorder.Recorder
	// 日志里记录执行的算子名称
	// 默认情况下，日志里记录的是算子的序号
	logOperatorName bool
	// 打印算子耗时的阈值
	// 当算子的耗时大于这个阈值，才会记录算子的耗时到日志
	logThreshold time.Duration
	// strictGeneratedOnly 为 true 时，仅允许执行 gorchc 生成算子
	strictGeneratedOnly bool
	// traceHooks 跟随 executor 生命周期固定，避免执行中动态切换。
	traceHooks TraceHooks

	p Processor
}

var (
	executorPool = pool.NewObjectPool(func(t *executor) {
		t.ctx = nil
		t.startName = ""
		t.timeout = 0
		t.recorder = nil
		t.logOperatorName = false
		t.logThreshold = defaultLogThreshold
		t.strictGeneratedOnly = false
		t.traceHooks = nil
		t.p = nil
	})
)

type executorOptions struct {
	ctx                 *Context
	startName           string
	processor           PrepareProcessor
	timeout             time.Duration
	strictGeneratedOnly bool
	traceHooks          TraceHooks
}

func newExecutor(opt executorOptions) *executor {
	s := executorPool.Get()
	s.ctx = opt.ctx
	s.startName = opt.startName
	s.p = opt.processor
	s.timeout = opt.timeout
	s.strictGeneratedOnly = opt.strictGeneratedOnly
	if s.timeout <= 0 {
		s.timeout = 8 * time.Second
	}
	if opt.traceHooks == nil {
		opt.traceHooks = getTraceHooks()
	}
	s.traceHooks = opt.traceHooks
	return s
}

func (s *executor) Inject(vals ...any) error {
	for i := range vals {
		if err := registerAny(s.ctx.container, vals[i], true); err != nil {
			return err
		}
	}

	return nil
}

func (s *executor) SetTimeout(to time.Duration) {
	if to > 0 {
		s.timeout = to
	}
}

func (s *executor) SetRecorder(r recorder.Recorder) {
	s.recorder = r
}

func (s *executor) SetLogOperatorName() {
	s.logOperatorName = true
}

func (s *executor) SetLogThreshold(t time.Duration) {
	s.logThreshold = t
}

func (s *executor) Execute(ctx context.Context) (retErr error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rctx := s.ctx
	rctx.traceHooks = s.traceHooks
	flowCtx, flowSpan := rctx.traceHooks.StartFlow(ctx, s.startName)
	rctx.ctx = flowCtx
	rctx.recorder = s.recorder
	if rctx.recorder == nil {
		rctx.recorder = recorder.New()
	}
	if s.logOperatorName {
		rctx.recorder.UseOperatorName(true)
	}
	rctx.strictGeneratedOnly = s.strictGeneratedOnly

	rctx.recorder.Start()
	defer func() {
		flowSpan.End(retErr)

		rctx.recorder.Stop()

		rs := rctx.recorder.Report()
		rctx.Logger.Info("EngineExecuteFinish",
			mlog.String("record", rs.Format(s.logThreshold)),
			mlog.String("total", rs.TotalCost.String()),
		)
		rctx.recorder.Reset()

		releaseContext(rctx)
		executorPool.Put(s)
	}()

	done := make(chan error)

	goRoutinePool.Go(func() {
		done <- s.p.Execute(rctx)
	})

	timer := time.NewTimer(s.timeout)
	select {
	case <-timer.C:
		rctx.Exit(errEngineExecuteTimeout)
		<-done
	case <-ctx.Done():
		rctx.Exit(ctx.Err())
		<-done
	case err := <-done:
		rctx.Exit(err)
	}

	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}

	retErr = rctx.err
	return retErr
}

// InjectTyped 向执行器容器注入强类型对象，避免 Inject(vals ...any) 的反射开销。
func InjectTyped[T any](e Executor, ins *T) error {
	exec, ok := e.(*executor)
	if !ok {
		return fmt.Errorf("inject typed error: unsupported executor type %s", reflect.TypeOf(e))
	}
	return registerTyped(exec.ctx.container, ins, true)
}
