package engine

import (
	"errors"
)

var (
	errRegisterInvalidIns = errors.New("register error: instance invalid")
	// errRegisterNotPtr 注册到container中的对象必须是一个指针类型的对象，否则报错
	errRegisterNotPtr = errors.New("register error: must be Ptr")

	errMutableInvalidIns = errors.New("register error: instance invalid")
	errMutableNotPtr     = errors.New("mutable error: must be Ptr")
	errNotWrapOperator   = errors.New("not wrap operator")
	errPrepareSwitchCase = errors.New("prepare switch case processor error")

	// errNotSwitchOperator 不是一个switch选择算子
	errNotSwitchOperator = errors.New("not switch operator")

	errEngineExecuteTimeout   = errors.New("engine execute timeout")
	errOperatorExecuteTimeout = errors.New("operator execute timeout")
)
