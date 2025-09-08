//go:build freebsd || openbsd || netbsd || dragonfly

package sysutil

import (
	"os/exec"
	"runtime"
	"strings"
)

// OsName system name. like runtime.GOOS.
// For Unix systems, this will be the actual GOOS value (freebsd, openbsd, netbsd, dragonfly)
var OsName = runtime.GOOS

// IsWin system. linux windows darwin
func IsWin() bool { return false }

// IsWindows system. linux windows darwin
func IsWindows() bool { return false }

// IsMac system
func IsMac() bool { return false }

// IsDarwin system
func IsDarwin() bool { return false }

// IsLinux system
func IsLinux() bool { return false }

// There are multiple possible providers to open a browser on Unix systems
// Similar to Linux, try xdg-open and other common browser launchers
var openBins = []string{"xdg-open", "x-www-browser", "www-browser", "firefox", "chrome", "chromium"}

// OpenURL Open file or browser URL
//
// Mac：
//
//	open 'https://github.com/inhere'
//
// Linux:
//
//	xdg-open URL
//	x-www-browser 'https://github.com/inhere'
//
// Windows:
//
//	cmd /c start https://github.com/inhere
//
// Unix (FreeBSD, OpenBSD, etc.):
//
//	Try xdg-open, x-www-browser, or fallback browsers
func OpenURL(URL string) error {
	for _, bin := range openBins {
		if _, err := exec.LookPath(bin); err == nil {
			return exec.Command(bin, URL).Run()
		}
	}

	return &exec.Error{Name: strings.Join(openBins, ","), Err: exec.ErrNotFound}
}
