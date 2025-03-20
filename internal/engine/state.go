package engine

import "fmt"

type HasOperatorState interface {
	OperatorState() OperatorStates
}

type OperatorStates interface {
	// Fatal 创建一个致命错误，返回这类错误会直接中断整个执行图
	Fatal(code int, msg string) State

	// FatalCause 创建一个致命错误，同时可以内嵌一个error对象
	FatalCause(code int, msg string) func(error) State

	// Info 创建一个提示状态码，仅在log中打印这个状态码
	Info(code int, msg string) State

	// InfoCause 创建一个提示状态码，仅在log中打印这个状态码；同时可以内嵌一个error对象
	InfoCause(code int, msg string) func(error) State
}

func NewOperatorStates(name string) OperatorStates {
	return &operatorState{
		name:   name,
		states: make(map[int]*status, 10),
	}
}

type operatorState struct {
	name   string
	states map[int]*status
}

func (os *operatorState) Fatal(code int, msg string) State {
	return os.newState(code, msg, true)
}

func (os *operatorState) FatalCause(code int, msg string) func(error) State {
	s := os.newState(code, msg, true)
	return func(err error) State {
		s.dynamicErr = err
		return s
	}
}

func (os *operatorState) Info(code int, msg string) State {
	return os.newState(code, msg, false)
}

func (os *operatorState) InfoCause(code int, msg string) func(error) State {
	s := os.newState(code, msg, false)
	return func(err error) State {
		s.dynamicErr = err
		return s
	}
}

func (os *operatorState) newState(code int, msg string, fatal bool) *status {
	if _, ok := os.states[code]; ok {
		panic(fmt.Errorf("%q status duplicate code: %d", os.name, code))
	}
	st := &status{code: code, msg: msg, fatal: fatal}
	os.states[code] = st
	return st
}

type State interface {
	// Error 实现error接口
	Error() string

	// Code 状态码
	Code() int
}

type status struct {
	code  int
	msg   string
	fatal bool

	dynamicErr error
}

func (s *status) Error() string {
	if s.dynamicErr != nil {
		return fmt.Sprintf("%s{%s}", s.msg, s.dynamicErr)
	}
	return s.msg
}

func (s *status) Code() int {
	return s.code
}
