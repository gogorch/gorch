package ast

import (
	"strconv"
	"strings"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type OperatorDirective struct {
	Pos

	Seq        int64
	PkgPath    string
	StructName string
	Name       string
}

type OperatorDirectiveHolder interface {
	AcceptOperatorDirective(*OperatorDirective) error
}

func (ros *OperatorDirective) AcceptStringLiteral(s string) error {
	return nil
}

func (ros *OperatorDirective) AcceptIntegerLiteral(val int64) error {
	ros.Seq = val
	return nil
}

func (ros *OperatorDirective) Description(sb *strings.Builder, blank string) {
	sb.WriteString(blank)
	sb.WriteString("OPERATOR(\"")
	sb.WriteString(ros.PkgPath)
	sb.WriteString("\", \"")
	sb.WriteString(ros.StructName)
	sb.WriteString("\"")
	if len(ros.Name) > 0 && ros.Name != ros.StructName {
		sb.WriteString(", \"")
		sb.WriteString(ros.Name)
		sb.WriteString("\"")
	}
	sb.WriteString(", ")
	sb.WriteString(strconv.FormatInt(ros.Seq, 10))
	sb.WriteString(")")
}

// EnterRegisterOperatorDirective is called when entering the registerOperatorDirective production.
func (s *GorchlangParserListener) EnterRegisterOperatorDirective(ctx *alr.RegisterOperatorDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ros := &OperatorDirective{Pos: ToPos(ctx, s.File)}

	if ctx.GetPkgPath() != nil {
		sval := ctx.GetPkgPath().GetText()
		if len(sval) <= 2 {
			err := &ParseError{
				Pos: ros.Pos,
				Msg: "register operator syntax error",
			}
			s.AddError(err)
			goto structName
		}
		ros.PkgPath = sval[1 : len(sval)-1]
	}

structName:
	if ctx.GetStructName() != nil {
		sval := ctx.GetStructName().GetText()
		if len(sval) <= 2 {
			err := &ParseError{
				Pos: ros.Pos,
				Msg: "register operator sturct empty",
			}
			s.AddError(err)
			goto operatorName
		}
		ros.StructName = sval[1 : len(sval)-1]
	}

operatorName:
	if ctx.GetOperatorName() != nil {
		sval := ctx.GetOperatorName().GetText()
		if len(sval) < 2 {
			err := &ParseError{
				Pos: ros.Pos,
				Msg: "register operator name empty",
			}
			s.AddError(err)
			goto last
		}
		ros.Name = sval[1 : len(sval)-1]
	}

	if len(ros.Name) == 0 {
		ros.Name = ros.StructName
	}

last:
	s.Stack.Push(ros)
}

// ExitRegisterOperatorDirective is called when exiting the registerOperatorDirective production.
func (s *GorchlangParserListener) ExitRegisterOperatorDirective(ctx *alr.RegisterOperatorDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	ros := s.Stack.Pop().(*OperatorDirective)

	holder := s.Stack.Peek().(OperatorDirectiveHolder)
	err := holder.AcceptOperatorDirective(ros)
	if err != nil {
		s.AddError(err)
	}
}

// Operators 全局的算子收集器，Primary的属性
type Operators struct {
	Sort   []*OperatorDirective
	Names  map[string]struct{}
	seqDup map[int64]struct{}
}

func (ods *Operators) Append(od *OperatorDirective) error {
	if ods.Names == nil {
		ods.Names = make(map[string]struct{}, 20)
	}

	if ods.seqDup == nil {
		ods.seqDup = make(map[int64]struct{}, 20)
	}

	if ods.Sort == nil {
		ods.Sort = make([]*OperatorDirective, 0, 20)
	}

	if _, has := ods.Names[od.Name]; has {
		return &ParseError{
			Pos: od.Pos,
			Msg: "register operator duplicate name",
			KV:  map[string]string{"name": od.Name},
		}
	}

	if _, has := ods.seqDup[od.Seq]; has {
		return &ParseError{
			Pos: od.Pos,
			Msg: "register operator duplicate seq",
			KV:  map[string]string{"seq": strconv.FormatInt(od.Seq, 10)},
		}
	}

	ods.seqDup[od.Seq] = struct{}{}
	ods.Names[od.Name] = struct{}{}
	ods.Sort = append(ods.Sort, od)

	return nil
}

func (ods *Operators) Has(n string) bool {
	_, ok := ods.Names[n]
	return ok
}

func (ods *Operators) Len() int {
	return len(ods.Sort)
}
