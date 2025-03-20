package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOperatorState(t *testing.T) {
	t.Run("duplicate status code", func(t *testing.T) {
		os := NewOperatorStates("test")
		os.Fatal(1, "test")
		assert.PanicsWithError(t, `"test" status duplicate code: 1`, func() {
			os.Fatal(1, "test")
		})
	})
}
