package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type OnFinishStmt struct {
	Pos
	ExedescStmt *ExedescStmt
}

type OnFinishStmtHolder interface {
	AcceptOnFinishStmt(*OnFinishStmt) error
}

func (ss *OnFinishStmt) AcceptExedescStmt(gs *ExedescStmt) error {
	ss.ExedescStmt = gs
	return nil
}

func (ofs *OnFinishStmt) Description(s *strings.Builder, blank string) {
	s.WriteString(blank)
	s.WriteString("ON_FINISH() {\n")
	ofs.ExedescStmt.Description(s, blank+"  ")
	s.WriteString("\n" + blank)
	s.WriteString("}\n\n")
}

func (ofs *OnFinishStmt) String() string {
	s := &strings.Builder{}
	s.Grow(20)

	return s.String()
}

// EnterOnFinishStmt is called when entering the onFinishStmt production.
func (s *GorchlangParserListener) EnterOnFinishStmt(ctx *alr.OnFinishStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ofd := &OnFinishStmt{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(ofd)
}

// ExitOnFinishStmt is called when exiting the onFinishStmt production.
func (s *GorchlangParserListener) ExitOnFinishStmt(ctx *alr.OnFinishStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ofd := s.Stack.Pop().(*OnFinishStmt)
	holder := s.Stack.Peek().(OnFinishStmtHolder)
	err := holder.AcceptOnFinishStmt(ofd)
	if err != nil {
		s.AddError(err)
	}
}
