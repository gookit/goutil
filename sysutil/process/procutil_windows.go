package process

import (
	"errors"
	"syscall"

	"golang.org/x/sys/windows"
)

const (
	processQueryLimitedInformation = 0x1000

	stillActive = 259
)

// Kill a process by pid
func Kill(pid int, signal syscall.Signal) error {
	return errors.New("not support")
}

// Exists check process running by given pid
func Exists(pid int) bool {
	h, err := windows.OpenProcess(processQueryLimitedInformation, false, uint32(pid))
	if err != nil {
		return false
	}

	var c uint32
	err = windows.GetExitCodeProcess(h, &c)
	_ = windows.Close(h)

	if err != nil {
		return c == stillActive
	}
	return true
}
