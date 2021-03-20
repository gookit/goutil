package cliutil_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestCurrentShell(t *testing.T) {
	path := cliutil.CurrentShell(true)

	if path != "" {
		assert.NotEmpty(t, path)
		assert.True(t, cliutil.HasShellEnv(path))

		path = cliutil.CurrentShell(false)
		assert.NotEmpty(t, path)
	}
}

func TestExecCmd(t *testing.T) {
	ret, err := cliutil.ExecCmd("echo", []string{"OK"})
	assert.NoError(t, err)
	// *nix: "OK\n" win: "OK\r\n"
	assert.Equal(t, "OK", strings.TrimSpace(ret))

	ret, err = cliutil.ExecCommand("echo", []string{"OK1"})
	assert.NoError(t, err)
	assert.Equal(t, "OK1", strings.TrimSpace(ret))

	ret, err = cliutil.ExecLine("echo OK2")
	assert.NoError(t, err)
	assert.Equal(t, "OK2", strings.TrimSpace(ret))

	ret, err = cliutil.ExecLine("echo OK3")
	assert.NoError(t, err)
	assert.Equal(t, "OK3", strings.TrimSpace(ret))
}

func TestShellExec(t *testing.T) {
	ret, err := cliutil.ShellExec("echo OK")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)

	ret, err = cliutil.ShellExec("echo OK", "bash")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)
}

func TestLineBuild(t *testing.T) {
	s := cliutil.LineBuild("myapp", []string{"-a", "val0", "arg0"})
	assert.Equal(t, "myapp -a val0 arg0", s)

	s = cliutil.BuildLine("./myapp", []string{
		"-a", "val0",
		"-m", "this is message",
		"arg0",
	})
	assert.Equal(t, `./myapp -a val0 -m "this is message" arg0`, s)
}

func TestParseLine(t *testing.T) {
	args := cliutil.ParseLine(`./app top sub -a ddd --xx "msg"`)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = cliutil.String2OSArgs(`./app top sub --msg "has inner 'quote'"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has inner 'quote'", args[4])

	// exception line string.
	args = cliutil.ParseLine(`./app top sub -a ddd --xx msg"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = cliutil.StringToOSArgs(`./app top sub -a ddd --xx "msg "text"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg text", args[6])
}
