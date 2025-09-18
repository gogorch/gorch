package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type SwitchDirective struct {
	Pos

	OperatorStmt   *OperatorStmt
	SwitchCaseList []*SwitchCaseDirective
	SwitchCases    map[string]*SwitchCaseDirective
}

type SwitchDirectiveHolder interface {
	AcceptSwitchDirective(*SwitchDirective) error
}

func (ss *SwitchDirective) AcceptOperatorStmt(os *OperatorStmt) error {
	ss.OperatorStmt = os
	return nil
}

func (ss *SwitchDirective) AcceptSwitchCaseDirective(sas *SwitchCaseDirective) error {
	_, ok := ss.SwitchCases[sas.CaseName]
	if ok {
		return &ParseError{
			Pos: sas.Pos,
			Msg: "duplicate switch case name",
			KV:  map[string]string{"name": sas.CaseName},
		}
	}
	ss.SwitchCaseList = append(ss.SwitchCaseList, sas)
	ss.SwitchCases[sas.CaseName] = sas
	return nil
}

func (ss *SwitchDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("SWITCH(\n")
	ss.OperatorStmt.Description(s, "")
	s.WriteString("){\n")
	scll := len(ss.SwitchCaseList)
	blank += " "
	for idx, sas := range ss.SwitchCaseList {
		sas.Description(s, blank)
		if idx < scll-1 {
			s.WriteString(",\n")
		}
	}
	s.WriteString("\n}\n")
}

func (*SwitchDirective) isLeafSnippet() {}
func (*SwitchDirective) isExedesc()     {}

// EnterSwitchDirective is called when entering the switchDirective production.
func (s *GorchlangParserListener) EnterSwitchDirective(ctx *alr.SwitchDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := &SwitchDirective{
		Pos:            ToPos(ctx, s.File),
		SwitchCases:    make(map[string]*SwitchCaseDirective, 2),
		SwitchCaseList: make([]*SwitchCaseDirective, 0, 2),
	}
	s.Stack.Push(ss)
}

// ExitSwitchDirective is called when exiting the switchDirective production.
func (s *GorchlangParserListener) ExitSwitchDirective(ctx *alr.SwitchDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := s.Stack.Pop().(*SwitchDirective)
	holder := s.Stack.Peek().(SwitchDirectiveHolder)
	err := holder.AcceptSwitchDirective(ss)
	if err != nil {
		s.AddError(err)
	}
}
