package cmdline_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestLineParser_Parse(t *testing.T) {
	args := cmdline.NewParser(`./app top sub -a ddd --xx "msg"`).Parse()
	assert.Len(t, args, 7)
	assert.Eq(t, "msg", args[6])

	args = cmdline.ParseLine(" ")
	assert.Len(t, args, 0)

	args = cmdline.ParseLine("./app")
	assert.Len(t, args, 1)

	p := cmdline.NewParser("./app sub ${A_ENV_VAR}")
	p.WithParseEnv()
	assert.True(t, p.ParseEnv)

	testutil.MockEnvValue("A_ENV_VAR", "env-value", func(nv string) {
		bin, args := p.BinAndArgs()
		assert.Len(t, args, 2)
		assert.Eq(t, "./app", bin)
		assert.Eq(t, "env-value", args[1])

		assert.NotEmpty(t, p.NewExecCmd())
	})

	p = cmdline.NewParser("./app sub ${A_ENV_VAR2}")
	testutil.MockEnvValue("A_ENV_VAR2", "env-value2", func(nv string) {
		args := p.AlsoEnvParse()
		assert.Len(t, args, 3)
		assert.Eq(t, "env-value2", args[2])
	})
}

func TestParseLine_Parse_withQuote(t *testing.T) {
	tests := []struct {
		line  string
		argN  int
		index int
		value string
	}{
		{
			line: `./app top sub -a ddd --xx "abc
def"`,
			argN: 7, index: 6, value: "abc\ndef",
		},
		{
			line: `./app top sub -a ddd --xx "abc
def ghi"`,
			argN: 7, index: 6, value: "abc\ndef ghi",
		},
		{
			line: `./app top sub --msg "has multi words"`,
			argN: 5, index: 4, value: "has multi words",
		},
		{
			line: `./app top sub --msg "has inner 'quote'"`,
			argN: 5, index: 4, value: "has inner 'quote'",
		},
		{
			line: `./app top sub --msg "'has' inner quote"`,
			argN: 5, index: 4, value: "'has' inner quote",
		},
		{
			line: `./app top sub --msg "has inner 'quote' words"`,
			argN: 5, index: 4, value: "has inner 'quote' words",
		},
		{
			line: `./app top sub --msg "has 'inner quote' words"`,
			argN: 5, index: 4, value: "has 'inner quote' words",
		},
		{
			line: `./app top sub --msg "has 'inner quote words'"`,
			argN: 5, index: 4, value: "has 'inner quote words'",
		},
		{
			line: `./app top sub --msg "'has inner quote' words"`,
			argN: 5, index: 4, value: "'has inner quote' words",
		},
	}

	for _, tt := range tests {
		args := cmdline.NewParser(tt.line).Parse()
		assert.Len(t, args, tt.argN)
		assert.Eq(t, tt.value, args[tt.index])
	}
}

func TestParseLine_longLine(t *testing.T) {
	line := "git log --pretty=format:'one two three'"
	args := cmdline.ParseLine(line)
	assert.Len(t, args, 3)
	assert.Eq(t, "--pretty=format:'one two three'", args[2])

	line = `git log --pretty=format:"one two three""`
	args = cmdline.ParseLine(line)
	assert.Len(t, args, 3)
	assert.Eq(t, `--pretty=format:"one two three""`, args[2])

	line = "git log --color --graph --pretty=format:'%Cred%h%Creset:%C(ul yellow)%d%Creset %s (%Cgreen%cr%Creset, %C(bold blue)%an%Creset)' --abbrev-commit -10"
	args = cmdline.ParseLine(line)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Eq(t, "--graph", args[3])
	assert.Eq(t, "--abbrev-commit", args[5])
}

func TestParseLine_errLine(t *testing.T) {
	// exception line string.
	args := cmdline.NewParser(`./app top sub -a ddd --xx msg"`).Parse()
	assert.Len(t, args, 7)
	assert.Eq(t, "msg\"", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg`)
	assert.Len(t, args, 7)
	assert.Eq(t, "msg", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg text`)
	assert.Len(t, args, 7)
	assert.Eq(t, "msg text", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg "text"`)
	assert.Len(t, args, 7)
	assert.Eq(t, "msg \"text", args[6])
}

func TestLineParser_BinAndArgs(t *testing.T) {
	p := cmdline.NewParser("git status")
	b, a := p.BinAndArgs()
	assert.Eq(t, "git", b)
	assert.Eq(t, "status", strings.Join(a, " "))

	p = cmdline.NewParser("git")
	b, a = p.BinAndArgs()
	assert.Eq(t, "git", b)
	assert.Empty(t, a)
}
