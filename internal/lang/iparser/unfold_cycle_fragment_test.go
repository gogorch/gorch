package iparser

import (
	"testing"

	"github.com/gorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

// TestNestFragmentGraph1 检查一层引用循环
func TestNestFragmentGraph1(t *testing.T) {
	var nestFragmentGraph1 = `
START("test"){
	UNFOLD("123") -> UNFOLD("456") -> p
}

FRAGMENT("123") {
	a -> UNFOLD("456")
}

FRAGMENT("456") {
	a -> UNFOLD("123")
}`
	reg := `
REGISTER("github.com/gorch") {
	OPERATOR("gorch/ops", "a", 1)
	OPERATOR("gorch/ops", "p", 2)
}
`
	RunMyTest(t, nestFragmentGraph1, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nfragment directive import cycle error: FRAGMENT(\"123\"):6:0, conflict FRAGMENT(\"456\"):10:0")
		assert.Nil(t, p)
	})
}

// TestNestFragmentGraph2 检查三层引用循环
func TestNestFragmentGraph2(t *testing.T) {
	var nestFragmentGraph2 = `
START("test"){
	UNFOLD("123")
}

FRAGMENT("123") {
	a1 -> UNFOLD("456")
}

FRAGMENT("456") {
	a1 -> UNFOLD("789")
}

FRAGMENT("789") {
	a1 -> UNFOLD("10")
}
FRAGMENT("10") {
	a1 -> UNFOLD("123")
}`
	reg := `
REGISTER("github.com/gorch") {
	OPERATOR("gorch/ops", "a", "a1", 1)
}
`
	RunMyTest(t, nestFragmentGraph2, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nfragment directive import cycle error: FRAGMENT(\"123\"):6:0, conflict FRAGMENT(\"10\"):17:0")
		assert.Nil(t, p)
	})
}
