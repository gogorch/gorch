package iparser

import (
	"fmt"
	"testing"

	"github.com/gorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

func TestFragmentNotFound(t *testing.T) {
	t.Run("startNotFoundFragment", func(t *testing.T) {
		var g = `
START("test") {
	a -> UNFOLD("123")
}
`
		reg := `
REGISTER("github.com/gorch") {
	OPERATOR("gorch/ops","a1", "a", 1)
}`
		RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
			assert.EqualError(t, e,
				fmt.Sprintf("selfcheck gorch gramer error\nfragment not found, file=%s(3:6) name=123", f))
			assert.Nil(t, p)
		})
	})

	t.Run("unfoldNotFoundFragment", func(t *testing.T) {
		var g = `
START("test") {
	a -> UNFOLD("123")
}
FRAGMENT("123") {
	a -> UNFOLD("456")
}
`
		reg := `
REGISTER("github.com/gorch") {
	OPERATOR("gorch/ops","a1", "a", 1)
	}`
		RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
			assert.EqualError(t, e,
				fmt.Sprintf("selfcheck gorch gramer error\nfragment not found, file=%s(6:6) name=456", f))
			assert.Nil(t, p)
		})
	})

	t.Run("waitNotFoundFragment", func(t *testing.T) {
		var registerSyntax = `
	START("name") {
		UNFOLD("123")
	}
	`
		RunMyTest(t, registerSyntax, "", func(p *ast.Primary, f string, e error) {
			assert.EqualError(t, e, fmt.Sprintf("selfcheck gorch gramer error\nfragment not found, file=%s(3:2) name=123", f))
			assert.Nil(t, p)
		})
	})
}
