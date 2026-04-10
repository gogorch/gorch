package engine

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type traceEvent struct {
	kind string
	name string
	err  string
}

type testTraceHooks struct {
	mu     sync.Mutex
	starts []traceEvent
	ends   []traceEvent
}

func (h *testTraceHooks) StartFlow(ctx context.Context, flow string) (context.Context, TraceSpan) {
	h.mu.Lock()
	h.starts = append(h.starts, traceEvent{kind: "flow", name: flow})
	h.mu.Unlock()
	return ctx, &testTraceSpan{hooks: h, kind: "flow", name: flow}
}

func (h *testTraceHooks) StartNode(ctx context.Context, node string, kind string) (context.Context, TraceSpan) {
	h.mu.Lock()
	h.starts = append(h.starts, traceEvent{kind: kind, name: node})
	h.mu.Unlock()
	return ctx, &testTraceSpan{hooks: h, kind: kind, name: node}
}

func (h *testTraceHooks) hasStart(kind, name string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, one := range h.starts {
		if one.kind == kind && one.name == name {
			return true
		}
	}
	return false
}

func (h *testTraceHooks) hasEnd(kind, name string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, one := range h.ends {
		if one.kind == kind && one.name == name {
			return true
		}
	}
	return false
}

func (h *testTraceHooks) endEvent(kind, name string) (traceEvent, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, one := range h.ends {
		if one.kind == kind && one.name == name {
			return one, true
		}
	}
	return traceEvent{}, false
}

type testTraceSpan struct {
	hooks *testTraceHooks
	kind  string
	name  string
}

func (s *testTraceSpan) End(err error) {
	evt := traceEvent{kind: s.kind, name: s.name}
	if err != nil {
		evt.err = err.Error()
	}
	s.hooks.mu.Lock()
	s.hooks.ends = append(s.hooks.ends, evt)
	s.hooks.mu.Unlock()
}

func TestTraceHooks(t *testing.T) {
	t.Run("flow_and_operator", func(t *testing.T) {
		hooks := &testTraceHooks{}
		SetTraceHooks(hooks)
		defer SetTraceHooks(nil)

		MyTestRun(t, TestValus{
			name: "test",
			g:    `START("test"){ NO_CHECK_MISS() ChangeValueOP(val=2) }`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})

		assert.True(t, hooks.hasStart("flow", "test"))
		assert.True(t, hooks.hasStart("operator", "ChangeValueOP"))
		assert.True(t, hooks.hasEnd("flow", "test"))
		assert.True(t, hooks.hasEnd("operator", "ChangeValueOP"))
	})

	t.Run("trace_directive", func(t *testing.T) {
		hooks := &testTraceHooks{}
		SetTraceHooks(hooks)
		defer SetTraceHooks(nil)

		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				TRACE("segment_1"){
					ChangeValueOP(val=2)
				}
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})

		assert.True(t, hooks.hasStart("trace", "segment_1"))
		assert.True(t, hooks.hasEnd("trace", "segment_1"))
		assert.True(t, hooks.hasStart("operator", "ChangeValueOP"))
		assert.True(t, hooks.hasEnd("operator", "ChangeValueOP"))
	})

	t.Run("trace_directive_with_error", func(t *testing.T) {
		hooks := &testTraceHooks{}
		SetTraceHooks(hooks)
		defer SetTraceHooks(nil)

		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				TRACE("segment_error"){
					TestOperatorState(status=1)
				}
			}`,
			assertFunc: func(res error, _ *BeChangeValue) {
				assert.Equal(t, fatal1, res)
			},
		})

		traceEnd, ok := hooks.endEvent("trace", "segment_error")
		assert.True(t, ok)
		assert.Equal(t, fatal1.Error(), traceEnd.err)

		operatorEnd, ok := hooks.endEvent("operator", "TestOperatorState")
		assert.True(t, ok)
		assert.Equal(t, fatal1.Error(), operatorEnd.err)

		flowEnd, ok := hooks.endEvent("flow", "test")
		assert.True(t, ok)
		assert.Equal(t, fatal1.Error(), flowEnd.err)
	})

	t.Run("noop_span_end", func(t *testing.T) {
		ctx, span := noopTraceHooks{}.StartFlow(context.Background(), "noop")
		assert.NotNil(t, ctx)
		span.End(nil)
	})
}
