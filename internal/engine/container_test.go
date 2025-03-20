package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
				assert.Nil(t, cter.RegisterIns(ins, false))

				val := cs.val()

				assert.Nil(t, cter.MutableIns(val))
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
		assert.Nil(t, cter.RegisterIns(&a, false))
		b := 456
		assert.EqualValues(t, "register error: duplicate type, error type: int", cter.RegisterIns(&b, false).Error())
		assert.Nil(t, cter.RegisterIns(&b, true))

		var c int
		assert.Nil(t, cter.MutableIns(&c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_string", func(t *testing.T) {
		a := "123"
		assert.Nil(t, cter.RegisterIns(&a, false))
		b := "456"
		assert.EqualValues(t, "register error: duplicate type, error type: string", cter.RegisterIns(&b, false).Error())
		assert.Nil(t, cter.RegisterIns(&b, true))

		var c string
		assert.Nil(t, cter.MutableIns(&c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_float32", func(t *testing.T) {
		a := float32(123)
		assert.Nil(t, cter.RegisterIns(&a, false))
		b := float32(456)
		assert.EqualValues(t, "register error: duplicate type, error type: float32", cter.RegisterIns(&b, false).Error())
		assert.Nil(t, cter.RegisterIns(&b, true))

		var c float32
		assert.Nil(t, cter.MutableIns(&c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_struct", func(t *testing.T) {
		a := MyStruct{name: "mystruct"}
		assert.Nil(t, cter.RegisterIns(&a, false))
		b := MyStruct{name: "mystruct1"}
		assert.EqualValues(t, "register error: duplicate type, error type: engine.MyStruct", cter.RegisterIns(&b, false).Error())
		assert.Nil(t, cter.RegisterIns(&b, true))

		var c MyStruct
		assert.Nil(t, cter.MutableIns(&c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_struct_ptr", func(t *testing.T) {
		a := &MyStruct{name: "mystruct"}
		assert.Nil(t, cter.RegisterIns(&a, false))
		b := &MyStruct{name: "mystruct1"}
		assert.EqualValues(t, "register error: duplicate type, error type: *engine.MyStruct", cter.RegisterIns(&b, false).Error())
		assert.Nil(t, cter.RegisterIns(&b, true))

		var c *MyStruct
		assert.Nil(t, cter.MutableIns(&c))
		assert.Equal(t, b, c)
	})

	t.Run("replace_interface", func(t *testing.T) {
		a := MyInterface(&MyStruct{name: "mystruct"})
		assert.Nil(t, cter.RegisterIns(&a, false))
		b := MyInterface(&MyStruct{name: "mystruct1"})
		assert.EqualValues(t, "register error: duplicate type, error type: engine.MyInterface", cter.RegisterIns(&b, false).Error())
		assert.Nil(t, cter.RegisterIns(&b, true))

		var c MyInterface
		assert.Nil(t, cter.MutableIns(&c))
		assert.Equal(t, b, c)
	})
}

func TestRegisterInsError(t *testing.T) {
	cter := newContainer()
	assert.EqualValues(t, errRegisterInvalidIns, cter.RegisterIns(nil, false))
	assert.EqualValues(t, errRegisterNotPtr, cter.RegisterIns(123, false))
}

func TestMutableInsError(t *testing.T) {
	cter := newContainer()
	assert.EqualValues(t, errMutableInvalidIns, cter.MutableIns(nil))
	assert.EqualValues(t, errMutableNotPtr, cter.MutableIns(123))

	var a int
	assert.Equal(t, "mutable error: ins not found, error type: int", cter.MutableIns(&a).Error())
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
