package cflag_test

import (
	"testing"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFlagArg_check(t *testing.T) {
	c := cflag.New()

	assert.PanicsErrMsg(t, func() {
		c.AddArg("", "", false, nil)
	}, "cflag: arg#0 name cannot be empty")

	assert.PanicsErrMsg(t, func() {
		c.AddArg("test", "", true, "def")
	}, "cflag: cannot set default value for 'required' arg: test")

	assert.PanicsMsg(t, func() {
		c.Arg("not-exist")
	}, "cflag: get not binding arg 'not-exist'")

	c.AddArg("test", "", true, nil)
	c.AddArg("test2", "arg2 desc", true, nil)

	arg := c.Arg("test")
	assert.Eq(t, 0, arg.Index)
	assert.Eq(t, "no description", arg.Desc)
	assert.Len(t, c.RemainArgs(), 0)

	arg = c.Arg("test2")
	assert.Eq(t, "arg2 desc", arg.Desc)
	assert.Eq(t, 1, arg.Index)
}

func TestFlagArg_HelpDesc(t *testing.T) {
	arg := cflag.NewArg("test", "desc for arg", false)
	assert.Eq(t, "Desc for arg", arg.HelpDesc())

	arg.Required = true
	assert.Eq(t, "<red>*</>Desc for arg", arg.HelpDesc())
}
