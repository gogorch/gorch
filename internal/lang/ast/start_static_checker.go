package ast

import "fmt"

type startStaticChecker struct {
	start            *StartDirective
	visitedFragments map[string]struct{}
}

func newStartStaticChecker(start *StartDirective) *startStaticChecker {
	return &startStaticChecker{
		start:            start,
		visitedFragments: make(map[string]struct{}),
	}
}

type startCheckCallback struct {
	onOperator func(*OperatorStmt) error
	onSerial   func(*SerialStmt) error
	onLoop     func(*LoopDirective) error
	onRetry    func(*RetryDirective) error
}

func (c *startStaticChecker) checkOperatorMiss(allOperators *Operators) error {
	return c.walkStart(startCheckCallback{
		onOperator: func(op *OperatorStmt) error {
			if allOperators == nil || allOperators.Names == nil {
				return fmt.Errorf("in start %q, operator %q not register", c.start.Name, op.Name)
			}
			if _, ok := allOperators.Names[op.Name]; !ok {
				return fmt.Errorf("in start %q, operator %q not register", c.start.Name, op.Name)
			}
			return nil
		},
	})
}

func (c *startStaticChecker) checkRequiredParams() error {
	return c.walkStart(startCheckCallback{
		onLoop: func(loop *LoopDirective) error {
			if loop.Args == nil || loop.Args.ArgMap == nil {
				return fmt.Errorf("START %q: LOOP missing required argument MAX_ITER", c.start.Name)
			}

			if arg, ok := loop.Args.ArgMap["MAX_ITER"]; ok && len(arg.IntegerValue) > 0 {
				if arg.IntegerValue[0] > 0 {
					return nil
				}
			}

			if arg, ok := loop.Args.ArgMap["maxIter"]; ok && len(arg.IntegerValue) > 0 {
				if arg.IntegerValue[0] > 0 {
					return nil
				}
			}

			return fmt.Errorf("START %q: LOOP missing required argument MAX_ITER", c.start.Name)
		},
		onRetry: func(retry *RetryDirective) error {
			if retry.Args == nil || retry.Args.ArgMap == nil {
				return fmt.Errorf("START %q: RETRY missing required argument MAX_TIMES", c.start.Name)
			}

			if arg, ok := retry.Args.ArgMap["MAX_TIMES"]; ok && len(arg.IntegerValue) > 0 {
				if arg.IntegerValue[0] > 0 {
					return nil
				}
			}

			if arg, ok := retry.Args.ArgMap["maxTimes"]; ok && len(arg.IntegerValue) > 0 {
				if arg.IntegerValue[0] > 0 {
					return nil
				}
			}

			return fmt.Errorf("START %q: RETRY missing required argument MAX_TIMES", c.start.Name)
		},
	})
}

func (c *startStaticChecker) checkUnreachableBranch() error {
	return c.walkStart(startCheckCallback{
		onSerial: func(serial *SerialStmt) error {
			for idx, leaf := range serial.LeafSnippetList {
				if c.leafAlwaysTerminate(leaf.Leaf) && idx < len(serial.LeafSnippetList)-1 {
					return fmt.Errorf("START %q: unreachable branch after BREAK()", c.start.Name)
				}
			}
			return nil
		},
	})
}

func (c *startStaticChecker) walkStart(callback startCheckCallback) error {
	if c.start.OnFinishStmt != nil && c.start.OnFinishStmt.ExedescStmt != nil {
		if err := c.walkExedesc(c.start.OnFinishStmt.ExedescStmt.ExedescStmt, callback); err != nil {
			return err
		}
	}

	if c.start.ExedescStmt == nil {
		return nil
	}

	return c.walkExedesc(c.start.ExedescStmt.ExedescStmt, callback)
}

func (c *startStaticChecker) walkExedesc(exedesc ExedescI, callback startCheckCallback) error {
	switch stmt := exedesc.(type) {
	case *OperatorStmt:
		if callback.onOperator != nil {
			return callback.onOperator(stmt)
		}
		return nil
	case *SerialStmt:
		if callback.onSerial != nil {
			if err := callback.onSerial(stmt); err != nil {
				return err
			}
		}
		for _, leaf := range stmt.LeafSnippetList {
			if err := c.walkLeafSnippet(leaf.Leaf, callback); err != nil {
				return err
			}
		}
		return nil
	case *ConcurrentStmt:
		for _, leaf := range stmt.LeafSnippetList {
			if err := c.walkLeafSnippet(leaf.Leaf, callback); err != nil {
				return err
			}
		}
		return nil
	case *WrapStmt:
		for _, op := range stmt.WrapOpStmtList {
			if callback.onOperator != nil {
				if err := callback.onOperator(op); err != nil {
					return err
				}
			}
		}
		if stmt.LeafSnippet == nil {
			return nil
		}
		return c.walkLeafSnippet(stmt.LeafSnippet.Leaf, callback)
	case *SwitchDirective:
		if stmt.OperatorStmt != nil && callback.onOperator != nil {
			if err := callback.onOperator(stmt.OperatorStmt); err != nil {
				return err
			}
		}
		for _, one := range stmt.SwitchCaseList {
			if one.LeafSnippet == nil {
				continue
			}
			if err := c.walkLeafSnippet(one.LeafSnippet.Leaf, callback); err != nil {
				return err
			}
		}
		return nil
	case *UnfoldDirective:
		return c.walkUnfold(stmt, callback)
	case *GoDirective:
		if stmt.ExedescStmt == nil {
			return nil
		}
		return c.walkExedesc(stmt.ExedescStmt.ExedescStmt, callback)
	case *LoopDirective:
		if callback.onLoop != nil {
			if err := callback.onLoop(stmt); err != nil {
				return err
			}
		}
		if stmt.ExedescStmt == nil {
			return nil
		}
		return c.walkExedesc(stmt.ExedescStmt.ExedescStmt, callback)
	case *RetryDirective:
		if callback.onRetry != nil {
			if err := callback.onRetry(stmt); err != nil {
				return err
			}
		}
		if stmt.ExedescStmt == nil {
			return nil
		}
		return c.walkExedesc(stmt.ExedescStmt.ExedescStmt, callback)
	case *TraceDirective:
		if stmt.ExedescStmt == nil {
			return nil
		}
		return c.walkExedesc(stmt.ExedescStmt.ExedescStmt, callback)
	case *UntilDirective:
		if stmt.OperatorStmt != nil && callback.onOperator != nil {
			return callback.onOperator(stmt.OperatorStmt)
		}
		return nil
	case *BreakDirective:
		return nil
	default:
		return nil
	}
}

func (c *startStaticChecker) walkLeafSnippet(leaf LeafSnippetI, callback startCheckCallback) error {
	switch stmt := leaf.(type) {
	case *OperatorStmt:
		if callback.onOperator != nil {
			return callback.onOperator(stmt)
		}
		return nil
	case *ConcurrentStmt:
		return c.walkExedesc(stmt, callback)
	case *SerialBracketStmt:
		if stmt.SerialStmt == nil {
			return nil
		}
		return c.walkExedesc(stmt.SerialStmt, callback)
	case *UnfoldDirective:
		return c.walkUnfold(stmt, callback)
	case *WrapBracketStmt:
		if stmt.WrapStmt == nil {
			return nil
		}
		return c.walkExedesc(stmt.WrapStmt, callback)
	case *GoDirective:
		return c.walkExedesc(stmt, callback)
	case *SwitchDirective:
		return c.walkExedesc(stmt, callback)
	case *LoopDirective:
		return c.walkExedesc(stmt, callback)
	case *RetryDirective:
		return c.walkExedesc(stmt, callback)
	case *TraceDirective:
		return c.walkExedesc(stmt, callback)
	case *UntilDirective:
		return c.walkExedesc(stmt, callback)
	case *BreakDirective:
		return nil
	case *SkipDirective:
		if stmt.OperatorStmt != nil && callback.onOperator != nil {
			return callback.onOperator(stmt.OperatorStmt)
		}
		return nil
	default:
		return nil
	}
}

func (c *startStaticChecker) walkUnfold(stmt *UnfoldDirective, callback startCheckCallback) error {
	if stmt.FragmentDirective == nil || stmt.FragmentDirective.ExedescStmt == nil {
		return nil
	}

	if _, ok := c.visitedFragments[stmt.FragmentDirective.Name]; ok {
		return nil
	}
	c.visitedFragments[stmt.FragmentDirective.Name] = struct{}{}

	return c.walkExedesc(stmt.FragmentDirective.ExedescStmt.ExedescStmt, callback)
}

func (sds *StartDirectives) checkRequiredParams() error {
	for _, name := range sds.Sort {
		sd := sds.Map[name]
		checker := newStartStaticChecker(sd)
		if err := checker.checkRequiredParams(); err != nil {
			return err
		}
	}
	return nil
}

func (sds *StartDirectives) checkUnreachableBranch() error {
	for _, name := range sds.Sort {
		sd := sds.Map[name]
		checker := newStartStaticChecker(sd)
		if err := checker.checkUnreachableBranch(); err != nil {
			return err
		}
	}
	return nil
}

func (c *startStaticChecker) leafAlwaysTerminate(leaf LeafSnippetI) bool {
	return c.leafAlwaysTerminateWithSeen(leaf, map[string]struct{}{})
}

func (c *startStaticChecker) leafAlwaysTerminateWithSeen(leaf LeafSnippetI, seen map[string]struct{}) bool {
	switch stmt := leaf.(type) {
	case *BreakDirective:
		return true
	case *SerialBracketStmt:
		if stmt.SerialStmt == nil {
			return false
		}
		return c.serialAlwaysTerminateWithSeen(stmt.SerialStmt, seen)
	case *UnfoldDirective:
		return c.unfoldAlwaysTerminateWithSeen(stmt, seen)
	case *SwitchDirective:
		return c.switchAlwaysTerminateWithSeen(stmt, seen)
	case *LoopDirective:
		if stmt.ExedescStmt == nil {
			return false
		}
		return c.exedescAlwaysTerminateWithSeen(stmt.ExedescStmt.ExedescStmt, seen)
	case *RetryDirective:
		if stmt.ExedescStmt == nil {
			return false
		}
		return c.exedescAlwaysTerminateWithSeen(stmt.ExedescStmt.ExedescStmt, seen)
	case *TraceDirective:
		if stmt.ExedescStmt == nil {
			return false
		}
		return c.exedescAlwaysTerminateWithSeen(stmt.ExedescStmt.ExedescStmt, seen)
	default:
		return false
	}
}

func (c *startStaticChecker) serialAlwaysTerminateWithSeen(serial *SerialStmt, seen map[string]struct{}) bool {
	for _, leaf := range serial.LeafSnippetList {
		if c.leafAlwaysTerminateWithSeen(leaf.Leaf, seen) {
			return true
		}
	}
	return false
}

func (c *startStaticChecker) exedescAlwaysTerminateWithSeen(exedesc ExedescI, seen map[string]struct{}) bool {
	switch stmt := exedesc.(type) {
	case *BreakDirective:
		return true
	case *SerialStmt:
		return c.serialAlwaysTerminateWithSeen(stmt, seen)
	case *UnfoldDirective:
		return c.unfoldAlwaysTerminateWithSeen(stmt, seen)
	case *SwitchDirective:
		return c.switchAlwaysTerminateWithSeen(stmt, seen)
	case *LoopDirective:
		if stmt.ExedescStmt == nil {
			return false
		}
		return c.exedescAlwaysTerminateWithSeen(stmt.ExedescStmt.ExedescStmt, seen)
	case *RetryDirective:
		if stmt.ExedescStmt == nil {
			return false
		}
		return c.exedescAlwaysTerminateWithSeen(stmt.ExedescStmt.ExedescStmt, seen)
	case *TraceDirective:
		if stmt.ExedescStmt == nil {
			return false
		}
		return c.exedescAlwaysTerminateWithSeen(stmt.ExedescStmt.ExedescStmt, seen)
	default:
		return false
	}
}

func (c *startStaticChecker) unfoldAlwaysTerminateWithSeen(unfold *UnfoldDirective, seen map[string]struct{}) bool {
	if unfold.FragmentDirective == nil || unfold.FragmentDirective.ExedescStmt == nil {
		return false
	}

	name := unfold.FragmentDirective.Name
	if _, ok := seen[name]; ok {
		return false
	}

	nextSeen := cloneSeen(seen)
	nextSeen[name] = struct{}{}
	return c.exedescAlwaysTerminateWithSeen(unfold.FragmentDirective.ExedescStmt.ExedescStmt, nextSeen)
}

func (c *startStaticChecker) switchAlwaysTerminateWithSeen(sw *SwitchDirective, seen map[string]struct{}) bool {
	if len(sw.SwitchCaseList) == 0 {
		return false
	}

	for _, cs := range sw.SwitchCaseList {
		if cs.LeafSnippet == nil {
			return false
		}
		if !c.leafAlwaysTerminateWithSeen(cs.LeafSnippet.Leaf, cloneSeen(seen)) {
			return false
		}
	}
	return true
}

func cloneSeen(src map[string]struct{}) map[string]struct{} {
	dst := make(map[string]struct{}, len(src))
	for k := range src {
		dst[k] = struct{}{}
	}
	return dst
}
