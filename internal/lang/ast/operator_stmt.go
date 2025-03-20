package ast

import (
	"fmt"
	"strings"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type OperatorStmt struct {
	Pos
	Name        string
	IgnoreError bool

	ArgsStmt *ArgsStmt

	WaitList []*WaitDirective
	sort     []any
}

type OperatorStmtHolder interface {
	AcceptOperatorStmt(*OperatorStmt) error
}

type OperatorStmtAppender interface {
	AppendOperatorStmt(*OperatorStmt) error
}

func (os *OperatorStmt) AcceptArgsStmt(args *ArgsStmt) error {
	os.ArgsStmt = args
	return nil
}

func (os *OperatorStmt) AcceptWaitDirective(bws *WaitDirective) error {
	os.WaitList = append(os.WaitList, bws)
	os.sort = append(os.sort, bws)
	return nil
}

func (*OperatorStmt) isLeafSnippet() {}
func (*OperatorStmt) isExedesc()     {}

func (os *OperatorStmt) Description(s *strings.Builder, blank string) {
	if os.IgnoreError {
		s.WriteString("@")
	}
	s.WriteString(os.Name)
	if os.ArgsStmt != nil || len(os.sort) > 0 {
		s.WriteString("(")
		if os.ArgsStmt != nil {
			os.ArgsStmt.Description(s, "")
		}

		if len(os.sort) > 0 {
			if os.ArgsStmt != nil {
				s.WriteString(", ")
			}
			for idx, one := range os.sort {
				switch v := one.(type) {
				case *WaitDirective:
					v.Description(s, "")
				}

				if idx < len(os.sort)-1 {
					s.WriteString(", ")
				}
			}

		}
		s.WriteString(")")
	}
}

func (os *OperatorStmt) AcceptIgnoreErrorFlag() error {
	os.IgnoreError = true
	return nil
}

func (os *OperatorStmt) String() string {
	builder := &strings.Builder{}

	os.Description(builder, "")
	return builder.String()
}

// EnterOperatorStmt is called when entering the operatorStmt production.
func (s *GorchlangParserListener) EnterOperatorStmt(ctx *alr.OperatorStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	os := &OperatorStmt{
		Name: ctx.GetOperatorName().GetText(),
		Pos:  ToPos(ctx, s.File),
	}
	s.Stack.Push(os)
}

// ExitOperatorStmt is called when exiting the operatorStmt production.
func (s *GorchlangParserListener) ExitOperatorStmt(ctx *alr.OperatorStmtContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	os := s.Stack.Pop().(*OperatorStmt)

	var err error
	if holder, ok := s.Stack.Peek().(OperatorStmtHolder); ok {
		err = holder.AcceptOperatorStmt(os)
	} else if appender, ok := s.Stack.Peek().(OperatorStmtAppender); ok {
		err = appender.AppendOperatorStmt(os)
	} else {
		err = fmt.Errorf("OperatorStmt %v not have accepter", os)
	}

	if err != nil {
		s.AddError(err)
	}
}
