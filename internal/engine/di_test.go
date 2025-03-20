package engine

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestSubStrut struct {
	Name string
}

func (tss *TestSubStrut) TestInterface() string {
	return tss.Name
}

type TestSubInterface interface {
	TestInterface() string
}

type TypeTest struct {
	String       *string         `inject:"optional" extract:"optional"`
	Int8         *int8           `inject:"optional" extract:"optional"`
	Int32        *int32          `inject:"optional" extract:"optional"`
	Int64        *int64          `inject:"optional" extract:"optional"`
	Int          *int            `inject:"optional" extract:"optional"`
	Float32      *float32        `inject:"optional" extract:"optional"`
	Float64      *float64        `inject:"optional" extract:"optional"`
	Uint8        *uint8          `inject:"optional" extract:"optional"`
	Uint32       *uint32         `inject:"optional" extract:"optional"`
	Uint64       *uint64         `inject:"optional" extract:"optional"`
	Uint         *uint           `inject:"optional" extract:"optional"`
	MapIntString *map[int]string `inject:"optional" extract:"optional"`
	MapStringAny *map[string]any `inject:"optional" extract:"optional"`
	SliceInt     *[]int          `inject:"optional" extract:"optional"`
	SliceAny     *[]any          `inject:"optional" extract:"optional"`
	SliceString  *[]string       `inject:"optional" extract:"optional"`

	Struct    *TestSubStrut    `inject:"optional" extract:"optional"`
	Interface TestSubInterface `inject:"optional" extract:"optional"`
}

func TestInject(t *testing.T) {
	t.Run("test_inject_interface", func(t *testing.T) {
		s := ditool[TypeTest]()
		c := newContainer()
		var tss TestSubInterface = &TestSubStrut{Name: "test"}
		c.RegisterIns(&tss, true)

		tt := &TypeTest{}
		err := s.inject(tt, c)
		assert.Nil(t, err)
		assert.Equal(t, "test", tt.Interface.TestInterface())
	})

	t.Run("test_inject_buildin_type", func(t *testing.T) {
		s := ditool[TypeTest]()
		c := newContainer()

		var vstring string = "test"
		var vint8 int8 = 1
		var vint32 int32 = 2
		var vint64 int64 = 3
		var vint int = 4
		var vfloat32 float32 = 5
		var vfloat64 float64 = 6
		var vuint8 uint8 = 7
		var vuint32 uint32 = 8
		var vuint64 uint64 = 9
		var vuint uint = 10
		var vmapIntString map[int]string = map[int]string{1: "1"}
		var vmapStringAny map[string]any = map[string]any{"1": 1}
		var vsliceInt []int = []int{1, 2, 3}
		var vsliceAny []any = []any{1, 2, 3}
		var vsliceString []string = []string{"1", "2", "3"}
		assert.Nil(t, c.RegisterIns(toPtr(&vstring), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vint8), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vint32), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vint64), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vint), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vfloat32), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vfloat64), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vuint8), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vuint32), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vuint64), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vuint), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vmapIntString), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vmapStringAny), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vsliceInt), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vsliceAny), true))
		assert.Nil(t, c.RegisterIns(toPtr(&vsliceString), true))
		tt := &TypeTest{}
		err := s.inject(tt, c)
		assert.Nil(t, err)
		assert.Equal(t, "test", *tt.String)
		assert.Equal(t, int8(1), *tt.Int8)
		assert.Equal(t, int32(2), *tt.Int32)
		assert.Equal(t, int64(3), *tt.Int64)
		assert.Equal(t, int(4), *tt.Int)
		assert.Equal(t, float32(5), *tt.Float32)
		assert.Equal(t, float64(6), *tt.Float64)
		assert.Equal(t, uint8(7), *tt.Uint8)
		assert.Equal(t, uint32(8), *tt.Uint32)
		assert.Equal(t, uint64(9), *tt.Uint64)
		assert.Equal(t, uint(10), *tt.Uint)
		assert.Equal(t, map[int]string{1: "1"}, *tt.MapIntString)
		assert.Equal(t, map[string]any{"1": 1}, *tt.MapStringAny)
		assert.Equal(t, []int{1, 2, 3}, *tt.SliceInt)
		assert.Equal(t, []any{1, 2, 3}, *tt.SliceAny)
		assert.Equal(t, []string{"1", "2", "3"}, *tt.SliceString)
	})

	t.Run("test_inject_struct", func(t *testing.T) {
		s := ditool[TypeTest]()
		c := newContainer()
		ss := &TestSubStrut{Name: "test"}
		assert.Nil(t, c.RegisterIns(&ss, true))
		tt := &TypeTest{}
		err := s.inject(tt, c)
		assert.Nil(t, err)
		assert.Equal(t, "test", tt.Struct.Name)
	})
}

func toPtr[T any](val T) *T {
	return &val
}

func TestExtract(t *testing.T) {
	t.Run("test_extract_interface", func(t *testing.T) {
		s := ditool[TypeTest]()
		c := newContainer()
		var tss TestSubInterface = &TestSubStrut{Name: "test"}
		tt := &TypeTest{Interface: tss}
		err := s.extract(tt, c)
		assert.Nil(t, err)

		// 通过getIns
		rv, has := c.getIns(reflect.TypeOf(&tss).Elem())
		assert.True(t, has)
		assert.Equal(t, "test", rv.Interface().(TestSubInterface).TestInterface())

		// 通过MutableIns
		var tss1 TestSubInterface
		assert.Nil(t, c.MutableIns(&tss1))
		assert.Equal(t, "test", tss1.TestInterface())

		// 通过inject tag
		tt1 := &TypeTest{}
		err = s.inject(tt1, c)
		assert.Nil(t, err)
		assert.Equal(t, "test", tt1.Interface.TestInterface())
	})

	t.Run("test_extract_buildin_type", func(t *testing.T) {
		s := ditool[TypeTest]()
		c := newContainer()
		var vstring string = "test"
		var vint8 int8 = 1
		var vint32 int32 = 2
		var vint64 int64 = 3
		var vint int = 4
		var vfloat32 float32 = 5
		var vfloat64 float64 = 6
		var vuint8 uint8 = 7
		var vuint32 uint32 = 8
		var vuint64 uint64 = 9
		var vuint uint = 10
		var vmapIntString map[int]string = map[int]string{1: "1"}
		var vmapStringAny map[string]any = map[string]any{"1": 1}
		var vsliceInt []int = []int{1, 2, 3}
		var vsliceAny []any = []any{1, 2, 3}
		var vsliceString []string = []string{"1", "2", "3"}

		tt := &TypeTest{
			String:       &vstring,
			Int8:         &vint8,
			Int32:        &vint32,
			Int64:        &vint64,
			Int:          &vint,
			Float32:      &vfloat32,
			Float64:      &vfloat64,
			Uint8:        &vuint8,
			Uint32:       &vuint32,
			Uint64:       &vuint64,
			Uint:         &vuint,
			MapIntString: &vmapIntString,
			MapStringAny: &vmapStringAny,
			SliceInt:     &vsliceInt,
			SliceAny:     &vsliceAny,
			SliceString:  &vsliceString,
		}
		err := s.extract(tt, c)
		assert.Nil(t, err)

		// 通过getIns
		checkGetIns(t, c, vstring, "test")
		checkGetIns(t, c, vint8, int8(1))
		checkGetIns(t, c, vint32, int32(2))
		checkGetIns(t, c, vint64, int64(3))
		checkGetIns(t, c, vint, int(4))
		checkGetIns(t, c, vfloat32, float32(5))
		checkGetIns(t, c, vfloat64, float64(6))
		checkGetIns(t, c, vuint8, uint8(7))
		checkGetIns(t, c, vuint32, uint32(8))
		checkGetIns(t, c, vuint64, uint64(9))
		checkGetIns(t, c, vuint, uint(10))
		checkGetIns(t, c, vmapIntString, map[int]string{1: "1"})
		checkGetIns(t, c, vmapStringAny, map[string]any{"1": 1})
		checkGetIns(t, c, vsliceInt, []int{1, 2, 3})
		checkGetIns(t, c, vsliceAny, []any{1, 2, 3})
		checkGetIns(t, c, vsliceString, []string{"1", "2", "3"})

		// 通过MutableIns
		checkMutableIns(t, c, vstring, "test")
		checkMutableIns(t, c, vint8, int8(1))
		checkMutableIns(t, c, vint32, int32(2))
		checkMutableIns(t, c, vint64, int64(3))
		checkMutableIns(t, c, vint, int(4))
		checkMutableIns(t, c, vfloat32, float32(5))
		checkMutableIns(t, c, vfloat64, float64(6))
		checkMutableIns(t, c, vuint8, uint8(7))
		checkMutableIns(t, c, vuint32, uint32(8))
		checkMutableIns(t, c, vuint64, uint64(9))
		checkMutableIns(t, c, vuint, uint(10))
		checkMutableIns(t, c, vmapIntString, map[int]string{1: "1"})
		checkMutableIns(t, c, vmapStringAny, map[string]any{"1": 1})
		checkMutableIns(t, c, vsliceInt, []int{1, 2, 3})
		checkMutableIns(t, c, vsliceAny, []any{1, 2, 3})
		checkMutableIns(t, c, vsliceString, []string{"1", "2", "3"})

		// 通过inject tag
		tt1 := &TypeTest{}
		err = s.inject(tt1, c)
		assert.Nil(t, err)
		assert.Equal(t, "test", *tt1.String)
		assert.Equal(t, int8(1), *tt1.Int8)
		assert.Equal(t, int32(2), *tt1.Int32)
		assert.Equal(t, int64(3), *tt1.Int64)
		assert.Equal(t, int(4), *tt1.Int)
		assert.Equal(t, float32(5), *tt1.Float32)
		assert.Equal(t, float64(6), *tt1.Float64)
		assert.Equal(t, uint8(7), *tt1.Uint8)
		assert.Equal(t, uint32(8), *tt1.Uint32)
		assert.Equal(t, uint64(9), *tt1.Uint64)
		assert.Equal(t, uint(10), *tt1.Uint)
		assert.Equal(t, map[int]string{1: "1"}, *tt1.MapIntString)
		assert.Equal(t, map[string]any{"1": 1}, *tt1.MapStringAny)
		assert.Equal(t, []int{1, 2, 3}, *tt1.SliceInt)
		assert.Equal(t, []any{1, 2, 3}, *tt1.SliceAny)
		assert.Equal(t, []string{"1", "2", "3"}, *tt1.SliceString)
	})

	t.Run("test_extract_struct", func(t *testing.T) {
		s := ditool[TypeTest]()
		c := newContainer()

		ss := &TestSubStrut{Name: "test"}
		tt := &TypeTest{Struct: ss}
		err := s.extract(tt, c)
		assert.Nil(t, err)

		// 通过getIns
		rv, has := c.getIns(reflect.TypeOf(&ss).Elem())
		assert.True(t, has)
		assert.Equal(t, "test", rv.Interface().(*TestSubStrut).Name)

		// 通过MutableIns
		var tss1 *TestSubStrut
		assert.Nil(t, c.MutableIns(&tss1))
		assert.Equal(t, "test", tss1.Name)

		// 通过inject tag
		tt1 := &TypeTest{}
		err = s.inject(tt1, c)
		assert.Nil(t, err)
		assert.Equal(t, "test", tt1.Struct.Name)
	})
}
func checkGetIns[T any](t *testing.T, c *container, val T, except any) {
	rv, has := c.getIns(reflect.TypeOf(toPtr(&val)).Elem())
	assert.True(t, has)
	assert.Equal(t, except, *rv.Interface().(*T))
}

func checkMutableIns[T any](t *testing.T, c *container, _ T, except any) {
	var ts *T
	assert.Nil(t, c.MutableIns(&ts))
	assert.Equal(t, except, *ts)
}
