package sysutil_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gookit/goutil/sysutil"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSysenv_common(t *testing.T) {
	ss := sysutil.EnvPaths()
	assert.NotEmpty(t, ss)
	assert.NotEmpty(t, sysutil.Environ())
	assert.NotEmpty(t, sysutil.EnvMapWith(nil))
	assert.NotEmpty(t, sysutil.EnvMapWith(map[string]string{"NEW_KEY": "value"}))

	ss = sysutil.SearchPath("go", 3)
	assert.NotEmpty(t, ss)
	// dump.P(ss)
	ss = sysutil.SearchPath("o", 3)
	assert.NotEmpty(t, ss)

	assert.Empty(t, sysutil.Getenv("NOT_EXISTS_ENV"))
	assert.Equal(t, "defVal", sysutil.Getenv("NOT_EXISTS_ENV", "defVal"))
}

func TestCurrentShell(t *testing.T) {
	path := sysutil.CurrentShell(true)

	if path != "" {
		assert.NotEmpty(t, path)
		assert.True(t, sysutil.HasShellEnv(path))

		path = sysutil.CurrentShell(false)
		assert.NotEmpty(t, path)
	}

	assert.NotEmpty(t, sysutil.Hostname())
	assert.True(t, sysutil.IsShellSpecialVar('$'))
	assert.True(t, sysutil.IsShellSpecialVar('@'))
	assert.False(t, sysutil.IsShellSpecialVar('a'))
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
		assert.NoErr(t, os.Unsetenv("MSYSTEM"))
		assert.False(t, sysutil.IsMSys())
	})
}

func TestIsConsole(t *testing.T) {
	is := assert.New(t)

	// IsConsole
	is.True(sysutil.IsConsole(os.Stdin))
	is.True(sysutil.IsConsole(os.Stdout))
	is.True(sysutil.IsConsole(os.Stderr))
	is.False(sysutil.IsConsole(&bytes.Buffer{}))
	ff, err := os.OpenFile("sysutil.go", os.O_WRONLY, 0)
	is.NoErr(err)
	is.False(sysutil.IsConsole(ff))
}

func TestFindExecutable(t *testing.T) {
	path, err := sysutil.Executable("echo")
	assert.NoErr(t, err)
	assert.NotEmpty(t, path)

	path, err = sysutil.FindExecutable("echo")
	assert.NoErr(t, err)
	assert.NotEmpty(t, path)

	assert.True(t, sysutil.HasExecutable("echo"))
}
