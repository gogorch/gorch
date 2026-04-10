package engine

import (
	"log"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 实现Releasable接口的测试结构体，其Release方法会触发panic
type PanicReleasable struct {
	Name string
}

func (p *PanicReleasable) Release() {
	panic("test panic in Release")
}

// 正常的Releasable实现，用于对比测试
type NormalReleasable struct {
	Name     string
	Released bool
}

func (n *NormalReleasable) Release() {
	n.Released = true
	log.Printf("NormalReleasable %s released\n", n.Name)
}

// 测试releaseInstances函数对panic的处理
func TestReleaseInstancesPanic(t *testing.T) {
	// 创建测试实例
	panicIns := &PanicReleasable{Name: "panic"}
	normalIns := &NormalReleasable{Name: "normal", Released: false}

	cter := newContainer()
	registerAny(cter, &panicIns, false)
	registerAny(cter, &normalIns, false)

	// 调用releaseInstances，这应该会触发panic但被捕获
	releaseContainer(cter)

	// 验证正常实例被正确释放
	assert.True(t, normalIns.Released, "normalIns should be released")

	// 如果程序能执行到这里，说明panic被成功捕获
	assert.True(t, true, "panic was successfully captured")
}

type MyStruct struct {
	name     string
	released bool
}

// 实现 Release 接口
type MyReleasableStruct struct {
	name     string
	released bool
}

func (m *MyReleasableStruct) Release() { m.released = true }

type MyInterface interface{}
type MyReleasableInterface interface{ Release() }

func TestReleaseInstances(t *testing.T) {
	cter := newContainer()

	// 基础类型
	a := 123
	registerAny(cter, &a, false)

	// 指针类型未实现 Release
	b := &MyStruct{name: "b"}
	registerAny(cter, &b, false)

	// 指针类型实现 Release
	c := &MyReleasableStruct{name: "c"}
	registerAny(cter, &c, false)

	// 接口类型未实现 Release
	var d MyInterface = &MyStruct{name: "d"}
	registerAny(cter, &d, false)

	// 接口类型实现 Release
	var e MyReleasableInterface = &MyReleasableStruct{name: "e"}
	registerAny(cter, &e, false)

	// 释放
	releaseInstances(cter.typedInstances)

	// 验证
	assert.Equal(t, false, b.released, "b should not be released")
	assert.Equal(t, true, c.released, "c should be released")
	assert.Equal(t, false, d.(*MyStruct).released, "d should not be released")
	assert.Equal(t, true, e.(*MyReleasableStruct).released, "e should be released")
}

func TestRegisterInsAndMutables(t *testing.T) {
	cter := newContainer()

	type MyStruct struct {
		name string
	}

	type MyInterface interface{}

	t.Run("no_replace", func(t *testing.T) {
		cases := []struct {
			name string
			ins  func() any
			val  func() any
		}{
			{name: "int", ins: func() any { a := 123; return &a }, val: func() any { var i int; return &i }},
			{name: "string", ins: func() any { a := "string"; return &a }, val: func() any { var s string; return &s }},
			{name: "float32", ins: func() any { a := float32(123.0); return &a }, val: func() any { var a float32; return &a }},
			{name: "float64", ins: func() any { a := float64(123.0); return &a }, val: func() any { var a float64; return &a }},
			{name: "mystruct", ins: func() any { a := MyStruct{name: "mystruct"}; return &a }, val: func() any { var a MyStruct; return &a }},
			{name: "mystruct_ptr", ins: func() any { a := &MyStruct{name: "mystruct"}; return &a }, val: func() any { var a *MyStruct; return &a }},
			{name: "my_interface", ins: func() any { a := MyInterface(MyStruct{name: "mystruct"}); return &a }, val: func() any { var a MyInterface; return &a }},
		}

		for ins := range cases {
			cs := cases[ins]
			t.Run(cs.name, func(t *testing.T) {
				ins := cs.ins()
				assert.Nil(t, registerAny(cter, ins, false))

				val := cs.val()

				assert.Nil(t, mutableAny(cter, val))
				assert.Equal(t, ins, val)
			})
		}
	})

}

func TestRegisterInsAndMutableRepalce(t *testing.T) {
	cter := newContainer()

	type MyStruct struct {
		name string
	}

	type MyInterface interface{}

	t.Run("replace_int", func(t *testing.T) {
		a := 123
		assert.Nil(t, registerAny(cter, &a, false))
		b := 456
		assert.EqualValues(t, "register error: duplicate type, error type: int", registerAny(cter, &b, false).Error())
		assert.Nil(t, registerAny(cter, &b, true))

		var c int
		assert.Nil(t, mutableAny(cter, &c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_string", func(t *testing.T) {
		a := "123"
		assert.Nil(t, registerAny(cter, &a, false))
		b := "456"
		assert.EqualValues(t, "register error: duplicate type, error type: string", registerAny(cter, &b, false).Error())
		assert.Nil(t, registerAny(cter, &b, true))

		var c string
		assert.Nil(t, mutableAny(cter, &c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_float32", func(t *testing.T) {
		a := float32(123)
		assert.Nil(t, registerAny(cter, &a, false))
		b := float32(456)
		assert.EqualValues(t, "register error: duplicate type, error type: float32", registerAny(cter, &b, false).Error())
		assert.Nil(t, registerAny(cter, &b, true))

		var c float32
		assert.Nil(t, mutableAny(cter, &c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_struct", func(t *testing.T) {
		a := MyStruct{name: "mystruct"}
		assert.Nil(t, registerAny(cter, &a, false))
		b := MyStruct{name: "mystruct1"}
		assert.EqualValues(t, "register error: duplicate type, error type: engine.MyStruct", registerAny(cter, &b, false).Error())
		assert.Nil(t, registerAny(cter, &b, true))

		var c MyStruct
		assert.Nil(t, mutableAny(cter, &c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_struct_ptr", func(t *testing.T) {
		a := &MyStruct{name: "mystruct"}
		assert.Nil(t, registerAny(cter, &a, false))
		b := &MyStruct{name: "mystruct1"}
		assert.EqualValues(t, "register error: duplicate type, error type: *engine.MyStruct", registerAny(cter, &b, false).Error())
		assert.Nil(t, registerAny(cter, &b, true))

		var c *MyStruct
		assert.Nil(t, mutableAny(cter, &c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_interface", func(t *testing.T) {
		a := MyInterface(&MyStruct{name: "mystruct"})
		assert.Nil(t, registerAny(cter, &a, false))
		b := MyInterface(&MyStruct{name: "mystruct1"})
		assert.EqualValues(t, "register error: duplicate type, error type: engine.MyInterface", registerAny(cter, &b, false).Error())
		assert.Nil(t, registerAny(cter, &b, true))

		var c MyInterface
		assert.Nil(t, mutableAny(cter, &c))
		assert.Equal(t, b, c)
	})
}

func TestRegisterInsError(t *testing.T) {
	cter := newContainer()
	assert.EqualValues(t, errRegisterInvalidIns, registerAny(cter, nil, false))
	assert.EqualValues(t, errRegisterNotPtr, registerAny(cter, 123, false))
}

func TestRegisterAndMutableTyped(t *testing.T) {
	cter := newContainer()

	t.Run("typed_basic", func(t *testing.T) {
		a := 123
		assert.Nil(t, registerTyped(cter, &a, false))

		var b int
		assert.Nil(t, mutableTyped(cter, &b))
		assert.Equal(t, a, b)
	})

	t.Run("typed_replace", func(t *testing.T) {
		a := "123"
		assert.Nil(t, registerTyped(cter, &a, true))

		b := "456"
		assert.EqualValues(t, "register error: duplicate type, error type: string", registerTyped(cter, &b, false).Error())
		assert.Nil(t, registerTyped(cter, &b, true))

		var c string
		assert.Nil(t, mutableTyped(cter, &c))
		assert.Equal(t, b, c)
	})
}

func TestTypedAndLegacyCompatibility(t *testing.T) {
	cter := newContainer()

	t.Run("register_typed_mutable_legacy", func(t *testing.T) {
		v := int64(11)
		assert.Nil(t, registerTyped(cter, &v, true))

		var got int64
		assert.Nil(t, mutableAny(cter, &got))
		assert.Equal(t, v, got)
	})

	t.Run("register_legacy_mutable_typed", func(t *testing.T) {
		v := float64(3.14)
		assert.Nil(t, registerAny(cter, &v, true))

		var got float64
		assert.Nil(t, mutableTyped(cter, &got))
		assert.Equal(t, v, got)
	})

	t.Run("typed_and_legacy_keep_latest_value", func(t *testing.T) {
		v := 10
		assert.Nil(t, registerTyped(cter, &v, true))

		v = 20

		var gotTyped int
		assert.Nil(t, mutableTyped(cter, &gotTyped))
		assert.Equal(t, 20, gotTyped)

		var gotAny int
		assert.Nil(t, mutableAny(cter, &gotAny))
		assert.Equal(t, 20, gotAny)
	})

	t.Run("legacy_and_typed_keep_latest_value", func(t *testing.T) {
		v := "v1"
		assert.Nil(t, registerAny(cter, &v, true))

		v = "v2"

		var gotTyped string
		assert.Nil(t, mutableTyped(cter, &gotTyped))
		assert.Equal(t, "v2", gotTyped)

		var gotAny string
		assert.Nil(t, mutableAny(cter, &gotAny))
		assert.Equal(t, "v2", gotAny)
	})
}

func TestTypedAPIError(t *testing.T) {
	cter := newContainer()
	assert.EqualValues(t, "register error: pointer is nil", registerTyped[int](cter, nil, false).Error())
	assert.EqualValues(t, "mutable error: pointer is nil", mutableTyped[int](cter, nil).Error())

	var v int
	assert.EqualValues(t, "mutable error: ins not found, error type: int", mutableTyped(cter, &v).Error())
}

func TestMutableInsError(t *testing.T) {
	cter := newContainer()
	assert.EqualValues(t, errMutableInvalidIns, mutableAny(cter, nil))
	assert.EqualValues(t, errMutableNotPtr, mutableAny(cter, 123))

	var a int
	assert.Equal(t, "mutable error: ins not found, error type: int", mutableAny(cter, &a).Error())
}

func TestNewContainer(t *testing.T) {
	cter := newContainer()
	containerPool.Put(cter)
}

func Test_getRoutine(t *testing.T) {
	cter := newContainer()
	testRoutine := cter.getRoutine("test")

	testRoutine1 := cter.getRoutine("test")
	assert.Equal(t, testRoutine, testRoutine1)
}

func TestContainerConcurrentAccess(t *testing.T) {
	cter := newContainer()
	var wg sync.WaitGroup

	// 先注册一些值
	for i := 0; i < 100; i++ {
		val := i
		assert.Nil(t, registerAny(cter, &val, true))
	}

	// 并发获取并验证值在0-99范围内
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var val int
			assert.Nil(t, mutableAny(cter, &val))
			assert.True(t, val >= 0 && val < 100, "value should be in registered range")
		}()
	}

	wg.Wait()
}

func TestContainerReset(t *testing.T) {
	cter := newContainer()
	a := 123
	assert.Nil(t, registerAny(cter, &a, false))
	b := "hello"
	assert.Nil(t, registerTyped(cter, &b, false))

	// 重置容器
	cter.reset()

	// 验证实例已被清空
	cter.mutex.RLock()
	assert.Equal(t, 0, len(cter.typedInstances))
	cter.mutex.RUnlock()
}

func TestSetIns(t *testing.T) {
	cter := newContainer()
	typ := reflect.TypeOf(123)
	val := reflect.ValueOf(123)

	// 测试替换
	assert.Nil(t, cter.setIns(typ, val, false))
	assert.Nil(t, cter.setIns(typ, val, true))

	// 测试不替换
	err := cter.setIns(typ, val, false)
	assert.Equal(t, "register error: duplicate type, error type: int", err.Error())
}

func TestContainerPoolReuse(t *testing.T) {
	// 从池中获取容器
	c1 := newContainer()
	containerPool.Put(c1)

	// 再次获取应该重用同一个实例
	c2 := newContainer()
	assert.Equal(t, c1, c2, "should reuse container from pool")

	// 验证重置后的状态
	c2.mutex.RLock()
	assert.Equal(t, 0, len(c2.typedInstances))
	c2.mutex.RUnlock()
}
