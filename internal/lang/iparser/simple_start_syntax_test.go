package iparser

import (
	"strings"
	"testing"

	"github.com/gorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

func strip(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

var registerOp = `
REGISTER("github.com/gorch") {
    OPERATOR("gorch/ops", "p", 1)
    OPERATOR("gorch/ops", "p1", 2)
    OPERATOR("gorch/ops", "test", "op", 3)
    OPERATOR("gorch/ops", "a", 4)
    OPERATOR("gorch/ops", "b", 5)
    OPERATOR("gorch/ops", "op1", 6)
    OPERATOR("gorch/ops", "c", 7)
    OPERATOR("gorch/ops", "switch_op", 8)
    OPERATOR("switch_op_nest", "switch_op_nest", 9)
}`

type gorchLangSimpleStartSyntaxCase struct {
	name string
	lang string
}

var gorchLangSimpleSyntaxCases = []gorchLangSimpleStartSyntaxCase{
	{
		name: "args",
		lang: `START("name", str1="123", str2=["123", "456"], int1=1, int2=[123, 456], bool1=true, bool2=[true, false]) {
			p(str1="123", str2=["123", "456"], int1=1, int2=[123, 456], bool1=true, bool2=[true, false])
				-> GO(a, "a")
				-> p1(name="123", b=123, c=true)
				-> p1(WAIT("a"))
				-> p1(str1="123", str2=["123", "456"], int1=1, int2=[123, 456], bool1=true, bool2=[true, false], WAIT("a"))
		}`,
	},
	{
		name: "single_serial",
		lang: `START("name") {
			p(name="123")
			-> p1(name="123", b=123, c=true)
			-> [p, p1]
			-> (p -> p1)
			-> UNFOLD("123")
			-> (p | p1 | (p -> p1))
			-> GO(p, "p1_go")
			-> SWITCH(dsfa) {
				CASE "123" => op
			}
		}
			FRAGMENT("123") {
			a -> b
		}`,
	},
	{
		name: "single_concurrent",
		lang: `START("name", abc="123", bb=1, cc=true) {
			[
				p(name="123"),
				p1(name="123", b=123, c=true),
				[p, p1],
				(p -> p1),
				UNFOLD("123"),
				(p | p1 | (p -> p1)),
				GO(p, "p1_go"),
				SWITCH(dsfa) {
					CASE "123" => op
				}
			]
		}
			FRAGMENT("123") {
			a -> b
		}`,
	},
	{
		name: "single_wrap_serial",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | a | (GO(b, "b") -> c(WAIT("b")))
		}`,
	},
	{
		name: "single_wrap_concurrent",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | a | [b, c] 
		}`,
	},
	{
		name: "single_wrap_wrap_bracket",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | a | (a | b)
		}`,
	},
	{
		name: "single_wrap_operator",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | a | a
		}`,
	},
	{
		name: "single_wrap_switch",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | a | SWITCH(switch_op) {
				CASE "123" => a
			}
		}`,
	},
	{
		name: "single_wrap_unfold",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | a | UNFOLD("123")
		}
			FRAGMENT("123") {
			a -> b
		}`,
	},
	{
		name: "single_op",
		lang: `START("name", abc="123", bb=1, cc=true) {p(name="123")}`,
	},
	{
		name: "single_switch",
		lang: `START("name", abc="123", bb=1, cc=true) {
			SWITCH(switch_op){
				CASE "123" => (op -> op1),
				CASE "456" => [a, b],
				CASE "789" => (a | a | b),
				CASE "10" => GO(a, "123"),
				CASE "11" => op,
				CASE "12" => SWITCH(switch_op_nest) {CASE "1" => op, CASE "2" => op}
			}
		}`,
	},
	{
		name: "single_op_with_on_err_on_finish_preset",
		lang: `START("name", abc="123", bb=1, cc=true) {
			ON_FINISH() { a }
			op
		}`,
	},
	{
		name: "single_unflod",
		lang: `START("name", abc="123", bb=1, cc=true) {
			UNFOLD("123")
		}
		FRAGMENT("123") {
			a -> b
		}
		`,
	},
}

func TestGorchSingleStartSyntax(t *testing.T) {
	for _, cs := range gorchLangSimpleSyntaxCases {
		t.Run("start_"+cs.name, func(t *testing.T) {
			RunMyTest(t, cs.lang, registerOp, func(p *ast.Primary, f string, e error) {
				assert.Nil(t, e)
				assert.Equal(t, strip(cs.lang+registerOp), strip(p.Description()))
			})
		})
	}
}
