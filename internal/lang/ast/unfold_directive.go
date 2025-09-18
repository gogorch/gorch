package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type UnfoldDirective struct {
	Pos
	Name string

	// FragmentDirective 解析最终阶段进行填充
	FragmentDirective *FragmentDirective
}

type UnfoldDirectiveHolder interface {
	AcceptUnfoldDirective(*UnfoldDirective) error
}

type UnfoldDirectiveAppender interface {
	AppendUnfoldDirective(us *UnfoldDirective) error
}

func (ias *UnfoldDirective) AcceptStringLiteral(val string) error {
	ias.Name = val
	return nil
}
func (*UnfoldDirective) isExedesc() {}

func (ias *UnfoldDirective) Description(s *strings.Builder, _ string) {
	s.WriteString("UNFOLD(\"")
	s.WriteString(ias.Name)
	s.WriteString("\")")
}

func (ias *UnfoldDirective) String() string {
	builder := &strings.Builder{}
	builder.Grow(20)
	builder.WriteString("UNFOLD(")
	builder.WriteString("\"")
	builder.WriteString(ias.Name)
	builder.WriteString("\"")
	builder.WriteString(")")
	return builder.String()
}

func (*UnfoldDirective) isLeafSnippet() {}

var _ LeafSnippetI = (*UnfoldDirective)(nil)

// EnterUnfoldDirective is called when entering the unfoldDirective production.
func (s *GorchlangParserListener) EnterUnfoldDirective(ctx *alr.UnfoldDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	us := &UnfoldDirective{Pos: ToPos(ctx, s.File)}
	s.Stack.Push(us)
}

// ExitUnfoldDirective is called when exiting the unfoldDirective production.
func (s *GorchlangParserListener) ExitUnfoldDirective(ctx *alr.UnfoldDirectiveContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	us := s.Stack.Pop().(*UnfoldDirective)
	holder := s.Stack.Peek().(UnfoldDirectiveHolder)
	err := holder.AcceptUnfoldDirective(us)
	if err != nil {
		s.AddError(err)
	}

	if s.UnfoldDirectiveAppender != nil {
		s.UnfoldDirectiveAppender.AppendUnfoldDirective(us)
	}
}

type UnfoldDirectives struct {
	Sort []string
	Map  map[string][]*UnfoldDirective
}

func (isl *UnfoldDirectives) Append(ss *UnfoldDirective) error {
	if isl.Map == nil {
		isl.Map = make(map[string][]*UnfoldDirective)
	}

	isl.Map[ss.Name] = append(isl.Map[ss.Name], ss)
	isl.Sort = append(isl.Sort, ss.Name)

	return nil
}
