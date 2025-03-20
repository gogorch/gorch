package pool

import (
	"bytes"
	"sync"
)

var (
	// BytesBufferPool bytes.Buffer对象池
	BytesBufferPool = NewObjectPool(func(b *bytes.Buffer) { b.Reset() })

	// WaitGroupPool sync.WaitGroup对象池
	WaitGroupPool = NewObjectPool(func(wg *sync.WaitGroup) {})

	// MutexPool sync.Mutex对象池
	MutexPool = NewObjectPool(func(m *sync.Mutex) {})

	// RWMutexPool sync.RWMutex对象池
	RWMutexPool = NewObjectPool(func(rw *sync.RWMutex) {})
)
