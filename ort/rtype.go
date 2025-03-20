package ort

import (
	"fmt"
	"reflect"
	"strings"
)

type OperatorRType struct {
	Typ         reflect.Type
	HasPrepare  bool
	Injects     []*Field
	Extracts    []*Field
	AllDIFields map[string]struct{}
}

type tmpPrepareOperator interface {
	Prepare() error
}

func AnalyzeOperator[T any]() *OperatorRType {
	s := new(T)
	rt := reflect.TypeOf(s).Elem()

	of := &OperatorRType{
		Typ:         rt,
		Injects:     []*Field{},
		Extracts:    []*Field{},
		AllDIFields: map[string]struct{}{},
	}

	for i := range rt.NumField() {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}

		if field.Tag == "" {
			continue
		}

		if field.Type.Kind() != reflect.Ptr &&
			field.Type.Kind() != reflect.Interface {
			panic(fmt.Sprintf("inject/extract object only support ptr or interface: %s.%s %s",
				rt.Name(), field.Name, field.Type))
		}
		of.structTagFromField(i, field)
	}

	if _, ok := any(s).(tmpPrepareOperator); ok {
		of.HasPrepare = true
	}

	return of
}

type Field struct {
	Typ reflect.Type

	Name     string
	Index    int
	Optional bool
	Replace  bool
}

func (a *OperatorRType) structTagFromField(idx int, field reflect.StructField) {
	for _, tagName := range []string{"inject", "extract"} {
		tag, has := field.Tag.Lookup(tagName)
		if !has {
			continue
		}
		st := &Field{Typ: field.Type, Index: idx, Name: field.Name}
		opts := strings.Split(tag, ",")
		for _, opt := range opts {
			switch opt {
			case "optional":
				st.Optional = true
			case "replace":
				st.Replace = true
			}
		}

		a.AllDIFields[st.Name] = struct{}{}
		if tagName == "inject" {
			a.Injects = append(a.Injects, st)
		} else if tagName == "extract" {
			a.Extracts = append(a.Extracts, st)
		}
	}
}
