package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gorch/gorch/internal/lang/iantlr/alr"
)

type FragmentDirective struct {
	Pos
	Name             string
	UnfoldDirectives UnfoldDirectives
	ExedescStmt      *ExedescStmt

	// expectExecuteOperators 整个FRAGMENT指令中预期内要执行的算子
	expectExecuteOperators map[string]struct{}

	// GoDirectiveMap 整个FRAGMENT指令中的所有GO指令对象都收集起来
	GoDirectiveMap map[string]*GoDirective

	// WaitDirectiveNames 整个FRAGMENT指令中的所有WAIT指令的名称都收集起来
	WaitDirectiveNames map[string]struct{}
}

type FragmentDirectiveHolder interface {
	AcceptFragmentDirective(*FragmentDirective) error
}

func (fs *FragmentDirective) AcceptStringLiteral(val string) error {
	fs.Name = val
	return nil
}

func (fs *FragmentDirective) AcceptExedescStmt(gs *ExedescStmt) error {
	fs.ExedescStmt = gs
	return nil
}

func (fs *FragmentDirective) AppendUnfoldDirective(us *UnfoldDirective) error {
	fs.UnfoldDirectives.Append(us)
	return nil
}

func (fs *FragmentDirective) AppendGoDirective(brs *GoDirective) error {
	if fs.GoDirectiveMap == nil {
		fs.GoDirectiveMap = make(map[string]*GoDirective)
	}
	if _, ok := fs.GoDirectiveMap[brs.GoName]; ok {
		return &ParseError{
			Pos: fs.Pos,
			Msg: "duplicate goDirective",
			KV:  map[string]string{"name": brs.GoName},
		}
	}
	fs.GoDirectiveMap[brs.GoName] = brs
	return nil
}

func (ss *FragmentDirective) AppendWaitDirectiveName(name string) error {
	if ss.WaitDirectiveNames == nil {
		ss.WaitDirectiveNames = make(map[string]struct{})
	}
	ss.WaitDirectiveNames[name] = struct{}{}
	return nil
}

func (ss *FragmentDirective) AcceptOperatorStmt(os *OperatorStmt) error {
	ss.expectExecuteOperators[os.Name] = struct{}{}
	return nil
}

func (fs *FragmentDirective) Description(sb *strings.Builder, blank string) {
	sb.WriteString("FRAGMENT(\"")
	sb.WriteString(fs.Name)
	sb.WriteString("\"){\n\n")
	blank += "  "
	fs.ExedescStmt.Description(sb, blank)
	sb.WriteString("\n\n}\n")
}

func (fs *FragmentDirective) String() string {
	builder := &strings.Builder{}
	builder.Grow(20)
	builder.WriteString("FRAGMENT(")
	builder.WriteString("\"")
	builder.WriteString(fs.Name)
	builder.WriteString("\"")
	builder.WriteString("):")
	builder.WriteString(strconv.Itoa(fs.StartLine))
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(fs.StartColunm))

	return builder.String()
}

// EnterFragmentDirective is called when entering the fragmentDirective production.
func (s *GorchlangParserListener) EnterFragmentDirective(ctx *alr.FragmentDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	fs := &FragmentDirective{
		Pos:                ToPos(ctx, s.File),
		GoDirectiveMap:     make(map[string]*GoDirective),
		WaitDirectiveNames: make(map[string]struct{}),
	}
	s.UnfoldDirectiveAppender = fs
	s.GoDirectiveAppender = fs
	s.WaitDirectiveNameAppender = fs
	s.Stack.Push(fs)
}

// ExitFragmentDirective is called when exiting the fragmentDirective production.
func (s *GorchlangParserListener) ExitFragmentDirective(ctx *alr.FragmentDirectiveContext) {
	s.UnfoldDirectiveAppender = nil
	s.GoDirectiveAppender = nil
	s.WaitDirectiveNameAppender = nil
	if len(s.ParseErrors) > 0 {
		return
	}

	fs := s.Stack.Pop().(*FragmentDirective)
	holder := s.Stack.Peek().(FragmentDirectiveHolder)
	err := holder.AcceptFragmentDirective(fs)
	if err != nil {
		s.AddError(err)
	}
}

type FragmentDirectives struct {
	Sort []string
	Map  map[string]*FragmentDirective
}

func NewFragmentDirectives() *FragmentDirectives {
	return &FragmentDirectives{
		Sort: make([]string, 0, 20),
		Map:  make(map[string]*FragmentDirective),
	}
}

func (fds *FragmentDirectives) Append(fs *FragmentDirective) error {
	if fds.Map == nil {
		fds.Map = make(map[string]*FragmentDirective)
	}

	_, has := fds.Map[fs.Name]

	if has {
		return &ParseError{
			Pos: fs.Pos,
			Msg: "duplicate fragmentDirective",
			KV:  map[string]string{"name": fs.Name},
		}
	}

	fds.Map[fs.Name] = fs
	fds.Sort = append(fds.Sort, fs.Name)

	return nil
}

func (fds *FragmentDirectives) Description(sb *strings.Builder, blank string) {
	for idx, fs := range fds.Sort {
		fds.Map[fs].Description(sb, blank)

		if idx < len(fds.Sort)-1 {
			sb.WriteString("\n")
		}
	}
}

func (fds *FragmentDirectives) SelfCheck() error {
	// 检查unfold引用fragment是否存在调用环
	if err := fds.checkUnfoldCycle(); err != nil {
		return err
	}
	return nil
}

// checkUnfoldCycle fragmentStmt循环引用检查
func (fds *FragmentDirectives) checkUnfoldCycle() error {
	for _, fsName := range fds.Sort {
		// 当前fragmentStmt
		curFs := fds.Map[fsName]

		if err := checkUnfoldCycle(fds, curFs, curFs); err != nil {
			return err
		}
	}

	return nil
}

// checkUnfoldCycle 递归进行循环引用检查
func checkUnfoldCycle(fds *FragmentDirectives, fd *FragmentDirective, checkFd *FragmentDirective) error {
	// 检查当前fragmentStmt的依赖
	for _, usName := range fd.UnfoldDirectives.Sort {
		// 找到依赖的fragmentStmt
		depFd, ok := fds.Map[usName]

		if !ok {
			continue
		}

		_, has := depFd.UnfoldDirectives.Map[checkFd.Name]
		if has {
			return fmt.Errorf("fragment directive import cycle error: "+
				"%s, conflict %s", checkFd, depFd)
		}

		if err := checkUnfoldCycle(fds, depFd, checkFd); err != nil {
			return err
		}
	}

	return nil
}
