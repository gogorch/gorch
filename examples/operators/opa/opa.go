package opa

import (
	"fmt"
	"time"

	"github.com/gogorch/gorch"
	"github.com/gogorch/gorch/internal/engine"
)

var (
	astatus      = gorch.NewOperatorStates("aOp")
	aOpFatal     = astatus.Fatal(1, "this is fatal error log")
	aOpNotice    = astatus.Info(2, "this is notice log")
	aOpOnFinish  = astatus.Info(3, "this is finish log")
	aOpSleep     = astatus.Info(4, "this is sleep log")
	aOpChangeVal = astatus.Info(5, "this is change val log")
)

type OperatorA0 struct {
	A *int64 `inject:""`
}

func (o *OperatorA0) Execute(ctx *engine.Context) error {
	if ctx.Arg("onfinish").Bool() {
		fmt.Println("onfinish")
		return aOpOnFinish
	}
	if ctx.Arg("fatal").Bool() {
		return aOpFatal
	}

	if ctx.Arg("panic").Bool() {
		panic("panic from aOp")
	}

	if s := ctx.Arg("sleep").Duration(); s > 0 {
		time.Sleep(s)
		return aOpSleep
	}
	if ctx.Has("changeTo") {
		*o.A = ctx.Arg("changeTo").Int64()
		return aOpChangeVal
	}
	return aOpNotice
}

type OperatorA1 struct {
	Name *string `inject:"optional"`
	Age  *int    `inject:"optional" extract:"replace"`
	Ax   *int32
}

func (o *OperatorA1) Execute(ctx *engine.Context) error {
	return nil
}
