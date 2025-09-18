package ast

import (
	"strings"

	alr "github.com/gogorch/gorch/internal/lang/iantlr"
)

type Primary struct {
	Starters  *StartDirectives
	Fragments *FragmentDirectives
	Registers *RegisterDirectives
}

func NewPrimary() *Primary {
	return &Primary{
		Starters:  NewStartDirectives(),
		Fragments: NewFragmentDirectives(),
		Registers: NewRegisterDirectives(),
	}
}

func (p *Primary) Description() string {
	s := new(strings.Builder)
	s.Grow(1024 * 30)
	p.Starters.Description(s, "")
	s.WriteString("\n")
	p.Fragments.Description(s, "")
	p.Registers.Description(s, "")
	return s.String()
}

func (p *Primary) AcceptRegisterDirective(rs *RegisterDirective) error {
	return p.Registers.Append(rs)
}

func (p *Primary) AcceptFragmentDirective(fs *FragmentDirective) error {
	return p.Fragments.Append(fs)
}

func (p *Primary) AcceptStartDirective(ss *StartDirective) error {
	return p.Starters.Append(ss)
}

// EnterPrimary is called when production program is entered.
func (s *GorchlangParserListener) EnterPrimary(ctx *alr.PrimaryContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	if v := s.Stack.Peek(); v != nil {
		pv, ok := v.(*Primary)
		if !ok || pv == nil {
			panic("please push a primary to stack")
		}
	}
}

// ExitPrimary is called when production program is exited.
func (s *GorchlangParserListener) ExitPrimary(ctx *alr.PrimaryContext) {
	if len(s.ParseErrors) > 0 {
		return
	}

	_ = s.Stack.Pop().(*Primary)
}

func (p *Primary) SelfCheck() error {
	// 将图上START、FRAGMENT指令中的UNFOLD指令用FRAGMENT填充
	if err := p.fullFragmentDirective(); err != nil {
		return err
	}

	if err := p.Fragments.SelfCheck(); err != nil {
		return err
	}

	if err := p.Starters.SelfCheck(); err != nil {
		return err
	}

	//  检查算子是否注册
	if err := p.Starters.checkOperatorMiss(p.Registers.AllOperators); err != nil {
		return err
	}

	return nil
}

// fullFragmentDirective 将Start/Fragment语句中的UnfoldStmt用FragmentStmt进行填充
func (p *Primary) fullFragmentDirective() error {
	for _, startStmtName := range p.Starters.Sort {
		startStmt := p.Starters.Map[startStmtName]

		for _, unfoldName := range startStmt.UnfoldDirectives.Sort {
			unfoldStmts := startStmt.UnfoldDirectives.Map[unfoldName]
			for _, unfoldStmt := range unfoldStmts {
				fragmentStmt, ok := p.Fragments.Map[unfoldName]
				if !ok {
					return &ParseError{
						Pos: unfoldStmt.Pos,
						Msg: "fragment not found",
						KV:  map[string]string{"name": unfoldName},
					}
				}

				unfoldStmt.FragmentDirective = fragmentStmt
			}
		}
	}

	for _, fragmentStmtName := range p.Fragments.Sort {
		fragmentStmt := p.Fragments.Map[fragmentStmtName]

		for _, unfoldName := range fragmentStmt.UnfoldDirectives.Sort {
			unfoldStmts := fragmentStmt.UnfoldDirectives.Map[unfoldName]
			for _, unfoldStmt := range unfoldStmts {
				fragmentStmt, ok := p.Fragments.Map[unfoldName]
				if !ok {
					return &ParseError{
						Pos: unfoldStmt.Pos,
						Msg: "fragment not found",
						KV:  map[string]string{"name": unfoldName},
					}
				}

				unfoldStmt.FragmentDirective = fragmentStmt
			}
		}
	}

	return nil
}
