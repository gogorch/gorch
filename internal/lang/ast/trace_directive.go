package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type TraceDirective struct {
	Pos
	Name        string
	ExedescStmt *ExedescStmt
}

type TraceDirectiveHolder interface {
	AcceptTraceDirective(*TraceDirective) error
}

func (td *TraceDirective) AcceptStringLiteral(name string) error {
	td.Name = name
	return nil
}

func (td *TraceDirective) AcceptExedescStmt(gs *ExedescStmt) error {
	td.ExedescStmt = gs
	return nil
}

func (td *TraceDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("TRACE(\"")
	s.WriteString(td.Name)
	s.WriteString("\") {\n")
	td.ExedescStmt.Description(s, blank+"  ")
	s.WriteString("\n")
	s.WriteString(blank)
	s.WriteString("}")
}

func (*TraceDirective) isLeafSnippet() {}
func (*TraceDirective) isExedesc()     {}

// EnterTraceDirective is called when entering the traceDirective production.
func (s *GorchlangParserListener) EnterTraceDirective(ctx *alr.TraceDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	td := &TraceDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(td)
}

// ExitTraceDirective is called when exiting the traceDirective production.
func (s *GorchlangParserListener) ExitTraceDirective(ctx *alr.TraceDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	td := s.Stack.Pop().(*TraceDirective)
	holder := s.Stack.Peek().(TraceDirectiveHolder)
	if err := holder.AcceptTraceDirective(td); err != nil {
		s.AddError(err)
	}
}
