package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type LoopDirective struct {
	Pos
	Args        *ArgsStmt
	ExedescStmt *ExedescStmt
}

type LoopDirectiveHolder interface {
	AcceptLoopDirective(*LoopDirective) error
}

func (ld *LoopDirective) AcceptArgsStmt(args *ArgsStmt) error {
	ld.Args = args
	return nil
}

func (ld *LoopDirective) AcceptExedescStmt(gs *ExedescStmt) error {
	ld.ExedescStmt = gs
	return nil
}

func (ld *LoopDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("LOOP(")
	if ld.Args != nil {
		ld.Args.Description(s, blank)
	}
	s.WriteString(") {\n")
	ld.ExedescStmt.Description(s, blank+"  ")
	s.WriteString("\n")
	s.WriteString(blank)
	s.WriteString("}")
}

func (*LoopDirective) isLeafSnippet() {}
func (*LoopDirective) isExedesc()     {}

// EnterLoopDirective is called when entering the loopDirective production.
func (s *GorchlangParserListener) EnterLoopDirective(ctx *alr.LoopDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ld := &LoopDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(ld)
}

// ExitLoopDirective is called when exiting the loopDirective production.
func (s *GorchlangParserListener) ExitLoopDirective(ctx *alr.LoopDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ld := s.Stack.Pop().(*LoopDirective)
	holder := s.Stack.Peek().(LoopDirectiveHolder)
	err := holder.AcceptLoopDirective(ld)
	if err != nil {
		s.AddError(err)
	}
}
