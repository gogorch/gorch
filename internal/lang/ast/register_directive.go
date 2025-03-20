package ast

import (
	"strings"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type RegisterDirective struct {
	Pos
	Pkg string

	Operators *Operators
}

type RegisterDirectiveHold interface {
	AcceptRegisterDirective(*RegisterDirective) error
}

func (rs *RegisterDirective) AcceptStringLiteral(val string) error {
	rs.Pkg = val
	return nil
}

func (rs *RegisterDirective) AcceptOperatorDirective(ors *OperatorDirective) error {
	return rs.Operators.Append(ors)
}

func (rs *RegisterDirective) Description(sb *strings.Builder, blank string) {
	sb.WriteString("REGISTER(\"")
	sb.WriteString(rs.Pkg)
	sb.WriteString("\"")
	sb.WriteString(") {\n")

	for _, one := range rs.Operators.Sort {
		one.Description(sb, "  ")
		sb.WriteString("\n")
	}

	sb.WriteString("}\n")
}

// EnterRegisterDirective is called when entering the registerDirective production.
func (s *GorchlangParserListener) EnterRegisterDirective(ctx *alr.RegisterDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	rs := &RegisterDirective{
		Pos: ToPos(ctx, s.File),
		Operators: &Operators{
			Sort:   make([]*OperatorDirective, 0),
			seqDup: make(map[int64]struct{}),
			Names:  make(map[string]struct{}),
		},
	}
	s.Stack.Push(rs)
}

// ExitRegisterDirective is called when exiting the registerDirective production.
func (s *GorchlangParserListener) ExitRegisterDirective(ctx *alr.RegisterDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	rs := s.Stack.Pop().(*RegisterDirective)
	holder := s.Stack.Peek().(RegisterDirectiveHold)
	err := holder.AcceptRegisterDirective(rs)
	if err != nil {
		s.AddError(err)
	}
}

type RegisterDirectives struct {
	Sort []*RegisterDirective

	AllOperators *Operators
}

func NewRegisterDirectives() *RegisterDirectives {
	return &RegisterDirectives{
		Sort: make([]*RegisterDirective, 0),
		AllOperators: &Operators{
			Sort:   make([]*OperatorDirective, 0),
			seqDup: make(map[int64]struct{}),
			Names:  make(map[string]struct{}),
		},
	}
}

func (rd *RegisterDirectives) Append(rs *RegisterDirective) error {
	rd.Sort = append(rd.Sort, rs)
	if rd.AllOperators == nil {
		rd.AllOperators = &Operators{
			Sort:   make([]*OperatorDirective, 0),
			seqDup: make(map[int64]struct{}),
			Names:  make(map[string]struct{}),
		}
	}
	for i := range rs.Operators.Sort {
		if err := rd.AllOperators.Append(rs.Operators.Sort[i]); err != nil {
			return err
		}
	}
	return nil
}

func (rd *RegisterDirectives) Description(sb *strings.Builder, blank string) {
	for idx := range rd.Sort {
		rd.Sort[idx].Description(sb, blank)
	}
}
