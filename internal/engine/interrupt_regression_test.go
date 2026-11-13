package engine

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 所有等待中断的 goroutine 都必须拿到同一个错误。
// 旧实现共享一个容量为 1 的 channel，只有一个等待者能读到错误，
// 其余等待者从已关闭的 channel 读到 nil。
func TestRegressionInterruptBroadcastAllWaiters(t *testing.T) {
	i := newInterrupt()
	defer releaseInterrupt(i)

	const n = 16
	chans := make([]<-chan error, n)
	for idx := 0; idx < n; idx++ {
		chans[idx] = i.Wait()
	}

	want := errors.New("async exit error")
	i.Exit(want)

	for idx, ch := range chans {
		assert.Equal(t, want, <-ch, "第 %d 个等待者必须收到中断错误", idx)
	}

	// 中断之后新的等待者同样要立刻拿到错误
	assert.Equal(t, want, <-i.Wait())
}

// 中断器被对象池复用后不能带着上一次的中断状态。
func TestRegressionInterruptPoolReset(t *testing.T) {
	i := newInterrupt()
	i.Exit(errors.New("first"))
	releaseInterrupt(i)

	j := newInterrupt()
	defer releaseInterrupt(j)
	assert.False(t, j.Exited())
	assert.Nil(t, j.Error())

	done := make(chan error, 1)
	go func() { done <- <-j.Wait() }()
	j.Exit(errors.New("second"))
	assert.Equal(t, errors.New("second"), <-done)
}

// 并发调用 Exit 只有第一次生效。
func TestRegressionInterruptExitOnlyOnce(t *testing.T) {
	i := newInterrupt()
	defer releaseInterrupt(i)

	ch := i.Wait()
	var wg sync.WaitGroup
	for idx := 0; idx < 8; idx++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			i.Exit(errors.New(string(rune('a' + idx))))
		}(idx)
	}
	wg.Wait()

	assert.Equal(t, i.Error(), <-ch)
	assert.NotNil(t, i.Error())
}
