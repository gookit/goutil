package cliutil

import (
	"os"
	"path"

	"golang.org/x/term"
)

// Workdir get
func Workdir() string {
	dir, _ := os.Getwd()
	return dir
}

// BinDir get
func BinDir() string {
	return path.Dir(os.Args[0])
}

// BinFile get
func BinFile() string {
	return os.Args[0]
}

// BinName get
func BinName() string {
	return path.Base(os.Args[0])
}

// exec: `stty -a 2>&1`
// const (
// mac: speed 9600 baud; 97 rows; 362 columns;
// macSttyMsgPattern = `(\d+)\s+rows;\s*(\d+)\s+columns;`
// linux: speed 38400 baud; rows 97; columns 362; line = 0;
// linuxSttyMsgPattern = `rows\s+(\d+);\s*columns\s+(\d+);`
// )
var terminalWidth, terminalHeight int

// GetTermSize for current console terminal.
func GetTermSize(refresh ...bool) (w int, h int) {
	if terminalWidth > 0 && len(refresh) > 0 && !refresh[0] {
		return terminalWidth, terminalHeight
	}

	var err error
	w, h, err = term.GetSize(syscallStdinFd())
	if err != nil {
		return
	}

	// cache result
	terminalWidth, terminalHeight = w, h
	return
}
