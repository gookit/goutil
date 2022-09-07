package cflag

import (
	"flag"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/gookit/color"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/stdutil"
	"github.com/gookit/goutil/strutil"
)

// App struct
type App struct {
	cmds map[string]*Cmd
	// names
	names []string

	Name string
	Desc string
	// NameWidth max width for command name
	NameWidth int
	// Version for app
	Version string

	// AfterHelpBuild hook
	AfterHelpBuild func(buf *strutil.Buffer)
}

// NewApp instance
func NewApp(fns ...func(app *App)) *App {
	app := &App{
		cmds: make(map[string]*Cmd),
		// NameWidth default value
		NameWidth: 12,
	}

	for _, fn := range fns {
		fn(app)
	}
	return app
}

// Add command(s) to app
func (a *App) Add(cmds ...*Cmd) {
	for _, cmd := range cmds {
		a.addCmd(cmd)
	}
}

func (a *App) addCmd(c *Cmd) {
	ln := len(c.Name)
	if ln == 0 {
		panic("command name cannot be empty")
	}

	if _, ok := a.cmds[c.Name]; ok {
		stdutil.Panicf("command name %s has been exists", c.Name)
	}

	a.names = append(a.names, c.Name)
	a.cmds[c.Name] = c
	a.NameWidth = mathutil.MaxInt(a.NameWidth, ln)

	// attach func
	if c.Func != nil {
		c.CFlags.Func = func(_ *CFlags) error {
			return c.Func(c)
		}
	}

	if c.OnAdd != nil {
		c.OnAdd(c)
	}
}

func (a *App) init() {
	if a.Name == "" {
		a.Name = path.Base(os.Args[0])
	}
}

// Run app by os.Args
func (a *App) Run() {
	err := a.RunWithArgs(os.Args[1:])
	if err != nil {
		cliutil.Errorln("ERROR:", err)
	}
}

// RunWithArgs run app by input args
func (a *App) RunWithArgs(args []string) error {
	a.init()

	if len(args) == 0 || args[0] == "" {
		return a.showHelp()
	}

	name := args[0]
	if name == "help" || name == "--help" || name == "-h" {
		return a.showHelp()
	}

	if name[0] == '-' {
		return fmt.Errorf("provide undefined flag %s", name)
	}

	cmd, ok := a.findCmd(name)
	if !ok {
		return fmt.Errorf("input not exists command %s", name)
	}

	return cmd.Parse(args[1:])
}

func (a *App) findCmd(name string) (*Cmd, bool) {
	cmd, ok := a.cmds[name]
	return cmd, ok
}

func (a *App) showHelp() error {
	bin := a.Name
	buf := strutil.NewBuffer()
	buf.QuietWritef("<cyan>%s</> - %s", bin, a.Desc)

	if a.Version != "" {
		buf.QuietWritef("(Version: <cyan>%s</>)", a.Version)
	}

	buf.QuietWritef("\n\n<comment>Usage:</> %s <green>COMMAND</> [--Options...] [...Arguments]\n", bin)

	buf.QuietWriteln("<comment>Options:</>:")
	buf.QuietWriteln("  <green>-h, --help</>     Display application help")
	buf.QuietWriteln("\n<comment>Commands</>:")

	sort.Strings(a.names)
	for _, name := range a.names {
		c := a.cmds[name]
		name := strutil.PadRight(name, " ", a.NameWidth)
		buf.QuietWritef("  <green>%s</>  %s\n", name, strutil.UpperFirst(c.Desc))
	}

	name := strutil.PadRight("help", " ", a.NameWidth)
	buf.QuietWritef("  <green>%s</>  Display application help\n", name)
	buf.QuietWritef("\nUse \"<cyan>%s COMMAND --help</>\" for about a command\n", bin)

	if a.AfterHelpBuild != nil {
		a.AfterHelpBuild(buf)
	}

	color.Print(buf.ResetAndGet())
	return nil
}

// Cmd struct
type Cmd struct {
	*CFlags
	Name  string
	Func  func(c *Cmd) error
	OnAdd func(c *Cmd)
}

// NewCmd instance
func NewCmd(name, desc string) *Cmd {
	fs := NewEmpty(func(c *CFlags) {
		c.Desc = desc
		c.FlagSet = flag.NewFlagSet(name, flag.ContinueOnError)
	})

	return &Cmd{
		Name:   name,
		CFlags: fs,
	}
}
