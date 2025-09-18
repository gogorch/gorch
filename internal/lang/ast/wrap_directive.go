package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type WrapStmt struct {
	Pos
	WrapOpStmtList []*OperatorStmt
	LeafSnippet    *LeafSnippet
}

type WrapDirectiveHolder interface {
	AcceptWrapDirective(*WrapStmt) error
}

func (ws *WrapStmt) AcceptOperatorStmt(os *OperatorStmt) error {
	ws.WrapOpStmtList = append(ws.WrapOpStmtList, os)
	return nil
}

func (ws *WrapStmt) AcceptLeafSnippet(ls *LeafSnippet) error {
	ws.LeafSnippet = ls
	return nil
}

func (ws *WrapStmt) Description(s *strings.Builder, blank string) {
	for _, os := range ws.WrapOpStmtList {
		os.Description(s, blank)
		s.WriteString(" | ")
	}

	ws.LeafSnippet.Description(s, blank)
}

func (*WrapStmt) isExedesc() {}

// EnterWrapStmt is called when production wrapStmt is entered.
func (s *GorchlangParserListener) EnterWrapStmt(ctx *alr.WrapStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	ws := &WrapStmt{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(ws)
}

// ExitWrapStmt is called when production wrapStmt is exited.
func (s *GorchlangParserListener) ExitWrapStmt(ctx *alr.WrapStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	ws := s.Stack.Pop().(*WrapStmt)
	holder := s.Stack.Peek().(WrapDirectiveHolder)
	err := holder.AcceptWrapDirective(ws)
	if err != nil {
		s.AddError(err)
	}
}

type WrapBracketStmt struct {
	WrapStmt    *WrapStmt
	IgnoreError bool
}

type WrapBracketStmtHolder interface {
	AcceptWrapBracketStmt(*WrapBracketStmt) error
}

func (wbs *WrapBracketStmt) AcceptWrapDirective(ws *WrapStmt) error {
	wbs.WrapStmt = ws
	return nil
}

func (wbs *WrapBracketStmt) Description(s *strings.Builder, blank string) {
	s.WriteString("(")
	wbs.WrapStmt.Description(s, blank)
	s.WriteString(")")
}

func (wbs *WrapBracketStmt) AcceptIgnoreErrorFlag() error {
	wbs.IgnoreError = true
	return nil
}

func (*WrapBracketStmt) isLeafSnippet() {}

var _ LeafSnippetI = (*WrapBracketStmt)(nil)

// EnterWrapBracketStmt is called when production wrapBracketStmt is entered.
func (s *GorchlangParserListener) EnterWrapBracketStmt(ctx *alr.WrapBracketStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	wbs := &WrapBracketStmt{}
	s.Stack.Push(wbs)
}

// ExitWrapBracketStmt is called when production wrapBracketStmt is exited.
func (s *GorchlangParserListener) ExitWrapBracketStmt(ctx *alr.WrapBracketStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	wbs := s.Stack.Pop().(*WrapBracketStmt)
	holder := s.Stack.Peek().(WrapBracketStmtHolder)
	err := holder.AcceptWrapBracketStmt(wbs)
	if err != nil {
		s.AddError(err)
	}
}
