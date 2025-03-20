package ast

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/gorch/gorch/internal/lang/utils"
)

type GorchlangParserListener struct {
	Stack *utils.Stack

	ParseErrors []error

	File string

	Primary *Primary

	// UnfoldDirectiveAppender UNFOLD指令收集器
	// 只有STRAT和FRAGMENT两个指令
	UnfoldDirectiveAppender UnfoldDirectiveAppender

	// GoDirectiveAppender GO指令收集器
	// 只支持STRAT和FRAGMENT两个指令
	GoDirectiveAppender GoDirectiveAppender

	// WaitDirectiveNameAppender WAIT指令收集器
	// 只支持STRAT和FRAGMENT两个指令
	WaitDirectiveNameAppender WaitDirectiveNameAppender

	// NoCheckMissDirectiveAcceptor NO_CHECK_OPERATOR_MISS指令接受者
	// 只支持START指令
	NoCheckMissDirectiveAcceptor NoCheckMissDirectiveAcceptor
}

func NewGorchlangParserListener() *GorchlangParserListener {
	return &GorchlangParserListener{
		Stack: utils.NewStack(),
	}
}

func (g *GorchlangParserListener) AddError(e error) {
	g.ParseErrors = append(g.ParseErrors, e)
}

// VisitTerminal is called when a terminal node is visited.
func (s *GorchlangParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *GorchlangParserListener) VisitErrorNode(node antlr.ErrorNode) {
	// fmt.Println("visitErrorNode", node.GetText())
}

// EnterEveryRule is called when any rule is entered.
func (s *GorchlangParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println("enter", ctx.GetRuleContext().GetText())
}

// ExitEveryRule is called when any rule is exited.
func (s *GorchlangParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Println("exit", ctx.GetRuleContext().GetText())
}

// var _ alr.gorchlangListener = (*gorchlangParserListener)(nil)
