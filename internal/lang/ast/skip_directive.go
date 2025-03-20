package ast

import (
	"strings"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type SkipDirective struct {
	Pos
	OperatorStmt *OperatorStmt
}

type SkipDirectiveHolder interface {
	AcceptSkipDirective(*SkipDirective) error
}

func (ss *SkipDirective) AcceptOperatorStmt(os *OperatorStmt) error {
	ss.OperatorStmt = os
	return nil
}

func (ss *SkipDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("SKIP(")
	ss.OperatorStmt.Description(s, "")
	s.WriteString(")")
}

func (*SkipDirective) isLeafSnippet() {}

var _ LeafSnippetI = (*SkipDirective)(nil)

// EnterSkipDirective is called when entering the skipDirective production.
func (s *GorchlangParserListener) EnterSkipDirective(ctx *alr.SkipDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := &SkipDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(ss)
}

// ExitSkipDirective is called when exiting the skipDirective production.
func (s *GorchlangParserListener) ExitSkipDirective(ctx *alr.SkipDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := s.Stack.Pop().(*SkipDirective)

	holder := s.Stack.Peek().(SkipDirectiveHolder)
	err := holder.AcceptSkipDirective(ss)
	if err != nil {
		s.AddError(err)
	}
}
