package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type UntilDirective struct {
	Pos
	OperatorStmt *OperatorStmt
}

type UntilDirectiveHolder interface {
	AcceptUntilDirective(*UntilDirective) error
}

func (ud *UntilDirective) AcceptOperatorStmt(os *OperatorStmt) error {
	ud.OperatorStmt = os
	return nil
}

func (ud *UntilDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("UNTIL(")
	ud.OperatorStmt.Description(s, blank)
	s.WriteString(")")
}

func (*UntilDirective) isLeafSnippet() {}
func (*UntilDirective) isExedesc()     {}

// EnterUntilDirective is called when entering the untilDirective production.
func (s *GorchlangParserListener) EnterUntilDirective(ctx *alr.UntilDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ud := &UntilDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(ud)
}

// ExitUntilDirective is called when exiting the untilDirective production.
func (s *GorchlangParserListener) ExitUntilDirective(ctx *alr.UntilDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ud := s.Stack.Pop().(*UntilDirective)
	holder := s.Stack.Peek().(UntilDirectiveHolder)
	err := holder.AcceptUntilDirective(ud)
	if err != nil {
		s.AddError(err)
	}
}
