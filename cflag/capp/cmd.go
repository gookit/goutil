package capp

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/basefn"
	"github.com/gookit/goutil/x/ccolor"
)

// CmdOptionFn for one command
type CmdOptionFn func(c *Cmd)

// Cmd for App
type Cmd struct {
	*cflag.CFlags
	init bool
	// sub commands set. only the group command has it.
	subs *cmdSet
	// the app that the command belongs to. may be nil on standalone use.
	app *App
	// parent group command. nil for the top command.
	parent *Cmd
	// hasPosArg is true on declared positional args by AddArg() or BindArg().
	hasPosArg bool

	Name string
	Desc string // desc for command, will sync set to CFlags.Desc
	// Aliases name for command
	Aliases []string
	// OnAdd hook func. fire on add to App
	//  - you can add some cli options or arguments.
	OnAdd func(c *Cmd)
	// Func for run command, will call after options parsed. will sync set to CFlags.Func
	Func func(c *Cmd) error
}

// WrapRunFunc wrap a no-params func as cmd run func
func WrapRunFunc(fn func() error) func(c *Cmd) error {
	return func(c *Cmd) error { return fn() }
}

// NewCmd create a Cmd instance
func NewCmd(name, desc string, runFunc ...func(c *Cmd) error) *Cmd {
	fs := cflag.NewEmpty(func(c *cflag.CFlags) {
		c.Desc = desc
		c.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)
	})

	cmd := &Cmd{Name: name, Desc: desc, CFlags: fs}
	if len(runFunc) > 0 {
		cmd.Func = runFunc[0]
	}
	return cmd
}

// WithConfigFn config cmd, alias of ConfigCmd()
func (c *Cmd) WithConfigFn(fns ...CmdOptionFn) *Cmd {
	return c.Config(fns...)
}

// Config the cmd. eg: bing flags
func (c *Cmd) Config(fns ...CmdOptionFn) *Cmd {
	for _, fn := range fns {
		if fn != nil {
			fn(c)
		}
	}
	return c
}

// QuickRun parse OS flags and run command, will auto handle error
func (c *Cmd) QuickRun() { c.MustParse(nil) }

// MustRun parse flags and run command. alias of MustParse()
func (c *Cmd) MustRun(args []string) { c.MustParse(args) }

// MustParse parse flags and run command, will auto handle error
func (c *Cmd) MustParse(args []string) {
	if err := c.Parse(args); err != nil {
		ccolor.Redln("ERROR:", err)
	}
}

// Parse flags and run command func
//
// If args is nil, will parse os.Args
func (c *Cmd) Parse(args []string) error {
	// fix: cmd.xxRun not exec Cmd.Func
	c.initCmd()
	return c.ensureCFlags().Parse(args)
}

func (c *Cmd) initCmd() {
	if c.init {
		return
	}
	c.init = true

	// attach handle func
	if c.Func == nil {
		return
	}

	// fix: init c.CFlags on not exist
	c.ensureCFlags()

	if len(c.Aliases) > 0 {
		c.CFlags.Desc += " (alias: " + strings.Join(c.Aliases, ", ") + ")"
	}
	c.CFlags.Func = func(_ *cflag.CFlags) error {
		return c.Func(c)
	}
}

// ensureCFlags ensure the c.CFlags is initialized.
func (c *Cmd) ensureCFlags() *cflag.CFlags {
	if c.CFlags == nil {
		c.CFlags = cflag.NewEmpty(func(cf *cflag.CFlags) {
			cf.Desc = c.Desc
			cf.FlagSet = flag.NewFlagSet(c.Name, flag.ContinueOnError)
		})
	}
	return c.CFlags
}

//
// region sub commands
// -----------------------------------

// Add sub command(s) to the command. panic if error.
//
// Usage:
//
//	group := capp.NewCmd("group", "this is a group command")
//	group.Add(
//		capp.NewCmd("sub1", "this is a sub command1"),
//		capp.NewCmd("sub2", "this is a sub command2"),
//	)
//	app.Add(group)
func (c *Cmd) Add(cmds ...*Cmd) {
	basefn.PanicErr(c.AddOrErr(cmds...))
}

// AddOrErr add sub command(s) to the command.
//
// NOTE:
//   - a group command can not declare positional args.
//   - the sub command Aliases should be set before add.
func (c *Cmd) AddOrErr(cmds ...*Cmd) error {
	if len(cmds) == 0 {
		return nil
	}

	if c.hasPosArg {
		return fmt.Errorf("group command %q cannot declare positional args", c.Name)
	}

	if c.subs == nil {
		c.subs = newCmdSet()
	}

	for _, sub := range cmds {
		sub.parent = c
		sub.setApp(c.app)

		if err := c.subs.add(sub); err != nil {
			return err
		}
	}

	// render the group help on show help. only install on user not set.
	if cf := c.ensureCFlags(); cf.HelpFunc == nil {
		cf.HelpFunc = c.renderGroupHelp
	}
	return nil
}

// AddArg binding for command, by position.
//
// NOTE: a group command (has sub commands) can not declare positional args.
func (c *Cmd) AddArg(name, desc string, requireDefaultArrayed ...any) {
	if c.HasSubs() {
		basefn.Panicf("group command %q cannot declare positional args", c.Name)
	}

	c.hasPosArg = true
	c.ensureCFlags().AddArg(name, desc, requireDefaultArrayed...)
}

// BindArg binding a named arg for command.
//
// NOTE: a group command (has sub commands) can not declare positional args.
func (c *Cmd) BindArg(arg *cflag.FlagArg) {
	if c.HasSubs() {
		basefn.Panicf("group command %q cannot declare positional args", c.Name)
	}

	c.hasPosArg = true
	c.ensureCFlags().BindArg(arg)
}

// Sub get sub command by name. will resolve the alias name.
func (c *Cmd) Sub(name string) (*Cmd, bool) {
	if c.subs == nil {
		return nil, false
	}
	return c.subs.find(name)
}

// HasSubs check the command has sub commands.
func (c *Cmd) HasSubs() bool { return c.subs.hasCommands() }

// Parent get the parent group command. nil for the top command.
func (c *Cmd) Parent() *Cmd { return c.parent }

// Path get the full command path. eg: "app session list"
func (c *Cmd) Path() string {
	names := []string{c.Name}
	for p := c.parent; p != nil; p = p.parent {
		names = append(names, p.Name)
	}

	// reverse to: top command -> current command
	for i, j := 0, len(names)-1; i < j; i, j = i+1, j-1 {
		names[i], names[j] = names[j], names[i]
	}

	path := strings.Join(names, " ")
	if c.app != nil && c.app.Name != "" {
		path = c.app.Name + " " + path
	}
	return path
}

// setApp set the app for the command and all its sub commands.
//
// NOTE: a command can only belong to one app.
func (c *Cmd) setApp(a *App) {
	c.app = a
	if c.subs != nil {
		c.subs.propagateApp(a)
	}
}

// renderGroupHelp render the help of a group command (has sub commands).
//
// It is used as the CFlags.HelpFunc, so it should keep the same signature.
func (c *Cmd) renderGroupHelp(cf *cflag.CFlags) {
	path := c.Path()
	buf := strutil.NewBuffer(512)

	buf.Printf("<cyan>%s</>\n\n", strutil.UpperFirst(c.getDesc()))
	buf.Printf("<comment>Usage:</> %s SUBCOMMAND [--Options...] [...Arguments]\n", path)
	buf.WriteStr("<comment>Options:</>\n")
	cf.RenderOptionsHelp(buf)
	buf.WriteStr1Nl("  <green>--help, -h</>" + strings.Repeat("    ", 4) + "Display command help")

	if c.subs.hasCommands() {
		buf.WriteStr1Nl("\n<comment>Commands:</>")
		c.subs.render(buf, c.subs.nameWidth)
	}

	buf.Printf("\nUse \"<cyan>%s SUBCOMMAND --help</>\" for about a command\n", path)

	if c.LongHelp != "" {
		buf.WriteStr1Nl("\n<comment>Help:</>")
		buf.WriteStr1Nl(strings.Trim(c.LongHelp, "\n"))
	}

	if c.Example != "" {
		buf.WriteStr1Nl("\n<comment>Examples:</>")
		buf.WriteStr1(strings.Trim(c.Example, "\n"))
	}

	binName := cf.Name()
	helpVars := map[string]string{
		"{{cmd}}":     binName,
		"{{command}}": binName,
		"{{binName}}": binName,
		"{{binFile}}": cf.BinFile(),
	}
	ccolor.Fprintln(cf.Output(), strutil.Replaces(buf.ResetAndGet(), helpVars))
}

func (c *Cmd) getDesc() string {
	if c.Desc != "" {
		return c.Desc
	}

	if c.CFlags == nil { // Cmd created by struct literal
		return ""
	}
	return c.CFlags.Desc
}

// WithAliases set aliases for command
func WithAliases(aliases ...string) CmdOptionFn {
	return func(c *Cmd) {
		c.Aliases = aliases
	}
}
