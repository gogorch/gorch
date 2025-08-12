package gorch

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testGraph = `
START("test"){
	ON_FINISH() { a(onfinish=true) }

	UNFOLD("test")
	-> SleepOp(sleep=10ms)
	-> b(changeTo=1)
	-> b(changeTo=12, WAIT("aa", timeout=10ms))
}

FRAGMENT("test"){
	@a(fatal=true) -> GO(a(changeTo=11), "aa")
}

REGISTER("github.com/gogorch"){
	OPERATOR("gorch/ops", "a", 1)
	OPERATOR("gorch/ops", "b", 2)
}
`

type aOp struct {
	A *int64 `inject:""`
}

var (
	astatus      = NewOperatorStates("aOp")
	aOpFatal     = astatus.Fatal(1, "this is fatal error log")
	aOpNotice    = astatus.Info(2, "this is notice log")
	aOpOnFinish  = astatus.Info(3, "this is finish log")
	aOpSleep     = astatus.Info(4, "this is sleep log")
	aOpChangeVal = astatus.Info(5, "this is change val log")
)

func (a *aOp) Execute(ctx *Context) error {
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
		*a.A = ctx.Arg("changeTo").Int64()
		return aOpChangeVal
	}
	return aOpNotice
}

func (a *aOp) Prepare() error {
	fmt.Println("aOp Prepare", a)
	return nil
	// return errors.New("aOp Prepare")
}

type bOp struct {
	A *int64 `inject:""`
}

var (
	bstatus      = NewOperatorStates("bOp")
	bOpNotice    = bstatus.Info(1, "this is notice log")
	bOpChangeVal = bstatus.Info(2, "this is change val log")
)

func (b *bOp) Execute(ctx *Context) error {
	if ctx.Has("changeTo") {
		*b.A = ctx.Arg("changeTo").Int64()
		return bOpChangeVal
	}
	return bOpNotice
}

func (b *bOp) Prepare() error {
	fmt.Println("bOp Prepare", b)
	return nil
}

// SleepOp 睡眠算子，在测试时睡眠指定时间
type SleepOp struct{}

func (s *SleepOp) Execute(ctx *Context) error {
	to := ctx.Arg("sleep").Duration()
	if to <= 0 {
		return errors.New("sleep operator must get sleep argument")
	}

	time.Sleep(to)
	return nil
}

func TestGraph(t *testing.T) {
	p, e := ParseDSL(testGraph)
	assert.Nil(t, e)

	RegOp[aOp]("a", 1)
	RegOp[bOp]("b", 2)
	RegOp[SleepOp]("SleepOp", 3)

	ctx := context.TODO()
	eng, err := New(p)
	// fmt.Println(err, "-1-1-1-1-")
	if err != nil {
		panic(err)
	}
	fmt.Println("start---")

	var a int64 = 123
	s := &a
	executor := eng.Start("test")
	// executor.SetTimeout(time.Millisecond)
	executor.Inject(&s)
	executor.SetLogOperatorName()
	err = executor.Execute(ctx)
	fmt.Println("a=", a)
	if err != nil {
		panic(err)
	}
}
