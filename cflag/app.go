package cflag

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/gookit/color"
	"github.com/gookit/goutil"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// App struct
type App struct {
	cmds  map[string]*Cmd
	names []string

	Name string
	Desc string
	// NameWidth max width for command name
	NameWidth  int
	HelpWriter io.Writer
	// Version for app
	Version string

	// AfterHelpBuild hook
	AfterHelpBuild func(buf *strutil.Buffer)
	// BeforeRun hook func
	//  - cmdArgs: input raw args for current command.
	//  - return false to stop run.
	BeforeRun func(c *Cmd, cmdArgs []string) bool
	// AfterRun hook func
	AfterRun func(c *Cmd, err error)
}

// NewApp instance
func NewApp(fns ...func(app *App)) *App {
	app := &App{
		cmds: make(map[string]*Cmd),
		// with default version
		Version: "0.0.1",
		// NameWidth default value
		NameWidth:  12,
		HelpWriter: os.Stdout,
	}

	for _, fn := range fns {
		fn(app)
	}
	return app
}

// Add command(s) to app.
//
// NOTE: command object should create use NewCmd()
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
		goutil.Panicf("command name %s has been exists", c.Name)
	}

	a.names = append(a.names, c.Name)
	a.cmds[c.Name] = c
	a.NameWidth = mathutil.MaxInt(a.NameWidth, ln)

	// attach handle func
	if c.Func != nil {
		// fix: init c.CFlags on not exist
		if c.CFlags == nil {
			c.CFlags = NewEmpty(func(cf *CFlags) {
				cf.Desc = c.Desc
				cf.FlagSet = flag.NewFlagSet(c.Name, flag.ContinueOnError)
			})
		}

		c.CFlags.Func = func(_ *CFlags) error {
			return c.Func(c)
		}
	}

	if c.OnAdd != nil {
		c.OnAdd(c)
	}
}

// Run app by os.Args
func (a *App) Run() {
	err := a.RunWithArgs(os.Args[1:])
	if err != nil {
		cliutil.Errorln("ERROR:", err)
		if Debug {
			fmt.Println(errorx.Newf("(debug mode)RUNTIME ERROR: %v", err))
		}
		os.Exit(1)
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
		return fmt.Errorf("provide undefined flag option %q", name)
	}

	cmd, ok := a.findCmd(name)
	if !ok {
		return fmt.Errorf("input not exists command %q", name)
	}

	cmdArgs := args[1:]
	if a.BeforeRun != nil && !a.BeforeRun(cmd, cmdArgs) {
		return nil
	}

	err := cmd.Parse(cmdArgs)

	// fire after run hook
	if a.AfterRun != nil {
		a.AfterRun(cmd, err)
	}
	return err
}

func (a *App) init() {
	if a.Name == "" {
		// fix: path.Base not support windows
		a.Name = filepath.Base(os.Args[0])
	}
}

func (a *App) findCmd(name string) (*Cmd, bool) {
	cmd, ok := a.cmds[name]
	return cmd, ok
}

func (a *App) showHelp() error {
	bin := a.Name
	buf := strutil.NewBuffer()
	buf.Printf("<cyan>%s</> - %s", bin, a.Desc)

	if a.Version != "" {
		buf.Printf("(Version: <cyan>%s</>)", a.Version)
	}

	buf.Printf("\n\n<comment>Usage:</> %s <green>COMMAND</> [--Options...] [...Arguments]\n", bin)

	buf.WriteStr1Nl("<comment>Options:</>")
	buf.WriteStr1Nl("  <green>-h, --help</>     Display application help")
	buf.WriteStr1Nl("\n<comment>Commands:</>")

	sort.Strings(a.names)
	for _, name := range a.names {
		c := a.cmds[name]
		name := strutil.PadRight(name, " ", a.NameWidth)
		buf.Printf("  <green>%s</>  %s\n", name, strutil.UpperFirst(c.getDesc()))
	}

	name := strutil.PadRight("help", " ", a.NameWidth)
	buf.Printf("  <green>%s</>  Display application help\n", name)
	buf.Printf("\nUse \"<cyan>%s COMMAND --help</>\" for about a command\n", bin)

	if a.AfterHelpBuild != nil {
		a.AfterHelpBuild(buf)
	}

	if a.HelpWriter == nil {
		a.HelpWriter = os.Stdout
	}

	color.Fprint(a.HelpWriter, buf.ResetAndGet())
	return nil
}

// Cmd struct
type Cmd struct {
	*CFlags
	Name  string
	Desc  string // desc for command, will sync to CFlags.Desc
	OnAdd func(c *Cmd)
	Func  func(c *Cmd) error
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

// Config the cmd
func (c *Cmd) Config(fn func(c *Cmd)) *Cmd {
	if fn != nil {
		fn(c)
	}
	return c
}

func (c *Cmd) getDesc() string {
	if c.CFlags.Desc != "" {
		return c.CFlags.Desc
	}
	return c.Desc
}
