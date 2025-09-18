package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type LeafSnippet struct {
	Leaf LeafSnippetI
}

type LeafSnippetHolder interface {
	AcceptLeafSnippet(*LeafSnippet) error
}

func (ls *LeafSnippet) AcceptOperatorStmt(os *OperatorStmt) error {
	ls.Leaf = os
	return nil
}

func (ls *LeafSnippet) AcceptGoDirective(brs *GoDirective) error {
	ls.Leaf = brs
	return nil
}

func (ls *LeafSnippet) AcceptSkipDirective(sd *SkipDirective) error {
	ls.Leaf = sd
	return nil
}

func (ls *LeafSnippet) AcceptUnfoldDirective(ud *UnfoldDirective) error {
	ls.Leaf = ud
	return nil
}

func (ls *LeafSnippet) AcceptConcurrentStmt(cd *ConcurrentStmt) error {
	ls.Leaf = cd
	return nil
}

func (ls *LeafSnippet) AcceptSwitchDirective(sd *SwitchDirective) error {
	ls.Leaf = sd
	return nil
}

func (ls *LeafSnippet) AcceptSerialBracketStmt(ss *SerialBracketStmt) error {
	ls.Leaf = ss
	return nil
}

func (ls *LeafSnippet) AcceptWrapBracketStmt(wbs *WrapBracketStmt) error {
	ls.Leaf = wbs
	return nil
}

func (ls *LeafSnippet) Description(s *strings.Builder, blank string) {
	ls.Leaf.Description(s, blank)
}

// EnterLeafSnippet is called when production leafSnippet is entered.
func (s *GorchlangParserListener) EnterLeafSnippet(ctx *alr.LeafSnippetContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	ls := &LeafSnippet{}
	s.Stack.Push(ls)
}

// ExitLeafSnippet is called when production leafSnippet is exited.
func (s *GorchlangParserListener) ExitLeafSnippet(ctx *alr.LeafSnippetContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ls := s.Stack.Pop().(*LeafSnippet)
	holder := s.Stack.Peek().(LeafSnippetHolder)
	err := holder.AcceptLeafSnippet(ls)
	if err != nil {
		s.AddError(err)
	}
}
