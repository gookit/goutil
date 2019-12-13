package sysutil_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestCurrentShell(t *testing.T) {
	path := sysutil.CurrentShell(true)
	assert.NotEmpty(t, path)

	if path != "" {
		path = sysutil.CurrentShell(false)
		assert.NotEmpty(t, path)
	}

	assert.True(t, sysutil.HasShellEnv("sh"))
}

func TestOS(t *testing.T) {
	if isw := sysutil.IsWin(); isw {
		assert.True(t, isw)
		assert.False(t, sysutil.IsMac())
		assert.False(t, sysutil.IsLinux())
	}

	if ism := sysutil.IsMac(); ism {
		assert.True(t, ism)
		assert.False(t, sysutil.IsWin())
		assert.False(t, sysutil.IsWindows())
		assert.False(t, sysutil.IsLinux())
		assert.False(t, sysutil.IsMSys())
	}

	if isl := sysutil.IsLinux(); isl {
		assert.True(t, isl)
		assert.False(t, sysutil.IsMac())
		assert.False(t, sysutil.IsWin())
		assert.False(t, sysutil.IsWindows())
		assert.False(t, sysutil.IsMSys())
	}

	// IsMSys
	testutil.MockEnvValue("MSYSTEM", "MINGW64", func(nv string) {
		assert.True(t, sysutil.IsMSys())
		// delete
		assert.NoError(t, os.Unsetenv("MSYSTEM"))
		assert.False(t, sysutil.IsMSys())
	})
}

func TestIsConsole(t *testing.T) {
	is := assert.New(t)

	is.True(sysutil.IsConsole(os.Stdout))
	is.False(sysutil.IsConsole(&bytes.Buffer{}))
}

func TestExecCmd(t *testing.T) {
	ret, err := sysutil.ExecCmd("echo", []string{"OK"})
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)

	ret, err = sysutil.QuickExec("echo OK")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)
}

func TestShellExec(t *testing.T) {
	ret, err := sysutil.ShellExec("echo OK")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)

	ret, err = sysutil.ShellExec("echo OK", "bash")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)
}

func TestProcessExists(t *testing.T) {
	pid := os.Getpid()

	assert.True(t, sysutil.ProcessExists(pid))
}
