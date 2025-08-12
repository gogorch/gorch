package engine

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gogorch/gorch/internal/lang/ast"
	"github.com/gogorch/gorch/mlog"
)

func createProcessor(stmt any) PrepareProcessor {
	switch astStmt := stmt.(type) {
	case *ast.StartDirective:
		return newStarterProcessor(astStmt)
	case *ast.FragmentDirective:
		return newFragmentProcessor(astStmt)
	case *ast.SerialStmt:
		return newSerialProcessor(astStmt)
	case *ast.SerialBracketStmt:
		return newSerialProcessor(astStmt.SerialStmt)
	case *ast.ConcurrentStmt:
		return newConcurrentprocess(astStmt)
	case *ast.WrapStmt:
		return newWrapProcessor(astStmt)
	case *ast.WrapBracketStmt:
		return newWrapProcessor(astStmt.WrapStmt)
	case *ast.SwitchDirective:
		return newSwitchProcess(astStmt)
	case *ast.SwitchCaseDirective:
		return newSwitchCaseProcess(astStmt)
	case *ast.GoDirective:
		return newGoProcessor(astStmt)
	case *ast.WaitDirective:
		return newWaitProcessor(astStmt)
	case *ast.OnFinishStmt:
		return newOnFinishProcessor(astStmt)
	case *ast.SkipDirective:
		return newSkipsProcessor(astStmt)
	case *ast.UnfoldDirective:
		return newUnfoldProcessor(astStmt)
	case *ast.OperatorStmt:
		return newProcessor(astStmt)
	default:
		panic(fmt.Sprintf("not support stmt type %T", astStmt))
	}
}

type Processor interface {
	Execute(*Context) error
}

type Preprocessor interface {
	Prepare() error
}

type PresetProcessor interface {
	Execute(*Context) error
}

type PrepareProcessor interface {
	Processor
	Preprocessor
}

type starterProcessor struct {
	*ast.StartDirective

	args *args

	exedescProc  PrepareProcessor
	onFinishProc PrepareProcessor
}

func newStarterProcessor(stmt *ast.StartDirective) *starterProcessor {
	sp := new(starterProcessor)
	sp.StartDirective = stmt
	return sp
}

func (s *starterProcessor) Prepare() error {
	s.args = fromArgsStmt(s.StartDirective.Args)
	s.exedescProc = createProcessor(s.ExedescStmt.ExedescStmt)
	if err := s.exedescProc.Prepare(); err != nil {
		return err
	}
	if s.OnFinishStmt != nil {
		s.onFinishProc = createProcessor(s.OnFinishStmt.ExedescStmt.ExedescStmt)
		if err := s.onFinishProc.Prepare(); err != nil {
			return err
		}
	}

	return nil
}

func (s *starterProcessor) Execute(ctx *Context) (err error) {
	newCtx := ctx.clone()

	defer func() {
		if err := recover(); err != nil {
			fields := RecoverPanic(err)
			fields = append(fields,
				mlog.String("panic", fmt.Sprintf("%v", err)))
			ctx.Logger.Error("engineExecutePanic", fields...)
		}
	}()

	defer func() {
		if s.onFinishProc != nil {
			if er1 := s.onFinishProc.Execute(newCtx); er1 != nil {
				ctx.Logger.Error("onFinishExecuteError", mlog.Error("error", er1))
			}
		}

		newCtx.putPool()
	}()

	err = s.exedescProc.Execute(newCtx)

	newCtx.waiter.Wait()
	return
}

type serialProcessor struct {
	*ast.SerialStmt
	processors []PrepareProcessor
}

func newSerialProcessor(stmt *ast.SerialStmt) *serialProcessor {
	sp := new(serialProcessor)
	sp.SerialStmt = stmt
	return sp
}

func (sp *serialProcessor) Prepare() error {
	sp.processors = make([]PrepareProcessor, len(sp.SerialStmt.LeafSnippetList))

	for idx, leaf := range sp.SerialStmt.LeafSnippetList {
		t := createProcessor(leaf.Leaf)
		sp.processors[idx] = t
		if e := t.Prepare(); e != nil {
			return errors.Join(e, errors.New("prepare serial processor error1"))
		}
	}

	return nil
}

func (sp *serialProcessor) Execute(ctx *Context) (err error) {
	newCtx := ctx.clone()
	defer newCtx.putPool()

	for _, processor := range sp.processors {
		if ctx.Exited() {
			continue
		}

		_ = processor.Execute(newCtx)

		if _, ok := processor.(*skipProcessor); ok && newCtx.skipSerial {
			break
		}
	}

	return nil
}

type skipProcessor struct {
	*ast.SkipDirective
	processor PrepareProcessor
}

func newSkipsProcessor(stmt *ast.SkipDirective) *skipProcessor {
	sp := new(skipProcessor)
	sp.SkipDirective = stmt
	return sp
}

func (sp *skipProcessor) Prepare() error {
	sp.processor = createProcessor(sp.OperatorStmt)
	return sp.processor.Prepare()
}

func (sp *skipProcessor) Execute(ctx *Context) error {
	newCtx := ctx.clone()
	defer newCtx.putPool()

	ret := sp.processor.Execute(newCtx)
	ctx.skipSerial = newCtx.skipSerial
	return ret
}

type concurrentProcessor struct {
	*ast.ConcurrentStmt

	processors []Processor
}

func newConcurrentprocess(stmt *ast.ConcurrentStmt) *concurrentProcessor {
	cp := new(concurrentProcessor)
	cp.ConcurrentStmt = stmt
	return cp
}

func (cp *concurrentProcessor) Prepare() error {
	cp.processors = make([]Processor, len(cp.LeafSnippetList))

	for idx, leaf := range cp.LeafSnippetList {
		t := createProcessor(leaf.Leaf)
		cp.processors[idx] = t
		if e := t.Prepare(); e != nil {
			return errors.Join(e, errors.New("prepare concurrent processor error"))
		}
	}

	return nil
}

func (cp *concurrentProcessor) Execute(ctx *Context) error {
	g := &sync.WaitGroup{}
	newCtx := ctx.clone()
	defer newCtx.putPool()

	for _, processor := range cp.processors {
		g.Add(1)

		goRoutinePool.Go(
			func(nctx *Context, t Processor) func() {
				return func() {
					defer g.Done()
					_ = t.Execute(nctx)
				}
			}(newCtx, processor))
	}
	g.Wait()
	return nil
}

type goProcessor struct {
	*ast.GoDirective
	processor PrepareProcessor
}

func newGoProcessor(stmt *ast.GoDirective) *goProcessor {
	gp := new(goProcessor)
	gp.GoDirective = stmt
	return gp
}

func (gp *goProcessor) Prepare() error {
	gp.processor = createProcessor(gp.ExedescStmt.ExedescStmt)
	return gp.processor.Prepare()
}

func (gp *goProcessor) Execute(ctx *Context) error {
	b := ctx.getRoutine(gp.GoName)
	newCtx := ctx.clone()

	ctx.waiter.Add(1)
	goRoutinePool.Go(func(nctx *Context, r *routine) func() {
		return func() {
			defer func() {
				nctx.waiter.Done()
				nctx.putPool()
			}()

			if err := r.start(func() {
				st := time.Now()
				_ = gp.processor.Execute(nctx)
				co := time.Since(st)
				nctx.Logger.AddInfo(mlog.Duration("routine_"+r.name, co))
			}); err != nil {
				nctx.Logger.Error("routine execute error", mlog.Error("error", err))
			}
		}
	}(newCtx, b))

	return nil
}

type waitProcessor struct {
	*ast.WaitDirective
}

func newWaitProcessor(stmt *ast.WaitDirective) *waitProcessor {
	wp := new(waitProcessor)
	wp.WaitDirective = stmt
	return wp
}

func (wp *waitProcessor) Prepare() error {
	return nil
}

func (wp *waitProcessor) Execute(ctx *Context) error {
	r := ctx.container.getRoutine(wp.GoName)

	timeout := wp.Timeout
	fromStart := false
	if wp.TotalTimeout > 0 {
		fromStart = true
		timeout = wp.TotalTimeout
	}

	err := r.wait(ctx.interrupt, timeout, fromStart, wp.NotCheckStart)
	if wp.IgnoreError && err != nil {
		ctx.Logger.Warn("wait routine error", mlog.Error("error", err), mlog.String("routine", wp.GoName))
		err = nil
	}

	return err
}

type onFinishProcessor struct {
	*ast.OnFinishStmt
	processor PrepareProcessor
}

func newOnFinishProcessor(stmt *ast.OnFinishStmt) *onFinishProcessor {
	return &onFinishProcessor{OnFinishStmt: stmt}
}

func (ofp *onFinishProcessor) Prepare() error {
	ofp.processor = createProcessor(ofp.ExedescStmt.ExedescStmt)
	return ofp.processor.Prepare()
}

func (ofp *onFinishProcessor) Execute(ctx *Context) error {
	newCtx := ctx.clone()
	defer newCtx.putPool()
	return ofp.processor.Execute(newCtx)
}

type operatorProcessor struct {
	*ast.OperatorStmt

	args *args

	operatorFactory OperatorFactory
	operatorSeq     string
	waitList        []Processor
	status          OperatorStates
}

func newProcessor(stmt *ast.OperatorStmt) *operatorProcessor {
	ot := new(operatorProcessor)
	ot.OperatorStmt = stmt
	return ot
}

var (
	operatorPrepareDup = map[Preprocessor]bool{}
)

func (op *operatorProcessor) Prepare() (err error) {
	op.args = fromArgsStmt(op.ArgsStmt)

	of, has := getOperatorFactory(op.OperatorStmt.Name)
	if !has {
		err = fmt.Errorf("operator %s not register, %s", op.Name, op.OperatorStmt)
		return
	}

	op.operatorSeq = of.Seq()
	op.operatorFactory = of

	nilo := of.NilProcessor()
	if preprocessor, ok := nilo.(Preprocessor); ok {
		if _, has := operatorPrepareDup[preprocessor]; !has {
			if err = preprocessor.Prepare(); err != nil {
				return err
			}
			operatorPrepareDup[preprocessor] = true
		}
	}

	if os, ok := nilo.(HasOperatorState); ok {
		op.status = os.OperatorState()
	}

	if len(op.OperatorStmt.WaitList) > 0 {
		op.waitList = make([]Processor, len(op.OperatorStmt.WaitList))
		for idx, waitStmt := range op.OperatorStmt.WaitList {
			wp := newWaitProcessor(waitStmt)
			op.waitList[idx] = wp
			if err := wp.Prepare(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (op *operatorProcessor) Execute(ctx *Context) (err error) {
	endRecord := ctx.recorder.RecordOperator(op.Name, op.operatorSeq)
	var statusCode int

	defer func() {
		endRecord(statusCode)

		if er := recover(); er != nil {
			err = fmt.Errorf("operator %s execute panic", op.Name)
			fields := RecoverPanic(er)
			fields = append(fields, mlog.String("panicOperator", op.Name))
			ctx.Logger.Error("operator execute panic", fields...)
		}

		if err != nil {
			if op.OperatorStmt.IgnoreError {
				ctx.Logger.Warn("operator execute error, ignore", mlog.String("operator", op.Name), mlog.Error("error", err))
				err = nil
			} else {
				ctx.Logger.Error("operator execute error", mlog.String("operator", op.Name), mlog.Error("error", err))
				ctx.Exit(err)
			}
		}
	}()

	if ctx.Exited() {
		return nil
	}

	newCtx := ctx.clone()
	defer func() {
		ctx.skipSerial = newCtx.skipSerial
		newCtx.putPool()
	}()
	newCtx.args = op.args
	newCtx.switchCase = ctx.switchCase
	newCtx.nextWrap = ctx.nextWrap

	for _, wait := range op.waitList {
		if werr := wait.Execute(newCtx); werr != nil {
			err = werr
			return
		}
	}

	if ctx.Exited() {
		return nil
	}

	of := op.operatorFactory
	if to := op.args.Arg("timeout").Duration(); to > 0 {
		timer := time.NewTimer(to)
		ch := make(chan error, 1)
		goRoutinePool.Go(func() {
			ch <- of.Execute(newCtx)
		})
		select {
		case err = <-ch:
		case <-timer.C:
			err = errOperatorExecuteTimeout
		case err = <-ctx.Wait():
		}
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
	} else {
		err = of.Execute(newCtx)
	}

	if status, ok := err.(*status); ok {
		statusCode = status.code
		if status.fatal {
			if op.OperatorStmt.IgnoreError {
				ctx.Logger.Info("operator execute return fatal error, but ignore", mlog.String("operator", op.Name), mlog.Error("error", err))
				err = nil
				return
			}
			ctx.Exit(status)
			return
		}
		err = nil
	}

	return nil
}

type switchProcessor struct {
	*ast.SwitchDirective

	processor   PrepareProcessor
	switchCases map[string]PrepareProcessor
}

func newSwitchProcess(stmt *ast.SwitchDirective) *switchProcessor {
	st := new(switchProcessor)
	st.SwitchDirective = stmt
	return st
}

func (wp *switchProcessor) Prepare() error {
	wp.processor = createProcessor(wp.OperatorStmt)
	if err := wp.processor.Prepare(); err != nil {
		return err
	}

	wp.switchCases = make(map[string]PrepareProcessor, len(wp.SwitchCaseList))
	for _, sc := range wp.SwitchCaseList {
		wp.switchCases[sc.CaseName] = createProcessor(sc)
		if e := wp.switchCases[sc.CaseName].Prepare(); e != nil {
			return errors.Join(e, errPrepareSwitchCase)
		}
	}
	return nil
}

func (sp *switchProcessor) Execute(ctx *Context) error {
	newCtx0 := ctx.clone()
	defer newCtx0.putPool()

	choicesCases := make([]Processor, 0, len(sp.switchCases))
	checkCaseDup := make(map[string]struct{}, len(sp.switchCases))
	newCtx0.switchCase = func(s ...string) error {
		for _, c := range s {
			if _, ok := checkCaseDup[c]; ok {
				return fmt.Errorf("duplicate switch case %q", c)
			} else {
				checkCaseDup[c] = struct{}{}
			}
			if v, ok := sp.switchCases[c]; ok {
				choicesCases = append(choicesCases, v)
			} else {
				return fmt.Errorf("not found switch case %q", c)
			}
		}

		return nil
	}

	if err := sp.processor.Execute(newCtx0); err != nil {
		return err
	}

	if ctx.Exited() {
		return nil
	}

	if len(choicesCases) == 0 {
		return fmt.Errorf("switch %q no matching switch case found", sp.SwitchDirective.OperatorStmt.Name)
	} else if len(choicesCases) == 1 {
		newCtx1 := ctx.clone()
		defer newCtx1.putPool()
		return choicesCases[0].Execute(newCtx1)
	}

	wait := &sync.WaitGroup{}

	newCtx2 := ctx.clone()
	defer newCtx2.putPool()

	for _, cp := range choicesCases {
		if ctx.Exited() {
			break
		}

		wait.Add(1)
		goRoutinePool.Go(func(nctx *Context, p Processor) func() {
			return func() {
				defer wait.Done()

				_ = p.Execute(nctx)
			}
		}(newCtx2, cp))
	}

	wait.Wait()
	return nil
}

type switchCaseProcessor struct {
	*ast.SwitchCaseDirective

	processor PrepareProcessor
}

func newSwitchCaseProcess(stmt *ast.SwitchCaseDirective) *switchCaseProcessor {
	scp := new(switchCaseProcessor)
	scp.SwitchCaseDirective = stmt
	return scp
}

func (scp *switchCaseProcessor) Prepare() error {
	scp.processor = createProcessor(scp.LeafSnippet.Leaf)
	scp.processor.Prepare()

	return nil
}

func (scp *switchCaseProcessor) Execute(ctx *Context) error {
	newCtx := ctx.clone()
	defer newCtx.putPool()
	return scp.processor.Execute(newCtx)
}

type unfoldProcessor struct {
	*ast.UnfoldDirective

	processor PrepareProcessor
}

func newUnfoldProcessor(stmt *ast.UnfoldDirective) *unfoldProcessor {
	up := new(unfoldProcessor)
	up.UnfoldDirective = stmt
	return up
}

func (up *unfoldProcessor) Prepare() error {
	up.processor = createProcessor(up.FragmentDirective)
	return up.processor.Prepare()
}

func (up *unfoldProcessor) Execute(ctx *Context) error {
	newCtx := ctx.clone()
	defer newCtx.putPool()
	return up.processor.Execute(newCtx)
}

type wrapLayer struct {
	proc PrepareProcessor
	next *wrapLayer
}

func (wl *wrapLayer) Prepare() error {
	next := wl
	for next != nil {
		err := next.proc.Prepare()
		if err != nil {
			return errors.Join(err, errors.New("prepare wrapLayer processor error"))
		}
		next = next.next
	}

	return nil
}

func (wl *wrapLayer) Execute(ctx *Context) error {
	newCtx := ctx.clone()
	defer newCtx.putPool()

	if wl.next != nil {
		newCtx.nextWrap = func() error {
			newCtx := ctx.clone()
			defer newCtx.putPool()
			return wl.next.Execute(newCtx)
		}
	}

	return wl.proc.Execute(newCtx)
}

type wrapProcessor struct {
	*ast.WrapStmt

	processor PrepareProcessor
}

func newWrapProcessor(stmt *ast.WrapStmt) *wrapProcessor {
	wp := new(wrapProcessor)
	wp.WrapStmt = stmt
	return wp
}

func (wp *wrapProcessor) Prepare() error {
	next := &wrapLayer{
		proc: createProcessor(wp.LeafSnippet.Leaf),
	}
	for i := len(wp.WrapOpStmtList) - 1; i >= 0; i-- {
		next = &wrapLayer{
			proc: createProcessor(wp.WrapOpStmtList[i]),
			next: next,
		}
	}

	wp.processor = next
	return wp.processor.Prepare()
}

func (wp *wrapProcessor) Execute(ctx *Context) error {
	newCtx := ctx.clone()
	defer newCtx.putPool()
	return wp.processor.Execute(newCtx)
}

type fragmentProcessor struct {
	*ast.FragmentDirective
	processor PrepareProcessor
}

func newFragmentProcessor(stmt *ast.FragmentDirective) *fragmentProcessor {
	fp := new(fragmentProcessor)
	fp.FragmentDirective = stmt
	return fp
}

func (fp *fragmentProcessor) Prepare() error {
	fp.processor = createProcessor(fp.ExedescStmt.ExedescStmt)
	return fp.processor.Prepare()
}

func (fp *fragmentProcessor) Execute(ctx *Context) error {
	newCtx := ctx.clone()
	defer newCtx.putPool()
	return fp.processor.Execute(newCtx)
}
