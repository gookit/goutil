// Package process Provide some process handle util functions
package process

import (
	"errors"
	"os"
	"os/exec"
	"syscall"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/strutil"
)

// PID get current process ID
func PID() int {
	return os.Getpid()
}

// Start starts a new process with the program, arguments and attributes
// specified by name, argv and attr.
//
// alias of os.StartProcess()
func Start(name string, argv []string, attr *os.ProcAttr) (*os.Process, error) {
	return os.StartProcess(name, argv, attr)
}

// ProcInfo looks for a running process by its pid.
//
// alias of os.FindProcess()
func ProcInfo(pid int) (*os.Process, error) {
	return os.FindProcess(pid)
}

// PIDByName get PID by process name match(by pgrep)
func PIDByName(keywords string) int {
	// pgrep keywords
	binFile := "pgrep"
	_, err := exec.LookPath(binFile)
	if err == nil {
		output, err := exec.Command(binFile, keywords).Output()
		if err != nil {
			return 0
		}

		return strutil.SafeInt(string(byteutil.FirstLine(output)))
	}

	return 0
}

// KillByName kill process by name match
func KillByName(keywords string, sig syscall.Signal) error {
	if pid := PIDByName(keywords); pid > 0 {
		return Kill(pid, sig)
	}
	return errors.New("not found process pid of " + keywords)
}

// StopProcessOption stop process option
type StopProcessOption struct {
	// Check if the process exists before stopping it
	CheckExist bool
	// Whether to force exit the process
	ForceKill bool
	// Whether to wait for the process to exit
	WaitExit bool
	// How long to wait for the process to exit
	ExitTimeout int
	// Signal to send to the process(non-win)
	Signal syscall.Signal
}
