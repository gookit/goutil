// +build windows

package sysutil

import (
	"errors"
	"syscall"
)

// Kill process by pid
func Kill(pid int, signal syscall.Signal) error {
	return errors.New("not support")
}

// ProcessExists check process exists by pid
func ProcessExists(pid int) bool {
	return false
}
