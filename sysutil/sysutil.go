package sysutil

import (
	"os"
)

// Workdir get
func Workdir() string {
	dir, _ := os.Getwd()
	return dir
}
