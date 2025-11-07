package capp

import (
	"flag"
	"strings"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/x/ccolor"
)

// CmdOptionFn for one command
type CmdOptionFn func(c *Cmd)

// Cmd for App
type Cmd struct {
	*cflag.CFlags
	init bool
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

// NewCmd instance
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
	return c.ConfigCmd(fns...)
}

// ConfigCmd the cmd. eg: bing flags
func (c *Cmd) ConfigCmd(fns ...CmdOptionFn) *Cmd {
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
	return c.CFlags.Parse(args)
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
	if c.CFlags == nil {
		c.CFlags = cflag.NewEmpty(func(cf *cflag.CFlags) {
			cf.Desc = c.Desc
			cf.FlagSet = flag.NewFlagSet(c.Name, flag.ContinueOnError)
		})
	}

	if len(c.Aliases) > 0 {
		c.CFlags.Desc += " (alias: " + strings.Join(c.Aliases, ", ") + ")"
	}
	c.CFlags.Func = func(_ *cflag.CFlags) error {
		return c.Func(c)
	}
}

func (c *Cmd) getDesc() string {
	if c.Desc != "" {
		return c.Desc
	}
	return c.CFlags.Desc
}

// WithAliases set aliases for command
func WithAliases(aliases ...string) CmdOptionFn {
	return func(c *Cmd) {
		c.Aliases = aliases
	}
}
