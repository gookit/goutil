package cflag_test

import (
	"testing"

	"github.com/gookit/goutil/cflag"
	"github.com/stretchr/testify/assert"
)

func TestFlagArg_check(t *testing.T) {
	c := cflag.New()

	assert.PanicsWithError(t, "cflag: arg#0 name cannot be empty", func() {
		c.AddArg("", "", false, nil)
	})

	assert.PanicsWithError(t, "cflag: cannot set default value for 'required' arg: test", func() {
		c.AddArg("test", "", true, "def")
	})

	assert.PanicsWithValue(t, "cflag: get not binding arg 'not-exist'", func() {
		c.Arg("not-exist")
	})

	c.AddArg("test", "", true, nil)

	arg := c.Arg("test")
	assert.Equal(t, "no description", arg.Desc)
	assert.Len(t, c.RemainArgs(), 0)
}

func TestFlagArg_HelpDesc(t *testing.T) {
	arg := cflag.NewArg("test", "desc for arg", false)
	assert.Equal(t, "Desc for arg", arg.HelpDesc())

	arg.Required = true
	assert.Equal(t, "<red>*</>Desc for arg", arg.HelpDesc())
}
