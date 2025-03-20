package iparser

import (
	"testing"

	"github.com/gorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

var registerOp1 = `
FRAGMENT("123") {
	a -> b
}

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

type gorchLangComplexStartSyntaxCase struct {
	name string
	lang string
}

var gorchLangComplexStartSyntaxCases = []gorchLangComplexStartSyntaxCase{
	{
		name: "serial",
		lang: `START("name", abc="123", bb=1, cc=true) {
			p(name="123") -> a -> [a, b]
			-> (c -> a)
			-> UNFOLD("123")
			-> (a | b)
			-> GO(a, "go_a")
			-> SWITCH(switch_op){
				CASE "123" => op,
				CASE "456" => (op -> op1)
			}
			-> b(WAIT("go_a"))
		}`,
	},
	{
		name: "concurrent",
		lang: `START("name", abc="123", bb=1, cc=true) {
			[
				p(name="123"),
				(p -> a),
				UNFOLD("123"),
				GO(a, "go_a"),
				SWITCH(switch_op){
					CASE "123" => op,
					CASE "456" => (op -> op1)
				},
				b(WAIT("go_a"))
			]
		}`,
	},
	{
		name: "wrap_serial",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | (GO(b, "b") -> c(WAIT("b")))
		}`,
	},
	{
		name: "wrap_concurrent",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | [a,b]
		}`,
	},
	{
		name: "wrap_unfold",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | UNFOLD("123")
		}`,
	},
	{
		name: "wrap_switch",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | SWITCH(switch_op){
				CASE "123" => op,
				CASE "456" => (op -> op1)
			}
		}`,
	},
	{
		name: "wrap_wrap_brack",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | (a | b)
		}`,
	},
	{
		name: "wrap_go",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | GO(a, "123")
		}`,
	},
	{
		name: "wrap_op",
		lang: `START("name", abc="123", bb=1, cc=true) {
			a | a
		}`,
	},
	{
		name: "switch",
		lang: `START("name", abc="123", bb=1, cc=true) {
			SWITCH(switch_op){
				CASE "123" => op,
				CASE "456" => (op -> op1),
				CASE "789" => [a, b],
				CASE "10" => SWITCH(switch_op_nest) {CASE "1" => op, CASE "2" => op},
				CASE "11" => UNFOLD("123"),
				CASE "12" => GO(a, "123"),
				CASE "14" => (a | b)
			}
		}`,
	},
}

func TestGorchComplexStartSyntax(t *testing.T) {
	for _, cs := range gorchLangComplexStartSyntaxCases {
		t.Run(cs.name, func(t *testing.T) {
			RunMyTest(t, cs.lang, registerOp1, func(p *ast.Primary, f string, e error) {
				assert.Nil(t, e)
				assert.Equal(t, strip(cs.lang+registerOp1), strip(p.Description()))
			})
		})
	}
}
