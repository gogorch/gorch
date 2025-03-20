package engine

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoutine(t *testing.T) {
	t.Run("not_started", func(t *testing.T) {
		r := &routine{
			name: "testRoutine",
		}

		i := newInterrupt()

		// 模拟未开始
		err := r.wait(i, time.Second, false, false)
		assert.Equal(t, errors.New("routine \"testRoutine\" not started"), err)
	})

	t.Run("started_and_timeout", func(t *testing.T) {
		r := &routine{
			name: "testRoutine",
		}
		i := newInterrupt()
		// 模拟已经开始
		// 等待超时
		r.startAt = time.Now()
		err := r.wait(i, time.Millisecond*20, false, false)
		assert.Equal(t, errors.New("wait routine \"testRoutine\" timeout"), err)
	})

	t.Run("started_and_execute_timeout", func(t *testing.T) {
		r := &routine{
			name: "testRoutine",
		}
		i := newInterrupt()
		// 模拟已经开始
		// 等待超时
		r.startAt = time.Now()
		time.Sleep(time.Millisecond * 100)
		err := r.wait(i, time.Millisecond*20, true, false)
		assert.Equal(t, errors.New("routine \"testRoutine\" execute timeout"), err)
	})

	t.Run("from_start_success", func(t *testing.T) {
		r := &routine{
			name:   "testRoutine",
			signal: make(chan struct{}, 1),
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
		err := r.wait(i, time.Millisecond*40, true, false)
		assert.Nil(t, nil, err)
	})

	t.Run("from_start_not_check_start", func(t *testing.T) {
		// 创建一个 routine 实例
		r := &routine{
			name: "testRoutine",
		}
		i := newInterrupt()
		// 等待过程中中断了
		r.startAt = zeroTime
		go func() {
			time.Sleep(time.Millisecond * 50)
			i.Exit(errors.New("interrupted"))
		}()
		r.startAt = time.Now()
		err := r.wait(i, time.Second, false, false)
		assert.Equal(t, errors.New("wait routine \"testRoutine\" interrupted"), err)
	})
}
