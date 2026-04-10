package engine

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/stretchr/testify/assert"
)

func registerTestOperator(t *testing.T) {
	assert.Nil(t, RegisterOperator[SleepOp]("SleepOp", 1))
	assert.Nil(t, RegisterOperator[PanicOp]("PanicOp", 2))
	assert.Nil(t, RegisterOperator[ChangeValueOP]("ChangeValueOP", 3))
	assert.Nil(t, RegisterOperator[NothingOp]("NothingOp", 4))
	assert.Nil(t, RegisterOperator[ContextInterrupt]("ContextInterrupt", 5))
	assert.Nil(t, RegisterOperator[TestOperatorState]("TestOperatorState", 6))
	assert.Nil(t, RegisterOperator[GoOperator]("GoOperator", 7))
	assert.Nil(t, RegisterOperator[SwitchOp]("SwitchOp", 8))
	assert.Nil(t, RegisterOperator[WrapAndChangeValOp]("WrapAndChangeValOp", 9))
	assert.Nil(t, RegisterOperator[SkipOp]("SkipOp", 10))
	assert.Nil(t, RegisterOperator[AddValueOp]("AddValueOp", 11))
	assert.Nil(t, RegisterOperator[UntilValueReachedOp]("UntilValueReachedOp", 12))
	assert.Nil(t, RegisterOperator[RetryFailThenPassOp]("RetryFailThenPassOp", 13))
}

type TestValus struct {
	name string
	g    string
	ctx  context.Context

	assertFunc    func(error, *BeChangeValue)
	engineTimeout time.Duration
}

func MyTestRun(t *testing.T, tv TestValus) {
	registerTestOperator(t)
	defer clean()
	p, e := iparser.Parse(tv.g)
	assert.Nil(t, e)

	ctx := context.TODO()
	if tv.ctx != nil {
		ctx = tv.ctx
	}
	eng, err := New(p)
	assert.Nil(t, err)

	val := &BeChangeValue{Val: 0}
	executor := eng.Start(tv.name)

	if tv.engineTimeout > 0 {
		executor.SetTimeout(tv.engineTimeout)
	}

	assert.Nil(t, executor.Inject(&val, &t))
	res := executor.Execute(ctx)
	tv.assertFunc(res, val)
}

// NothingOp 啥也不干的算子
type NothingOp struct{}

func (s *NothingOp) Execute(ctx *Context) error {
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

// PanicOp 肯定会触发panic的算子
type PanicOp struct{}

func (p *PanicOp) Execute(ctx *Context) error {
	panic("panic from PanicOp")
}

type BeChangeValue struct {
	mu  sync.Mutex
	Val int64
}

func (b *BeChangeValue) Get() int64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Val
}

func (b *BeChangeValue) Set(v int64) {
	b.mu.Lock()
	b.Val = v
	b.mu.Unlock()
}

func (b *BeChangeValue) Add(delta int64) int64 {
	b.mu.Lock()
	b.Val += delta
	v := b.Val
	b.mu.Unlock()
	return v
}

// ChangeValueOP 改变值的算子
type ChangeValueOP struct {
	TestT         *testing.T     `inject:""`
	BeChangeValue *BeChangeValue `inject:""`
}

func (c *ChangeValueOP) Execute(ctx *Context) error {
	if !ctx.Has("val") {
		return errors.New("ChangeValueOP operator must get val argument")
	}

	if ctx.Has("if") {
		if ctx.Arg("if").Int64() != c.BeChangeValue.Get() {
			return errors.New("ChangeValueOP operator if condition not equal")
		}
	}

	c.BeChangeValue.Set(ctx.Arg("val").Int64())
	return nil
}

type ContextInterrupt struct {
}

func (s *ContextInterrupt) Execute(ctx *Context) error {
	to := ctx.Arg("sleep").Duration()
	if to <= 0 {
		return errors.New("ContextInterrupt operator must get sleep argument")
	}
	i := ctx.Interrupt()
	go func() {
		time.Sleep(to)
		i.Exit(errors.New("exit by ContextInterrupt"))
	}()
	return nil
}

type TestOperatorState struct {
}

var (
	testOperatorState = NewOperatorStates("TestOperatorState")
	fatal1            = testOperatorState.Fatal(1, "fatal1")
	fatal2Cause       = testOperatorState.FatalCause(2, "fatal2")
	info3             = testOperatorState.Info(3, "info3")
	info4Cause        = testOperatorState.InfoCause(4, "info4")
)

func (s *TestOperatorState) Execute(ctx *Context) error {
	if !ctx.Has("status") {
		return errors.New("TestOperatorState operator must get status argument")
	}
	status := ctx.Arg("status").Int64()

	switch status {
	case 1:
		return fatal1
	case 2:
		return fatal2Cause(errors.New("fatal2WithErr"))
	case 3:
		return info3
	case 4:
		return info4Cause(errors.New("info4WithErr"))
	}
	return nil
}

type GoOperator struct {
	TestT         *testing.T     `inject:""`
	BeChangeValue *BeChangeValue `inject:""`
}

func (g *GoOperator) Execute(ctx *Context) error {
	if !ctx.Has("sleep") {
		return errors.New("GoOperator operator must get sleep argument")
	}

	if !ctx.Has("val") {
		return errors.New("GoOperator operator must get val argument")
	}

	time.Sleep(ctx.Arg("sleep").Duration())

	g.BeChangeValue.Set(ctx.Arg("val").Int64())
	return nil
}

type SwitchOp struct{}

func (s *SwitchOp) Execute(ctx *Context) error {
	if !ctx.Has("case") {
		return errors.New("SwitchOp operator must get case argument")
	}

	c := ctx.Arg("case").StrList()
	return ctx.Switch(c...)
}

type WrapAndChangeValOp struct {
	BeChangeValue *BeChangeValue `inject:""`
}

func (s *WrapAndChangeValOp) Execute(ctx *Context) error {
	if !ctx.Has("val") {
		return errors.New("WrapAndChangeValOp operator must get val argument")
	}

	s.BeChangeValue.Set(ctx.Arg("val").Int64())

	if ctx.Has("noNext") && ctx.Arg("noNext").Bool() {
		return nil
	}
	return ctx.Next()
}

type SkipOp struct{}

func (s *SkipOp) Execute(ctx *Context) error {
	if !ctx.Has("skip") {
		return errors.New("SkipOp operator must get skip argument")
	}

	if ctx.Arg("skip").Bool() {
		return ctx.SkipSerial()
	}

	return nil
}

type AddValueOp struct {
	BeChangeValue *BeChangeValue `inject:""`
}

func (a *AddValueOp) Execute(ctx *Context) error {
	if !ctx.Has("delta") {
		return errors.New("AddValueOp operator must get delta argument")
	}
	a.BeChangeValue.Add(ctx.Arg("delta").Int64())
	return nil
}

type UntilValueReachedOp struct {
	BeChangeValue *BeChangeValue `inject:""`
}

func (u *UntilValueReachedOp) Execute(ctx *Context) error {
	if !ctx.Has("target") {
		return errors.New("UntilValueReachedOp operator must get target argument")
	}
	return ctx.SetLoopUntil(u.BeChangeValue.Get() >= ctx.Arg("target").Int64())
}

type RetryFailThenPassOp struct {
	BeChangeValue *BeChangeValue `inject:""`
}

func (r *RetryFailThenPassOp) Execute(ctx *Context) error {
	if !ctx.Has("failTimes") {
		return errors.New("RetryFailThenPassOp operator must get failTimes argument")
	}

	try := r.BeChangeValue.Add(1)
	if try <= ctx.Arg("failTimes").Int64() {
		return errors.New("retry failed")
	}

	return nil
}
