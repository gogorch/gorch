package main

import (
	"context"
	"fmt"

	"github.com/gogorch/gorch"
	_ "github.com/gogorch/gorch/gorchc/testoutput/gen"
)

func main() {
	p, e := gorch.ParseDSLInDir("./conf")
	if e != nil {
		panic(e)
	}

	eng, err := gorch.New(p)
	if err != nil {
		panic(err)
	}

	executor := eng.Start("test")
	ctx := context.TODO()

	age := gorch.MustPtr(10)
	name := gorch.MustPtr("test")
	executor.Inject(&age)
	executor.Inject(&name)

	res := executor.Execute(ctx)
	fmt.Println(res)
}
