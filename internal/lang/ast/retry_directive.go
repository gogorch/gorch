package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type RetryDirective struct {
	Pos
	Args        *ArgsStmt
	ExedescStmt *ExedescStmt
}

type RetryDirectiveHolder interface {
	AcceptRetryDirective(*RetryDirective) error
}

func (rd *RetryDirective) AcceptArgsStmt(args *ArgsStmt) error {
	rd.Args = args
	return nil
}

func (rd *RetryDirective) AcceptExedescStmt(gs *ExedescStmt) error {
	rd.ExedescStmt = gs
	return nil
}

func (rd *RetryDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("RETRY(")
	if rd.Args != nil {
		rd.Args.Description(s, blank)
	}
	s.WriteString(") {\n")
	rd.ExedescStmt.Description(s, blank+"  ")
	s.WriteString("\n")
	s.WriteString(blank)
	s.WriteString("}")
}

func (*RetryDirective) isLeafSnippet() {}
func (*RetryDirective) isExedesc()     {}

// EnterRetryDirective is called when entering the retryDirective production.
func (s *GorchlangParserListener) EnterRetryDirective(ctx *alr.RetryDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	rd := &RetryDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(rd)
}

// ExitRetryDirective is called when exiting the retryDirective production.
func (s *GorchlangParserListener) ExitRetryDirective(ctx *alr.RetryDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	rd := s.Stack.Pop().(*RetryDirective)
	holder := s.Stack.Peek().(RetryDirectiveHolder)
	err := holder.AcceptRetryDirective(rd)
	if err != nil {
		s.AddError(err)
	}
}
