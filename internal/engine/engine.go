package engine

import (
	"context"
	"time"

	"github.com/gorch/gorch/internal/lang/ast"
	"github.com/gorch/gorch/mlog"
	"github.com/gorch/gorch/pool"
	"github.com/gorch/gorch/recorder"
)

type Engine interface {
	Start(string) Executor
}

type engine struct {
	starters map[string]PrepareProcessor
}

func (eng *engine) Start(key string) Executor {
	st, has := eng.starters[key]
	if has {
		return newExecutor(NewContext(), st, 8*time.Second)
	}

	return nil
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
	SetCostThreshold(t time.Duration)

	// 执行算子
	Execute(ctx context.Context) error
}

type executor struct {
	ctx *Context

	// 执行超时时间
	timeout time.Duration
	// 记录算子的执行结果
	recorder recorder.Recorder
	// 日志里记录执行的算子名称
	// 默认情况下，日志里记录的是算子的序号
	logOperatorName bool
	// 打印算子耗时的阈值
	// 当算子的耗时大于这个阈值，才会记录算子的耗时到日志
	costThreshold time.Duration

	p Processor
}

var (
	executorPool = pool.NewObjectPool(func(t *executor) {
		t.ctx = nil
		t.timeout = 0
		t.recorder = nil
		t.logOperatorName = false
		t.costThreshold = 0
		t.p = nil
	})
)

func newExecutor(ctx *Context, p PrepareProcessor, to time.Duration) *executor {
	s := executorPool.Get()
	s.ctx = ctx
	s.p = p
	s.timeout = to
	return s
}

func (s *executor) Inject(vals ...any) error {
	for i := range vals {
		if err := s.ctx.RegisterIns(vals[i], true); err != nil {
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

func (s *executor) SetCostThreshold(t time.Duration) {
	s.costThreshold = t
}

func (s *executor) Execute(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rctx := s.ctx
	rctx.ctx = ctx
	rctx.recorder = s.recorder
	if rctx.recorder == nil {
		rctx.recorder = recorder.NewRecorder()
	}
	if s.logOperatorName {
		rctx.recorder.SetLogOperatorName()
	}
	if s.costThreshold > 0 {
		rctx.recorder.SetCostThreshold(s.costThreshold)
	}

	rctx.recorder.Start()
	defer func() {
		rctx.recorder.Stop()

		rs := rctx.recorder.OperatorRecordStr()
		rctx.Logger.AddInfo(mlog.String("record", rs))
		rctx.Logger.AddInfo(mlog.String("total", rctx.recorder.TotalCost()))
		rctx.recorder.PutPool()

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

	return rctx.err
}
