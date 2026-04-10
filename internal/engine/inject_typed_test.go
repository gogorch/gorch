package engine

import (
	"context"
	"testing"

	"github.com/gogorch/gorch/internal/lang/iparser"
	"github.com/stretchr/testify/assert"
)

func TestInjectTyped(t *testing.T) {
	clean()
	defer clean()
	registerTestOperator(t)

	p, err := iparser.Parse(`START("test"){ NO_CHECK_MISS() ChangeValueOP(val=7) }`)
	assert.Nil(t, err)

	eng, err := New(p)
	assert.Nil(t, err)
	exe := eng.Start("test")
	assert.NotNil(t, exe)

	val := &BeChangeValue{Val: 0}
	assert.Nil(t, InjectTyped(exe, &val))
	assert.Nil(t, InjectTyped(exe, &t))
	assert.Nil(t, exe.Execute(context.Background()))
	assert.Equal(t, int64(7), val.Val)
}
