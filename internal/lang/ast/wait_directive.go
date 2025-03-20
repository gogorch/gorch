package ast

import (
	"strings"
	"time"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type WaitDirective struct {
	Pos
	GoName string

	// Timeout 本次最长等待的时间
	Timeout time.Duration

	// TotalTimeout 总共等待的时间，从Routine开始算起
	TotalTimeout time.Duration

	// NotCheckStart 是否检查routine是否已经启动
	// 如果为false，在Wait时，如果还没启动直接返回；否则会等待routine启动
	NotCheckStart bool

	// 是否忽略错误
	IgnoreError bool

	argsStmt *ArgsStmt
}

type WaitDirectiveHolder interface {
	AcceptWaitDirective(bws *WaitDirective) error
}

type WaitDirectiveNameAppender interface {
	AppendWaitDirectiveName(name string) error
}

func (bws *WaitDirective) AcceptStringLiteral(val string) error {
	bws.GoName = val
	return nil
}

func (bws *WaitDirective) AcceptArgsStmt(args *ArgsStmt) error {
	bws.argsStmt = args

	if v, ok := args.ArgMap["timeout"]; ok && len(v.DurationValue) > 0 {
		bws.Timeout = v.DurationValue[0]
	}

	if v, ok := args.ArgMap["totalTimeout"]; ok && len(v.DurationValue) > 0 {
		bws.TotalTimeout = v.DurationValue[0]
	}

	if v, ok := args.ArgMap["notCheckStart"]; ok && len(v.BooleanValue) > 0 {
		bws.NotCheckStart = v.BooleanValue[0]
	}

	return nil
}

func (bws *WaitDirective) AcceptIgnoreErrorFlag() error {
	bws.IgnoreError = true
	return nil
}

func (bws *WaitDirective) Description(s *strings.Builder, blank string) {
	s.WriteString("WAIT(\"")
	s.WriteString(bws.GoName)
	s.WriteString("\"")
	if bws.argsStmt != nil {
		s.WriteString(", ")
		bws.argsStmt.Description(s, blank)
	}
	s.WriteString(")")
}

func (brs *WaitDirective) isLeafSnippet() {}

var _ LeafSnippetI = (*WaitDirective)(nil)

// EnterWaitDirective is called when entering the waitDirective production.
func (s *GorchlangParserListener) EnterWaitDirective(ctx *alr.WaitDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	bws := &WaitDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(bws)
}

// ExitWaitDirective is called when exiting the waitDirective production.
func (s *GorchlangParserListener) ExitWaitDirective(ctx *alr.WaitDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	bws := s.Stack.Pop().(*WaitDirective)
	holder := s.Stack.Peek().(WaitDirectiveHolder)
	err := holder.AcceptWaitDirective(bws)
	if err != nil {
		s.AddError(err)
	}

	if err := s.WaitDirectiveNameAppender.AppendWaitDirectiveName(bws.GoName); err != nil {
		s.AddError(err)
	}
}
