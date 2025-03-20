package pool

import (
	"sync"
)

// ObjectPool 对象池接口
// 通过NewObjectPool创建一个对象池
type ObjectPool[T any] interface {
	Get() *T

	Put(*T)
}

type objectPool[T any] struct {
	*sync.Pool

	creator func() any
	reset   func(*T)
}

func NewObjectPool[T any](reset func(*T)) ObjectPool[T] {
	ff := func() any {
		return new(T)
	}
	return &objectPool[T]{
		&sync.Pool{New: ff}, ff, reset,
	}
}

func (op *objectPool[T]) Get() *T {
	s := op.Pool.Get().(*T)
	if op.reset != nil {
		op.reset(s)
	}
	return s
}

func (op *objectPool[T]) Put(obj *T) {
	if obj == nil {
		return
	}
	op.Pool.Put(obj)
}
