package cmdline_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestLineParser_Parse(t *testing.T) {
	args := cmdline.NewParser(`./app top sub -a ddd --xx "msg"`).Parse()
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = cmdline.NewParser(`./app top sub -a ddd --xx "abc
def"`).Parse()
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "abc\ndef", args[6])

	args = cmdline.NewParser(`./app top sub -a ddd --xx "abc
def ghi"`).Parse()
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "abc\ndef ghi", args[6])

	args = cmdline.NewParser(`./app top sub --msg "has multi words"`).Parse()
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has multi words", args[4])

	args = cmdline.NewParser(`./app top sub --msg "has inner 'quote'"`).Parse()
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has inner 'quote'", args[4])

	args = cmdline.NewParser(`./app top sub --msg "'has' inner quote"`).Parse()
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "'has' inner quote", args[4])

	args = cmdline.NewParser(`./app top sub --msg "has inner 'quote' words"`).Parse()
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has inner 'quote' words", args[4])

	args = cmdline.ParseLine(`./app top sub --msg "has 'inner quote' words"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has 'inner quote' words", args[4])

	args = cmdline.ParseLine(`./app top sub --msg "has 'inner quote words'"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has 'inner quote words'", args[4])

	args = cmdline.ParseLine(`./app top sub --msg "'has inner quote' words"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "'has inner quote' words", args[4])

	args = cmdline.ParseLine(" ")
	assert.Len(t, args, 0)

	args = cmdline.ParseLine("./app")
	assert.Len(t, args, 1)
}

func TestParseLine_errLine(t *testing.T) {
	// exception line string.
	args := cmdline.NewParser(`./app top sub -a ddd --xx msg"`).Parse()
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg text`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg text", args[6])

	args = cmdline.ParseLine(`./app top sub -a ddd --xx "msg "text"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg text", args[6])
}

func TestLineParser_BinAndArgs(t *testing.T) {
	p := cmdline.NewParser("git status")
	b, a := p.BinAndArgs()
	assert.Equal(t, "git", b)
	assert.Equal(t, "status", strings.Join(a, " "))

	p = cmdline.NewParser("git")
	b, a = p.BinAndArgs()
	assert.Equal(t, "git", b)
	assert.Empty(t, a)
}
