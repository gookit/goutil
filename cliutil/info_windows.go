package cliutil

import "syscall"

// on Windows, must convert 'syscall.Stdin' to int
func syscallStdinFd() int {
	return int(syscall.Stdin)
}
