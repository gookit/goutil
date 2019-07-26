// Package cliutil provides some util functions for CLI
package cliutil

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"
	
	"github.com/gookit/goutil/strutil"
)

// QuickExec quick exec an simple command line
func QuickExec(cmdLine string, workDir ...string) (string, error) {
	ss := strutil.Split(cmdLine)

	return ExecCmd(ss[0], ss[1:], workDir...)
}

// ExecCmd a CLI bin file and return output.
// usage:
// 	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	// create a new Cmd instance
	cmd := exec.Command(binName, args...)
	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}

	bs, err := cmd.Output()
	return string(bs), err
}

// ExecCommand alias of the ExecCmd()
func ExecCommand(binName string, args []string, workDir ...string) (string, error) {
	return ExecCmd(binName, args, workDir...)
}

// ShellExec exec command by shell
// cmdStr eg. "ls -al"
func ShellExec(cmdStr string, shells ...string) (string, error) {
	shell := "/bin/sh"

	if len(shells) > 0 {
		shell = shells[0]
	}

	// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(shell, "-c", cmdStr)

	// 读取io.Writer类型的cmd.Stdout，
	// 再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run执行c包含的命令，并阻塞直到完成。
	// 这里stdout被取出，cmd.Wait()无法正确获取 stdin,stdout,stderr，则阻塞在那了
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}

// CurrentShell get current used shell env file. eg "/bin/zsh" "/bin/bash"
func CurrentShell(onlyName bool) (path string) {
	path, err := ShellExec("echo $SHELL")
	if err != nil {
		return ""
	}

	path = strings.TrimSpace(path)
	if onlyName && len(path) > 0 {
		path = filepath.Base(path)
	}
	return
}

// HasShellEnv has shell env check.
// Usage:
// 	HasShellEnv("sh")
// 	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	// can also use: "echo $0"
	out, err := ShellExec("echo OK", "", shell)
	if err != nil {
		return false
	}

	return strings.TrimSpace(out) == "OK"
}
