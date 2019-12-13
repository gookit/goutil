package envutil

import (
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/gookit/goutil/sysutil"
)

// IsWin system. linux windows darwin
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// IsMac system
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux system
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsConsole check out is console env. alias of the sysutil.IsConsole()
func IsConsole(out io.Writer) bool {
	return sysutil.IsConsole(out)
}

// IsMSys msys(MINGW64) env. alias of the sysutil.IsMSys()
func IsMSys() bool {
	return sysutil.IsMSys()
}

// HasShellEnv has shell env check.
// Usage:
// 	HasShellEnv("sh")
// 	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	return sysutil.HasShellEnv(shell)
}

// IsSupportColor check console is support color.
// supported: linux, mac, or windows's ConEmu, Cmder, putty, git-bash.exe
// not support: windows cmd, powerShell
func IsSupportColor() bool {
	// "TERM=xterm"  support color
	// "TERM=xterm-vt220" support color
	// "TERM=xterm-256color" support color
	// "TERM=cygwin" don't support color
	if strings.Contains(os.Getenv("TERM"), "xterm") {
		return true
	}

	// like on ConEmu software, e.g "ConEmuANSI=ON"
	if os.Getenv("ConEmuANSI") == "ON" {
		return true
	}

	// like on ConEmu software, e.g "ANSICON=189x2000 (189x43)"
	if os.Getenv("ANSICON") != "" {
		return true
	}

	return false
}

// IsSupport256Color is support 256 color
func IsSupport256Color() bool {
	// "TERM=xterm-256color"
	return strings.Contains(os.Getenv("TERM"), "256color")
}
