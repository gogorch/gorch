package engine

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetryDirective(t *testing.T) {
	t.Run("retry_success", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				RETRY(MAX_TIMES=3){
					RetryFailThenPassOp(failTimes=2)
				}
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(3), val.Val)
			},
		})
	})

	t.Run("retry_exhausted", func(t *testing.T) {
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				RETRY(MAX_TIMES=2){
					RetryFailThenPassOp(failTimes=3)
				}
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Equal(t, errors.New("retry failed"), res)
				assert.Equal(t, int64(2), val.Val)
			},
		})
	})

	t.Run("retry_with_interval", func(t *testing.T) {
		start := time.Now()
		MyTestRun(t, TestValus{
			name: "test",
			g: `START("test"){
				NO_CHECK_MISS()
				RETRY(MAX_TIMES=2, interval=20ms){
					RetryFailThenPassOp(failTimes=1)
				}
			}`,
			assertFunc: func(res error, val *BeChangeValue) {
				assert.Nil(t, res)
				assert.Equal(t, int64(2), val.Val)
				assert.GreaterOrEqual(t, time.Since(start), 18*time.Millisecond)
			},
		})
	})
}
