// Code generated from /Users/zhangming/gopath/src/github.com/period331/jagep/internal/lang/iantlr/jagelang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // jagelang

import "github.com/antlr4-go/antlr/v4"

// BasejagelangListener is a complete listener for a parse tree produced by jagelangParser.
type BasejagelangListener struct{}

var _ jagelangListener = &BasejagelangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasejagelangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasejagelangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasejagelangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasejagelangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasejagelangListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasejagelangListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterStartDirective is called when production startDirective is entered.
func (s *BasejagelangListener) EnterStartDirective(ctx *StartDirectiveContext) {}

// ExitStartDirective is called when production startDirective is exited.
func (s *BasejagelangListener) ExitStartDirective(ctx *StartDirectiveContext) {}

// EnterFragmentDirective is called when production fragmentDirective is entered.
func (s *BasejagelangListener) EnterFragmentDirective(ctx *FragmentDirectiveContext) {}

// ExitFragmentDirective is called when production fragmentDirective is exited.
func (s *BasejagelangListener) ExitFragmentDirective(ctx *FragmentDirectiveContext) {}

// EnterRegisterDirective is called when production registerDirective is entered.
func (s *BasejagelangListener) EnterRegisterDirective(ctx *RegisterDirectiveContext) {}

// ExitRegisterDirective is called when production registerDirective is exited.
func (s *BasejagelangListener) ExitRegisterDirective(ctx *RegisterDirectiveContext) {}

// EnterOnFinishStmt is called when production onFinishStmt is entered.
func (s *BasejagelangListener) EnterOnFinishStmt(ctx *OnFinishStmtContext) {}

// ExitOnFinishStmt is called when production onFinishStmt is exited.
func (s *BasejagelangListener) ExitOnFinishStmt(ctx *OnFinishStmtContext) {}

// EnterExedescStmt is called when production exedescStmt is entered.
func (s *BasejagelangListener) EnterExedescStmt(ctx *ExedescStmtContext) {}

// ExitExedescStmt is called when production exedescStmt is exited.
func (s *BasejagelangListener) ExitExedescStmt(ctx *ExedescStmtContext) {}

// EnterLeafSnippet is called when production leafSnippet is entered.
func (s *BasejagelangListener) EnterLeafSnippet(ctx *LeafSnippetContext) {}

// ExitLeafSnippet is called when production leafSnippet is exited.
func (s *BasejagelangListener) ExitLeafSnippet(ctx *LeafSnippetContext) {}

// EnterOperatorStmt is called when production operatorStmt is entered.
func (s *BasejagelangListener) EnterOperatorStmt(ctx *OperatorStmtContext) {}

// ExitOperatorStmt is called when production operatorStmt is exited.
func (s *BasejagelangListener) ExitOperatorStmt(ctx *OperatorStmtContext) {}

// EnterSerialStmt is called when production serialStmt is entered.
func (s *BasejagelangListener) EnterSerialStmt(ctx *SerialStmtContext) {}

// ExitSerialStmt is called when production serialStmt is exited.
func (s *BasejagelangListener) ExitSerialStmt(ctx *SerialStmtContext) {}

// EnterSerialBracketStmt is called when production serialBracketStmt is entered.
func (s *BasejagelangListener) EnterSerialBracketStmt(ctx *SerialBracketStmtContext) {}

// ExitSerialBracketStmt is called when production serialBracketStmt is exited.
func (s *BasejagelangListener) ExitSerialBracketStmt(ctx *SerialBracketStmtContext) {}

// EnterConcurrentStmt is called when production concurrentStmt is entered.
func (s *BasejagelangListener) EnterConcurrentStmt(ctx *ConcurrentStmtContext) {}

// ExitConcurrentStmt is called when production concurrentStmt is exited.
func (s *BasejagelangListener) ExitConcurrentStmt(ctx *ConcurrentStmtContext) {}

// EnterUnfoldDirective is called when production unfoldDirective is entered.
func (s *BasejagelangListener) EnterUnfoldDirective(ctx *UnfoldDirectiveContext) {}

// ExitUnfoldDirective is called when production unfoldDirective is exited.
func (s *BasejagelangListener) ExitUnfoldDirective(ctx *UnfoldDirectiveContext) {}

// EnterGoDirective is called when production goDirective is entered.
func (s *BasejagelangListener) EnterGoDirective(ctx *GoDirectiveContext) {}

// ExitGoDirective is called when production goDirective is exited.
func (s *BasejagelangListener) ExitGoDirective(ctx *GoDirectiveContext) {}

// EnterWaitDirective is called when production waitDirective is entered.
func (s *BasejagelangListener) EnterWaitDirective(ctx *WaitDirectiveContext) {}

// ExitWaitDirective is called when production waitDirective is exited.
func (s *BasejagelangListener) ExitWaitDirective(ctx *WaitDirectiveContext) {}

// EnterNoCheckMissDirective is called when production noCheckMissDirective is entered.
func (s *BasejagelangListener) EnterNoCheckMissDirective(ctx *NoCheckMissDirectiveContext) {}

// ExitNoCheckMissDirective is called when production noCheckMissDirective is exited.
func (s *BasejagelangListener) ExitNoCheckMissDirective(ctx *NoCheckMissDirectiveContext) {}

// EnterSkipDirective is called when production skipDirective is entered.
func (s *BasejagelangListener) EnterSkipDirective(ctx *SkipDirectiveContext) {}

// ExitSkipDirective is called when production skipDirective is exited.
func (s *BasejagelangListener) ExitSkipDirective(ctx *SkipDirectiveContext) {}

// EnterWrapStmt is called when production wrapStmt is entered.
func (s *BasejagelangListener) EnterWrapStmt(ctx *WrapStmtContext) {}

// ExitWrapStmt is called when production wrapStmt is exited.
func (s *BasejagelangListener) ExitWrapStmt(ctx *WrapStmtContext) {}

// EnterWrapBracketStmt is called when production wrapBracketStmt is entered.
func (s *BasejagelangListener) EnterWrapBracketStmt(ctx *WrapBracketStmtContext) {}

// ExitWrapBracketStmt is called when production wrapBracketStmt is exited.
func (s *BasejagelangListener) ExitWrapBracketStmt(ctx *WrapBracketStmtContext) {}

// EnterSwitchDirective is called when production switchDirective is entered.
func (s *BasejagelangListener) EnterSwitchDirective(ctx *SwitchDirectiveContext) {}

// ExitSwitchDirective is called when production switchDirective is exited.
func (s *BasejagelangListener) ExitSwitchDirective(ctx *SwitchDirectiveContext) {}

// EnterSwitchCaseDirective is called when production switchCaseDirective is entered.
func (s *BasejagelangListener) EnterSwitchCaseDirective(ctx *SwitchCaseDirectiveContext) {}

// ExitSwitchCaseDirective is called when production switchCaseDirective is exited.
func (s *BasejagelangListener) ExitSwitchCaseDirective(ctx *SwitchCaseDirectiveContext) {}

// EnterRegisterOperatorDirective is called when production registerOperatorDirective is entered.
func (s *BasejagelangListener) EnterRegisterOperatorDirective(ctx *RegisterOperatorDirectiveContext) {
}

// ExitRegisterOperatorDirective is called when production registerOperatorDirective is exited.
func (s *BasejagelangListener) ExitRegisterOperatorDirective(ctx *RegisterOperatorDirectiveContext) {}

// EnterArgStmt is called when production argStmt is entered.
func (s *BasejagelangListener) EnterArgStmt(ctx *ArgStmtContext) {}

// ExitArgStmt is called when production argStmt is exited.
func (s *BasejagelangListener) ExitArgStmt(ctx *ArgStmtContext) {}

// EnterArgsStmt is called when production argsStmt is entered.
func (s *BasejagelangListener) EnterArgsStmt(ctx *ArgsStmtContext) {}

// ExitArgsStmt is called when production argsStmt is exited.
func (s *BasejagelangListener) ExitArgsStmt(ctx *ArgsStmtContext) {}

// EnterIgnoreError is called when production ignoreError is entered.
func (s *BasejagelangListener) EnterIgnoreError(ctx *IgnoreErrorContext) {}

// ExitIgnoreError is called when production ignoreError is exited.
func (s *BasejagelangListener) ExitIgnoreError(ctx *IgnoreErrorContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BasejagelangListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BasejagelangListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BasejagelangListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BasejagelangListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BasejagelangListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BasejagelangListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterDurationLiteral is called when production durationLiteral is entered.
func (s *BasejagelangListener) EnterDurationLiteral(ctx *DurationLiteralContext) {}

// ExitDurationLiteral is called when production durationLiteral is exited.
func (s *BasejagelangListener) ExitDurationLiteral(ctx *DurationLiteralContext) {}
