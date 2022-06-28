package stdutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/stdutil"
	"github.com/stretchr/testify/assert"
)

func TestPanicIfErr(t *testing.T) {
	stdutil.DiscardE(nil)

	stdutil.PanicIf(nil)
	stdutil.PanicIfErr(nil)

	assert.Panics(t, func() {
		stdutil.PanicIf(errorx.Raw("a error"))
	})
	assert.Panics(t, func() {
		stdutil.PanicIfErr(errorx.Raw("a error"))
	})
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

func TestGoVersion(t *testing.T) {
	assert.NotEmpty(t, stdutil.GoVersion())
}
