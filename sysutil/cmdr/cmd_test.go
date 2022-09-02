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

	c.OnBefore(func(c *cmdr.Cmd) {
		assert.Eq(t, "ls -l -h -a ./", c.Cmdline())
	})

	out := c.SafeOutput()
	fmt.Println(out)
	assert.NotEmpty(t, out)
	assert.NotEmpty(t, cmdr.OutputLines(out))
	assert.NotEmpty(t, cmdr.FirstLine(out))
}
