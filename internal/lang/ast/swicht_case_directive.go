package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type SwitchCaseDirective struct {
	Pos
	LeafSnippet *LeafSnippet

	CaseName string
}

type SwitchCaseDirectiveHolder interface {
	AcceptSwitchCaseDirective(*SwitchCaseDirective) error
}

func (sas *SwitchCaseDirective) AcceptLeafSnippet(ls *LeafSnippet) error {
	sas.LeafSnippet = ls
	return nil
}

func (sas *SwitchCaseDirective) AcceptStringLiteral(val string) error {
	sas.CaseName = val
	return nil
}

func (sas *SwitchCaseDirective) Description(s *strings.Builder, blank string) {
	s.WriteString(blank)
	s.WriteString("CASE ")
	s.WriteString("\"")
	s.WriteString(sas.CaseName)
	s.WriteString("\" => ")
	sas.LeafSnippet.Description(s, blank)
}

// EnterSwitchCaseDirective is called when entering the switchCaseDirective production.
func (s *GorchlangParserListener) EnterSwitchCaseDirective(ctx *alr.SwitchCaseDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	sas := &SwitchCaseDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(sas)
}

// ExitSwitchCaseDirective is called when exiting the switchCaseDirective production.
func (s *GorchlangParserListener) ExitSwitchCaseDirective(ctx *alr.SwitchCaseDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	sas := s.Stack.Pop().(*SwitchCaseDirective)
	holder := s.Stack.Peek().(SwitchCaseDirectiveHolder)
	err := holder.AcceptSwitchCaseDirective(sas)
	if err != nil {
		s.AddError(err)
	}
}
