// Code generated from gorchlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package iantlr // gorchlang
import "github.com/antlr4-go/antlr/v4"

// BasegorchlangListener is a complete listener for a parse tree produced by gorchlangParser.
type BasegorchlangListener struct{}

var _ gorchlangListener = &BasegorchlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegorchlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegorchlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegorchlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegorchlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasegorchlangListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasegorchlangListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterStartDirective is called when production startDirective is entered.
func (s *BasegorchlangListener) EnterStartDirective(ctx *StartDirectiveContext) {}

// ExitStartDirective is called when production startDirective is exited.
func (s *BasegorchlangListener) ExitStartDirective(ctx *StartDirectiveContext) {}

// EnterFragmentDirective is called when production fragmentDirective is entered.
func (s *BasegorchlangListener) EnterFragmentDirective(ctx *FragmentDirectiveContext) {}

// ExitFragmentDirective is called when production fragmentDirective is exited.
func (s *BasegorchlangListener) ExitFragmentDirective(ctx *FragmentDirectiveContext) {}

// EnterRegisterDirective is called when production registerDirective is entered.
func (s *BasegorchlangListener) EnterRegisterDirective(ctx *RegisterDirectiveContext) {}

// ExitRegisterDirective is called when production registerDirective is exited.
func (s *BasegorchlangListener) ExitRegisterDirective(ctx *RegisterDirectiveContext) {}

// EnterOnFinishStmt is called when production onFinishStmt is entered.
func (s *BasegorchlangListener) EnterOnFinishStmt(ctx *OnFinishStmtContext) {}

// ExitOnFinishStmt is called when production onFinishStmt is exited.
func (s *BasegorchlangListener) ExitOnFinishStmt(ctx *OnFinishStmtContext) {}

// EnterExedescStmt is called when production exedescStmt is entered.
func (s *BasegorchlangListener) EnterExedescStmt(ctx *ExedescStmtContext) {}

// ExitExedescStmt is called when production exedescStmt is exited.
func (s *BasegorchlangListener) ExitExedescStmt(ctx *ExedescStmtContext) {}

// EnterLeafSnippet is called when production leafSnippet is entered.
func (s *BasegorchlangListener) EnterLeafSnippet(ctx *LeafSnippetContext) {}

// ExitLeafSnippet is called when production leafSnippet is exited.
func (s *BasegorchlangListener) ExitLeafSnippet(ctx *LeafSnippetContext) {}

// EnterOperatorStmt is called when production operatorStmt is entered.
func (s *BasegorchlangListener) EnterOperatorStmt(ctx *OperatorStmtContext) {}

// ExitOperatorStmt is called when production operatorStmt is exited.
func (s *BasegorchlangListener) ExitOperatorStmt(ctx *OperatorStmtContext) {}

// EnterSerialStmt is called when production serialStmt is entered.
func (s *BasegorchlangListener) EnterSerialStmt(ctx *SerialStmtContext) {}

// ExitSerialStmt is called when production serialStmt is exited.
func (s *BasegorchlangListener) ExitSerialStmt(ctx *SerialStmtContext) {}

// EnterSerialBracketStmt is called when production serialBracketStmt is entered.
func (s *BasegorchlangListener) EnterSerialBracketStmt(ctx *SerialBracketStmtContext) {}

// ExitSerialBracketStmt is called when production serialBracketStmt is exited.
func (s *BasegorchlangListener) ExitSerialBracketStmt(ctx *SerialBracketStmtContext) {}

// EnterConcurrentStmt is called when production concurrentStmt is entered.
func (s *BasegorchlangListener) EnterConcurrentStmt(ctx *ConcurrentStmtContext) {}

// ExitConcurrentStmt is called when production concurrentStmt is exited.
func (s *BasegorchlangListener) ExitConcurrentStmt(ctx *ConcurrentStmtContext) {}

// EnterUnfoldDirective is called when production unfoldDirective is entered.
func (s *BasegorchlangListener) EnterUnfoldDirective(ctx *UnfoldDirectiveContext) {}

// ExitUnfoldDirective is called when production unfoldDirective is exited.
func (s *BasegorchlangListener) ExitUnfoldDirective(ctx *UnfoldDirectiveContext) {}

// EnterGoDirective is called when production goDirective is entered.
func (s *BasegorchlangListener) EnterGoDirective(ctx *GoDirectiveContext) {}

// ExitGoDirective is called when production goDirective is exited.
func (s *BasegorchlangListener) ExitGoDirective(ctx *GoDirectiveContext) {}

// EnterWaitDirective is called when production waitDirective is entered.
func (s *BasegorchlangListener) EnterWaitDirective(ctx *WaitDirectiveContext) {}

// ExitWaitDirective is called when production waitDirective is exited.
func (s *BasegorchlangListener) ExitWaitDirective(ctx *WaitDirectiveContext) {}

// EnterNoCheckMissDirective is called when production noCheckMissDirective is entered.
func (s *BasegorchlangListener) EnterNoCheckMissDirective(ctx *NoCheckMissDirectiveContext) {}

// ExitNoCheckMissDirective is called when production noCheckMissDirective is exited.
func (s *BasegorchlangListener) ExitNoCheckMissDirective(ctx *NoCheckMissDirectiveContext) {}

// EnterSkipDirective is called when production skipDirective is entered.
func (s *BasegorchlangListener) EnterSkipDirective(ctx *SkipDirectiveContext) {}

// ExitSkipDirective is called when production skipDirective is exited.
func (s *BasegorchlangListener) ExitSkipDirective(ctx *SkipDirectiveContext) {}

// EnterWrapStmt is called when production wrapStmt is entered.
func (s *BasegorchlangListener) EnterWrapStmt(ctx *WrapStmtContext) {}

// ExitWrapStmt is called when production wrapStmt is exited.
func (s *BasegorchlangListener) ExitWrapStmt(ctx *WrapStmtContext) {}

// EnterWrapBracketStmt is called when production wrapBracketStmt is entered.
func (s *BasegorchlangListener) EnterWrapBracketStmt(ctx *WrapBracketStmtContext) {}

// ExitWrapBracketStmt is called when production wrapBracketStmt is exited.
func (s *BasegorchlangListener) ExitWrapBracketStmt(ctx *WrapBracketStmtContext) {}

// EnterSwitchDirective is called when production switchDirective is entered.
func (s *BasegorchlangListener) EnterSwitchDirective(ctx *SwitchDirectiveContext) {}

// ExitSwitchDirective is called when production switchDirective is exited.
func (s *BasegorchlangListener) ExitSwitchDirective(ctx *SwitchDirectiveContext) {}

// EnterSwitchCaseDirective is called when production switchCaseDirective is entered.
func (s *BasegorchlangListener) EnterSwitchCaseDirective(ctx *SwitchCaseDirectiveContext) {}

// ExitSwitchCaseDirective is called when production switchCaseDirective is exited.
func (s *BasegorchlangListener) ExitSwitchCaseDirective(ctx *SwitchCaseDirectiveContext) {}

// EnterRegisterOperatorDirective is called when production registerOperatorDirective is entered.
func (s *BasegorchlangListener) EnterRegisterOperatorDirective(ctx *RegisterOperatorDirectiveContext) {
}

// ExitRegisterOperatorDirective is called when production registerOperatorDirective is exited.
func (s *BasegorchlangListener) ExitRegisterOperatorDirective(ctx *RegisterOperatorDirectiveContext) {
}

// EnterArgStmt is called when production argStmt is entered.
func (s *BasegorchlangListener) EnterArgStmt(ctx *ArgStmtContext) {}

// ExitArgStmt is called when production argStmt is exited.
func (s *BasegorchlangListener) ExitArgStmt(ctx *ArgStmtContext) {}

// EnterArgsStmt is called when production argsStmt is entered.
func (s *BasegorchlangListener) EnterArgsStmt(ctx *ArgsStmtContext) {}

// ExitArgsStmt is called when production argsStmt is exited.
func (s *BasegorchlangListener) ExitArgsStmt(ctx *ArgsStmtContext) {}

// EnterIgnoreError is called when production ignoreError is entered.
func (s *BasegorchlangListener) EnterIgnoreError(ctx *IgnoreErrorContext) {}

// ExitIgnoreError is called when production ignoreError is exited.
func (s *BasegorchlangListener) ExitIgnoreError(ctx *IgnoreErrorContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BasegorchlangListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BasegorchlangListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BasegorchlangListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BasegorchlangListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BasegorchlangListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BasegorchlangListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterDurationLiteral is called when production durationLiteral is entered.
func (s *BasegorchlangListener) EnterDurationLiteral(ctx *DurationLiteralContext) {}

// ExitDurationLiteral is called when production durationLiteral is exited.
func (s *BasegorchlangListener) ExitDurationLiteral(ctx *DurationLiteralContext) {}
