// Package capp provides a simple command line application build.
//
//  - Support add multiple commands
//  - Support add aliases for command
package capp

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/basefn"
	"github.com/gookit/goutil/x/ccolor"
)

// App struct
type App struct {
	*cflag.CFlags // save global flags
	// added commands
	names []string
	cmds  map[string]*Cmd
	cmdAs maputil.Aliases

	Name string
	Desc string
	// Version for app
	Version string
	// NameWidth max width for command name
	NameWidth  int
	HelpWriter io.Writer

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
//	 app := capp.New(func(app *cflag.App) {})
//	 app.Name = "mycli"
//	 app.Version = "0.0.1"
//	 app.Desc = "mycli is a command line tool"
func New(fns ...func(app *App)) *App {
	// global flags for app
	gfs := cflag.NewEmpty(func(cf *cflag.CFlags) {
		cf.FlagSet = flag.NewFlagSet("app-flags", flag.ContinueOnError)
	})

	app := &App{
		CFlags: gfs,
		cmds:   make(map[string]*Cmd),
		cmdAs: make(maputil.Aliases),
		// with default version
		Version: "0.0.1",
		// NameWidth default value
		NameWidth:  12,
		HelpWriter: os.Stdout,
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
	ln := len(c.Name)
	if ln == 0 {
		return errors.New("command name cannot be empty")
	}

	if _, ok := a.cmds[c.Name]; ok {
		return fmt.Errorf("command name %s has been exists", c.Name)
	}

	a.names = append(a.names, c.Name)
	a.cmds[c.Name] = c

	// add aliases
	if len(c.Aliases) > 0 {
		ln += len(strings.Join(c.Aliases, ", ")) + 4
		a.cmdAs.AddAliases(c.Name, c.Aliases)
	}
	a.NameWidth = mathutil.MaxInt(a.NameWidth, ln)

	// attach handle func
	c.initCmd()

	if c.OnAdd != nil {
		c.OnAdd(c)
	}
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
		return a.showHelp() // stop run.
	} else if err != nil {
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

	cmdArgs := args[1:]
	if a.BeforeRun != nil && !a.BeforeRun(cmd, cmdArgs) {
		return nil
	}

	// parse command flags and execute func.
	err := cmd.Parse(cmdArgs)

	// fire after run hook
	if a.AfterRun != nil {
		a.AfterRun(cmd, err)
	}
	return err
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
	if err = a.DoParse(args); err != nil {
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

	// resolve alias
	name = a.cmdAs.ResolveAlias(name)
	cmd, ok := a.cmds[name]
	return cmd, ok
}

//
// region Show help
// -----------------------------------

func (a *App) showHelp() error {
	bin := a.Name
	buf := strutil.NewBuffer(512)

	buf.Printf("<cyan>%s</> - %s", bin, a.Desc)
	if a.Version != "" {
		buf.Printf("(Version: <cyan>%s</>)", a.Version)
	}

	buf.Printf("\n\n<comment>Usage:</> %s <green>COMMAND</> [--Options...] [...Arguments]\n", bin)

	buf.WriteStr1Nl("<comment>Options:</>")
	buf.WriteStr1Nl("  <green>--help, -h</>" + strings.Repeat("    ", 4) + "Display application help")
	if a.CFlags != nil {
		a.RenderOptionsHelp(buf)
	}

	buf.WriteStr1Nl("\n<comment>Commands:</>")
	sort.Strings(a.names)

	gaMap := a.cmdAs.GroupAliases()
	for _, name := range a.names {
		c := a.cmds[name]
		if len(gaMap[name]) > 0 {
			name = name + ", " + strings.Join(gaMap[name], ", ")
		}

		name = strutil.PadRight(name, " ", a.NameWidth)
		buf.Printf("  <green>%s</>    %s\n", name, strutil.UpperFirst(c.getDesc()))
	}

	name := strutil.PadRight("help", " ", a.NameWidth)
	buf.Printf("  <green>%s</>    Display application help\n", name)
	buf.Printf("\nUse \"<cyan>%s COMMAND --help</>\" for about a command\n", bin)

	if a.AfterHelpBuild != nil {
		a.AfterHelpBuild(buf)
	}

	if a.HelpWriter == nil {
		a.HelpWriter = os.Stdout
	}

	ccolor.Fprint(a.HelpWriter, buf.ResetAndGet())
	return nil
}

func isHelp(s string) bool {
	return s == "help" || s == "--help" || s == "-h"
}
