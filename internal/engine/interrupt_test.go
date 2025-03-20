package engine

import (
	"errors"
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
}
