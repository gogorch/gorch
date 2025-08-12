package engine

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/gogorch/gorch/internal/ort"
)

type operatorFields ort.OperatorRType

func ditool[T any]() *operatorFields {
	of := ort.AnalyzeOperator[T]()
	g := (*operatorFields)(unsafe.Pointer(of))
	return g
}

func (a *operatorFields) inject(op any, c *container) error {
	opRVal := reflect.ValueOf(op).Elem()
	for _, ifields := range a.Injects {
		v, ok := c.getIns(ifields.Typ)
		if ok {
			opRVal.Field(ifields.Index).Set(v)
		} else {
			if ifields.Optional {
				continue
			}
			return fmt.Errorf("%q inject error: %s not found", a.Typ, ifields.Typ.String())
		}
	}
	return nil
}

func (a *operatorFields) extract(op any, c *container) error {
	opRVal := reflect.ValueOf(op).Elem()
	for _, efields := range a.Extracts {
		f := opRVal.Field(efields.Index)
		if f.IsNil() && !efields.Optional {
			continue
		}
		if err := c.setIns(efields.Typ, f, efields.Replace); err != nil {
			return fmt.Errorf("%q extract error: %s", a.Typ, err)
		}
	}
	return nil
}
