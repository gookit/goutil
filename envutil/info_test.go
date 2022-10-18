package envutil_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestOS(t *testing.T) {
	if isw := envutil.IsWin(); isw {
		assert.True(t, isw)
		assert.False(t, envutil.IsMac())
		assert.False(t, envutil.IsLinux())
	}

	if ism := envutil.IsMac(); ism {
		assert.True(t, ism)
		assert.False(t, envutil.IsWin())
		assert.False(t, envutil.IsLinux())
		assert.False(t, envutil.IsMSys())
		assert.False(t, envutil.IsWSL())
		assert.False(t, envutil.IsWindows())
	}

	if isl := envutil.IsLinux(); isl {
		assert.True(t, isl)
		assert.False(t, envutil.IsMac())
		assert.False(t, envutil.IsWin())
		assert.False(t, envutil.IsWindows())
		assert.False(t, envutil.IsMSys())
		assert.False(t, envutil.IsWSL())
	}
}

func TestIsConsole(t *testing.T) {
	is := assert.New(t)

	is.True(envutil.IsConsole(os.Stdout))
	is.False(envutil.IsConsole(&bytes.Buffer{}))

	is.IsType(true, envutil.IsTerminal(os.Stdout.Fd()))
	is.IsType(true, envutil.StdIsTerminal())

	is.True(envutil.HasShellEnv("sh"))
}

func TestIsGithubActions(t *testing.T) {
	is := assert.New(t)

	testutil.MockEnvValue("GITHUB_ACTIONS", "true", func(nv string) {
		is.Eq("true", nv)
		is.True(envutil.IsGithubActions())
	})
}

func TestIsSupportColor(t *testing.T) {
	is := assert.New(t)

	// clear all OS env
	testutil.ClearOSEnv()
	defer testutil.RevertOSEnv()

	// IsSupport256Color
	is.False(envutil.IsSupport256Color())

	// ConEmuANSI
	testutil.MockEnvValue("ConEmuANSI", "ON", func(nv string) {
		is.Eq("ON", nv)
		is.True(envutil.IsSupportColor())
	})

	// ANSICON
	testutil.MockEnvValue("ANSICON", "189x2000 (189x43)", func(_ string) {
		is.True(envutil.IsSupportColor())
	})

	// "COLORTERM=truecolor"
	testutil.MockEnvValue("COLORTERM", "truecolor", func(_ string) {
		is.True(envutil.IsSupportTrueColor())
	})

	// TERM
	testutil.MockEnvValue("TERM", "screen-256color", func(_ string) {
		is.True(envutil.IsSupportColor())
	})

	// TERM
	testutil.MockEnvValue("TERM", "", func(_ string) {
		is.False(envutil.IsSupportColor())
	})

	is.NoErr(os.Setenv("TERM", "xterm-vt220"))
	is.True(envutil.IsSupportColor())
}
