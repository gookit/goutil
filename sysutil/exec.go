package sysutil

import "github.com/gookit/goutil/cliutil"

// QuickExec quick exec an simple command line
func QuickExec(cmdLine string, workDir ...string) (string, error) {
	return cliutil.QuickExec(cmdLine, workDir...)
}

// ExecCmd a CLI bin file and return output.
// Usage:
// 	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	return cliutil.ExecCmd(binName, args, workDir...)
}

