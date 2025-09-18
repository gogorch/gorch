package ast

import (
	"fmt"
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type ArgsStmt struct {
	Pos
	ArgList []*ArgStmt
	ArgMap  map[string]*ArgStmt
}

type ArgsStmtHolder interface {
	AcceptArgsStmt(args *ArgsStmt) error
}

func (as *ArgsStmt) AcceptArgStmt(arg *ArgStmt) error {
	if _, ok := as.ArgMap[arg.Key]; ok {
		return fmt.Errorf("arg %q already exists", arg.Key)
	}

	as.ArgList = append(as.ArgList, arg)

	if as.ArgMap == nil {
		as.ArgMap = make(map[string]*ArgStmt, 2)
	}

	as.ArgMap[arg.Key] = arg
	return nil
}

func (args *ArgsStmt) Description(s *strings.Builder, blank string) {
	for idx, a := range args.ArgList {
		s.WriteString(a.Description(""))
		if idx < len(args.ArgList)-1 {
			s.WriteString(", \n")
		}
	}
}

// EnterArgsStmt is called when production arg is entered.
func (s *GorchlangParserListener) EnterArgsStmt(ctx *alr.ArgsStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	a := &ArgsStmt{
		Pos:     ToPos(ctx, s.File),
		ArgList: make([]*ArgStmt, 0),
	}
	s.Stack.Push(a)
}

// ExitArgsStmt is called when production arg is exited.
func (s *GorchlangParserListener) ExitArgsStmt(ctx *alr.ArgsStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	expr := s.Stack.Pop().(*ArgsStmt)
	argHolder := s.Stack.Peek().(ArgsStmtHolder)
	err := argHolder.AcceptArgsStmt(expr)
	if err != nil {
		s.AddError(err)
	}
}
