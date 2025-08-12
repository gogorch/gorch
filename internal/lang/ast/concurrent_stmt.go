package ast

import (
	"strings"

	"github.com/gogorch/gorch/internal/lang/iantlr/alr"
)

type ConcurrentStmt struct {
	Pos
	IgnoreError     bool
	LeafSnippetList []*LeafSnippet
}

type ConcurrentStmtHolder interface {
	AcceptConcurrentStmt(*ConcurrentStmt) error
}

func (cs *ConcurrentStmt) AcceptLeafSnippet(ls *LeafSnippet) error {
	cs.LeafSnippetList = append(cs.LeafSnippetList, ls)
	return nil
}

func (*ConcurrentStmt) isExedesc() {}

func (cs *ConcurrentStmt) Description(s *strings.Builder, blank string) {
	lsll := len(cs.LeafSnippetList)
	s.WriteString("[")
	for idx, ls := range cs.LeafSnippetList {
		ls.Description(s, blank)
		if idx < lsll-1 {
			if lsll > 3 {
				s.WriteString("\n")
				s.WriteString(blank)
			}
			s.WriteString(", ")
		}
	}
	s.WriteString("]")
}

func (cs *ConcurrentStmt) AcceptIgnoreErrorFlag() error {
	cs.IgnoreError = true
	return nil
}

func (*ConcurrentStmt) isLeafSnippet() {}

var _ LeafSnippetI = (*ConcurrentStmt)(nil)

// EnterConcurrent is called when production concurrent is entered.
func (s *GorchlangParserListener) EnterConcurrentStmt(ctx *alr.ConcurrentStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	cs := &ConcurrentStmt{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(cs)
}

// ExitConcurrent is called when production concurrent is exited.
func (s *GorchlangParserListener) ExitConcurrentStmt(ctx *alr.ConcurrentStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	cs := s.Stack.Pop().(*ConcurrentStmt)
	holder := s.Stack.Peek().(ConcurrentStmtHolder)
	err := holder.AcceptConcurrentStmt(cs)
	if err != nil {
		s.AddError(err)
	}
}
