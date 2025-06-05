// Package termenv provides detect color support of the current terminal.
// And with some utils for terminal env.
package termenv

import (
	"fmt"
	"os"
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
