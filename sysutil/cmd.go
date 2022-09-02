package sysutil

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Cmd struct
type Cmd struct {
	*exec.Cmd

	BeforeExec func(c *Cmd)
}

// NewCmd instance
func NewCmd(bin string, args ...string) *Cmd {
	return &Cmd{
		Cmd: exec.Command(bin, args...),
	}
}

// Cmdline to command line
func (c *Cmd) Cmdline() string {
	b := new(strings.Builder)
	// b.WriteString(c.Path)

	for i, a := range c.Args {
		if i > 0 {
			b.WriteByte(' ')
		}

		if strings.ContainsRune(a, '"') {
			b.WriteString(fmt.Sprintf(`'%s'`, a))
		} else if a == "" || strings.ContainsRune(a, '\'') || strings.ContainsRune(a, ' ') {
			b.WriteString(fmt.Sprintf(`"%s"`, a))
		} else {
			b.WriteString(a)
		}
	}
	return b.String()
}

// OnBefore exec add hook
func (c *Cmd) OnBefore(fn func(c *Cmd)) *Cmd {
	c.BeforeExec = fn
	return c
}

// WithWorkDir returns the current object
func (c *Cmd) WithWorkDir(dir string) *Cmd {
	c.Dir = dir
	return c
}

// OutputToStd output to OS stdout and error
func (c *Cmd) OutputToStd() *Cmd {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

// WithStdin returns the current argument
func (c *Cmd) WithStdin(in *os.File) *Cmd {
	c.Stdin = in
	return c
}

// WithOutput returns the current argument
func (c *Cmd) WithOutput(out *os.File, errOut *os.File) *Cmd {
	c.Stdout = out
	if errOut != nil {
		c.Stderr = errOut
	}
	return c
}

// WithArg add args and returns the current object. alias of the WithArg()
func (c *Cmd) WithArg(args ...string) *Cmd {
	c.Args = append(c.Args, args...)
	return c
}

// AddArg add args and returns the current object
func (c *Cmd) AddArg(args ...string) *Cmd {
	return c.WithArg(args...)
}

// Argf add arg and returns the current object.
func (c *Cmd) Argf(format string, args ...interface{}) *Cmd {
	c.Args = append(c.Args, fmt.Sprintf(format, args...))
	return c
}

// WithArgf add arg and returns the current object. alias of the Argf()
func (c *Cmd) WithArgf(format string, args ...interface{}) *Cmd {
	return c.Argf(format, args...)
}

// ArgIf add arg and returns the current object
func (c *Cmd) ArgIf(arg string, exprOk bool) *Cmd {
	if exprOk {
		c.Args = append(c.Args, arg)
	}
	return c
}

// WithArgIf add arg and returns the current object
func (c *Cmd) WithArgIf(arg string, exprOk bool) *Cmd {
	return c.ArgIf(arg, exprOk)
}

// AddArgs for the git. alias of WithArgs()
func (c *Cmd) AddArgs(args []string) *Cmd {
	return c.WithArgs(args)
}

// WithArgs for the git
func (c *Cmd) WithArgs(args []string) *Cmd {
	if len(args) > 0 {
		c.Args = append(c.Args, args...)
	}
	return c
}

// WithArgsIf add arg and returns the current object
func (c *Cmd) WithArgsIf(args []string, exprOk bool) *Cmd {
	if exprOk && len(args) > 0 {
		c.Args = append(c.Args, args...)
	}
	return c
}

// ResetArgs for git
func (c *Cmd) ResetArgs() {
	c.Args = make([]string, 0)
}

// -------------------------------------------------
// run command
// -------------------------------------------------

// GoCmd get exec.Cmd
func (c *Cmd) GoCmd() *exec.Cmd {
	return c.Cmd
}

// Success run and return whether success
func (c *Cmd) Success() bool {
	return c.Run() == nil
}

// SafeLines run and return output as lines
func (c *Cmd) SafeLines() []string {
	ss, _ := c.OutputLines()
	return ss
}

// OutputLines run and return output as lines
func (c *Cmd) OutputLines() ([]string, error) {
	out, err := c.Output()
	if err != nil {
		return nil, err
	}
	return OutputLines(out), err
}

// SafeOutput run and return output
func (c *Cmd) SafeOutput() string {
	out, err := c.Output()
	if err != nil {
		return ""
	}
	return out
}

// Output run and return output
func (c *Cmd) Output() (string, error) {
	if c.BeforeExec != nil {
		c.BeforeExec(c)
	}

	output, err := c.Output()
	return string(output), err
}

// CombinedOutput run and return output, will combine stderr and stdout output
func (c *Cmd) CombinedOutput() (string, error) {
	if c.BeforeExec != nil {
		c.BeforeExec(c)
	}

	output, err := c.CombinedOutput()
	return string(output), err
}

// MustRun a command. will panic on error
func (c *Cmd) MustRun() {
	if err := c.Run(); err != nil {
		panic(err)
	}
}

// Run runs command with `Exec` on platforms except Windows
// which only supports `Spawn`
func (c *Cmd) Run() error {
	if c.BeforeExec != nil {
		c.BeforeExec(c)
	}
	return c.Cmd.Run()

	// if IsWindows() {
	// 	return c.Spawn()
	// }
	// return c.Exec()
}

// Spawn runs command with spawn(3)
// func (c *Cmd) Spawn() error {
// 	return c.Cmd.Run()
// }
//
// // Exec runs command with exec(3)
// // Note that Windows doesn't support exec(3): http://golang.org/src/pkg/syscall/exec_windows.go#L339
// func (c *Cmd) Exec() error {
// 	binary, err := exec.LookPath(c.Path)
// 	if err != nil {
// 		return &exec.Error{
// 			Name: c.Path,
// 			Err:  errorx.Newf("%s not found in the system", c.Path),
// 		}
// 	}
//
// 	args := []string{binary}
// 	args = append(args, c.Args...)
//
// 	if c.BeforeExec != nil {
// 		c.BeforeExec(c)
// 	}
// 	return syscall.Exec(binary, args, os.Environ())
// }

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
