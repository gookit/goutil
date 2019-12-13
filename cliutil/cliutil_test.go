package cliutil_test

import (
	"testing"

	"github.com/gookit/goutil/cliutil"
	"github.com/stretchr/testify/assert"
)

func TestCurrentShell(t *testing.T) {
	path := cliutil.CurrentShell(true)
	assert.NotEmpty(t, path)

	if path != "" {
		path = cliutil.CurrentShell(false)
		assert.NotEmpty(t, path)
	}

	assert.True(t, cliutil.HasShellEnv("sh"))
}

func TestExecCmd(t *testing.T) {
	ret, err := cliutil.ExecCmd("echo", []string{"OK"})
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)

	ret, err = cliutil.ExecCommand("echo", []string{"OK1"})
	assert.NoError(t, err)
	assert.Equal(t, "OK1\n", ret)

	ret, err = cliutil.QuickExec("echo OK")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)
}

func TestShellExec(t *testing.T) {
	ret, err := cliutil.ShellExec("echo OK")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)

	ret, err = cliutil.ShellExec("echo OK", "bash")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)
}
