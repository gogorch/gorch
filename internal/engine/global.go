package engine

import (
	"time"

	"github.com/gogorch/gorch/mlog"
	"github.com/gogorch/gorch/pool"
)

var (
	// printLogger print日志，默认的日志打印接口
	printLogger mlog.Logger = mlog.NewPrintLogger()

	// goRoutinePool goroutine池对象
	goRoutinePool pool.GoRoutinePool = new(pool.NotPool)

	// 全局日志对象，默认使用printLogger
	globalLogger mlog.Logger = printLogger

	zeroTime = time.Time{}
)

// SetGoroutinePool 允许从外部设置一个goroutine池对象
func SetGoroutinePool(pool pool.GoRoutinePool) {
	goRoutinePool = pool
}

// RegisterOperator 注册算子对象工厂
// 参数:
//   - name: 算子名称，用于唯一标识算子，不能重复。
//   - seq: 算子的序号，同样全局唯一，不能重复。
func RegisterOperator[T any](name string, seq uint) error {
	return globalOperatorContainer.registerOperator(name, newOperatorFactory[T](seq))
}

// clean 清理所有对象
func clean() {
	globalOperatorContainer = &operatorFactoryContainer{
		namesMap: make(map[string]OperatorFactory),
		seqsMap:  make(map[string]string),
	}
}

func SetDefaultLogger(l mlog.Logger) {
	globalLogger = l
}
