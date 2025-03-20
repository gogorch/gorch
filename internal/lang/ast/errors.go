package ast

import (
	"fmt"
	"strings"
)

type ParseError struct {
	Pos
	Msg string
	KV  map[string]string
}

func (pe *ParseError) Error() string {
	return fmt.Sprintf("%s, file=%s(%d:%d) %s",
		pe.Msg, pe.File, pe.Pos.StartLine, pe.StartColunm,
		func() string {
			b := &strings.Builder{}
			for k, v := range pe.KV {
				b.WriteString(k)
				b.WriteString("=")
				b.WriteString(v)
				b.WriteString(", ")
			}
			return strings.TrimRight(b.String(), ", ")
		}())
}

var ()
