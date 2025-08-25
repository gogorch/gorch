package gorch

import (
	"github.com/gogorch/gorch/internal/engine"
	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/gogorch/gorch/internal/ort"
	"github.com/gogorch/gorch/mlog"
)

type Context = engine.Context
type Processor = engine.Processor
type Logger = mlog.Logger

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

// AnalyzeOperator 分析泛型类型 T 的结构体信息，提取注入和提取字段，并判断是否实现了 Prepare 方法。
// 返回一个包含分析结果的 OperatorRType 指针。
func AnalyzeOperator[T any]() *ort.OperatorRType {
	return ort.AnalyzeOperator[T]()
}
