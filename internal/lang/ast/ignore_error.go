package ast

import "github.com/gogorch/gorch/internal/lang/iantlr/alr"

type IgnoreErrorHolder interface {
	AcceptIgnoreErrorFlag() error
}

// EnterIgnoreError is called when production ignoreError is entered.
func (s *GorchlangParserListener) EnterIgnoreError(ctx *alr.IgnoreErrorContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
}

// ExitIgnoreError is called when production ignoreError is exited.
func (s *GorchlangParserListener) ExitIgnoreError(ctx *alr.IgnoreErrorContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	holder := s.Stack.Peek().(IgnoreErrorHolder)
	err := holder.AcceptIgnoreErrorFlag()
	if err != nil {
		s.AddError(err)
	}
}
