package engine

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorch/gorch/pool"
)

// Routine 异步执行的对象
// 在dsl中，使用 `GO(operator, "name")` 发起一次异步执行
// 发起之后，就可以通过 `WAIT` 指令等待routine的结果。示例：`operator(arg=1, WAIT("name"))`
type Routine interface {
	// BlockCost 方法用于获取等待routine完成而阻塞当前goroutine的耗时
	BlockCost() time.Duration
}

type routine struct {
	mutex sync.RWMutex
	// name routine名称，必须全局唯一
	name string
	// result routine的执行结果，只能是error或者nil
	// 框架内部会用这个属性==nil来判断是否
	signal chan struct{}
	// startAt routine开始执行的时间
	startAt time.Time
	// stopAt routine结束执行的时间
	stopAt time.Time
	// blockCost routine等待阻塞的时间
	blockCost time.Duration
}

var (
	routinePool = pool.NewObjectPool(func(r *routine) {
		r.name = ""
		r.startAt = zeroTime
		r.stopAt = zeroTime
		r.blockCost = 0
		r.signal = make(chan struct{}, 1)
	})
)

func newRoutine(name string) *routine {
	r := routinePool.Get()
	r.name = name
	return r
}

func (r *routine) start(fn func()) (err error) {
	r.mutex.Lock()
	if r.startAt.IsZero() {
		r.startAt = time.Now()
		r.mutex.Unlock()
		func() {
			if er := recover(); er != nil {
				err = fmt.Errorf("routine %s execute panic", r.name)
			}
			defer func() {
				r.mutex.Lock()
				r.signal <- struct{}{}
				r.stopAt = time.Now()
				r.mutex.Unlock()
			}()
			fn()
		}()
		return
	}

	r.mutex.Unlock()
	err = fmt.Errorf("routine %q is started", r.name)
	return
}

// wait 方法用于等待一个 routine 完成执行，并且可以设置超时时间。
// 参数:
//   - i: 一个 interrupt 对象，用于在等待过程中接收中断信号。
//   - timeout: 等待的超时时间。
//   - fromStart: 是否从 routine 开始执行时开始计时。
//   - notCheckStart: 不检查 routine 是否已经开始执行。
//
// 返回值:
//   - error: 如果等待超时或者 routine 执行出错，返回相应的错误信息。
func (r *routine) wait(i *interrupt, timeout time.Duration, fromStart bool, notCheckStart bool) (err error) {
	r.mutex.RLock()
	if r.startAt.IsZero() {
		if !notCheckStart {
			r.mutex.RUnlock()
			return fmt.Errorf("routine %q not started", r.name)
		}
	}
	// routine已经执行完成，直接返回
	if !r.stopAt.IsZero() {
		r.mutex.RUnlock()
		return nil
	}

	if fromStart && !r.startAt.IsZero() {
		elapsed := time.Since(r.startAt)
		if elapsed > timeout {
			r.mutex.RUnlock()
			return fmt.Errorf("routine %q execute timeout", r.name)
		}
		timeout = timeout - elapsed
	}
	r.mutex.RUnlock()

	now := time.Now()
	if timeout > 0 {
		timer := time.NewTimer(timeout)

		select {
		case <-r.signal:
		case <-i.Wait():
			err = fmt.Errorf("wait routine %q interrupted", r.name)
		case <-timer.C:
			r.mutex.Lock()
			r.blockCost = time.Since(time.Now())
			r.mutex.Unlock()
			err = fmt.Errorf("wait routine %q timeout", r.name)
		}

		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
	} else {
		select {
		case <-r.signal:
		case <-i.Wait():
			err = fmt.Errorf("wait routine %q interrupted", r.name)
		}
	}

	r.blockCost += time.Since(now)
	return err
}

// BlockCost 方法用于获取等待的时间
func (r *routine) BlockCost() time.Duration {
	return r.blockCost
}

var _ Routine = (*routine)(nil)
