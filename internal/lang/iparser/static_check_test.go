package iparser

import (
	"testing"

	"github.com/gogorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

func TestStaticCheckOperatorMissNested(t *testing.T) {
	g := `
	START("name") {
		a -> SWITCH(route) {
			CASE "hit" => a
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nin start \"name\", operator \"route\" not register")
		assert.Nil(t, p)
	})
}

func TestStaticCheckLoopRequiredParam(t *testing.T) {
	g := `
	START("name") {
		LOOP(retry=3) {
			a
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": LOOP missing required argument MAX_ITER")
		assert.Nil(t, p)
	})
}

func TestStaticCheckRetryRequiredParam(t *testing.T) {
	g := `
	START("name") {
		RETRY(interval=1ms) {
			a
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": RETRY missing required argument MAX_TIMES")
		assert.Nil(t, p)
	})
}

func TestStaticCheckUnreachableBranch(t *testing.T) {
	g := `
	START("name") {
		LOOP(MAX_ITER=3) {
			a -> BREAK() -> b
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
		OPERATOR("gorch/ops", "b", 2)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": unreachable branch after BREAK()")
		assert.Nil(t, p)
	})
}

func TestStaticCheckDeadGo(t *testing.T) {
	g := `
	START("name") {
		GO(a, "task_a")
		-> a
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": GO(\"task_a\") not find matched WAIT directive")
		assert.Nil(t, p)
	})
}

func TestStaticCheckUnreachableBranchNestedSerialBracket(t *testing.T) {
	g := `
	START("name") {
		(a -> BREAK())
		-> b
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
		OPERATOR("gorch/ops", "b", 2)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": unreachable branch after BREAK()")
		assert.Nil(t, p)
	})
}

func TestStaticCheckUnreachableBranchUnfold(t *testing.T) {
	g := `
	START("name") {
		UNFOLD("f")
		-> b
	}

	FRAGMENT("f") {
		a -> BREAK()
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
		OPERATOR("gorch/ops", "b", 2)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": unreachable branch after BREAK()")
		assert.Nil(t, p)
	})
}

func TestStaticCheckUnreachableBranchSwitchAllBreak(t *testing.T) {
	g := `
	START("name") {
		SWITCH(route) {
			CASE "a" => BREAK(),
			CASE "b" => (a -> BREAK())
		}
		-> b
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
		OPERATOR("gorch/ops", "b", 2)
		OPERATOR("gorch/ops", "route", 3)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": unreachable branch after BREAK()")
		assert.Nil(t, p)
	})
}

func TestStaticCheckRequiredParamAliasPass(t *testing.T) {
	g := `
	START("name") {
		LOOP(maxIter=2) {
			a
		}
		-> RETRY(maxTimes=2, interval=1ms) {
			a
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.Nil(t, e)
		assert.NotNil(t, p)
	})
}

func TestStaticCheckLoopRequiredParamInvalidValue(t *testing.T) {
	g := `
	START("name") {
		LOOP(MAX_ITER=0) {
			a
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": LOOP missing required argument MAX_ITER")
		assert.Nil(t, p)
	})
}

func TestStaticCheckRetryRequiredParamInvalidValue(t *testing.T) {
	g := `
	START("name") {
		RETRY(MAX_TIMES=0) {
			a
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": RETRY missing required argument MAX_TIMES")
		assert.Nil(t, p)
	})
}

func TestStaticCheckDeadGoNestedTrace(t *testing.T) {
	g := `
	START("name") {
		TRACE("segment") {
			GO(a, "task_a") -> a
		}
		-> a
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": GO(\"task_a\") not find matched WAIT directive")
		assert.Nil(t, p)
	})
}

func TestStaticCheckDeadGoInUnfold(t *testing.T) {
	g := `
	START("name") {
		UNFOLD("f")
	}

	FRAGMENT("f") {
		GO(a, "task_a") -> a
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": GO(\"task_a\") not find matched WAIT directive")
		assert.Nil(t, p)
	})
}

func TestStaticCheckWaitMissingNestedRetry(t *testing.T) {
	g := `
	START("name") {
		RETRY(MAX_TIMES=2) {
			a(WAIT("task_a"))
		}
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "a", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": WAIT(\"task_a\") not find matched GO directive")
		assert.Nil(t, p)
	})
}

func TestStaticCheckUnreachableAfterRetryBreak(t *testing.T) {
	g := `
	START("name") {
		RETRY(MAX_TIMES=2) {
			BREAK()
		}
		-> b
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "b", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": unreachable branch after BREAK()")
		assert.Nil(t, p)
	})
}

func TestStaticCheckUnreachableAfterTraceBreak(t *testing.T) {
	g := `
	START("name") {
		TRACE("segment") {
			BREAK()
		}
		-> b
	}
	`
	reg := `
	REGISTER("github.com/gogorch") {
		OPERATOR("gorch/ops", "b", 1)
	}
	`

	RunMyTest(t, g, reg, func(p *ast.Primary, f string, e error) {
		assert.EqualError(t, e, "selfcheck gorch gramer error\nSTART \"name\": unreachable branch after BREAK()")
		assert.Nil(t, p)
	})
}
