package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipOperator(t *testing.T) {
	t.Run("skip", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
				SKIP(SkipOp(skip=true)) -> ChangeValueOP(val=2) 
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(0), val.Val)
			},
		})
	})
	t.Run("no_skip", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){ 
				NO_CHECK_MISS() 
				SKIP(SkipOp(skip=false)) -> ChangeValueOP(val=2) 
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})
}
