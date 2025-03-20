package tpls

import (
	_ "embed"

	"github.com/gorch/gorch/ort"
)

//go:embed operator_import.tpl
var OperatorImportTpl string

//go:embed operator_middle_code.tpl
var OperatorMiddleCodeTpl string

type PkgOperator struct {
	PkgAlias string
	Name     string
	Struct   string
	Seq      string
	RType    *ort.OperatorRType
}

type PkgOperators struct {
	PkgAlias  string
	Operators map[int]PkgOperator
}

//go:embed operator.tpl
var OperatorTpl string
