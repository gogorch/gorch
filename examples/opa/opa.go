package opa

import "github.com/gogorch/gorch/internal/engine"

type OperatorA0 struct {
	Name *string `inject:""`
	Age  *int    `inject:"optional" extract:"replace"`
	Ax   *int32
}

func (o *OperatorA0) Execute(ctx *engine.Context) error {
	return nil
}

type OperatorA1 struct {
	Name *string `inject:"optional"`
	Age  *int    `inject:"optional" extract:"replace"`
	Ax   *int32
}

func (o *OperatorA1) Execute(ctx *engine.Context) error {
	return nil
}
