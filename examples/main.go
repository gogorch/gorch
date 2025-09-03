package main

import (
	"context"
	"fmt"

	"github.com/gogorch/gorch"
	_ "github.com/gogorch/gorch/examples/gen"
)

func main() {
	p, e := gorch.ParseDSLInDir("conf")
	if e != nil {
		panic(e)
	}

	ctx := context.TODO()
	eng, err := gorch.New(p)
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
