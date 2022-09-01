// Package cliutil provides some util functions for CLI
package cliutil

import (
	"os"
	"path"
	"strings"

	"github.com/gookit/goutil/cliutil/cmdline"
	"github.com/gookit/goutil/internal/comfunc"
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

// QuickExec quick exec a simple command line
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
//
//	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	return comfunc.ExecCmd(binName, args, workDir...)
}

// ExecCommand alias of the ExecCmd()
func ExecCommand(binName string, args []string, workDir ...string) (string, error) {
	return comfunc.ExecCmd(binName, args, workDir...)
}

// ShellExec exec command by shell
//
// Usage:
// ret, err := cliutil.ShellExec("ls -al")
func ShellExec(cmdLine string, shells ...string) (string, error) {
	return comfunc.ShellExec(cmdLine, shells...)
}

// CurrentShell get current used shell env file. eg "/bin/zsh" "/bin/bash"
func CurrentShell(onlyName bool) (path string) {
	return comfunc.CurrentShell(onlyName)
}

// HasShellEnv has shell env check.
//
// Usage:
//
//	HasShellEnv("sh")
//	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	return comfunc.HasShellEnv(shell)
}

// Workdir get
func Workdir() string {
	dir, _ := os.Getwd()
	return dir
}

// BinDir get
func BinDir() string {
	return path.Dir(os.Args[0])
}

// BinFile get
func BinFile() string {
	return os.Args[0]
}

// BinName get
func BinName() string {
	return path.Base(os.Args[0])
}

// BuildOptionHelpName for render flag help
func BuildOptionHelpName(names []string) string {
	var sb strings.Builder

	size := len(names) - 1
	for i, name := range names {
		sb.WriteByte('-')
		if len(name) > 1 {
			sb.WriteByte('-')
		}

		sb.WriteString(name)
		if i < size {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

// ShellQuote quote a string on contains ', ", SPACE
func ShellQuote(s string) string {
	var quote byte
	if strings.ContainsRune(s, '"') {
		quote = '\''
	} else if s == "" || strings.ContainsRune(s, '\'') || strings.ContainsRune(s, ' ') {
		quote = '"'
	}

	if quote > 0 {
		ln := len(s) + 2
		bs := make([]byte, ln, ln)

		bs[0] = quote
		bs[ln-1] = quote
		if ln > 2 {
			copy(bs[1:ln-1], s)
		}

		s = string(bs)
	}
	return s
}

// OutputLines split output to lines
func OutputLines(output string) []string {
	output = strings.TrimSuffix(output, "\n")
	if output == "" {
		return nil
	}
	return strings.Split(output, "\n")
}

// FirstLine from command output
func FirstLine(output string) string {
	if i := strings.Index(output, "\n"); i >= 0 {
		return output[0:i]
	}
	return output
}
