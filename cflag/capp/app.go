// Package capp provides a simple command line application build.
//
//   - Support add multiple commands
//   - Support add aliases for command
package capp

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/basefn"
	"github.com/gookit/goutil/x/ccolor"
)

// CFlags alias for cflag.CFlags
type CFlags = cflag.CFlags

// App struct
type App struct {
	*cflag.CFlags // save global flags
	// added commands
	commands *cmdSet

	Name string
	Desc string
	// Version for app
	Version string
	// NameWidth max width for command name
	NameWidth int

	// OnAppFlagParsed hook func
	OnAppFlagParsed func(app *App) bool
	// AfterHelpBuild hook
	AfterHelpBuild func(buf *strutil.Buffer)

	// BeforeRun each command hook func
	//  - cmdArgs: input raw args for current command.
	//  - return false to stop run.
	BeforeRun func(c *Cmd, cmdArgs []string) bool
	// AfterRun command hook func
	AfterRun func(c *Cmd, err error)
}

// New App instance
//
// Usage:
//
//	app := capp.New(func(app *cflag.App) {})
//	app.Name = "mycli"
//	app.Version = "0.0.1"
//	app.Desc = "mycli is a command line tool"
func New(fns ...func(app *App)) *App {
	// global flags for app
	gfs := cflag.NewEmpty(func(cf *cflag.CFlags) {
		cf.FlagSet = flag.NewFlagSet("app-flags", flag.ContinueOnError)
	})

	app := &App{
		CFlags:   gfs,
		commands: newCmdSet(),
		// with default version
		Version: "0.0.1",
		// NameWidth default value
		NameWidth: defaultNameWidth,
	}

	return app.WithConfigFn(fns...)
}

// NewApp instance. alias of New()
func NewApp(fns ...func(app *App)) *App { return New(fns...) }

// NewWith name and desc and option functions
func NewWith(name, version, desc string, fns ...func(app *App)) *App {
	app := NewApp(fns...)
	app.Name = name
	app.Desc = desc
	app.Version = version
	return app
}

// WithConfigFn config app
func (a *App) WithConfigFn(fns ...func(app *App)) *App {
	for _, fn := range fns {
		fn(a)
	}
	return a
}

// Add command(s) to app. panic if error.
//
// NOTE: command object should create use NewCmd()
//
// Usage:
//
//	app.Add(
//		cflag.NewCmd("cmd1", "desc1"),
//		cflag.NewCmd("cmd2", "desc2"),
//	)
//
// Or:
//
//	app.Add(cflag.NewCmd("cmd1", "desc1"))
//	app.Add(cflag.NewCmd("cmd2", "desc2"))
func (a *App) Add(cmds ...*Cmd) {
	basefn.PanicErr(a.AddOrErr(cmds...))
}

// AddOrErr add command(s) to app.
func (a *App) AddOrErr(cmds ...*Cmd) error {
	for _, cmd := range cmds {
		if err := a.addCmd(cmd); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) addCmd(c *Cmd) error {
	// NOTE: a command can only belong to one app.
	c.setApp(a)

	if err := a.commands.add(c); err != nil {
		return err
	}

	// sync the exported NameWidth with the added command
	a.NameWidth = mathutil.MaxInt(a.NameWidth, a.commands.nameLen(c))
	return nil
}

//
// region Run with args
// -----------------------------------

// Run app by os.Args
func (a *App) Run() {
	err := a.RunWithArgs(os.Args[1:])
	if err != nil {
		cflag.DebugMsg("app run error: %v", err)
		ccolor.Errorln("ERROR:", err)
		os.Exit(1)
	}
}

// RunWithArgs run app by input args
func (a *App) RunWithArgs(args []string) error {
	// init for run
	a.init()

	if showHelp, err := a.preRun(args); showHelp {
		return a.showHelp() // show help and stop run.
	} else if err != nil {
		// manual stop run
		if errors.Is(err, cflag.ErrStopRun) {
			cflag.DebugMsg("app flagParse return ErrStopRun, stop continue run.")
			return nil
		}
		return err
	}

	// fire onAppFlagParsed hook
	if a.OnAppFlagParsed != nil && !a.OnAppFlagParsed(a) {
		cflag.DebugMsg("app onAppFlagParsed return false, stop continue run.")
		return nil
	}

	// update args after parse global flags
	args = a.RemainArgs()
	if len(args) == 0 {
		return a.showHelp()
	}

	// first as command name
	cmd, ok := a.findCmd(args[0])
	if !ok {
		return fmt.Errorf("input not exists command %q", args[0])
	}

	// parse command flags and execute func.
	return a.runCmd(cmd, args[1:])
}

// runCmd run the command by args. will descend into the sub commands.
func (a *App) runCmd(c *Cmd, args []string) error {
	// let the command write to the app output, on its own output not set.
	a.syncOutput(c)

	// leaf command: same as before
	if !c.HasSubs() {
		return a.execCmd(c, args, func() error { return c.Parse(args) })
	}

	// group command: parse its own options, stop on first argument.
	//
	// NOTE: only parse once. parse again will make the repeatable flags append twice.
	if err := c.Prepare(); err != nil {
		return err
	}

	if err := c.DoParseStopOnArg(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			c.ShowHelp()
			return nil
		}
		if errors.Is(err, cflag.ErrStopRun) {
			return nil
		}
		return err
	}

	// TIP: not RemainArgs(). AddArg on a group command will bind the sub command name.
	posArgs := c.Args()
	if len(posArgs) > 0 {
		if isHelp(posArgs[0]) { // help is a reserved name
			c.ShowHelp()
			return nil
		}

		sub, ok := c.subs.find(posArgs[0])
		if !ok {
			return fmt.Errorf("command %q not exists in %q", posArgs[0], c.Path())
		}
		return a.runCmd(sub, posArgs[1:])
	}

	// no positional args: run itself or show the group help
	if c.Func == nil {
		c.ShowHelp()
		return nil
	}

	// same as the tail of cflag.CFlags.Parse(), but not parse the args again.
	return a.execCmd(c, args, func() error {
		if c.CFlags.BeforeRun != nil && !c.CFlags.BeforeRun(c.CFlags) {
			return nil
		}
		return c.Func(c)
	})
}

// execCmd run the command func, and fire the app hooks.
//
// NOTE: the app hooks only fire on the command that really runs.
func (a *App) execCmd(c *Cmd, cmdArgs []string, run func() error) error {
	if a.BeforeRun != nil && !a.BeforeRun(c, cmdArgs) {
		return nil
	}

	err := run()

	// fire after run hook
	if a.AfterRun != nil {
		a.AfterRun(c, err)
	}
	return err
}

// syncOutput make the command write to the app output, when the command did not
// set its own output. so `app.SetOutput(w)` also takes effect on the command help,
// the flag package uses os.Stderr on the output not set.
func (a *App) syncOutput(c *Cmd) {
	if c.CFlags == nil || a.CFlags == nil {
		return
	}

	out := a.Output()
	if out != nil && out != os.Stderr && c.CFlags.Output() == os.Stderr {
		c.CFlags.SetOutput(out)
	}
}

func (a *App) preRun(args []string) (showHelp bool, err error) {
	// prepare
	if err = a.Prepare(); err != nil {
		return false, err
	}

	// empty args or help flag
	if len(args) == 0 || args[0] == "" || isHelp(args[0]) {
		return true, nil
	}

	// parse global flags
	if err = a.DoParseStopOnArg(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return true, nil // ignore help error
		}
	}

	// rArgs := a.RemainArgs()
	// if len(rArgs) == 0 || isHelp(rArgs[0]) {
	// 	return true, nil
	// }
	return
}

func (a *App) init() {
	if a.Name == "" {
		// fix: path.Base not support windows
		a.Name = filepath.Base(os.Args[0])
	}
}

func (a *App) findCmd(name string) (*Cmd, bool) {
	if name[0] == '-' {
		return nil, false
	}

	return a.commands.find(name)
}

//
// region Show help
// -----------------------------------

func (a *App) showHelp() error {
	bin := a.Name
	buf := strutil.NewBuffer(512)

	buf.Printf("<cyan>%s</> - %s", bin, a.Desc)
	if a.Version != "" {
		buf.Printf(" Version: <cyan>%s</>", a.Version)
	}

	buf.Printf("\n\n<comment>Usage:</> %s <green>COMMAND</> [--Options...] [...Arguments]\n", bin)

	buf.WriteStr1Nl("<comment>Options:</>")
	buf.WriteStr1Nl("  <green>--help, -h</>" + strings.Repeat("    ", 4) + "Display application help")
	if a.CFlags != nil {
		a.RenderOptionsHelp(buf)
	}

	buf.WriteStr1Nl("\n<comment>Commands:</>")
	a.commands.render(buf, a.NameWidth)

	name := strutil.PadRight("help", " ", a.NameWidth)
	buf.Printf("  <green>%s</>    Display application help\n", name)
	buf.Printf("\nUse \"<cyan>%s COMMAND --help</>\" for about a command\n", bin)

	if a.LongHelp != "" {
		buf.WriteStr1Nl("\n" + a.LongHelp)
	}

	if a.AfterHelpBuild != nil {
		a.AfterHelpBuild(buf)
	}

	ccolor.Fprint(a.Output(), buf.ResetAndGet())
	return nil
}

func isHelp(s string) bool {
	return s == "help" || s == "--help" || s == "-h"
}
