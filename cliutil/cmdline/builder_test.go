package cmdline_test

import (
	"testing"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/gookit/goutil/testutil/assert"
)

func TestLineBuild(t *testing.T) {
	s := cmdline.LineBuild("myapp", []string{"-a", "val0", "arg0"})

	assert.Eq(t, "myapp -a val0 arg0", s)

	// case: empty string
	b := cmdline.NewBuilder("myapp", "-a", "")

	assert.Eq(t, 11, b.Len())
	assert.Eq(t, `myapp -a ""`, b.String())

	b.Reset()
	assert.Eq(t, 0, b.Len())

	// case: add first
	b.AddArg("myapp")
	assert.Eq(t, `myapp`, b.String())

	b.AddArgs("-a", "val0")
	assert.Eq(t, "myapp -a val0", b.String())

	// case: contains `"`
	b.Reset()
	b.AddArgs("myapp", "-a", `"val0"`)
	assert.Eq(t, `myapp -a '"val0"'`, b.String())
	b.Reset()
	b.AddArgs("myapp", "-a", `the "val0" of option`)
	assert.Eq(t, `myapp -a 'the "val0" of option'`, b.String())

	// case: contains `'`
	b.Reset()
	b.AddArgs("myapp", "-a", `'val0'`)
	assert.Eq(t, `myapp -a "'val0'"`, b.String())
	b.Reset()
	b.AddArgs("myapp", "-a", `the 'val0' of option`)
	assert.Eq(t, `myapp -a "the 'val0' of option"`, b.String())
}

func TestLineBuild_hasQuote(t *testing.T) {
	line := "git log --pretty=format:'one two three'"
	args := cmdline.ParseLine(line)
	// dump.P(args)
	assert.Eq(t, line, cmdline.LineBuild("", args))
}
