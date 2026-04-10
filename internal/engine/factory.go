package engine

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	globalOperatorContainer = &operatorFactoryContainer{
		namesMap: make(map[string]OperatorFactory),
		seqsMap:  make(map[string]string),
	}
)

// getOperatorFactory 获得算子工厂对象
func getOperatorFactory(name string) (OperatorFactory, bool) {
	return globalOperatorContainer.getOperatorFactory(name)
}

type operatorFactoryContainer struct {
	mutex sync.RWMutex
	// namesMap map[算子名]算子对象工厂
	namesMap map[string]OperatorFactory

	// seqsMap map[算子序号]算子名称
	seqsMap map[string]string
}

func (oc *operatorFactoryContainer) registerOperator(name string, of OperatorFactory) error {
	oc.mutex.Lock()
	defer oc.mutex.Unlock()

	_, has := oc.namesMap[name]
	if has {
		return fmt.Errorf("register operator %q error: name conflict", name)
	}

	seq := of.Seq()

	_, has = oc.seqsMap[seq]
	if has {
		return fmt.Errorf("register operator %q error: seq conflict %s", name, seq)
	}

	oc.namesMap[name] = of
	oc.seqsMap[seq] = name
	return nil
}

func (oc *operatorFactoryContainer) getOperatorFactory(name string) (OperatorFactory, bool) {
	oc.mutex.RLock()
	defer oc.mutex.RUnlock()
	of, has := oc.namesMap[name]
	return of, has
}

type OperatorFactory interface {
	Seq() string
	NilProcessor() Processor

	Execute(*Context) error
}

type operatorFactory[T any] struct {
	*sync.Pool
	seq       string
	adp       *operatorFields
	generated bool
}

var (
	fallbackDIExecCount   atomic.Uint64
	generatedExecCount    atomic.Uint64
	strictRejectExecCount atomic.Uint64
)

// newOperatorFactory 创建算子工厂对象
func newOperatorFactory[T any](seq uint) OperatorFactory {
	ff := func() any {
		return new(T)
	}
	of := &operatorFactory[T]{
		&sync.Pool{New: ff},
		strconv.FormatUint(uint64(seq), 10),
		nil,
		isGeneratedOperatorType[T](),
	}
	if !of.generated {
		of.adp = analyzeOperatorFields[T]()
	}
	return of
}

// IsGenerateOperatorCode 标记生成的代码
// 默认情况下，算子的依赖注入和对象导出是通过反射实现的
// 生成的代码，需要实现IsGenerateOperatorCode接口，这样生成的代码可以不需要反射
type IsGenerateOperatorCode interface {
	IsGenerateOperatorCode()
}

func isGeneratedOperatorType[T any]() bool {
	_, ok := any((*T)(nil)).(IsGenerateOperatorCode)
	return ok
}

func (of *operatorFactory[T]) Execute(ctx *Context) (err error) {
	op := of.Pool.Get()
	defer of.Pool.Put(op)

	// 非生成的代码，直接用反射执行对象注入和对象导出
	if !of.generated {
		if ctx.strictGeneratedOnly {
			strictRejectExecCount.Add(1)
			return fmt.Errorf("operator %T must implement IsGenerateOperatorCode in strict mode", op)
		}
		fallbackDIExecCount.Add(1)
		if err = of.adp.inject(op, ctx.container); err != nil {
			return
		}
		defer func() {
			if er0 := of.adp.extract(op, ctx.container); er0 != nil {
				if err == nil {
					err = er0
				} else {
					err = errors.Join(err, er0)
				}
			}
		}()
	} else {
		generatedExecCount.Add(1)
	}

	err = op.(Processor).Execute(ctx)
	return
}

func (of *operatorFactory[T]) NilProcessor() Processor {
	return any((*T)(nil)).(Processor)
}

func (of *operatorFactory[T]) Seq() string {
	return of.seq
}
