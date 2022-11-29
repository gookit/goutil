//go:build !windows

package cliutil

import (
	"syscall"
)

func syscallStdinFd() int {
	return syscall.Stdin
}
