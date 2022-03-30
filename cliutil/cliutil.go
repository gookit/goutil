// Package cliutil provides some util functions for CLI
package cliutil

import (
	"os"
	"path"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/gookit/goutil/sysutil"
)

// LineBuild build command line string by given args.
func LineBuild(binFile string, args []string) string {
	return cmdline.NewBuilder(binFile, args...).String()
}

// BuildLine build command line string by given args.
func BuildLine(binFile string, args []string) string {
	return cmdline.NewBuilder(binFile, args...).String()
}

// String2OSArgs parse input command line text to os.Args
func String2OSArgs(line string) []string {
	return cmdline.NewParser(line).Parse()
}

// StringToOSArgs parse input command line text to os.Args
func StringToOSArgs(line string) []string {
	return cmdline.NewParser(line).Parse()
}

// ParseLine input command line text. alias of the StringToOSArgs()
func ParseLine(line string) []string {
	return cmdline.NewParser(line).Parse()
}

// QuickExec quick exec an simple command line
func QuickExec(cmdLine string, workDir ...string) (string, error) {
	return sysutil.ExecLine(cmdLine, workDir...)
}

// ExecLine quick exec an command line string
func ExecLine(cmdLine string, workDir ...string) (string, error) {
	return sysutil.ExecLine(cmdLine, workDir...)
}

// ExecCmd a CLI bin file and return output.
//
// Usage:
// 	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	return sysutil.ExecCmd(binName, args, workDir...)
}

// ExecCommand alias of the ExecCmd()
func ExecCommand(binName string, args []string, workDir ...string) (string, error) {
	return sysutil.ExecCmd(binName, args, workDir...)
}

// ShellExec exec command by shell
//
// Usage:
// ret, err := cliutil.ShellExec("ls -al")
func ShellExec(cmdLine string, shells ...string) (string, error) {
	return sysutil.ShellExec(cmdLine, shells...)
}

// CurrentShell get current used shell env file. eg "/bin/zsh" "/bin/bash"
func CurrentShell(onlyName bool) (path string) {
	return sysutil.CurrentShell(onlyName)
}

// HasShellEnv has shell env check.
//
// Usage:
// 	HasShellEnv("sh")
// 	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	return sysutil.HasShellEnv(shell)
}

// Workdir get
func Workdir() string {
	dir, _ := os.Getwd()
	return dir
}

// BinDir get
func BinDir() string {
	binFile := os.Args[0]
	return path.Dir(binFile)
}

// BinFile get
func BinFile() string {
	return os.Args[0]
}
