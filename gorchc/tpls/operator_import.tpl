// Package operatorimport 引入外部代码库
// 本文件应该保存到你的代码库中，用来在go.mod中固定依赖库的版本
// Code generated by "gorchc"; DO NOT EDIT.
package gen_operators


import (
	{{- range $pkg, $operators := .Pkgs}}
	_ "{{$pkg}}"
	{{- end}}
)