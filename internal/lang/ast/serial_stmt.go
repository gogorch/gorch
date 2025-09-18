package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type SerialStmt struct {
	Pos
	LeafSnippetList []*LeafSnippet
}

type SerialStmtHolder interface {
	AcceptSerialStmt(*SerialStmt) error
}

func (ss *SerialStmt) AcceptLeafSnippet(ls *LeafSnippet) error {
	ss.LeafSnippetList = append(ss.LeafSnippetList, ls)
	return nil
}

func (ss *SerialStmt) AcceptSkipDirective(ls *SkipDirective) error {
	ss.LeafSnippetList = append(ss.LeafSnippetList, &LeafSnippet{Leaf: ls})
	return nil
}

func (*SerialStmt) isExedesc() {}

func (ss *SerialStmt) Description(s *strings.Builder, blank string) {
	lsll := len(ss.LeafSnippetList)
	for idx, ls := range ss.LeafSnippetList {
		ls.Description(s, "")
		if idx < lsll-1 {
			if lsll > 3 {
				s.WriteString("\n")
				s.WriteString(blank)
			} else {
				s.WriteString(" ")
			}
			s.WriteString("-> ")
		}
	}
}

// EnterSerial is called when production serial is entered.
func (s *GorchlangParserListener) EnterSerialStmt(ctx *alr.SerialStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := &SerialStmt{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(ss)
}

// ExitSerial is called when production serial is exited.
func (s *GorchlangParserListener) ExitSerialStmt(ctx *alr.SerialStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := s.Stack.Pop().(*SerialStmt)

	holder := s.Stack.Peek().(SerialStmtHolder)
	err := holder.AcceptSerialStmt(ss)
	if err != nil {
		s.AddError(err)
	}
}

type SerialBracketStmt struct {
	IgnoreError bool
	SerialStmt  *SerialStmt
}

type SerialBracketStmtHolder interface {
	AcceptSerialBracketStmt(*SerialBracketStmt) error
}

func (sbs *SerialBracketStmt) AcceptSerialStmt(ss *SerialStmt) error {
	sbs.SerialStmt = ss
	return nil
}

func (sbs *SerialBracketStmt) Description(s *strings.Builder, blank string) {
	s.WriteString("(")
	sbs.SerialStmt.Description(s, blank)
	s.WriteString(")")
}

func (sbs *SerialBracketStmt) AcceptIgnoreErrorFlag() error {
	sbs.IgnoreError = true
	return nil
}

func (*SerialBracketStmt) isLeafSnippet() {}

var _ LeafSnippetI = (*SerialBracketStmt)(nil)

// EnterSerialBracketStmt is called when production serialBracketStmt is entered.
func (s *GorchlangParserListener) EnterSerialBracketStmt(ctx *alr.SerialBracketStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := &SerialBracketStmt{}
	s.Stack.Push(ss)
}

// ExitSerialBracketStmt is called when production serialBracketStmt is exited.
func (s *GorchlangParserListener) ExitSerialBracketStmt(ctx *alr.SerialBracketStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	sbs := s.Stack.Pop().(*SerialBracketStmt)

	holder := s.Stack.Peek().(SerialBracketStmtHolder)
	err := holder.AcceptSerialBracketStmt(sbs)
	if err != nil {
		s.AddError(err)
	}
}
