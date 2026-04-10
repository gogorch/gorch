package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoopDirective(t *testing.T) {
	t.Run("max_iter", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				LOOP(MAX_ITER=3){
					AddValueOp(delta=1)
				}
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(3), val.Val)
			},
		})
	})

	t.Run("until_break_loop", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				LOOP(MAX_ITER=10){
					UNTIL(UntilValueReachedOp(target=3))
					-> AddValueOp(delta=1)
				}
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(3), val.Val)
			},
		})
	})

	t.Run("break_directive", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				LOOP(MAX_ITER=10){
					AddValueOp(delta=1)
					-> BREAK()
				}
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(1), val.Val)
			},
		})
	})

	t.Run("break_outside_loop", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				BREAK()
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, errNotLoopDirective, res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})
}
