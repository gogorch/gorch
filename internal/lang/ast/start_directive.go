package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type StartDirective struct {
	Pos

	Name string
	// Args Start指令的参数
	Args *ArgsStmt

	// OnFinishStmt Start指令内定义结束执行的事件，总是会被调用
	OnFinishStmt *OnFinishStmt

	// ExedescStmt 执行描述体
	ExedescStmt *ExedescStmt

	// NoCheckMiss 不校验算子缺少
	NoCheckMiss bool

	// expectExecuteOperators 预期内要执行的算子
	// 整个START指令中（含UNFOLD指令内）包含的所有算子都收集起来
	expectExecuteOperators map[string]struct{}

	// goDirectiveMap 整个START指令中（含UNFOLD指令内）包含的所有GoDirective都收集起来
	goDirectiveMap map[string]*GoDirective

	// WaitDirectiveMap 整个START指令中（含UNFOLD指令内）的所有WAIT指令的名称都收集起来
	// 用来校验是否与GoDirective匹配
	WaitDirectiveMap map[string]struct{}

	// UnfoldDirectives 待展开的FRAGMENT（不含子级UNFOLD指令内）
	UnfoldDirectives UnfoldDirectives
}

type StartStmtHolder interface {
	AcceptStartDirective(*StartDirective) error
}

type NoCheckMissDirectiveAcceptor interface {
	AcceptNoCheckMissDirective() error
}

func (ss *StartDirective) AcceptArgsStmt(args *ArgsStmt) error {
	ss.Args = args
	return nil
}

func (ss *StartDirective) AcceptStringLiteral(val string) error {
	ss.Name = val
	return nil
}

func (ss *StartDirective) AcceptExedescStmt(gs *ExedescStmt) error {
	ss.ExedescStmt = gs
	if os, ok := gs.ExedescStmt.(*OperatorStmt); ok {
		ss.expectExecuteOperators[os.Name] = struct{}{}
	}
	return nil
}

func (ss *StartDirective) AcceptOnFinishStmt(es *OnFinishStmt) error {
	ss.OnFinishStmt = es
	return nil
}

func (ss *StartDirective) AppendUnfoldDirective(us *UnfoldDirective) error {
	ss.UnfoldDirectives.Append(us)
	return nil
}

func (ss *StartDirective) AcceptNoCheckMissDirective() error {
	ss.NoCheckMiss = true
	return nil
}

func (ss *StartDirective) AppendGoDirective(brs *GoDirective) error {
	if ss.goDirectiveMap == nil {
		ss.goDirectiveMap = make(map[string]*GoDirective)
	}
	if _, ok := ss.goDirectiveMap[brs.GoName]; ok {
		return &ParseError{
			Pos: brs.Pos,
			Msg: "duplicate goDirective",
			KV:  map[string]string{"name": brs.GoName},
		}
	}
	ss.goDirectiveMap[brs.GoName] = brs
	return nil
}

func (ss *StartDirective) AcceptOperatorStmt(os *OperatorStmt) error {
	ss.expectExecuteOperators[os.Name] = struct{}{}
	return nil
}

func (ss *StartDirective) AppendWaitDirectiveName(name string) error {
	if ss.WaitDirectiveMap == nil {
		ss.WaitDirectiveMap = make(map[string]struct{})
	}
	ss.WaitDirectiveMap[name] = struct{}{}
	return nil
}

func (ss *StartDirective) String() string {
	builder := &strings.Builder{}
	builder.Grow(20)
	builder.WriteString("START(")
	builder.WriteString("\"")
	builder.WriteString(ss.Name)
	builder.WriteString("\"")
	builder.WriteString("):")
	builder.WriteString(strconv.Itoa(ss.StartLine))
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(ss.StartColunm))
	return builder.String()
}

func (ss *StartDirective) Description(sb *strings.Builder, blank string) {
	sb.WriteString("START(\"")
	sb.WriteString(ss.Name)
	sb.WriteString("\"")
	if ss.Args != nil {
		sb.WriteString(",\n")
		ss.Args.Description(sb, blank)
	}
	sb.WriteString("){\n\n")

	if ss.OnFinishStmt != nil {
		ss.OnFinishStmt.Description(sb, "  ")
	}

	if ss.NoCheckMiss {
		sb.WriteString("\n")
		sb.WriteString("NO_CHECK_MISS()")
		sb.WriteString("\n")
	}

	ss.ExedescStmt.Description(sb, "  ")
	sb.WriteString("\n\n}\n")
}

// EnterStartDirective is called when entering the startDirective production.
func (s *GorchlangParserListener) EnterStartDirective(ctx *alr.StartDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := &StartDirective{
		Pos:                    ToPos(ctx, s.File),
		goDirectiveMap:         make(map[string]*GoDirective),
		WaitDirectiveMap:       make(map[string]struct{}),
		expectExecuteOperators: make(map[string]struct{}),
	}
	s.UnfoldDirectiveAppender = ss
	s.GoDirectiveAppender = ss
	s.WaitDirectiveNameAppender = ss
	s.NoCheckMissDirectiveAcceptor = ss
	s.Stack.Push(ss)
}

// ExitStartDirective is called when exiting the startDirective production.
func (s *GorchlangParserListener) ExitStartDirective(ctx *alr.StartDirectiveContext) {
	s.UnfoldDirectiveAppender = nil
	s.GoDirectiveAppender = nil
	s.WaitDirectiveNameAppender = nil
	s.NoCheckMissDirectiveAcceptor = nil
	if len(s.ParseErrors) > 0 {
		return
	}

	ss := s.Stack.Pop().(*StartDirective)
	holder := s.Stack.Peek().(StartStmtHolder)
	err := holder.AcceptStartDirective(ss)
	if err != nil {
		s.AddError(err)
	}
}

// EnterNoCheckMissDirective is called when entering the noCheckOperatorMissDirective production.
func (s *GorchlangParserListener) EnterNoCheckMissDirective(c *alr.NoCheckMissDirectiveContext) {
}

// ExitNoCheckMissDirective is called when exiting the noCheckMissDirective production.
func (s *GorchlangParserListener) ExitNoCheckMissDirective(c *alr.NoCheckMissDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}
	s.NoCheckMissDirectiveAcceptor.AcceptNoCheckMissDirective()
}

type StartDirectives struct {
	Sort []string
	Map  map[string]*StartDirective
}

func NewStartDirectives() *StartDirectives {
	return &StartDirectives{
		Sort: make([]string, 0, 20),
		Map:  make(map[string]*StartDirective),
	}
}

func (sds *StartDirectives) Append(ss *StartDirective) error {
	if sds.Map == nil {
		sds.Map = make(map[string]*StartDirective)
	}

	_, has := sds.Map[ss.Name]

	if has {
		return &ParseError{
			Pos: ss.Pos,
			Msg: "duplicate startDirective",
			KV:  map[string]string{"name": ss.Name},
		}
	}

	sds.Map[ss.Name] = ss
	sds.Sort = append(sds.Sort, ss.Name)

	return nil
}

func (sds *StartDirectives) Description(sb *strings.Builder, blank string) {
	for idx, sn := range sds.Sort {
		sds.Map[sn].Description(sb, blank)

		if idx < len(sds.Sort)-1 {
			sb.WriteString("\n")
		}
	}
}

// checkGoDirectiveDuplicate 检查在一个START指令内，是否存在重复命名的GO指令
func (sds *StartDirectives) checkGoDirectiveDuplicate() error {
	for _, sd := range sds.Map {
		if e := checkFragmentGoDuplicate(sd.UnfoldDirectives, sd.goDirectiveMap); e != nil {
			return fmt.Errorf("start `%q` has %s", sd.Name, e.Error())
		}
	}

	return nil
}

func (sds *StartDirectives) SelfCheck() error {
	// 将图上START指令中的GO指令，递归的结合FRAGMENT，收集到一个START指令内，并检查是否存在冲突
	if err := sds.checkGoDirectiveDuplicate(); err != nil {
		return err
	}

	// 将图上START、FRAGMENT指令中的WAIT指令，匹配相应的GO指令，检查是否匹配
	if err := sds.checkWaitDirective(); err != nil {
		return err
	}
	return nil
}

func checkFragmentGoDuplicate(uds UnfoldDirectives, gdm map[string]*GoDirective) error {
	for _, ud := range uds.Map {
		fragmentStmt := ud[0].FragmentDirective
		if e := checkFragmentGoDuplicate(fragmentStmt.UnfoldDirectives, fragmentStmt.GoDirectiveMap); e != nil {
			return e
		}

		for _, fgd := range fragmentStmt.GoDirectiveMap {
			if _, ok := gdm[fgd.GoName]; ok {
				return &ParseError{
					Pos: fgd.Pos,
					Msg: "duplicate goDirective",
					KV:  map[string]string{"name": fgd.GoName},
				}
			}
			gdm[fgd.GoName] = fgd
		}
	}

	return nil
}

func (sds *StartDirectives) checkWaitDirective() error {
	for _, sd := range sds.Map {
		checkWaitDirective(sd.UnfoldDirectives, sd.WaitDirectiveMap)
	}

	for _, sd := range sds.Map {
		for wn := range sd.WaitDirectiveMap {
			if _, ok := sd.goDirectiveMap[wn]; !ok {
				return fmt.Errorf("START %q: WAIT(%q) not find matched GO directive", sd.Name, wn)
			}
		}
	}

	return nil
}

func (sds *StartDirectives) checkOperatorMiss(allOperators *Operators) error {
	for _, sd := range sds.Map {
		if sd.NoCheckMiss {
			continue
		}

		// 检查算子是否注册
		for i := range sd.expectExecuteOperators {
			if _, ok := allOperators.Names[i]; !ok {
				return fmt.Errorf("in start %q, operator %q not register", sd.Name, i)
			}
		}
	}
	return nil
}

func checkWaitDirective(uds UnfoldDirectives, gdm map[string]struct{}) {
	for _, ud := range uds.Map {
		fragmentStmt := ud[0].FragmentDirective
		checkWaitDirective(fragmentStmt.UnfoldDirectives, fragmentStmt.WaitDirectiveNames)

		for fgd := range fragmentStmt.WaitDirectiveNames {
			gdm[fgd] = struct{}{}
		}
	}
}
