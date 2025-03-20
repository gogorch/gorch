package pool

// GoRoutinePool 协程池
type GoRoutinePool interface {
	Go(func())
}

// NotPool 非goroutine池
type NotPool struct {
}

func (s *NotPool) Go(fn func()) {
	go fn()
}
