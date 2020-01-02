package envutil_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
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
	}

	if isl := envutil.IsLinux(); isl {
		assert.True(t, isl)
		assert.False(t, envutil.IsMac())
		assert.False(t, envutil.IsWin())
		assert.False(t, envutil.IsMSys())
	}
}

func TestIsConsole(t *testing.T) {
	is := assert.New(t)

	is.True(envutil.IsConsole(os.Stdout))
	is.False(envutil.IsConsole(&bytes.Buffer{}))

	is.True(envutil.HasShellEnv("sh"))
}

func TestIsSupportColor(t *testing.T) {
	is := assert.New(t)

	// IsSupport256Color
	oldVal := os.Getenv("TERM")
	_ = os.Unsetenv("TERM")
	// is.False(envutil.IsSupportColor())
	is.False(envutil.IsSupport256Color())

	// ConEmuANSI
	testutil.MockEnvValue("ConEmuANSI", "ON", func(nv string) {
		is.Equal("ON", nv)
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

	is.NoError(os.Setenv("TERM", "xterm-vt220"))
	is.True(envutil.IsSupportColor())
	// revert
	if oldVal != "" {
		is.NoError(os.Setenv("TERM", oldVal))
	} else {
		is.NoError(os.Unsetenv("TERM"))
	}
}
