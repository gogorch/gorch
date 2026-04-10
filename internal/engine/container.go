package engine

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/gogorch/gorch/mlog"
	"github.com/gogorch/gorch/pool"
)

var (
	containerPool = pool.NewObjectPool(func(t *container) {
		t.reset()
	})
)

// container 容器对象
// 每执行一次，都会创建或者从池里获取到一个新的容器对象用来保存当次执行所有要传递的对象
// 在执行结束后，会将容器对象放回池中，供下次使用
type container struct {
	typedInstances map[reflect.Type]any
	routines       map[string]*routine
	mutex          sync.RWMutex
}

func newContainer() *container {
	return containerPool.Get()
}

func (cter *container) reset() {
	// 【修复空指针】防止对象池回调时访问未初始化的 map
	// 原来的问题：对象池在重置 container 时，map 字段可能为 nil，
	// 直接调用 clear() 会导致 panic
	if cter.typedInstances == nil {
		cter.typedInstances = make(map[reflect.Type]any)
	}
	if cter.routines == nil {
		cter.routines = make(map[string]*routine)
	}

	// 【修复竞态条件】使用互斥锁保护 map 操作，防止并发访问冲突
	// 原来的问题：主线程在 releaseContainer() 中清理 map，
	// 而 goroutine 同时在 container.getIns() 中读取 map，导致数据竞争
	cter.mutex.Lock()
	defer cter.mutex.Unlock()

	clear(cter.typedInstances)
	clear(cter.routines)
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
	releaseInstances(cter.typedInstances)

	// 释放容器中的协程对象
	releaseRoutines(cter.routines)

	// 重置容器内容，避免对象池复用脏数据。
	cter.reset()

	// 将容器放回对象池
	containerPool.Put(cter)
}

// releaseInstances 释放容器中的实例对象，若对象实现了 Releasable 接口，则调用其 Release 方法。
func releaseInstances(instances map[reflect.Type]any) {
	for t, ins := range instances {
		rv := reflect.ValueOf(ins)
		if !rv.IsValid() || rv.Kind() != reflect.Ptr || rv.IsNil() {
			continue
		}
		val := rv.Elem()
		if !val.IsValid() || (val.Kind() == reflect.Ptr && val.IsNil()) {
			continue
		}
		if val.CanInterface() {
			if p, ok := val.Interface().(Releasable); ok && p != nil {
				func() {
					defer func() {
						if err := recover(); err != nil {
							fields := RecoverPanic(err)
							fields = append(fields,
								mlog.Any("err", err),
								mlog.Reflect("ins", val),
								mlog.ReflectType("type", t))

							globalLogger.Error("panicWhenReleasingInstance", fields...)
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

// registerAny 注册一个对象到当前调用的对象容器中，对象必须是一个指针类型的对象，否则会报错。
// 该函数仅用于动态类型场景（如 Executor.Inject）。
func registerAny(cter *container, ins any, replace bool) error {
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
	if cter.typedInstances == nil {
		cter.typedInstances = make(map[reflect.Type]any)
	}
	if _, has := cter.typedInstances[rTyp]; has && !replace {
		return fmt.Errorf("register error: duplicate type, error type: %s", rTyp)
	}
	// typed fast-path 缓存原始指针，保证读取到最新值。
	cter.typedInstances[rTyp] = ins
	return nil
}

// registerTyped 注册一个类型化对象到当前调用的对象容器中。
// 与 registerAny 对比：
// 1. 避免对调用参数进行反射解析
// 2. 为 mutableTyped 提供更快的读取路径
func registerTyped[T any](cter *container, ins *T, replace bool) error {
	if ins == nil {
		return fmt.Errorf("register error: pointer is nil")
	}

	rTyp := reflect.TypeFor[T]()

	cter.mutex.Lock()
	defer cter.mutex.Unlock()

	if cter.typedInstances == nil {
		cter.typedInstances = make(map[reflect.Type]any)
	}
	if _, has := cter.typedInstances[rTyp]; has && !replace {
		return fmt.Errorf("register error: duplicate type, error type: %s", rTyp)
	}

	// 直接缓存原始指针，保持读取语义简单一致。
	cter.typedInstances[rTyp] = ins

	return nil
}

// mutableAny 获取一个对象值，对象必须是一个指针类型的对象，否则会报错。
func mutableAny(cter *container, val any) error {
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
	if cter.typedInstances == nil {
		return fmt.Errorf("mutable error: ins not found, error type: %s", rTyp)
	}
	raw, has := cter.typedInstances[rTyp]
	if !has {
		return fmt.Errorf("mutable error: ins not found, error type: %s", rTyp)
	}
	rvIns := reflect.ValueOf(raw)
	if !rvIns.IsValid() || rvIns.Kind() != reflect.Ptr || rvIns.IsNil() {
		return fmt.Errorf("mutable error: ins not found, error type: %s", rTyp)
	}
	v := rvIns.Elem()
	if !v.IsValid() || !v.Type().AssignableTo(rTyp) {
		return fmt.Errorf("mutable error: ins not found, error type: %s", rTyp)
	}
	rVal.Set(v)
	return nil
}

// mutableTyped 获取一个类型化对象值，优先走 typed fast-path。
func mutableTyped[T any](cter *container, val *T) error {
	if val == nil {
		return fmt.Errorf("mutable error: pointer is nil")
	}

	rTyp := reflect.TypeFor[T]()

	cter.mutex.RLock()
	defer cter.mutex.RUnlock()

	if cter.typedInstances != nil {
		if raw, has := cter.typedInstances[rTyp]; has {
			if p, ok := raw.(*T); ok && p != nil {
				*val = *p
				return nil
			}
		}
	}

	return fmt.Errorf("mutable error: ins not found, error type: %s", rTyp)
}

// getIns 根据类型获取一个对象值
func (cter *container) getIns(rt reflect.Type) (val reflect.Value, has bool) {
	invalidValue := reflect.Value{}
	cter.mutex.RLock()
	defer cter.mutex.RUnlock()
	if cter.typedInstances == nil {
		return invalidValue, false
	}
	raw, ok := cter.typedInstances[rt]
	if !ok {
		return invalidValue, false
	}
	rvIns := reflect.ValueOf(raw)
	if !rvIns.IsValid() || rvIns.Kind() != reflect.Ptr || rvIns.IsNil() {
		return invalidValue, false
	}
	v := rvIns.Elem()
	if !v.IsValid() || !v.Type().AssignableTo(rt) {
		return invalidValue, false
	}
	return v, true
}

// setIns 设置一个对象值，本函数是在算子执行完成之后，将算子的extract属性设置到容器中
// 所以本函数是一个内部函数，不应该被外部调用
func (cter *container) setIns(rt reflect.Type, val reflect.Value, replace bool) error {
	cter.mutex.Lock()
	defer cter.mutex.Unlock()
	if cter.typedInstances == nil {
		cter.typedInstances = make(map[reflect.Type]any)
	}
	if _, has := cter.typedInstances[rt]; has && !replace {
		return fmt.Errorf("register error: duplicate type, error type: %s", rt)
	}

	if !val.IsValid() {
		delete(cter.typedInstances, rt)
		return nil
	}
	holder := reflect.New(rt)
	holder.Elem().Set(val)
	cter.typedInstances[rt] = holder.Interface()
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
