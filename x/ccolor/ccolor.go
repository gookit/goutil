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

	"github.com/gookit/goutil/internal/checkfn"
)

// Level is the color level supported by a terminal.
type Level uint8

const (
	LevelNone Level = iota // not support color
	Level16                // 16(4bit) color supported
	Level256               // 256(8bit) color supported
	LevelTrue              // support true(rgb) color
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
	// value of os color render and display
	//
	// NOTICE:
	// if ENV: NO_COLOR is not empty, will disable color render.
	noColor = os.Getenv("NO_COLOR") == ""

	// last error
	lastErr error
	// output the default io.Writer message print
	output io.Writer = os.Stdout
	// match color codes
	codeRegex = regexp.MustCompile(CodeExpr)
)

// cache color check values
var (
	colorLevel   Level
	supportColor bool
)

// CheckColorSupport on the system and terminal
func CheckColorSupport() bool {
	supportColor = false
	colorLevel = LevelNone

	// check is in the terminal
	if !isTerminal() {
		return false
	}

	if checkfn.IsSupportTrueColor() {
		supportColor = true
		colorLevel = LevelTrue
	} else if checkfn.IsSupport256Color() {
		supportColor = true
		colorLevel = Level256
	} else if checkfn.IsSupportColor() {
		supportColor = true
		colorLevel = Level16
	}

	// disable color by os ENV
	if noColor {
		supportColor = false
	}
	return supportColor
}

// ColorLevel value.
func ColorLevel() Level { return colorLevel }

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

var backOldVal bool

// ForceEnableColor setting value. TIP: use for unit testing.
//
// Usage:
//
//	ccolor.ForceEnableColor()
//	defer ccolor.RevertColorSupport()
func ForceEnableColor() {
	noColor = false
	backOldVal = supportColor
	// force enables color
	supportColor = true
	// return colorLevel
}

// RevertColorSupport value
func RevertColorSupport() {
	supportColor = backOldVal
	noColor = os.Getenv("NO_COLOR") == ""
}
