package engine

import (
	"fmt"
	"sync"
	"time"

	"github.com/gogorch/gorch/pool"
)

// WaitConfig 等待配置，用于配置routine等待行为
type WaitConfig struct {
	// Timeout 等待超时时间，从开始等待时计算
	// 如果为0，表示无限等待
	Timeout time.Duration

	// TotalTimeout 总超时时间，从routine开始执行时计算
	// 如果设置了此值，会覆盖Timeout的行为
	TotalTimeout time.Duration

	// AllowNotStarted 是否允许等待未启动的routine
	// 如果为false，等待未启动的routine会立即返回错误
	// 如果为true，会等待routine启动后再开始计时
	AllowNotStarted bool
}

// WaitMode 等待模式枚举
type WaitMode int

const (
	// WaitModeTimeout 超时等待模式 - 从开始等待时计时
	WaitModeTimeout WaitMode = iota
	// WaitModeTotalTimeout 总超时模式 - 从routine开始执行时计时
	WaitModeTotalTimeout
	// WaitModeInfinite 无限等待模式
	WaitModeInfinite
)

// String 返回等待模式的字符串描述
func (wm WaitMode) String() string {
	switch wm {
	case WaitModeTimeout:
		return "timeout"
	case WaitModeTotalTimeout:
		return "total-timeout"
	case WaitModeInfinite:
		return "infinite"
	default:
		return "unknown"
	}
}

// GetWaitMode 根据配置获取等待模式
func (wc *WaitConfig) GetWaitMode() WaitMode {
	if wc.TotalTimeout > 0 {
		return WaitModeTotalTimeout
	}
	if wc.Timeout > 0 {
		return WaitModeTimeout
	}
	return WaitModeInfinite
}

// GetEffectiveTimeout 获取有效的超时时间
func (wc *WaitConfig) GetEffectiveTimeout() time.Duration {
	if wc.TotalTimeout > 0 {
		return wc.TotalTimeout
	}
	return wc.Timeout
}

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
	done chan struct{}
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
		r.done = make(chan struct{})
	})
)

func newRoutine(name string) *routine {
	r := routinePool.Get()
	r.name = name
	return r
}

func (r *routine) start(fn func()) (err error) {
	r.mutex.Lock()
	// 检查 routine 是否已经开始
	if !r.startAt.IsZero() {
		r.mutex.Unlock()
		return fmt.Errorf("routine %q is started", r.name)
	}
	r.startAt = time.Now()
	// 在解锁后再执行函数，减少锁的持有时间
	r.mutex.Unlock()

	done := make(chan struct{}, 1)
	go func() {
		defer func() {
			if er := recover(); er != nil {
				err = fmt.Errorf("routine %s execute panic: %v", r.name, er)
			}
			done <- struct{}{}
		}()
		fn()
	}()

	go func() {
		<-done
		r.mutex.Lock()
		r.stopAt = time.Now()
		r.mutex.Unlock()
		// 关闭done channel通知所有等待者
		close(r.done)
	}()

	return nil
}

// wait 方法用于等待一个 routine 完成执行
// 参数:
//   - i: interrupt 对象，用于接收中断信号
//   - config: 等待配置，包含超时设置和行为控制
//
// 返回值:
//   - error: 等待过程中的错误信息
func (r *routine) wait(i *interrupt, config *WaitConfig) (err error) {
	waitStart := time.Now()
	waitMode := config.GetWaitMode()
	effectiveTimeout := config.GetEffectiveTimeout()

	// 首先进行状态检查
	r.mutex.RLock()
	// 检查routine是否已经完成
	if !r.stopAt.IsZero() {
		r.mutex.RUnlock()
		return nil
	}

	// 检查routine是否已经开始执行
	if r.startAt.IsZero() && !config.AllowNotStarted {
		r.mutex.RUnlock()
		return fmt.Errorf("routine %q not started", r.name)
	}

	// 计算超时时间
	var hasTimeout bool
	var deadline time.Time

	if waitMode != WaitModeInfinite {
		hasTimeout = true
		if waitMode == WaitModeTotalTimeout && !r.startAt.IsZero() {
			deadline = r.startAt.Add(effectiveTimeout)
			// 检查是否已经超时
			if time.Now().After(deadline) {
				r.mutex.RUnlock()
				return fmt.Errorf("routine %q execute timeout (total: %v)", r.name, effectiveTimeout)
			}
		} else {
			deadline = waitStart.Add(effectiveTimeout)
		}
	}
	r.mutex.RUnlock()

	// 等待逻辑
	if hasTimeout {
		// 有超时的等待
		remainingTimeout := time.Until(deadline)
		if remainingTimeout <= 0 {
			// 已经超时，进行非阻塞检查
			select {
			case <-r.done:
				return nil
			case <-i.Wait():
				return fmt.Errorf("wait routine %q interrupted", r.name)
			default:
				return fmt.Errorf("wait routine %q timeout (%v, mode: %s)",
					r.name, effectiveTimeout, waitMode)
			}
		}

		timer := time.NewTimer(remainingTimeout)
		defer timer.Stop()

		select {
		case <-r.done:
			// 正常完成，记录阻塞时间
			r.mutex.Lock()
			r.blockCost += time.Since(waitStart)
			r.mutex.Unlock()
			return nil
		case <-i.Wait():
			return fmt.Errorf("wait routine %q interrupted", r.name)
		case <-timer.C:
			// 超时，记录阻塞时间
			r.mutex.Lock()
			r.blockCost += time.Since(waitStart)
			r.mutex.Unlock()
			return fmt.Errorf("wait routine %q timeout (%v, mode: %s)",
				r.name, effectiveTimeout, waitMode)
		}
	} else {
		// 无限等待
		select {
		case <-r.done:
			// 正常完成，记录阻塞时间
			r.mutex.Lock()
			r.blockCost += time.Since(waitStart)
			r.mutex.Unlock()
			return nil
		case <-i.Wait():
			return fmt.Errorf("wait routine %q interrupted", r.name)
		}
	}
}

// BlockCost 方法用于获取等待的时间
func (r *routine) BlockCost() time.Duration {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.blockCost
}

var _ Routine = (*routine)(nil)
