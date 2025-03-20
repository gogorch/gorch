package pool

import "sync"

// SlicePool 切片对象池
type SlicePool[T any] interface {
	Get() []T
	Put([]T)
}

type slicePool[T any] struct {
	defaultLen int
	*sync.Pool
}

func NewSlicePool[T any](defaultLen, defaultCap int) SlicePool[T] {
	s := func() any {
		s := make([]T, defaultLen, defaultCap)
		return &s
	}

	return &slicePool[T]{defaultLen, &sync.Pool{New: s}}
}

func (sp *slicePool[T]) Get() []T {
	s := sp.Pool.Get().(*[]T)
	return (*s)[:sp.defaultLen]
}

func (sp *slicePool[T]) Put(s []T) {
	if s == nil {
		return
	}
	sp.Pool.Put(&s)
}
