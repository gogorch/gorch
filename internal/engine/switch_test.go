package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwitchOperatro(t *testing.T) {
	t.Run("no_switch_error", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
					SWITCH(SwitchOp(case=[])) {
						CASE "changeValueTo1" => ChangeValueOP(val=1),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
						CASE "changeValueTo3" => ChangeValueOP(val=3)
					}
			}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, `switch "SwitchOp" no matching switch case found`, res.Error())
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("switch_change_to_1", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
					SWITCH(SwitchOp(case=["changeValueTo1"])) {
						CASE "changeValueTo1" => ChangeValueOP(val=1),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
						CASE "changeValueTo3" => ChangeValueOP(val=3)
					}
			}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(1), val.Val)
			},
		})
	})

	t.Run("switch_2_cache", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
					SWITCH(SwitchOp(case=["changeValueTo2", "changeValueTo1"])) {
						CASE "changeValueTo1" => ChangeValueOP(val=1),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
						CASE "changeValueTo3" => ChangeValueOP(val=3)
					}
			}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
			},
		})
	})

	t.Run("switch_directive_not_switch_op", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
					SWITCH(NothingOp()) {
						CASE "changeValueTo1" => ChangeValueOP(val=1),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
						CASE "changeValueTo3" => ChangeValueOP(val=3)
					}
			}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, `switch "NothingOp" no matching switch case found`, res.Error())
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("switch_serial_change_to_1", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
					SWITCH(SwitchOp(case=["changeValueTo1"])) {
						CASE "changeValueTo1" => (
							ChangeValueOP(val=2) -> ChangeValueOP(val=1)
						),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
						CASE "changeValueTo3" => ChangeValueOP(val=3)
					}
			}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(1), val.Val)
			},
		})
	})

	t.Run("switch_duplicate_case", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
					SWITCH(SwitchOp(case=["changeValueTo1", "changeValueTo1"])) {
						CASE "changeValueTo1" => ChangeValueOP(val=1),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
						CASE "changeValueTo3" => ChangeValueOP(val=3)
					}
			}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.NotNil(t, res)
				assert.Equal(t, `duplicate switch case "changeValueTo1"`, res.Error())
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})

	t.Run("switch_not_found_case", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
					SWITCH(SwitchOp(case=["111"])) {
						CASE "changeValueTo1" => ChangeValueOP(val=1),
						CASE "changeValueTo2" => ChangeValueOP(val=2),
						CASE "changeValueTo3" => ChangeValueOP(val=3)
					}
			}
				`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, `not found switch case "111"`, res.Error())
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})
}
