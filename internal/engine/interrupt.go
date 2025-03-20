package engine

import (
	"sync"

	"github.com/gorch/gorch/pool"
)

// 中断器，每次执行引擎时自动创建的对象，方便用户可以随时终止程序执行
// 执行引擎在执行过程中可以随时调用Exit()方法，向所有Wait()的goroutine发送中断信号
type Interrupt interface {
	// 对执行引擎发起中断
	// 允许传入nil，表示发起无异常中断
	// 当执行引擎收到中断信号后，所有Wait的goroutine都会收到中断信号
	// 传入nil时，Wait()返回nil，传入错误时，Wait()返回相应错误
	// 一个中断器只能中断一次，重复中断无效
	Exit(err error)

	// 等待中断器的中断信号
	Wait() <-chan error

	// 获取中断器的错误信息，Exit(nil)时，返回也是nil
	Error() error

	// 判断是否已经中断
	Exited() bool
}

type interrupt struct {
	notify chan error
	err    error
	exit   bool
	mutex  sync.RWMutex
}

var (
	interruptPool = pool.NewObjectPool(func(i *interrupt) {
		i.err = nil
		i.exit = false
		i.notify = make(chan error, 1)
	})
)

func newInterrupt() *interrupt {
	return interruptPool.Get()
}

func releaseInterrupt(i *interrupt) {
	if i != nil {
		interruptPool.Put(i)
	}
}

func (i *interrupt) Exit(err error) {
	i.mutex.Lock()
	if !i.exit {
		i.exit = true
		i.err = err
		i.notify <- err
		close(i.notify)
	}
	i.mutex.Unlock()
}

func (i *interrupt) Wait() <-chan error {
	if i.Exited() {
		cpy := make(chan error, 1)
		cpy <- i.err
		return cpy
	}
	return i.notify
}

func (i *interrupt) Error() error {
	i.mutex.RLock()
	err := i.err
	i.mutex.RUnlock()
	return err
}

func (im *interrupt) Exited() bool {
	im.mutex.RLock()
	exit := im.exit
	im.mutex.RUnlock()
	return exit
}
