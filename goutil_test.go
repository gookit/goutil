package goutil_test

import (
	"fmt"
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

func TestPanicIfErr(t *testing.T) {
	goutil.PanicIfErr(nil)
}

func TestGetCallStacks(t *testing.T) {
	msg := goutil.GetCallStacks(false)
	fmt.Println(string(msg))

	fmt.Println("-------------full stacks-------------")
	msg = goutil.GetCallStacks(true)
	fmt.Println(string(msg))
}
