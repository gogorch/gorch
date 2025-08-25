package opb

import (
	"fmt"

	"github.com/gogorch/gorch/internal/engine"
)

type OperatorB0 struct {
}

func (o *OperatorB0) Execute(ctx *engine.Context) error {
	return nil
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
