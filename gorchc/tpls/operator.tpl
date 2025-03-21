// Package {{.Package}} 生成的算子代码
// 本文件不应该保存在你的代码库中，而是应该在线上编译前产出
// Code generated by "gorchc"; DO NOT EDIT.
package {{.Package}}

import (
    "fmt"
    {{- range $pkg, $operators := .Pkgs}}
	{{$operators.PkgAlias}} "{{$pkg}}"
    {{- end}}
    "github.com/gorch/gorch"
)
var _ = fmt.Printf
{{- range $pkg, $operators := .Pkgs}}
{{- range $idx0, $operator := $operators.Operators}}
type {{$operator.Name}}{{$operator.Struct}} {{$operator.PkgAlias}}.{{$operator.Struct}}
{{- if $operator.RType.AllDIFields}}
func(o *{{$operator.Name}}{{$operator.Struct}}) Execute(ctx *gorch.Context) (err error) {
    {{- range $idx1, $injectItem := $operator.RType.Injects}}
    {{- if $injectItem.Optional}}
    _ = ctx.MutableIns(&o.{{$injectItem.Name}})
    {{- else}}
    if er0 := ctx.MutableIns(&o.{{$injectItem.Name}}); er0 != nil {
        return fmt.Errorf("inject {{$pkg}}.{{$operator.Struct}}.{{$injectItem.Name}} error: %v", er0)
    }
    {{- end}}
    {{- end}}
    {{- if $operator.RType.Extracts}}
    defer func() {
    {{- range $idx2, $extractItem := $operator.RType.Extracts}}
        {{- if $extractItem.Optional}}
        _ = ctx.RegisterIns(&o.{{$extractItem.Name}}, {{if $extractItem.Replace}}true{{else}}false{{end}})
        {{- else}}
        if er0 := ctx.RegisterIns(&o.{{$extractItem.Name}}, {{if $extractItem.Replace}}true{{else}}false{{end}}); er0 != nil {
            err = fmt.Errorf("extract {{$pkg}}.{{$operator.Struct}}.{{$extractItem.Name}} error: %v", er0)
            return
        }
        {{- end}}
        {{- range $name, $_ := $operator.RType.AllDIFields}}
        o.{{$name}} = nil
        {{- end}}
    {{- end}}
    }()
    {{- end}}
    return (*{{$operator.PkgAlias}}.{{$operator.Struct}})(o).Execute(ctx)
}
{{- else}}
func(o *{{$operator.Name}}{{$operator.Struct}}) Execute(ctx *gorch.Context) (err error) {return (*{{$operator.PkgAlias}}.{{$operator.Struct}})(o).Execute(ctx)}
{{- end}}
{{- if $operator.RType.HasPrepare}}
func(o *{{$operator.Name}}{{$operator.Struct}}) Prepare() error {return (*{{$operator.PkgAlias}}.{{$operator.Struct}})(o).Prepare()}
{{- end}}
func(o *{{$operator.Name}}{{$operator.Struct}}) IsGenerateOperatorCode() {}
{{- end}}
{{- end}}

func init() {
	must(
    {{- range $pkg, $operators := .Pkgs}}
    {{- range $idx0, $operator := $operators.Operators}}
        gorch.RegOp[{{$operator.Name}}{{$operator.Struct}}]("{{$operator.Name}}", {{$operator.Seq}}),
    {{- end}}
    {{- end}}
    )
}

func must(errs ...error) {
    hasErr := false
 	for _, e := range errs {
 		if e != nil {
            hasErr = true
 			fmt.Println(e)
 		}
 	}
    if hasErr {
        panic("register operator has error")
    }
}