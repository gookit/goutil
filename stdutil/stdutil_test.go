package stdutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/stdutil"
	"github.com/stretchr/testify/assert"
)

func TestFuncName(t *testing.T) {
	name := stdutil.FuncName(stdutil.PkgName)
	assert.Equal(t, "github.com/gookit/stdutil.PkgName", name)

	name = stdutil.FuncName(stdutil.PanicIfErr)
	assert.Equal(t, "github.com/gookit/stdutil.PanicIfErr", name)
}

func TestPkgName(t *testing.T) {
	name := stdutil.PkgName(stdutil.FuncName(stdutil.PanicIfErr))
	assert.Equal(t, "github.com/gookit/stdutil", name)
}

func TestPanicIfErr(t *testing.T) {
	stdutil.PanicIfErr(nil)
}

func TestPanicf(t *testing.T) {
	assert.Panics(t, func() {
		stdutil.Panicf("hi %s", "inhere")
	})
}

func TestGetCallStacks(t *testing.T) {
	msg := stdutil.GetCallStacks(false)
	fmt.Println(string(msg))

	fmt.Println("-------------full stacks-------------")
	msg = stdutil.GetCallStacks(true)
	fmt.Println(string(msg))
}
