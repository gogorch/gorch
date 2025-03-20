package engine

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/gorch/gorch/pool"
)

var (
	containerPool = pool.NewObjectPool(func(t *container) {
		t.reset()
	})

	invalidValue = reflect.Value{}
)

// container 容器对象
// 每执行一次，都会创建或者从池里获取到一个新的容器对象用来保存当次执行所有要传递的对象
// 在执行结束后，会将容器对象放回池中，供下次使用
type container struct {
	instances map[reflect.Type]reflect.Value
	routines  map[string]*routine
	mutex     *sync.RWMutex
}

func newContainer() *container {
	ih := containerPool.Get()
	ih.mutex = pool.RWMutexPool.Get()
	return ih
}

func (cter *container) reset() {
	clear(cter.instances)
	clear(cter.routines)
	cter.mutex = pool.RWMutexPool.Get()
}

// PutPoolAble 支持将对象放回池中的对象
type PutPoolAble interface {
	PutPool()
}

func releaseContainer(cter *container) {
	// 如果对象支持PutPool方法，则调用PutPool方法，对象可以由业务自行决定是否放回池中
	// 这里要加锁吗？不用，已经是引擎执行完的状态，不需要再考虑并发问题
	for _, i := range cter.instances {
		if i.IsNil() {
			continue
		}
		if i.CanInterface() {
			if p, ok := i.Interface().(PutPoolAble); ok && p != nil {
				p.PutPool()
			}
		}
	}

	for _, r := range cter.routines {
		routinePool.Put(r)
	}
	pool.RWMutexPool.Put(cter.mutex)
	containerPool.Put(cter)
}

// RegisterIns 注册一个对象到当前调用的对象容器中，对象必须是一个指针类型的对象，否则会报错

// 注册基础类型对象：
//
//	var a int = 123
//	cter.RegisterIns(&a, true)
//
// 注册自定义类型对象：
//
//	var a *MyStruct
//	cter.RegisterIns(&a, true)

// 注册接口：
//
//	var a MyInterface
//	cter.RegisterIns(&a, true)
func (cter *container) RegisterIns(ins any, replace bool) error {
	rvVal := reflect.ValueOf(ins)

	if !rvVal.IsValid() {
		return errRegisterInvalidIns
	}

	rvType := rvVal.Type()

	if rvType.Kind() != reflect.Ptr {
		return errRegisterNotPtr
	}

	rVal := rvVal.Elem()
	rTyp := rVal.Type()

	cter.mutex.Lock()
	if cter.instances == nil {
		cter.instances = make(map[reflect.Type]reflect.Value)
	}
	if _, has := cter.instances[rTyp]; has && !replace {
		cter.mutex.Unlock()
		return fmt.Errorf("register error: duplicate type, error type: %s", rTyp)
	}

	cter.instances[rTyp] = rVal
	cter.mutex.Unlock()

	return nil
}

// MutableIns 获取一个对象值，对象必须是一个指针类型的对象，否则会报错
func (cter *container) MutableIns(val any) error {
	rvVal := reflect.ValueOf(val)

	if !rvVal.IsValid() {
		return errMutableInvalidIns
	}

	rvType := rvVal.Type()

	if rvType.Kind() != reflect.Ptr {
		return errMutableNotPtr
	}

	rVal := rvVal.Elem()
	rTyp := rVal.Type()

	cter.mutex.RLock()
	if cter.instances == nil {
		cter.instances = make(map[reflect.Type]reflect.Value)
	}
	rvIns, has := cter.instances[rTyp]
	cter.mutex.RUnlock()
	if !has {
		return fmt.Errorf("mutable error: ins not found, error type: %s", rTyp)
	}

	rVal.Set(rvIns)
	return nil
}

// getIns 根据类型获取一个对象值
func (cter *container) getIns(rt reflect.Type) (val reflect.Value, has bool) {
	cter.mutex.RLock()
	defer cter.mutex.RUnlock()
	if cter.instances == nil {
		cter.instances = make(map[reflect.Type]reflect.Value)
	}
	rvIns, has := cter.instances[rt]
	if has {
		return rvIns, true
	}
	return invalidValue, false
}

// setIns 设置一个对象值，本函数是在算子执行完成之后，将算子的extract属性设置到容器中
// 所以本函数是一个内部函数，不应该被外部调用
func (cter *container) setIns(rt reflect.Type, val reflect.Value, replace bool) error {
	cter.mutex.Lock()
	defer cter.mutex.Unlock()
	if cter.instances == nil {
		cter.instances = make(map[reflect.Type]reflect.Value)
	}
	if _, has := cter.instances[rt]; has && !replace {
		return fmt.Errorf("register error: duplicate type, error type: %s", rt)
	}

	cter.instances[rt] = val
	return nil
}

func (cter *container) getRoutine(name string) *routine {
	var r *routine
	cter.mutex.RLock()
	if cter.routines != nil {
		r, has := cter.routines[name]
		if has {
			cter.mutex.RUnlock()
			return r
		}
	}
	cter.mutex.RUnlock()

	cter.mutex.Lock()
	if cter.routines == nil {
		cter.routines = make(map[string]*routine)
	}
	r = newRoutine(name)
	cter.routines[name] = r
	cter.mutex.Unlock()

	return r
}
