package engine

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/gogorch/gorch/pool"
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

// Releasable 支持将对象放回池中的对象
type Releasable interface {
	Release()
}

// releaseContainer 负责将容器及其包含的对象放回相应的对象池。
// 在放回对象前，会检查对象是否实现了 Releasable 接口，若实现则调用其 Release 方法。
func releaseContainer(cter *container) {
	if cter == nil {
		return
	}

	// 释放容器中的实例对象
	releaseInstances(cter.instances)

	// 释放容器中的协程对象
	releaseRoutines(cter.routines)

	// 释放互斥锁
	pool.RWMutexPool.Put(cter.mutex)

	// 重置容器
	cter.reset()

	// 将容器放回对象池
	containerPool.Put(cter)
}

// releaseInstances 释放容器中的实例对象，若对象实现了 Releasable 接口，则调用其 Release 方法。
func releaseInstances(instances map[reflect.Type]reflect.Value) {
	for _, ins := range instances {
		// 检查值是否有效且不是零值
		if !ins.IsValid() || (ins.Kind() == reflect.Ptr && ins.IsNil()) {
			continue
		}
		if ins.CanInterface() {
			if p, ok := ins.Interface().(Releasable); ok && p != nil {
				func() {
					defer func() {
						if err := recover(); err != nil {
							log.Printf("[releaseInstances] panic when releasing instance: %v\n", err)
						}
					}()
					p.Release()
				}()
			}
		}
	}
}

// releaseRoutines 释放容器中的协程对象，将其放回协程池。
func releaseRoutines(routines map[string]*routine) {
	for _, r := range routines {
		if r != nil {
			routinePool.Put(r)
		}
	}
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
	if rvVal.IsNil() {
		return fmt.Errorf("register error: pointer is nil")
	}
	rVal := rvVal.Elem()
	rTyp := rVal.Type()

	cter.mutex.Lock()
	defer cter.mutex.Unlock()
	if cter.instances == nil {
		cter.instances = make(map[reflect.Type]reflect.Value)
	}
	if _, has := cter.instances[rTyp]; has && !replace {
		return fmt.Errorf("register error: duplicate type, error type: %s", rTyp)
	}
	cter.instances[rTyp] = rVal
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
	if rvVal.IsNil() {
		return fmt.Errorf("mutable error: pointer is nil")
	}
	rVal := rvVal.Elem()
	rTyp := rVal.Type()

	cter.mutex.RLock()
	defer cter.mutex.RUnlock()
	if cter.instances == nil {
		cter.instances = make(map[reflect.Type]reflect.Value)
	}
	rvIns, has := cter.instances[rTyp]
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
		var has bool
		r, has = cter.routines[name]
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
