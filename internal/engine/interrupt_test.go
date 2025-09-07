package engine

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInterrupt(t *testing.T) {
	t.Run("error_exit", func(t *testing.T) {
		i := newInterrupt()
		// 有错误中断
		err := errors.New("test error")
		i.Exit(err)
		assert.Equal(t, errors.New("test error"), i.Error())
		assert.True(t, i.Exited(), "should be exited")
	})

	t.Run("nil_exit", func(t *testing.T) {
		i := newInterrupt()
		// 无错误中断
		i.Exit(nil)
		assert.Equal(t, nil, i.Error())
	})

	t.Run("async_wait_error_exit", func(t *testing.T) {
		i := newInterrupt()
		// 无错误中断
		go func() {
			time.Sleep(time.Millisecond * 50)
			i.Exit(errors.New("async exit error"))
		}()
		assert.Equal(t, errors.New("async exit error"), <-i.Wait())
		assert.Equal(t, errors.New("async exit error"), <-i.Wait()) // 多次wait，返回同一个错误
		assert.Equal(t, errors.New("async exit error"), <-i.Wait())
	})

	t.Run("async_wait_nil_exit", func(t *testing.T) {
		i := newInterrupt()
		// 无错误中断
		go func() {
			time.Sleep(time.Millisecond * 50)
			i.Exit(nil)
		}()
		assert.Equal(t, nil, <-i.Wait())
		assert.Equal(t, nil, <-i.Wait())
		assert.Equal(t, nil, <-i.Wait())
	})

	t.Run("concurrent_interrupt", func(t *testing.T) {
		i := newInterrupt()
		var wg sync.WaitGroup

		// 并发中断
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				i.Exit(errors.New("concurrent error"))
			}(j)
		}

		// 并发等待
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := <-i.Wait()
				assert.Equal(t, errors.New("concurrent error"), err)
			}()
		}

		wg.Wait()
	})

	t.Run("interrupt_pool_reuse", func(t *testing.T) {
		// Get interrupt from pool
		i1 := newInterrupt()
		interruptPool.Put(i1)

		// Get again should reuse the same instance
		i2 := newInterrupt()
		assert.Equal(t, i1, i2, "should reuse interrupt from pool")
		assert.False(t, i2.Exited(), "should not be exited after reuse")
	})

	t.Run("double_exit", func(t *testing.T) {
		i := newInterrupt()
		i.Exit(errors.New("first error"))
		i.Exit(errors.New("second error")) // Should be ignored
		assert.Equal(t, errors.New("first error"), i.Error())
	})
}
