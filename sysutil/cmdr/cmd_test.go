package cmdr_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/sysutil/cmdr"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewCmd(t *testing.T) {
	c := cmdr.NewCmd("ls").
		WithArg("-l").
		WithArgs([]string{"-h"}).
		AddArg("-a").
		AddArgf("%s", "./")

	assert.Eq(t, "ls", c.BinName())
	assert.Eq(t, "ls", c.IDString())
	assert.StrContains(t, "ls", c.BinOrPath())
	assert.NotContains(t, c.OnlyArgs(), "ls")

	c.OnBefore(func(c *cmdr.Cmd) {
		assert.Eq(t, "ls -l -h -a ./", c.Cmdline())
	})

	// SafeOutput
	out := c.SafeOutput()
	fmt.Println(out)
	assert.NotEmpty(t, out)
	assert.NotEmpty(t, cmdr.OutputLines(out))
	assert.NotEmpty(t, cmdr.FirstLine(out))

	c.ResetArgs()
	assert.Len(t, c.Args, 1)
	assert.Empty(t, c.OnlyArgs())
}

func TestCmdRun(t *testing.T) {
	t.Run("AllOutput", func(t *testing.T) {
		c := cmdr.NewCmd("echo", "OK")
		assert.Eq(t, "OK", c.Args[1])
		output, err := c.AllOutput()
		assert.NoErr(t, err)
		assert.Eq(t, "OK\n", output)
	})
}
