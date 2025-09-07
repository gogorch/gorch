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
		r := newRoutine("testRoutine")

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
		r := newRoutine("testRoutine")
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

	t.Run("concurrent_routines", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				r := newRoutine("testRoutine")
				i := newInterrupt()
				err := r.start(func() {
					time.Sleep(time.Millisecond * 10)
				})
				assert.Nil(t, err)

				config := &WaitConfig{
					Timeout: time.Millisecond * 20,
				}
				err = r.wait(i, config)
				assert.Nil(t, err)
			}(i)
		}
		wg.Wait()
	})

	t.Run("block_cost_measurement", func(t *testing.T) {
		r := newRoutine("testRoutine")
		i := newInterrupt()

		start := time.Now()
		err := r.start(func() {
			time.Sleep(time.Millisecond * 50)
		})
		assert.Nil(t, err)

		config := &WaitConfig{
			Timeout: time.Millisecond * 100,
		}
		err = r.wait(i, config)
		assert.Nil(t, err)

		blockCost := r.BlockCost()
		actualCost := time.Since(start)
		assert.True(t, blockCost > 0, "block cost should be positive")
		assert.True(t, blockCost <= actualCost, "block cost should not exceed actual cost")
	})

	t.Run("routine_pool_reuse", func(t *testing.T) {
		// Get routine from pool
		r1 := newRoutine("testRoutine1")
		routinePool.Put(r1)

		// Get again should reuse the same instance
		r2 := newRoutine("testRoutine2")
		assert.Equal(t, r1, r2, "should reuse routine from pool")
		assert.Equal(t, "testRoutine2", r2.name, "name should be updated")
	})

	t.Run("panic_recovery", func(t *testing.T) {
		r := newRoutine("testRoutine")
		assert.NotNil(t, r, "routine should not be nil")

		err := r.start(func() {
			panic("test panic") // 直接panic，不捕获
		})
		assert.NotNil(t, err, "should return error for panic")
		if err != nil {
			assert.Contains(t, err.Error(), "routine testRoutine execute panic: test panic")
		}
	})
}
