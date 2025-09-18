package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type GoDirective struct {
	Pos
	ExedescStmt *ExedescStmt

	GoName string
}

type GoDirectiveHolder interface {
	AcceptGoDirective(*GoDirective) error
}

type GoDirectiveAppender interface {
	AppendGoDirective(brs *GoDirective) error
}

func (ss *GoDirective) AcceptExedescStmt(gs *ExedescStmt) error {
	ss.ExedescStmt = gs
	return nil
}

func (brs *GoDirective) AcceptStringLiteral(val string) error {
	brs.GoName = val
	return nil
}

func (brs *GoDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("GO(")
	brs.ExedescStmt.Description(s, "")
	s.WriteString(", \"")
	s.WriteString(brs.GoName)
	s.WriteString("\")")
}

func (*GoDirective) isLeafSnippet() {}
func (*GoDirective) isExedesc()     {}

// EnterGoDirective is called when entering the goDirective production.
func (s *GorchlangParserListener) EnterGoDirective(ctx *alr.GoDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	gd := &GoDirective{Pos: ToPos(ctx, s.File)}

	s.Stack.Push(gd)
}

// ExitGoDirective is called when exiting the goDirective production.
func (s *GorchlangParserListener) ExitGoDirective(ctx *alr.GoDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	gd := s.Stack.Pop().(*GoDirective)
	holder := s.Stack.Peek().(GoDirectiveHolder)
	err := holder.AcceptGoDirective(gd)
	if err != nil {
		s.AddError(err)
	}
	if s.GoDirectiveAppender != nil {
		if err := s.GoDirectiveAppender.AppendGoDirective(gd); err != nil {
			s.AddError(err)
		}
	}
}
