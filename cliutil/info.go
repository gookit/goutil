package cliutil

import (
	"os"
	"path/filepath"

	"github.com/gookit/goutil/x/termenv"
)

// Workdir get
func Workdir() string {
	dir, _ := os.Getwd()
	return dir
}

// BinDir get
func BinDir() string {
	return filepath.Dir(os.Args[0])
}

// BinFile get
func BinFile() string {
	return os.Args[0]
}

// BinName get
func BinName() string {
	return filepath.Base(os.Args[0])
}

// GetTermSize for current console terminal.
func GetTermSize(refresh ...bool) (w int, h int) {
	w, h = termenv.GetTermSize(refresh...)
	return
}
