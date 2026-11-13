package engine

import (
	"sync"

	"github.com/gogorch/gorch/pool"
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
	// waiters 所有正在等待中断信号的 goroutine 各自的通道。
	// 之前所有等待者共享一个容量为 1 的 chan，Exit 时只有一个等待者能拿到错误，
	// 其余等待者从已关闭的通道读到 nil，并发场景下表现为偶发空错误。
	waiters []chan error
	err     error
	exit    bool
	mutex   sync.Mutex
}

var (
	interruptPool = pool.NewObjectPool(func(i *interrupt) {
		i.err = nil
		i.exit = false
		i.waiters = nil
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
	if i.exit {
		i.mutex.Unlock()
		return
	}
	i.exit = true
	i.err = err

	waiters := i.waiters
	i.waiters = nil
	i.mutex.Unlock()

	// 广播给所有等待者（通道容量为 1，写入不会阻塞）
	for _, ch := range waiters {
		ch <- err
		close(ch)
	}
}

func (i *interrupt) Wait() <-chan error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	ch := make(chan error, 1)
	if i.exit {
		ch <- i.err
		close(ch)
		return ch
	}

	i.waiters = append(i.waiters, ch)
	return ch
}

func (i *interrupt) Error() error {
	i.mutex.Lock()
	err := i.err
	i.mutex.Unlock()
	return err
}

func (im *interrupt) Exited() bool {
	im.mutex.Lock()
	exit := im.exit
	im.mutex.Unlock()
	return exit
}
