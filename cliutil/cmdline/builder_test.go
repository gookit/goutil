package cmdline_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/gookit/goutil/testutil/assert"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", `""`},
		{" ", `" "`},
		{"one", `one`},
		{"one'", `one'`},
		{"one ", `"one "`},
		{" one ", `" one "`},
		{"one two", `"one two"`},
		{"one two three four", `"one two three four"`},
		{`the "val0" of option`, `'the "val0" of option'`},
		{`--pretty=format:'one two three'`, `--pretty=format:'one two three'`},
		{`--pretty=format:"one two three"`, `--pretty=format:"one two three"`},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, cmdline.Quote(tt.in))
	}

	fmt.Println(strconv.Quote(`one`))
	fmt.Println(strconv.Quote(`one two`))
	fmt.Println(strconv.Quote(`"one two" three four`))
	fmt.Println(strconv.Quote(`'one two' three four`))
}

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

	// addAny
	b.Reset()
	b.AddAny("myapp", "-a", "val0")
	assert.Eq(t, `myapp -a val0`, b.String())

	// case: contains `"`
	b.Reset()
	b.AddArgs("myapp", "-a", `"val0"`)
	assert.Eq(t, `myapp -a "val0"`, b.String())
	b.Reset()
	b.AddArgs("myapp", "-a", `the "val0" of option`)
	assert.Eq(t, `myapp -a 'the "val0" of option'`, b.String())

	// case: contains `'`
	b.Reset()
	b.AddArgs("myapp", "-a", `'val0'`)
	assert.Eq(t, `myapp -a 'val0'`, b.ResetGet())

	b.AddArgs("myapp", "-a", `the 'val0' of option`)
	assert.Eq(t, `myapp -a "the 'val0' of option"`, b.String())
	b.Reset()
}

func TestLineBuild_special(t *testing.T) {
	var b = cmdline.LineBuilder{}

	tests := []struct {
		args []string
		want string
	}{
		{
			[]string{"myapp", `one "two three"`},
			`myapp 'one "two three"'`,
		},
		{
			[]string{"myapp", `"one two" three`},
			`myapp '"one two" three'`,
		},
	}

	for _, tt := range tests {
		b.AddArray(tt.args)
		assert.Eq(t, tt.want, b.ResetGet())
	}
}

func TestLineBuild_complex(t *testing.T) {
	var b = cmdline.LineBuilder{}

	t.Run("case01", func(t *testing.T) {
		b.AddArgs("git", "log", `--pretty=format:"one two three"`)
		assert.Eq(t, `git log --pretty=format:"one two three"`, b.ResetGet())
	})

	t.Run("case02", func(t *testing.T) {
		b.AddArgs("git", "log", `--pretty`, `format:"one two three"`)
		assert.Eq(t, `git log --pretty format:"one two three"`, b.ResetGet())
	})
}

func TestLineBuild_hasQuote(t *testing.T) {
	line := "git log --pretty=format:'one two three'"
	args := cmdline.ParseLine(line)
	// dump.P(args)
	assert.Len(t, args, 3)
	assert.Eq(t, line, cmdline.LineBuild("", args))
}
