package engine

import (
	"context"
	"sync"
)

// TraceSpan 表示一次可结束的追踪片段。
// 具体实现可由外部适配器提供（例如 OpenTelemetry Span）。
type TraceSpan interface {
	End(err error)
}

// TraceHooks 提供引擎生命周期的可插拔追踪钩子。
// gorch 内部不依赖任何具体 tracing 实现。
type TraceHooks interface {
	StartFlow(ctx context.Context, flow string) (context.Context, TraceSpan)
	StartNode(ctx context.Context, node string, kind string) (context.Context, TraceSpan)
}

type noopTraceHooks struct{}

type noopTraceSpan struct{}

func (noopTraceHooks) StartFlow(ctx context.Context, _ string) (context.Context, TraceSpan) {
	return ctx, noopTraceSpan{}
}

func (noopTraceHooks) StartNode(ctx context.Context, _, _ string) (context.Context, TraceSpan) {
	return ctx, noopTraceSpan{}
}

func (noopTraceSpan) End(error) {}

var (
	traceHooksMu sync.RWMutex
	traceHooks   TraceHooks = noopTraceHooks{}
)

// SetTraceHooks 设置全局 TraceHooks。传入 nil 时回退到 no-op。
func SetTraceHooks(hooks TraceHooks) {
	traceHooksMu.Lock()
	if hooks == nil {
		traceHooks = noopTraceHooks{}
	} else {
		traceHooks = hooks
	}
	traceHooksMu.Unlock()
}

func getTraceHooks() TraceHooks {
	traceHooksMu.RLock()
	h := traceHooks
	traceHooksMu.RUnlock()
	return h
}
