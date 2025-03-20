package iparser

import (
	"testing"

	"github.com/gorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

func TestStartCollectGoAndWaitSuc(t *testing.T) {
	var g = `
	START("888") {
		GO(a, "789") -> UNFOLD("456")
	}

	FRAGMENT("456"){
		a -> UNFOLD("789")
	}

	FRAGMENT("789"){
		a -> a(b=true, WAIT("789"))
	} `

	reg := `REGISTER("github.com/gorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.Nil(t, e)
		assert.Equal(t, map[string]struct{}{"789": {}}, p.Starters.Map["888"].WaitDirectiveMap)
	})
}

func TestStartCollectGoAndWaitNotMatch(t *testing.T) {
	var g = `
	START("888") {
		GO(a, "789") -> UNFOLD("456")
	}

	START("999") {
		UNFOLD("456")
	}

	FRAGMENT("456"){
		a -> UNFOLD("789")
	}

	FRAGMENT("789"){
		a -> a(b=true, WAIT("789"))
	}`
	reg := `
	REGISTER("github.com/gorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"999\": WAIT(\"789\") not find matched GO directive")
		assert.Nil(t, p)
	})
}
