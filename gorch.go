package gorch

import (
	"github.com/gogorch/gorch/internal/engine"
	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/gogorch/gorch/internal/ort"
	"github.com/gogorch/gorch/mlog"
)

type Context = engine.Context
type Engine = engine.Engine
type Executor = engine.Executor
type Processor = engine.Processor
type RuntimeStats = engine.RuntimeStats
type Logger = mlog.Logger
type TraceHooks = engine.TraceHooks
type TraceSpan = engine.TraceSpan

var (
	// New 创建执行引擎
	New = engine.New

	// SetGoroutinePool 设置goroutine池
	SetGoroutinePool = engine.SetGoroutinePool

	// ParseDSL 解析gorch文件, 返回*ast.Primary对象
	// *ast.Primary用于在NewEngine函数创建Engine对象
	ParseDSL = iparser.Parse

	// ParseDSLInDir 解析包含gorch文件的目录，将目录内所有的gorch文件依次解析到同一个*ast.Primary对象中
	// 最后返回*ast.Primary对象
	ParseDSLInDir = iparser.ParseDirectory

	// NewOperatorStates 创建算子状态集合
	NewOperatorStates = engine.NewOperatorStates

	// SetDefaultLogger 设置默认的日志对象
	SetDefaultLogger = engine.SetDefaultLogger

	// SetTraceHooks 设置引擎全局追踪钩子（用于接入外部 tracing 实现，如 OpenTelemetry）。
	SetTraceHooks = engine.SetTraceHooks
)

// RegOp RegisterOperator 注册算子
// 使用方法：RegOp[OperatorStruct]("OperatorName", 1)
//
//	使用泛型的方式调用，在Register函数内部自动创建工厂函数
//
// name: 算子名称，可以和struct名称不同，全局唯一
// seq: 算子序号，全局唯一
func RegOp[T any](name string, seq uint) error {
	return engine.RegisterOperator[T](name, seq)
}

// MustPtr 返回指针
//
//	使用方法：MustPtr(1)
func MustPtr[T any](val T) *T {
	return &val
}

// RegisterTyped 注册强类型对象到当前执行上下文容器中。
func RegisterTyped[T any](ctx *Context, ins *T, replace bool) error {
	return engine.RegisterTyped(ctx, ins, replace)
}

// MutableTyped 从当前执行上下文容器中读取强类型对象。
func MutableTyped[T any](ctx *Context, val *T) error {
	return engine.MutableTyped(ctx, val)
}

// InjectTyped 向执行器注入强类型对象，避免 Inject(vals ...any) 的反射开销。
func InjectTyped[T any](executor Executor, ins *T) error {
	return engine.InjectTyped(executor, ins)
}

// SetStrictGeneratedOnly 设置执行引擎严格模式（实例级配置）。
func SetStrictGeneratedOnly(eng Engine, enabled bool) {
	if eng != nil {
		eng.SetStrictGeneratedOnly(enabled)
	}
}

// SnapshotRuntimeStats 返回运行时观测数据快照。
func SnapshotRuntimeStats() RuntimeStats {
	return engine.SnapshotRuntimeStats()
}

// AnalyzeOperator 分析算子结构体的 DI 元信息（供 gorchc 代码生成使用）。
func AnalyzeOperator[T any]() *ort.OperatorRType {
	return ort.AnalyzeOperator[T]()
}
