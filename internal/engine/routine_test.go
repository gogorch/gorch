package engine

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoutine(t *testing.T) {
	t.Run("not_started", func(t *testing.T) {
		r := &routine{
			name:  "testRoutine",
			done:  make(chan struct{}),
			mutex: sync.RWMutex{},
		}

		i := newInterrupt()

		// 模拟未开始，不允许等待未启动的routine
		config := &WaitConfig{
			Timeout:         time.Second,
			AllowNotStarted: false,
		}
		err := r.wait(i, config)
		assert.Equal(t, errors.New("routine \"testRoutine\" not started"), err)
	})

	t.Run("started_and_timeout", func(t *testing.T) {
		r := &routine{
			name:  "testRoutine",
			done:  make(chan struct{}),
			mutex: sync.RWMutex{},
		}
		i := newInterrupt()
		// 模拟已经开始
		// 等待超时
		r.startAt = time.Now()
		config := &WaitConfig{
			Timeout: time.Millisecond * 20,
		}
		err := r.wait(i, config)
		assert.Contains(t, err.Error(), "wait routine \"testRoutine\" timeout")
	})

	t.Run("started_and_execute_timeout", func(t *testing.T) {
		r := &routine{
			name:  "testRoutine",
			done:  make(chan struct{}),
			mutex: sync.RWMutex{},
		}
		i := newInterrupt()
		// 模拟已经开始
		// 等待超时（总超时模式）
		r.startAt = time.Now()
		time.Sleep(time.Millisecond * 100)
		config := &WaitConfig{
			TotalTimeout: time.Millisecond * 20,
		}
		err := r.wait(i, config)
		assert.Contains(t, err.Error(), "routine \"testRoutine\" execute timeout")
	})

	t.Run("from_start_success", func(t *testing.T) {
		r := &routine{
			name:  "testRoutine",
			done:  make(chan struct{}),
			mutex: sync.RWMutex{},
		}
		i := newInterrupt()
		// 模拟已经开始
		// 从启动计算，等待成功
		go func() {
			r.start(func() {
				time.Sleep(time.Millisecond * 20)
			})
		}()
		time.Sleep(time.Millisecond * 2)
		config := &WaitConfig{
			TotalTimeout: time.Millisecond * 40,
		}
		err := r.wait(i, config)
		assert.Nil(t, err)
	})

	t.Run("from_start_not_check_start", func(t *testing.T) {
		// 创建一个 routine 实例
		r := &routine{
			name:  "testRoutine",
			done:  make(chan struct{}),
			mutex: sync.RWMutex{},
		}
		i := newInterrupt()
		// 等待过程中中断了
		r.startAt = zeroTime
		go func() {
			time.Sleep(time.Millisecond * 50)
			i.Exit(errors.New("interrupted"))
		}()
		r.startAt = time.Now()
		config := &WaitConfig{
			Timeout: time.Second,
		}
		err := r.wait(i, config)
		assert.Equal(t, errors.New("wait routine \"testRoutine\" interrupted"), err)
	})
}
