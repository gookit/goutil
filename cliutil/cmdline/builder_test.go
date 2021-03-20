package cmdline_test

import (
	"testing"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/stretchr/testify/assert"
)

func TestLineBuild(t *testing.T) {
	s := cmdline.LineBuild("myapp", []string{"-a", "val0", "arg0"})

	assert.Equal(t, "myapp -a val0 arg0", s)

	// case: empty string
	b := cmdline.NewBuilder("myapp", "-a", "")

	assert.Equal(t, 11, b.Len())
	assert.Equal(t, `myapp -a ""`, b.String())

	b.Reset()
	assert.Equal(t, 0, b.Len())

	// case: add first
	b.AddArg("myapp")
	assert.Equal(t, `myapp`, b.String())

	b.AddArgs("-a", "val0")
	assert.Equal(t, "myapp -a val0", b.String())

	// case: contains `"`
	b.Reset()
	b.AddArgs("myapp", "-a", `"val0"`)
	assert.Equal(t, `myapp -a '"val0"'`, b.String())
	b.Reset()
	b.AddArgs("myapp", "-a", `the "val0" of option`)
	assert.Equal(t, `myapp -a 'the "val0" of option'`, b.String())

	// case: contains `'`
	b.Reset()
	b.AddArgs("myapp", "-a", `'val0'`)
	assert.Equal(t, `myapp -a "'val0'"`, b.String())
	b.Reset()
	b.AddArgs("myapp", "-a", `the 'val0' of option`)
	assert.Equal(t, `myapp -a "the 'val0' of option"`, b.String())
}
