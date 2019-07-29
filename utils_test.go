package goutil_test

import (
	"testing"

	"github.com/gookit/goutil"
	"github.com/stretchr/testify/assert"
)

func TestFuncName(t *testing.T) {
	name := goutil.FuncName(goutil.PkgName)
	assert.Equal(t, "github.com/gookit/goutil.PkgName", name)
}

func TestPkgName(t *testing.T) {
	name := goutil.PkgName()
	assert.Equal(t, "goutil", name)
}
