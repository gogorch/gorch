package iparser

import (
	"testing"

	"github.com/gogorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

// TestNestFragmentGraph2 检查二层引用循环
func TestFullFragment(t *testing.T) {
	g := `
START("test"){
	UNFOLD("123") -> UNFOLD("456")
}

FRAGMENT("123") {
	UNFOLD("456")
}

FRAGMENT("456") {
	a1 -> UNFOLD("789")
}

FRAGMENT("789") {
	a1
}`
	reg := `REGISTER("github.com/gogorch") {
	OPERATOR("gorch/ops", "a", "a1", 1)
}`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.Nil(t, e)
		assert.Equal(t, 3, len(p.Fragments.Map))
		assert.Equal(t, 2, len(p.Starters.Map["test"].UnfoldDirectives.Map))
		assert.Equal(t, "123", p.Starters.Map["test"].UnfoldDirectives.Map["123"][0].Name)
		assert.Equal(t, "456", p.Starters.Map["test"].UnfoldDirectives.Map["456"][0].Name)

		assert.Equal(t, "123", p.Starters.Map["test"].UnfoldDirectives.Map["123"][0].FragmentDirective.Name)
		assert.Equal(t, "456", p.Starters.Map["test"].UnfoldDirectives.Map["456"][0].FragmentDirective.Name)
	})
}

func TestFullRepeatFragment(t *testing.T) {
	g := `
START("test"){
	UNFOLD("123") -> UNFOLD("123")
}

FRAGMENT("123") {
	UNFOLD("456")
}

FRAGMENT("456") {
	a1 -> UNFOLD("789")
}

FRAGMENT("789") {
	a1
}`
	reg := `REGISTER("github.com/gogorch") {
	OPERATOR("gorch/ops", "a", "a1", 1)
}`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.Nil(t, e)
		assert.Equal(t, 3, len(p.Fragments.Map))
		assert.Equal(t, 1, len(p.Starters.Map["test"].UnfoldDirectives.Map))
		assert.Equal(t, "123", p.Starters.Map["test"].UnfoldDirectives.Map["123"][0].Name)
		assert.Equal(t, "123", p.Starters.Map["test"].UnfoldDirectives.Map["123"][1].Name)

		assert.Equal(t, "123", p.Starters.Map["test"].UnfoldDirectives.Map["123"][0].FragmentDirective.Name)
		assert.Equal(t, "123", p.Starters.Map["test"].UnfoldDirectives.Map["123"][1].FragmentDirective.Name)
	})
}
