package iparser

import (
	"testing"

	"github.com/gorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

type gorchLangSimpleFragmentSyntaxCase struct {
	name string
	lang string
}

var gorchLangSimpleFragmentSyntaxCases = []gorchLangSimpleFragmentSyntaxCase{
	{
		name: "single_serial",
		lang: `FRAGMENT("name") {p(name="123") -> p1(name="123", b=123, c=true)}`,
	},
	{
		name: "single_concurrent",
		lang: `FRAGMENT("name") {[p(name="123"),p1(name="123", b=123, c=true)]}`,
	},
	{
		name: "single_op",
		lang: `FRAGMENT("name") {p(name="123")}`,
	},
	{
		name: "single_wrap",
		lang: `FRAGMENT("name") {
			a | (b -> c)
		}`,
	},
	{
		name: "single_unflod",
		lang: `FRAGMENT("0") {
			UNFOLD("123")
		}
		FRAGMENT("123") {
			a -> b
		}
		`,
	},
}

func TestGorchSingleFragmentSyntax(t *testing.T) {
	for _, cs := range gorchLangSimpleFragmentSyntaxCases {
		t.Run("fragment_"+cs.name, func(t *testing.T) {
			RunMyTest(t, cs.lang, registerOp, func(p *ast.Primary, f string, e error) {
				assert.Nil(t, e)
				assert.Equal(t, strip(cs.lang+registerOp), strip(p.Description()))
			})
		})
	}
}
