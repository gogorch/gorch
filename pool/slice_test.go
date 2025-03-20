package pool

import (
	"fmt"
	"testing"
)

func TestNewSlicePool(t *testing.T) {
	sp := NewSlicePool[int](1, 10)
	for i := 0; i < 100; i++ {
		s := sp.Get()
		s = append(s, i)
		sp.Put(s)
		fmt.Println(s)
	}
	// for i := 0; i < 10; i++ {
	// 	s := sp.Get()
	// 	fmt.Println(s)
	// }
	// fmt.Println(s, (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}
