package cmdr

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/gookit/goutil/internal/comfunc"
)

// Cmd struct
type Cmd struct {
	*exec.Cmd
	// Name of the command
	Name string
	// RunBefore hook
	RunBefore func(c *Cmd)
	// RunAfter hook
	RunAfter func(c *Cmd, err error)
}

// WrapGoCmd instance
func WrapGoCmd(cmd *exec.Cmd) *Cmd {
	return &Cmd{Cmd: cmd}
}

// NewCmd instance
func NewCmd(bin string, args ...string) *Cmd {
	return &Cmd{
		Cmd: exec.Command(bin, args...),
	}
}

// IDString of the command
func (c *Cmd) IDString() string {
	if c.Name != "" {
		return c.Name
	}

	if len(c.Args) > 0 {
		return c.Args[0]
	}
	return c.Path
}

// Cmdline to command line
func (c *Cmd) Cmdline() string {
	return comfunc.Cmdline(c.Args)
}

// OnBefore exec add hook
func (c *Cmd) OnBefore(fn func(c *Cmd)) *Cmd {
	c.RunBefore = fn
	return c
}

// OnAfter exec add hook
func (c *Cmd) OnAfter(fn func(c *Cmd, err error)) *Cmd {
	c.RunAfter = fn
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
func (c *Cmd) WithStdin(in io.Reader) *Cmd {
	c.Stdin = in
	return c
}

// WithOutput returns the current argument
func (c *Cmd) WithOutput(out, errOut io.Writer) *Cmd {
	c.Stdout = out
	if errOut != nil {
		c.Stderr = errOut
	}
	return c
}

// AddArg add args and returns the current object
func (c *Cmd) AddArg(args ...string) *Cmd { return c.WithArg(args...) }

// WithArg add args and returns the current object. alias of the WithArg()
func (c *Cmd) WithArg(args ...string) *Cmd {
	c.Args = append(c.Args, args...)
	return c
}

// AddArgf add args and returns the current object. alias of the WithArgf()
func (c *Cmd) AddArgf(format string, args ...interface{}) *Cmd { return c.WithArgf(format, args...) }

// WithArgf add arg and returns the current object
func (c *Cmd) WithArgf(format string, args ...interface{}) *Cmd {
	c.Args = append(c.Args, fmt.Sprintf(format, args...))
	return c
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
func (c *Cmd) AddArgs(args []string) *Cmd { return c.WithArgs(args) }

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
	if len(c.Args) > 0 {
		c.Args = c.Args[0:1]
	} else {
		c.Args = c.Args[:0]
	}
}

// -------------------------------------------------
// run command
// -------------------------------------------------

// GoCmd get exec.Cmd
func (c *Cmd) GoCmd() *exec.Cmd { return c.Cmd }

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
	if c.RunBefore != nil {
		c.RunBefore(c)
	}

	output, err := c.Cmd.Output()

	if c.RunAfter != nil {
		c.RunAfter(c, err)
	}
	return string(output), err
}

// CombinedOutput run and return output, will combine stderr and stdout output
func (c *Cmd) CombinedOutput() (string, error) {
	if c.RunBefore != nil {
		c.RunBefore(c)
	}

	output, err := c.Cmd.CombinedOutput()

	if c.RunAfter != nil {
		c.RunAfter(c, err)
	}
	return string(output), err
}

// MustRun a command. will panic on error
func (c *Cmd) MustRun() {
	if err := c.Run(); err != nil {
		panic(err)
	}
}

// FlushRun runs command and flush output to stdout
func (c *Cmd) FlushRun() error {
	c.OutputToStd()
	return c.Run()
}

// Run runs command
func (c *Cmd) Run() error {
	if c.RunBefore != nil {
		c.RunBefore(c)
	}

	// do running
	err := c.Cmd.Run()

	if c.RunAfter != nil {
		c.RunAfter(c, err)
	}
	return err

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
// 	return syscall.Exec(binary, args, os.Environ())
// }
