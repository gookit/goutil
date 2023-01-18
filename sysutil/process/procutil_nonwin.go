//go:build !windows

package process

import "syscall"

// Kill a process by pid
func Kill(pid int, signal syscall.Signal) error {
	return syscall.Kill(pid, signal)
}
