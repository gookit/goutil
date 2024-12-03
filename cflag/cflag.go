// Package cflag Wraps and extends go `flag.FlagSet` to build simple command line applications
//
// - Support auto render a pretty help panel
//
// - Allow to add shortcuts for flag option
//
// - Allow binding named arguments
//
// - Allow set required for argument or option
//
// - Allow set validator for argument or option
package cflag

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil"
	"github.com/gookit/goutil/basefn"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/stdio"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/strutil"
)

// Debug mode
var Debug = envutil.GetBool("CFLAG_DEBUG")

// SetDebug mode
func SetDebug(open bool) {
	Debug = open
}

// CFlags wrap and extends the go flag.FlagSet
//
// eg:
//
//	// Can be set required and shorts on desc:
//	// format1: desc;required
//	cmd.IntVar(&age, "age", 0, "your age;true")
//	// format2: desc;required;shorts
//	cmd.IntVar(&age, "age", 0, "your age;true;a")
type CFlags struct {
	*flag.FlagSet
	// bound options.
	bindOpts map[string]*FlagOpt
	// shortcuts map for options. eg: n -> name
	shortcuts map[string]string

	// argWidth max width value
	argWidth int
	// bound arguments.
	bindArgs []*FlagArg
	// the argument name to index map.
	argNames map[string]int
	// remainArgs after binding args
	remainArgs []string

	// Desc command description
	Desc string
	// Version command version number
	Version string
	// Example command usage examples
	Example string
	// LongHelp custom help
	LongHelp string
	// Func handler for the command
	Func func(c *CFlags) error
}

// New create new instance.
//
// Usage:
//
//	cmd := cflag.New(func(c *cflag.CFlags) {
//		c.Version = "0.1.2"
//		c.Desc = "this is my cli tool"
//	})
//
//	// binding opts and args
//
//	cmd.Parse(nil)
func New(fns ...func(c *CFlags)) *CFlags {
	return NewEmpty(func(c *CFlags) {
		c.FlagSet = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	}).WithConfigFn(fns...)
}

// NewEmpty instance.
func NewEmpty(fns ...func(c *CFlags)) *CFlags {
	c := &CFlags{
		argWidth:  12,
		shortcuts: make(map[string]string),
		bindOpts:  make(map[string]*FlagOpt),
		argNames:  make(map[string]int, 4),
	}

	return c.WithConfigFn(fns...)
}

/*************************************************************
 * config command flags
 *************************************************************/

// WithDesc for command
func WithDesc(desc string) func(c *CFlags) {
	return func(c *CFlags) {
		c.Desc = desc
	}
}

// WithVersion for command
func WithVersion(version string) func(c *CFlags) {
	return func(c *CFlags) {
		c.Version = version
	}
}

// WithConfigFn for command
func (c *CFlags) WithConfigFn(fns ...func(c *CFlags)) *CFlags {
	for _, fn := range fns {
		fn(c)
	}
	return c
}

// AddValidator for a flag option
func (c *CFlags) AddValidator(name string, fn OptCheckFn) {
	c.ConfigOpt(name, func(opt *FlagOpt) {
		opt.Validator = fn
	})
}

// ConfigOpt for a flag option
func (c *CFlags) ConfigOpt(name string, fn func(opt *FlagOpt)) {
	if c.Lookup(name) == nil {
		goutil.Panicf("cflag: option '%s' is not registered", name)
	}

	// init on not exist
	if _, ok := c.bindOpts[name]; !ok {
		c.bindOpts[name] = &FlagOpt{}
	}

	fn(c.bindOpts[name])
}

// AddShortcuts for option flag
func (c *CFlags) AddShortcuts(name string, shorts ...string) {
	c.addShortcuts(name, shorts)
	c.ConfigOpt(name, func(opt *FlagOpt) {
		opt.Shortcuts = append(opt.Shortcuts, shorts...)
	})
}

// addShortcuts for option flag
func (c *CFlags) addShortcuts(name string, shorts []string) {
	for _, short := range shorts {
		if regName, ok := c.shortcuts[short]; ok {
			goutil.Panicf("cflag: shortcut '%s' has been used by option '%s'", short, regName)
		}

		c.shortcuts[short] = name
	}
}

// AddArg binding for command
func (c *CFlags) AddArg(name, desc string, required bool, value any) {
	arg := &FlagArg{
		Name:  name,
		Desc:  desc,
		Value: structs.NewValue(value),
		// required
		Required: required,
	}

	c.BindArg(arg)
}

// BindArg for command
func (c *CFlags) BindArg(arg *FlagArg) {
	arg.Index = len(c.bindArgs)

	// check arg info
	basefn.PanicErr(arg.check())

	if _, ok := c.argNames[arg.Name]; ok {
		basefn.Panicf("cflag: arg '%s' have been registered", arg.Name)
	}

	// register
	c.bindArgs = append(c.bindArgs, arg)
	c.argWidth = mathutil.MaxInt(c.argWidth, len(arg.Name))
	c.argNames[arg.Name] = arg.Index
}

/*************************************************************
 * parse command flags
 *************************************************************/

// QuickRun parse OS flags and run command, will auto handle error
func (c *CFlags) QuickRun() {
	c.MustParse(nil)
}

// MustRun parse flags and run command. alias of MustParse()
func (c *CFlags) MustRun(args []string) {
	c.MustParse(args)
}

// MustParse parse flags and run command, will auto handle error
func (c *CFlags) MustParse(args []string) {
	err := c.Parse(args)
	if err != nil {
		cliutil.Redln("ERROR:", err)
	}
}

// Parse flags and run command func
//
// If args is nil, will parse os.Args
func (c *CFlags) Parse(args []string) error {
	if args == nil {
		args = os.Args[1:]
	}

	defer func() {
		if err := recover(); err != nil {
			cliutil.Errorln("ERROR:", err)
			if Debug {
				fmt.Println(errorx.Newf("(debug mode)RECOVERD PARSE ERROR: %v", err))
			}
		}
	}()

	// prepare
	if err := c.prepare(); err != nil {
		return err
	}

	// do parsing
	if err := c.doParse(args); err != nil {
		if err == flag.ErrHelp {
			return nil // ignore help error
		}
		return err
	}

	// call func
	if c.Func != nil {
		return c.Func(c)
	}
	return nil
}

func (c *CFlags) prepare() error {
	// dont use flag output.
	c.SetOutput(io.Discard)

	// parse flag usage string
	c.VisitAll(func(f *flag.Flag) {
		if regName, ok := c.shortcuts[f.Name]; ok {
			goutil.Panicf("cflag: name '%s' has been as shortcut by '%s'", f.Name, regName)
		}

		f.Usage = c.parseFlagUsage(f.Name, f.Usage)
	})

	// custom something
	c.FlagSet.Usage = c.ShowHelp
	return nil
}

// do parse flag.Usage string.
func (c *CFlags) parseFlagUsage(name, usage string) string {
	opt, ok := c.bindOpts[name]
	if !ok {
		c.bindOpts[name] = &FlagOpt{}
		opt = c.bindOpts[name]
	}

	desc := strings.Trim(usage, "; ")
	if !strings.ContainsRune(desc, ';') {
		return strutil.UpperFirst(desc)
	}

	// format: desc;required OR desc;required;shorts
	parts := strutil.SplitNTrimmed(desc, ";", 3)
	if ln := len(parts); ln > 1 {
		// required
		if bl, err := strutil.Bool(parts[1]); err == nil && bl {
			desc = "<red>*</>" + strutil.UpperFirst(parts[0])
			opt.Required = true
		} else {
			desc = strutil.UpperFirst(parts[0])
		}

		// shortcuts
		if ln > 2 && len(parts[2]) > 0 {
			opt.Shortcuts = SplitShortcut(parts[2])
			c.addShortcuts(name, opt.Shortcuts)
		}
	}

	return desc
}

// do parse and validate
func (c *CFlags) doParse(args []string) error {
	if len(c.shortcuts) > 0 && len(args) > 0 {
		args = ReplaceShorts(args, c.shortcuts)
	}

	// do parsing
	if err := c.FlagSet.Parse(args); err != nil {
		return err
	}

	// check option values
	if err := c.checkBindOpts(); err != nil {
		return err
	}

	return c.bindParsedArgs()
}

// check bind option flags
func (c *CFlags) checkBindOpts() error {
	for name, opt := range c.bindOpts {
		fv := c.Lookup(name).Value
		if opt.Required && fv.String() == "" {
			return errorx.Rawf("flag option '%s' is required", name)
		}

		if opt.Validator == nil {
			continue
		}

		// call validator
		if fg, ok := fv.(flag.Getter); ok {
			err := opt.Validator(fg.Get())
			if err != nil {
				return errorx.Rawf("flag option '%s': %s", name, err.Error())
			}
		}
	}
	return nil
}

// desc for command
func (c *CFlags) bindParsedArgs() error {
	args := c.Args()
	argN := len(args) - 1

	var lastIdx int
	for _, arg := range c.bindArgs {
		name := arg.Name
		if arg.Index > argN {
			if arg.Required {
				return errorx.Rawf("argument '%s'(#%d) is required", name, arg.Index)
			}
			break
		}

		lastIdx++
		val := args[arg.Index]
		if arg.Required && val == "" {
			return errorx.Rawf("argument '%s'(#%d) is required", name, arg.Index)
		}

		arg.V = val
	}

	// collect remain args
	if lastIdx < argN {
		c.remainArgs = args[lastIdx:]
	}
	return nil
}

// Arg get by bind name
func (c *CFlags) Arg(name string) *FlagArg {
	idx, ok := c.argNames[name]
	if !ok {
		goutil.Panicf("cflag: get not binding arg '%s'", name)
	}
	return c.bindArgs[idx]
}

// RemainArgs get
func (c *CFlags) RemainArgs() []string { return c.remainArgs }

// Name for command
func (c *CFlags) Name() string {
	return filepath.Base(c.FlagSet.Name())
}

// BinFile path for command
func (c *CFlags) BinFile() string { return c.FlagSet.Name() }

/*************************************************************
 * render command help
 *************************************************************/

// desc for command
func (c *CFlags) helpDesc() string {
	desc := strutil.UpperFirst(c.Desc)

	if c.Version != "" {
		desc += "(v" + c.Version + ")"
	}
	return desc
}

// ShowHelp for command
func (c *CFlags) ShowHelp() {
	c.showHelp(nil)
}

// show help for command
func (c *CFlags) showHelp(err error) {
	binName := c.Name()
	helpVars := map[string]string{
		"{{cmd}}":     binName,
		"{{command}}": binName,
		"{{binName}}": binName,
		"{{binFile}}": c.BinFile(),
	}

	buf := new(strutil.Buffer)

	if err != nil {
		buf.Printf("<error>ERROR:</> %s\n", err.Error())
	} else {
		buf.Printf("<cyan>%s</>\n\n", c.helpDesc())
	}

	buf.Printf("<comment>Usage:</> %s [--Options...] [...CliArgs]\n", binName)
	buf.WriteStr("<comment>Options:</>\n")

	// render options help
	c.renderOptionsHelp(buf)

	if len(c.bindArgs) > 0 {
		buf.WriteStr1("\n<comment>CliArgs:</>\n")
		for _, arg := range c.bindArgs {
			buf.Printf(
				"  <green>%s</>   %s\n",
				strutil.PadRight(arg.Name, " ", c.argWidth),
				arg.HelpDesc(),
			)
		}
	}

	if c.LongHelp != "" {
		buf.WriteStr1Nl("\n<comment>Help:</>")
		buf.WriteStr1Nl(strings.Trim(c.LongHelp, "\n"))
	}

	if c.Example != "" {
		buf.WriteStr1Nl("\n<comment>Examples:</>")
		buf.WriteStr1(strings.Trim(c.Example, "\n"))
	}

	color.Println(strutil.Replaces(buf.String(), helpVars))
}

// ShowOptionsHelp prints, to standard error unless configured otherwise, the
// default values of all defined command-line flags in the set. See the
// documentation for the global function PrintDefaults for more information.
//
// from flag.PrintDefaults
func (c *CFlags) renderOptionsHelp(buf *strutil.Buffer) {
	c.VisitAll(func(opt *flag.Flag) {
		var b strings.Builder

		mate := c.bindOpts[opt.Name]
		stdio.QuietFprintf(&b, "  <info>%s</>", mate.HelpName(opt.Name))

		typName, usage := flag.UnquoteUsage(opt)
		if len(typName) > 0 {
			b.WriteString(" ")
			b.WriteString(typName)
		}

		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if b.Len() <= 32 { // space, space, '-', 'x'.
			b.WriteString("\t")
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			b.WriteString("\n          \t")
		}
		b.WriteString(strings.ReplaceAll(usage, "\n", "\n          \t"))

		// put quotes on the string value
		if isZero, isStr := IsZeroValue(opt, opt.DefValue); !isZero {
			if isStr {
				stdio.QuietFprintf(&b, " (default <magentaB>%q</>)", opt.DefValue)
			} else {
				stdio.QuietFprintf(&b, " (default <magentaB>%v</>)", opt.DefValue)
			}
		}

		b.WriteByte('\n')
		buf.WriteStr1(b.String())
	})
}
