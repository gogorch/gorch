package ast

import (
	"strings"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type ExedescStmt struct {
	Pos

	ExedescStmt ExedescI
}

type ExedescStmtHolder interface {
	AcceptExedescStmt(*ExedescStmt) error
}

func (gs *ExedescStmt) AcceptSerialStmt(ss *SerialStmt) error {
	gs.ExedescStmt = ss
	return nil
}

func (gs *ExedescStmt) AcceptConcurrentStmt(cs *ConcurrentStmt) error {
	gs.ExedescStmt = cs
	return nil
}

func (gs *ExedescStmt) AcceptWrapDirective(ws *WrapStmt) error {
	gs.ExedescStmt = ws
	return nil
}

func (gs *ExedescStmt) AcceptWrapBracketStmt(ws *WrapBracketStmt) error {
	gs.ExedescStmt = ws.WrapStmt
	return nil
}

func (gs *ExedescStmt) AcceptOperatorStmt(os *OperatorStmt) error {
	gs.ExedescStmt = os
	return nil
}

func (gs *ExedescStmt) AcceptSwitchDirective(sd *SwitchDirective) error {
	gs.ExedescStmt = sd
	return nil
}

func (gs *ExedescStmt) AcceptUnfoldDirective(ud *UnfoldDirective) error {
	gs.ExedescStmt = ud
	return nil
}

func (gs *ExedescStmt) Description(sb *strings.Builder, blank string) {
	sb.WriteString(blank)
	gs.ExedescStmt.Description(sb, blank)
}

// EnterExedescStmt is called when production exedescStmt is entered.
func (s *GorchlangParserListener) EnterExedescStmt(ctx *alr.ExedescStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	gs := &ExedescStmt{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(gs)
}

// ExitExedescStmt is called when production exedescStmt is exited.
func (s *GorchlangParserListener) ExitExedescStmt(ctx *alr.ExedescStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	gs := s.Stack.Pop().(*ExedescStmt)
	holder := s.Stack.Peek().(ExedescStmtHolder)
	err := holder.AcceptExedescStmt(gs)
	if err != nil {
		s.AddError(err)
	}
}
