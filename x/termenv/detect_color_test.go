package termenv_test

import (
	"runtime"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/termenv"
)

func TestCommon(t *testing.T) {
	termenv.ForceEnableColor()
	defer termenv.RevertColorSupport()

	assert.NotEmpty(t, termenv.TermColorLevel())
	assert.True(t, termenv.IsSupportColor())
	assert.True(t, termenv.IsSupport256Color())
	assert.False(t, termenv.NoColor())
	assert.False(t, termenv.IsSupportTrueColor())

	termenv.DisableColor()
	assert.True(t, termenv.NoColor())
	assert.False(t, termenv.IsSupportColor())
}

func TestDetectColorLevel(t *testing.T) {
	is := assert.New(t)
	defer termenv.RevertColorSupport()

	// "COLORTERM=truecolor"
	testutil.MockOsEnvByText("COLORTERM=truecolor", func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColorTrue, level)
		is.True(termenv.IsSupportColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportTrueColor())
	})

	// "FORCE_COLOR=on"
	testutil.MockOsEnvByText("FORCE_COLOR=on", func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor16, level)
		is.True(termenv.IsSupportColor())
		is.False(termenv.IsSupport256Color())
		is.False(termenv.IsSupportTrueColor())
	})

	// TERMINAL_EMULATOR=JetBrains-JediTerm
	testutil.MockOsEnvByText(`
TERM=xterm-256color
TERMINAL_EMULATOR=JetBrains-JediTerm
ZSH_TMUX_TERM=screen-256color
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColorTrue, level)
		is.True(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})
}

func TestDetectColorLevel_unix(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows")
		return
	}
	is := assert.New(t)
	defer termenv.RevertColorSupport()

	// no TERM env
	testutil.MockOsEnvByText("NO=none", func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColorNone, level)
		is.False(termenv.IsSupportTrueColor())
		is.False(termenv.IsSupport256Color())
		is.False(termenv.IsSupportColor())
	})

	testutil.MockOsEnvByText("TERM=not-exist-value", func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor16, level)
		is.False(termenv.IsSupportTrueColor())
		is.False(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	testutil.MockOsEnvByText("TERM=xterm", func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	testutil.MockOsEnvByText("TERM=screen-256color", func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	testutil.MockOsEnvByText("TERM=not-exist-256color", func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	testutil.MockOsEnvByText("WSL_DISTRO_NAME=Debian", func() {
		termenv.SetDebugMode(true)
		level := termenv.DetectColorLevel()
		termenv.SetDebugMode(false)
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColorNone, level)
		is.False(termenv.IsSupportTrueColor())
		is.False(termenv.IsSupport256Color())
		is.False(termenv.IsSupportColor())
	})

	// TERM_PROGRAM=Terminus
	testutil.MockOsEnvByText(`
TERMINUS_PLUGINS=
TERM=xterm-256color
TERM_PROGRAM=Terminus
ZSH_TMUX_TERM=screen-256color
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColorTrue, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// -------- tests on macOS ---------

	// TERM_PROGRAM=Apple_Terminal
	testutil.MockOsEnvByText(`
TERM_PROGRAM=Apple_Terminal
TERM=xterm-256color
TERM_PROGRAM_VERSION=433
TERM_SESSION_ID=F17907FE-DCA5-488D-829B-7AFA8B323753
ZSH_TMUX_TERM=screen-256color
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		// fmt.Println(os.Environ())
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app
	testutil.MockOsEnvByText(`
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=3.4.5beta1
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
TERM=xterm-256color
ZSH_TMUX_TERM=screen-256color
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColorTrue, level)
		is.True(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app invalid version
	testutil.MockOsEnvByText(`
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=xx.beta
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
TERM=xterm-256color
ZSH_TMUX_TERM=screen-256color
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app no version env
	testutil.MockOsEnvByText(`
ITERM_PROFILE=Default
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
TERM=xterm-256color
ZSH_TMUX_TERM=screen-256color
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// -------- tests on linux ---------
}

func TestDetectColorLevel_screen(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skip on windows")
		return
	}
	is := assert.New(t)
	defer termenv.RevertColorSupport()

	// COLORTERM=truecolor
	testutil.MockOsEnvByText(`
TERM=screen
COLORTERM=truecolor
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// TERM_PROGRAM=Apple_Terminal use screen
	testutil.MockOsEnvByText(`
TERM_PROGRAM=Apple_Terminal
TERM=screen
TERM_PROGRAM_VERSION=433
TERM_SESSION_ID=F17907FE-DCA5-488D-829B-7AFA8B323753
ZSH_TMUX_TERM=screen-256color
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		// fmt.Println(os.Environ())
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app use screen
	testutil.MockOsEnvByText(`
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
LC_TERMINAL_VERSION=3.4.5beta1
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=3.4.5beta1
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
ZSH_TMUX_TERM=screen
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// TERM_PROGRAM=Terminus use screen
	testutil.MockOsEnvByText(`
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
TERMINUS_PLUGINS=
TERM_PROGRAM=Terminus
ZSH_TMUX_TERM=screen
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})

	// TERMINAL_EMULATOR=JetBrains-JediTerm use screen
	testutil.MockOsEnvByText(`
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
TERMINAL_EMULATOR=JetBrains-JediTerm
ZSH_TMUX_TERM=screen
`, func() {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)

		is.Equal(termenv.TermColor256, level)
		is.False(termenv.IsSupportTrueColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportColor())
	})
}

func TestDetectColorLevel_windows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("skip on NOT windows")
		return
	}

	is := assert.New(t)
	defer termenv.RevertColorSupport()

	// ConEmuANSI
	testutil.MockEnvValue("ConEmuANSI", "ON", func(_ string) {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)
		is.Equal(termenv.TermColorTrue, level)

		is.True(termenv.IsSupportColor())
		is.True(termenv.IsSupport256Color())
		is.True(termenv.IsSupportTrueColor())
	})

	// WSL_DISTRO_NAME=Debian
	testutil.MockEnvValue("WSL_DISTRO_NAME", "Debian", func(_ string) {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)

		is.True(termenv.IsSupportColor())
	})

	// ANSICON
	testutil.MockEnvValue("ANSICON", "189x2000 (189x43)", func(_ string) {
		level := termenv.DetectColorLevel()
		termenv.SetColorLevel(level)

		is.True(termenv.IsSupportColor())
		// is.Equal(termenv.TermColor256, level)
		// is.Equal("TERM=xterm-256color", SupColorMark())
	})
}
