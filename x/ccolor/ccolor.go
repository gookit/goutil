// Package ccolor is a simple color render library for terminal.
// Its main code is the code that is extracted and simplified from gookit/color,
//
// TIP:
//
//	If you want to render with richer colors, use the https://github.com/gookit/color package.
package ccolor

import (
	"io"
	"os"
	"regexp"

	"github.com/gookit/goutil/x/termenv"
)

// color render templates
//
// ESC 操作的表示:
//
//	"\033"(Octal 8进制) = "\x1b"(Hexadecimal 16进制) = 27 (10进制)
const (
	// StartSet chars
	StartSet = "\x1b["
	// ResetSet close all properties.
	ResetSet = "\x1b[0m"
	// SettingTpl string.
	SettingTpl = "\x1b[%sm"
	// FullColorTpl for build color code
	FullColorTpl = "\x1b[%sm%s\x1b[0m"
	// CodeSuffix string for color code.
	CodeSuffix = "[0m"
)

// CodeExpr regex to clear color codes eg "\033[1;36mText\x1b[0m"
const CodeExpr = `\033\[[\d;?]+m`

var (
	// last error
	lastErr error
	// output the default io.Writer message print
	output io.Writer = os.Stdout
	// match color codes
	codeRegex = regexp.MustCompile(CodeExpr)
)

// Level value of current terminal.
func Level() termenv.ColorLevel { return termenv.TermColorLevel() }

// SetOutput set output writer
func SetOutput(w io.Writer) { output = w }

// LastErr info
func LastErr() error {
	defer func() {
		lastErr = nil
	}()
	return lastErr
}

//
// ---------------- for testing ----------------
//

// ForceEnableColor setting value. TIP: use for unit testing.
//
// Usage:
//
//	ccolor.ForceEnableColor()
//	defer ccolor.RevertColorSupport()
func ForceEnableColor() {
	termenv.ForceEnableColor()
}

// RevertColorSupport value
func RevertColorSupport() {
	termenv.RevertColorSupport()
}
