package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	t.Run("wrap_next_change_value", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
					NO_CHECK_MISS() 

					SWITCH(SwitchOp(case=["changeValueTo1"])) {
						CASE "changeValueTo1" => UNFOLD("fragment"),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
					}
				}
				FRAGMENT("fragment"){
					WrapAndChangeValOp(val=1) | ChangeValueOP(val=20, if=1)
				}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(20), val.Val)
			},
		})
	})

	t.Run("wrap_no_next", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
					NO_CHECK_MISS() 

					SWITCH(SwitchOp(case=["changeValueTo1"])) {
						CASE "changeValueTo1" => UNFOLD("fragment"),
						CASE "changeValueTo2" => ChangeValueOP(val=2)
					}
				}
				FRAGMENT("fragment"){
					WrapAndChangeValOp(noNext=true, val=11) | ChangeValueOP(val=20, if=1)
				}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(11), val.Val)
			},
		})
	})
}
