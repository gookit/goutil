package process

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/gookit/goutil/sysutil"
	"golang.org/x/sys/windows"
)

const (
	processQueryLimitedInformation = 0x1000

	stillActive = 259
)

// Kill a process by pid. use taskkill on windows
//
// CMD example:
//
//	taskkill /pid 1234
//	taskkill /pid 1234 /f
func Kill(pid int, signal syscall.Signal) error {
	taskKill := windows.NewLazySystemDLL("taskkill")
	proc := taskKill.NewProc("TaskKill")
	_, _, err := proc.Call(uintptr(pid), uintptr(signal), 0)
	return err
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

// ExistsByName Determine whether a process exists based on its name(by tasklist)
//
// Usage:
//
//	ExistsByName("MyApp.exe")
//	// Fuzzy match by input name
//	ExistsByName("MyApp", true)
func ExistsByName(name string, fuzzyMatch ...bool) bool {
	// 按名称模糊匹配
	if len(fuzzyMatch) > 0 && fuzzyMatch[0] {
		out, err := sysutil.ShellExec("tasklist | findstr \""+name+"\" /NH", "cmd")
		if err != nil {
			return false
		}
		return strings.Contains(out, name)
	}

	out, err := sysutil.ExecCmd("tasklist", []string{"/FI", fmt.Sprintf("IMAGENAME eq %s", name), "/NH"})
	// out, err := sysutil.ShellExec("tasklist /FI \"IMAGENAME eq "+name+"\" /NH", "cmd") // shell执行有问题
	if err != nil {
		return false
	}
	return strings.Contains(out, name)
}

// StopByName Stop process based on process name(by taskkill).
//
// return (exists, output, error). check error to see if the process exists
//
// Usage:
//
//	StopByName("MyApp.exe")
func StopByName(name string, option ...*StopProcessOption) (bool, string, error) {
	opt := &StopProcessOption{}
	if len(option) > 0 && option[0] != nil {
		opt = option[0]
	}

	// 1. 检查进程是否存在
	if opt.CheckExist {
		if !ExistsByName(name) {
			return false, "", nil
		}
	}

	// cmd: taskkill /IM name.exe /F
	args := []string{"/IM", name}
	if opt.ForceKill {
		args = append(args, "/F")
	}
	out, err := sysutil.ExecCmd("taskkill", args)
	return true, out, err
}
