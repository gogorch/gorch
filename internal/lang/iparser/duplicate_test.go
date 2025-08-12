package iparser

import (
	"fmt"
	"testing"

	"github.com/gogorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

func TestRegisterOperatorDupName(t *testing.T) {
	var reg = `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
		OPERATOR("gorch/ops", "b", "a", 2)
	}
	`
	RunMyTest(t, reg, "", func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e,
			fmt.Sprintf("parse gorch file error\nregister operator duplicate name, file=%s(4:2) name=a",
				f))
		assert.Nil(t, p)
	})
}

func TestRegisterOperatorDupName1(t *testing.T) {
	var reg = `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}

	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "b", "a", 2)
	}
	`
	RunMyTest(t, reg, "", func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e,
			fmt.Sprintf("parse gorch file error\nregister operator duplicate name, file=%s(7:2) name=a",
				f))
		assert.Nil(t, p)
	})
}

func TestRegisterDupSeq(t *testing.T) {
	var reg = `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
		OPERATOR("gorch/ops", "b", 1)
	}
	`
	RunMyTest(t, reg, "", func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e,
			fmt.Sprintf("parse gorch file error\nregister operator duplicate seq, file=%s(4:2) seq=1",
				f))
		assert.Nil(t, p)
	})
}

func TestFragmentDuplicate(t *testing.T) {
	t.Run("duplicateFragment", func(t *testing.T) {
		var registerSyntax = `
		FRAGMENT("123") {
			a
		}
		FRAGMENT("123") {
			b
		}
		REGISTER("github.com/gogorch") {
		}
		`
		RunMyTest(t, registerSyntax, "", func(p *ast.Primary, f string, e error) {
			assert.EqualError(t, e, fmt.Sprintf("parse gorch file error\nduplicate fragmentDirective, file=%s(5:2) name=123", f))
			assert.Nil(t, p)
		})
	})
}

func TestSwitchCaseDup(t *testing.T) {
	var registerSyntax = `
	START("123") {
		SWITCH(a){
			CASE "123" => a,
			CASE "123" => (a -> a)
		}
	}`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, registerSyntax, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, fmt.Sprintf("parse gorch file error\nduplicate switch case name, file=%s(5:3) name=123", f))
		assert.Nil(t, p)
	})
}

func TestStartDup(t *testing.T) {
	var g = `
	START("123") {
		SWITCH(a){
			CASE "123" => a,
			CASE "456" => (a -> a)
		}
	}
	START("123") {
		a
	}`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, fmt.Sprintf("parse gorch file error\nduplicate startDirective, file=%s(8:1) name=123", f))
		assert.Nil(t, p)
	})
}

func TestStartGoDirectiveDup(t *testing.T) {
	var g = `
	START("123") {
		GO(a, "123") -> GO(a, "123")
	}`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, fmt.Sprintf("parse gorch file error\nduplicate goDirective, file=%s(3:18) name=123", f))
		assert.Nil(t, p)
	})
}

func TestFragmentGoDirectiveDup(t *testing.T) {
	var g = `
	START("123") {
		UNFOLD("456")
	}
	FRAGMENT("456"){
		GO(a, "123") -> GO(a, "123")
	}`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, fmt.Sprintf("parse gorch file error\nduplicate goDirective, file=%s(5:1) name=123", f))
		assert.Nil(t, p)
	})
}

func TestGoDirectiveDup(t *testing.T) {
	var registerSyntax = `
	START("123") {
		UNFOLD("456") -> GO(a, "123")
	}
	FRAGMENT("456"){
		a -> UNFOLD("789")
	}
	FRAGMENT("789"){
		a -> UNFOLD("10")
	}
	FRAGMENT("10"){
		a -> GO(a, "123")
	}`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, registerSyntax, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, fmt.Sprintf("selfcheck gorch gramer error\nstart `\"123\"` has duplicate goDirective, file=%s(12:7) name=123", f))
		assert.Nil(t, p)
	})
}

func TestGoDirectiveDup1(t *testing.T) {
	var g = `
	START("123") {
		UNFOLD("456")
	}
	FRAGMENT("456"){
		a -> UNFOLD("789") -> GO(a, "123")
	}
	FRAGMENT("789"){
		a -> UNFOLD("10")
	}
	FRAGMENT("10"){
		a -> GO(a, "123")
	}`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, fmt.Sprintf("selfcheck gorch gramer error\nstart `\"123\"` has duplicate goDirective, file=%s(12:7) name=123", f))
		assert.Nil(t, p)
	})
}

func TestGoDirectiveDup2(t *testing.T) {
	var g = `
	START("999") {
		UNFOLD("789")
	}

	START("123") {
		UNFOLD("456")
	}

	FRAGMENT("456"){
		a -> UNFOLD("789") -> GO(a, "123")
	}
	FRAGMENT("789"){
		a -> UNFOLD("10")
	}
	FRAGMENT("10"){
		a -> GO(a, "123")
	}`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "p", "a", 1)
	}
	`
	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, fmt.Sprintf("selfcheck gorch gramer error\nstart `\"123\"` has duplicate goDirective, file=%s(17:7) name=123", f))
		assert.Nil(t, p)
	})
}
