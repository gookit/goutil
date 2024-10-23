package sysutil

import (
	"bytes"
	"os/exec"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/gookit/goutil/internal/checkfn"
	"github.com/gookit/goutil/sysutil/cmdr"
)

// NewCmd instance
func NewCmd(bin string, args ...string) *cmdr.Cmd {
	return cmdr.NewCmd(bin, args...)
}

// FlushExec command, will flush output to stdout,stderr
func FlushExec(bin string, args ...string) error {
	return cmdr.NewCmd(bin, args...).FlushRun()
}

// QuickExec quick exec a simple command line
func QuickExec(cmdLine string, workDir ...string) (string, error) {
	return ExecLine(cmdLine, workDir...)
}

// ExecLine quick exec a command line string.
//
//	NOTE: not support | or ; in cmdLine
func ExecLine(cmdLine string, workDir ...string) (string, error) {
	p := cmdline.NewParser(cmdLine)

	// create a new Cmd instance
	cmd := p.NewExecCmd()
	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}

	bs, err := cmd.Output()
	return string(bs), err
}

// ExecCmd a command and return output.
//
// Usage:
//
//	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	// create a new Cmd instance
	cmd := exec.Command(binName, args...)
	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}

	bs, err := cmd.Output()
	return string(bs), err
}

// ShellExec exec command by shell cmdLine.
//
// shells e.g. "/bin/sh", "bash", "cmd", "cmd.exe", "powershell", "powershell.exe", "pwsh", "pwsh.exe"
//
//	eg: ShellExec("ls -al")
func ShellExec(cmdLine string, shells ...string) (string, error) {
	// shell := "/bin/sh"
	shell := "sh"
	if len(shells) > 0 {
		shell = shells[0]
	}

	// "-c" for bash,sh,zsh shell
	mark := "-c"

	// special for Windows shell
	if IsWindows() {
		// use cmd.exe, mark is "/c"
		if checkfn.StringsContains([]string{"cmd", "cmd.exe"}, shell) {
			mark = "/c"
		} else if checkfn.StringsContains([]string{"powershell", "powershell.exe", "pwsh", "pwsh.exe"}, shell) {
			// "-Command" for powershell
			mark = "-Command"
		}
	}

	var out bytes.Buffer
	cmd := exec.Command(shell, mark, cmdLine)
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}
	return out.String(), nil
}
