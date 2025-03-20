// Code generated from gorchlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package alr // gorchlang
import "github.com/antlr4-go/antlr/v4"

// gorchlangListener is a complete listener for a parse tree produced by gorchlangParser.
type gorchlangListener interface {
	antlr.ParseTreeListener

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterStartDirective is called when entering the startDirective production.
	EnterStartDirective(c *StartDirectiveContext)

	// EnterFragmentDirective is called when entering the fragmentDirective production.
	EnterFragmentDirective(c *FragmentDirectiveContext)

	// EnterRegisterDirective is called when entering the registerDirective production.
	EnterRegisterDirective(c *RegisterDirectiveContext)

	// EnterOnFinishStmt is called when entering the onFinishStmt production.
	EnterOnFinishStmt(c *OnFinishStmtContext)

	// EnterExedescStmt is called when entering the exedescStmt production.
	EnterExedescStmt(c *ExedescStmtContext)

	// EnterLeafSnippet is called when entering the leafSnippet production.
	EnterLeafSnippet(c *LeafSnippetContext)

	// EnterOperatorStmt is called when entering the operatorStmt production.
	EnterOperatorStmt(c *OperatorStmtContext)

	// EnterSerialStmt is called when entering the serialStmt production.
	EnterSerialStmt(c *SerialStmtContext)

	// EnterSerialBracketStmt is called when entering the serialBracketStmt production.
	EnterSerialBracketStmt(c *SerialBracketStmtContext)

	// EnterConcurrentStmt is called when entering the concurrentStmt production.
	EnterConcurrentStmt(c *ConcurrentStmtContext)

	// EnterUnfoldDirective is called when entering the unfoldDirective production.
	EnterUnfoldDirective(c *UnfoldDirectiveContext)

	// EnterGoDirective is called when entering the goDirective production.
	EnterGoDirective(c *GoDirectiveContext)

	// EnterWaitDirective is called when entering the waitDirective production.
	EnterWaitDirective(c *WaitDirectiveContext)

	// EnterNoCheckMissDirective is called when entering the noCheckMissDirective production.
	EnterNoCheckMissDirective(c *NoCheckMissDirectiveContext)

	// EnterSkipDirective is called when entering the skipDirective production.
	EnterSkipDirective(c *SkipDirectiveContext)

	// EnterWrapStmt is called when entering the wrapStmt production.
	EnterWrapStmt(c *WrapStmtContext)

	// EnterWrapBracketStmt is called when entering the wrapBracketStmt production.
	EnterWrapBracketStmt(c *WrapBracketStmtContext)

	// EnterSwitchDirective is called when entering the switchDirective production.
	EnterSwitchDirective(c *SwitchDirectiveContext)

	// EnterSwitchCaseDirective is called when entering the switchCaseDirective production.
	EnterSwitchCaseDirective(c *SwitchCaseDirectiveContext)

	// EnterRegisterOperatorDirective is called when entering the registerOperatorDirective production.
	EnterRegisterOperatorDirective(c *RegisterOperatorDirectiveContext)

	// EnterArgStmt is called when entering the argStmt production.
	EnterArgStmt(c *ArgStmtContext)

	// EnterArgsStmt is called when entering the argsStmt production.
	EnterArgsStmt(c *ArgsStmtContext)

	// EnterIgnoreError is called when entering the ignoreError production.
	EnterIgnoreError(c *IgnoreErrorContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterDurationLiteral is called when entering the durationLiteral production.
	EnterDurationLiteral(c *DurationLiteralContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitStartDirective is called when exiting the startDirective production.
	ExitStartDirective(c *StartDirectiveContext)

	// ExitFragmentDirective is called when exiting the fragmentDirective production.
	ExitFragmentDirective(c *FragmentDirectiveContext)

	// ExitRegisterDirective is called when exiting the registerDirective production.
	ExitRegisterDirective(c *RegisterDirectiveContext)

	// ExitOnFinishStmt is called when exiting the onFinishStmt production.
	ExitOnFinishStmt(c *OnFinishStmtContext)

	// ExitExedescStmt is called when exiting the exedescStmt production.
	ExitExedescStmt(c *ExedescStmtContext)

	// ExitLeafSnippet is called when exiting the leafSnippet production.
	ExitLeafSnippet(c *LeafSnippetContext)

	// ExitOperatorStmt is called when exiting the operatorStmt production.
	ExitOperatorStmt(c *OperatorStmtContext)

	// ExitSerialStmt is called when exiting the serialStmt production.
	ExitSerialStmt(c *SerialStmtContext)

	// ExitSerialBracketStmt is called when exiting the serialBracketStmt production.
	ExitSerialBracketStmt(c *SerialBracketStmtContext)

	// ExitConcurrentStmt is called when exiting the concurrentStmt production.
	ExitConcurrentStmt(c *ConcurrentStmtContext)

	// ExitUnfoldDirective is called when exiting the unfoldDirective production.
	ExitUnfoldDirective(c *UnfoldDirectiveContext)

	// ExitGoDirective is called when exiting the goDirective production.
	ExitGoDirective(c *GoDirectiveContext)

	// ExitWaitDirective is called when exiting the waitDirective production.
	ExitWaitDirective(c *WaitDirectiveContext)

	// ExitNoCheckMissDirective is called when exiting the noCheckMissDirective production.
	ExitNoCheckMissDirective(c *NoCheckMissDirectiveContext)

	// ExitSkipDirective is called when exiting the skipDirective production.
	ExitSkipDirective(c *SkipDirectiveContext)

	// ExitWrapStmt is called when exiting the wrapStmt production.
	ExitWrapStmt(c *WrapStmtContext)

	// ExitWrapBracketStmt is called when exiting the wrapBracketStmt production.
	ExitWrapBracketStmt(c *WrapBracketStmtContext)

	// ExitSwitchDirective is called when exiting the switchDirective production.
	ExitSwitchDirective(c *SwitchDirectiveContext)

	// ExitSwitchCaseDirective is called when exiting the switchCaseDirective production.
	ExitSwitchCaseDirective(c *SwitchCaseDirectiveContext)

	// ExitRegisterOperatorDirective is called when exiting the registerOperatorDirective production.
	ExitRegisterOperatorDirective(c *RegisterOperatorDirectiveContext)

	// ExitArgStmt is called when exiting the argStmt production.
	ExitArgStmt(c *ArgStmtContext)

	// ExitArgsStmt is called when exiting the argsStmt production.
	ExitArgsStmt(c *ArgsStmtContext)

	// ExitIgnoreError is called when exiting the ignoreError production.
	ExitIgnoreError(c *IgnoreErrorContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitDurationLiteral is called when exiting the durationLiteral production.
	ExitDurationLiteral(c *DurationLiteralContext)
}
