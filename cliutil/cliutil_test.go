package cliutil_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/cliutil"
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

	ret, err = cliutil.QuickExec("echo OK2")
	assert.NoError(t, err)
	assert.Equal(t, "OK2", strings.TrimSpace(ret))
}

func TestShellExec(t *testing.T) {
	ret, err := cliutil.ShellExec("echo OK")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)

	ret, err = cliutil.ShellExec("echo OK", "bash")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)
}
