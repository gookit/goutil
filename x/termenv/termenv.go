package termenv

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/goutil/internal/comfunc"
)

var (
	lastErr error
	// debug mode
	debugMode bool
	// support color of current terminal
	supportColor bool

	// value of os color render and display
	//
	// NOTICE:
	// if ENV: NO_COLOR is not empty, will disable color render.
	noColor = os.Getenv("NO_COLOR") != ""

	// the color support level for current terminal
	// needVTP - need enable VTP, only for Windows OS
	colorLevel, needVTP = detectTermColorLevel()
)

// TermColorLevel returns the color support level for the current terminal.
func TermColorLevel() ColorLevel { return colorLevel }

// SetDebugMode sets debug mode.
func SetDebugMode(enable bool) { debugMode = enable }

// LastErr returns the last error.
func LastErr() error {
	defer func() {
		lastErr = nil // reset on get
	}()
	return lastErr
}

// CurrentShell get current used shell env file.
//
// eg "/bin/zsh" "/bin/bash".
// if onlyName=true, will return "zsh", "bash"
func CurrentShell(onlyName bool) string {
	return comfunc.CurrentShell(onlyName)
}

// HasShellEnv has shell env check.
//
// Usage:
//
//	HasShellEnv("sh")
//	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	// can also use: "echo $0"
	out, err := shellExec("echo OK", shell)
	if err != nil {
		return false
	}
	return strings.TrimSpace(out) == "OK"
}

// IsShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
func IsShellSpecialVar(c uint8) bool {
	switch c {
	case '*', '#', '$', '@', '!', '?', '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

func debugf(tpl string, v ...any) {
	if debugMode {
		fmt.Printf("TERM_ENV: "+tpl+"\n", v...)
	}
}

func setLastErr(err error) {
	if err != nil {
		debugf("last error: %v", err)
		lastErr = err
	}
}
