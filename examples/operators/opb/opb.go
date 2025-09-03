package opb

import (
	"fmt"

	"github.com/gogorch/gorch"
	"github.com/gogorch/gorch/internal/engine"
)

var (
	bstatus      = gorch.NewOperatorStates("bOp")
	bOpNotice    = bstatus.Info(1, "this is notice log")
	bOpChangeVal = bstatus.Info(2, "this is change val log")
)

type OperatorB0 struct {
	A *int64 `inject:""`
}

func (o *OperatorB0) Execute(ctx *engine.Context) error {
	if ctx.Has("changeTo") {
		*o.A = ctx.Arg("changeTo").Int64()
		return bOpChangeVal
	}
	return bOpNotice
}

type OperatorB1 struct {
	Age *int `extract:"replace"`
}

func (o *OperatorB1) Execute(ctx *engine.Context) error {
	return nil
}

func (o *OperatorB1) Prepare() error {
	fmt.Println("Prepare from OperatorB1")
	return nil
}
