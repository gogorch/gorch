package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gogorch/gorch/internal/lang/iparser"
)

// g 生成所有代码到指定目录
var g = flag.String("g", "", "generate all code to directory")

// i 生成import算子代码
var i = flag.String("i", "", "generate import code to directory")

// p 仅打印解析结果
var p = flag.Bool("p", false, "print parsed dsl")

// d 目录
var d = flag.String("d", "", "set gorch dsl directory")

// go go命令
var goCmd = flag.String("go", "go", "set go command")

func main() {
	flag.Parse()
	if (*g == "" && !*p && *i == "") || *d == "" {
		flag.Usage()
		return
	}

	astl, err := iparser.ParseDirectory(*d)
	if err != nil {
		log.Fatal("parse dsl error: ", err)
	}
	if *p {
		fmt.Println(astl.Description())
		return
	}

	if *i != "" {
		err := (&codeGenerator{
			p:     astl,
			todir: *i,
			goCmd: *goCmd,
		}).importOperator()
		if err != nil {
			log.Fatal("generate code error: ", err)
		}
		return
	}

	if *g != "" {
		err := (&codeGenerator{
			p:     astl,
			todir: *g,
			goCmd: *goCmd,
		}).releaseOperator()
		if err != nil {
			log.Fatal("generate code error: ", err)
		}
		return
	}
}
