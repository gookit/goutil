//go:build windows

package ccolor

import (
	"golang.org/x/sys/windows"
)

func init() {
	// terminal support color. don't need open virtual process
	if CheckColorSupport() || noColor {
		return
	}

	// use virtual terminal process on Windows
	handle, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return
	}

	var mode uint32
	if err = windows.GetConsoleMode(handle, &mode); err != nil {
		return
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	if err = windows.SetConsoleMode(handle, mode); err != nil {
		return
	}

	// set support color is true
	supportColor = true
	colorLevel = LevelTrue
}
