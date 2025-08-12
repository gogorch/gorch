package iparser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gogorch/gorch/internal/lang/ast"
	"github.com/stretchr/testify/assert"
)

func RunMyTest(t *testing.T, g string, reg string, f func(*ast.Primary, string, error)) {
	t.Run("BySplitFileInDir", func(t *testing.T) {
		p, e := BySplitFileInDir(t, g, reg)
		f(p, "logic.gorch", e)
	})
	t.Run("ByOneFileInDir", func(t *testing.T) {
		p, e := ByOneFileInDir(t, g, reg)
		f(p, "logic.gorch", e)
	})
	t.Run("ByString", func(t *testing.T) {
		p, e := ByString(t, g, reg)
		f(p, "", e)
	})
}

func BySplitFileInDir(t *testing.T, g string, reg string) (*ast.Primary, error) {
	dir := t.TempDir()
	t.Cleanup(func() { assert.Nil(t, os.RemoveAll(dir)) })

	file := filepath.Join(dir, "logic.gorch")
	assert.Nil(t, os.WriteFile(file, []byte(g), 0755))
	regFile := filepath.Join(dir, "reg.gorch")
	assert.Nil(t, os.WriteFile(regFile, []byte(reg), 0755))

	return ParseDirectory(dir)
}

func ByOneFileInDir(t *testing.T, g string, reg string) (*ast.Primary, error) {
	dir := t.TempDir()
	t.Cleanup(func() { assert.Nil(t, os.RemoveAll(dir)) })
	file := filepath.Join(dir, "logic.gorch")
	assert.Nil(t, os.WriteFile(file, []byte(g+"\n"+reg), 0755))
	return ParseDirectory(dir)
}

func ByString(t *testing.T, g string, reg string) (*ast.Primary, error) {
	return Parse(g + "\n" + reg)
}
