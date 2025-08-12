package iparser

import (
	"testing"

	"github.com/gogorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

func TestMissOperator(t *testing.T) {
	t.Run("notFoundOperator", func(t *testing.T) {
		var g = `
	START("name") {
		b
	}
	`
		RunMyTest(t, g, "", func(p *ast.Primary, f string, e error) {
			assert.EqualError(t, e, "selfcheck gorch gramer error\nin start \"name\", operator \"b\" not register")
			assert.Nil(t, p)
		})

	})

	t.Run("useNoCheckMissDirective", func(t *testing.T) {
		var registerSyntax = `
	START("name") {
		NO_CHECK_MISS()
		b
	}
	`
		RunMyTest(t, registerSyntax, "", func(p *ast.Primary, f string, e error) {
			assert.Nil(t, e)
			assert.Equal(t, true, p.Starters.Map["name"].NoCheckMiss)
		})
	})
}
