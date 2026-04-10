package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type BreakDirective struct {
	Pos
}

type BreakDirectiveHolder interface {
	AcceptBreakDirective(*BreakDirective) error
}

func (bd *BreakDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("BREAK()")
}

func (*BreakDirective) isLeafSnippet() {}
func (*BreakDirective) isExedesc()     {}

// EnterBreakDirective is called when entering the breakDirective production.
func (s *GorchlangParserListener) EnterBreakDirective(ctx *alr.BreakDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	bd := &BreakDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(bd)
}

// ExitBreakDirective is called when exiting the breakDirective production.
func (s *GorchlangParserListener) ExitBreakDirective(ctx *alr.BreakDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	bd := s.Stack.Pop().(*BreakDirective)
	holder := s.Stack.Peek().(BreakDirectiveHolder)
	err := holder.AcceptBreakDirective(bd)
	if err != nil {
		s.AddError(err)
	}
}
