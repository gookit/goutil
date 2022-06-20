package cflag

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/strutil"
)

// FlagArg struct
type FlagArg struct {
	// Value for the flag argument
	*structs.Value
	// Name of the argument
	Name string
	// Desc arg description
	Desc string
	// Index of the argument
	Index int
	// Required argument
	Required bool
	// Value for the flag argument
	// Value string
	// wrapped the default value
	// val *strutil.Value
}

// check string
func (a *FlagArg) check() error {
	if a.Required && a.V != nil {
		return errorx.Rawf("cannot set default value for 'required' arg: %s", a.Name)
	}

	return nil
}

// HelpDesc string
func (a *FlagArg) HelpDesc() string {
	desc := a.Desc
	if desc == "" {
		desc = "no description"
	}

	desc = strutil.UpperFirst(desc)
	if a.Required {
		desc = "<red>*</>" + desc
	}

	if a.V != nil {
		desc += "(Default:" + a.String() + ")"
	}
	return desc
}

// CFlags struct
type CFlags struct {
	*flag.FlagSet
	// aliases for option flags
	// aliases structs.Aliases
	required map[string]int8
	// argVals  maputil.Data
	argWidth int
	// bind arguments.
	bindArgs map[string]*FlagArg
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

// NewCFlags instance.
func NewCFlags(fns ...func(c *CFlags)) *CFlags {
	return NewEmptyCFlags(func(c *CFlags) {
		c.FlagSet = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	}).WithConfigFn(fns...)
}

// NewEmptyCFlags instance.
func NewEmptyCFlags(fns ...func(c *CFlags)) *CFlags {
	c := &CFlags{
		argWidth: 12,
		required: make(map[string]int8),
		bindArgs: make(map[string]*FlagArg),
	}

	return c.WithConfigFn(fns...)
}

// WithConfigFn for command
func (c *CFlags) WithConfigFn(fns ...func(c *CFlags)) *CFlags {
	for _, fn := range fns {
		fn(c)
	}
	return c
}

// WithDesc for command
func (c *CFlags) WithDesc(desc string) *CFlags {
	c.Desc = desc
	return c
}

// WithFunc for command
func (c *CFlags) WithFunc(fn func(c *CFlags) error) *CFlags {
	c.Func = fn
	return c
}

// AddArg binding for command
func (c *CFlags) AddArg(name, desc string, required bool, value interface{}) {
	arg := &FlagArg{
		Name:  name,
		Desc:  desc,
		Index: len(c.bindArgs),
		Value: structs.NewValue(value),
		// required
		Required: required,
	}

	err := arg.check()
	if err != nil {
		panic(err)
	}

	c.bindArgs[name] = arg
}

// MustParse for command
func (c *CFlags) MustParse(args []string) {
	err := c.Parse(args)
	if err != nil {
		cliutil.Redln("ERROR:", err)
	}
}

// Parse for command
func (c *CFlags) Parse(args []string) error {
	defer func() {
		if err := recover(); err != nil {
			cliutil.Errorln("ERROR:", err)
		}
	}()

	// prepare
	if err := c.prepare(); err != nil {
		return err
	}

	// do parsing
	err := c.FlagSet.Parse(args)
	if err != nil {
		if err == flag.ErrHelp {
			return nil // ignore help error
		}
		return err
	}

	// check required
	if err := c.checkRequired(); err != nil {
		return err
	}

	// do binding args
	if err := c.bindParsedArgs(); err != nil {
		return err
	}

	// call func
	if c.Func != nil {
		return c.Func(c)
	}
	return nil
}

func (c *CFlags) prepare() error {
	c.VisitAll(func(f *flag.Flag) {
		name := f.Name
		desc := strings.Trim(f.Usage, "; ")

		if strings.ContainsRune(desc, ';') {
			// format: desc;shorts;required
			// format: desc;required
			first, second := strutil.MustCut(desc, ";")
			if bl, err := strutil.Bool(second); err == nil && bl {
				desc = first
				c.required[name] = 1
			}
		}

		f.Usage = strutil.UpperFirst(desc)
	})

	// custom something
	c.FlagSet.Usage = c.ShowHelp

	return nil
}

// check required option flags
func (c *CFlags) checkRequired() error {
	for name, _ := range c.required {
		opt := c.Lookup(name)
		if opt.Value.String() == "" {
			return errorx.Rawf("flag option '%s' is required", AddPrefix(name))
		}
	}
	return nil
}

// desc for command
func (c *CFlags) bindParsedArgs() error {
	args := c.Args()
	argN := len(args) - 1

	var lastIdx int
	for name, arg := range c.bindArgs {
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
	arg, ok := c.bindArgs[name]
	if !ok {
		panic("cflag: get not binding argument: " + name)
	}
	return arg
}

// RemainArgs get
func (c *CFlags) RemainArgs() []string {
	return c.remainArgs
}

// Name for command
func (c *CFlags) Name() string {
	return path.Base(c.FlagSet.Name())
}

// BinFile path for command
func (c *CFlags) BinFile() string {
	return c.FlagSet.Name()
}

// desc for command
func (c *CFlags) helpDesc() string {
	desc := strutil.UpperFirst(c.Desc)

	if c.Version != "" {
		desc += "v(" + c.Version + ")"
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
		"{{command}}": binName,
		"{{binName}}": binName,
		"{{binFile}}": c.BinFile(),
	}

	out := c.Output()
	buf := new(strutil.Buffer)
	c.SetOutput(buf)
	if err != nil {
		buf.QuietWritef("<error>ERROR:</> %s\n", err.Error())
	} else {
		buf.QuietWritef("<cyan>%s</>\n\n", c.helpDesc())
	}

	buf.QuietWritef("<comment>Usage:</> %s [--Options...] [...Arguments]\n", binName)
	buf.QuietWriteString("<comment>Options:</>\n")
	// c.FlagSet.PrintDefaults()
	c.ShowOptionsHelp()

	if len(c.bindArgs) > 0 {
		buf.QuietWriteString("\n<comment>Arguments:</>\n")
		for name, arg := range c.bindArgs {
			buf.QuietWritef("  <green>%s</>   %s\n", strutil.PadRight(name, " ", c.argWidth), arg.HelpDesc())
		}
	}

	if c.LongHelp != "" {
		buf.QuietWriteln("<comment>Help:</>")
		buf.QuietWriteln(c.LongHelp)
	}

	if c.Example != "" {
		buf.QuietWriteln("<comment>Examples:</>")
		buf.QuietWriteln(c.Example)
	}

	color.Println(strutil.Replaces(buf.String(), helpVars))
	c.SetOutput(out) // revert output
}

// ShowOptionsHelp prints, to standard error unless configured otherwise, the
// default values of all defined command-line flags in the set. See the
// documentation for the global function PrintDefaults for more information.
//
// from flag.PrintDefaults
func (c *CFlags) ShowOptionsHelp() {
	c.VisitAll(func(opt *flag.Flag) {
		var b strings.Builder
		// Two spaces before -; see next two comments.
		if len(opt.Name) > 1 {
			_, _ = fmt.Fprintf(&b, "  <info>--%s</>", opt.Name)
		} else {
			_, _ = fmt.Fprintf(&b, "  <info>-%s</>", opt.Name)
		}

		name, usage := flag.UnquoteUsage(opt)
		if len(name) > 0 {
			b.WriteString(" ")
			b.WriteString(name)
		}

		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if b.Len() <= 4 { // space, space, '-', 'x'.
			b.WriteString("\t")
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			b.WriteString("\n    \t")
		}
		b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))

		if isZero, isStr := IsZeroValue(opt, opt.DefValue); !isZero {
			if isStr {
				// put quotes on the value
				_, _ = fmt.Fprintf(&b, " (default <magentaB>%q</>)", opt.DefValue)
			} else {
				_, _ = fmt.Fprintf(&b, " (default <magentaB>%v</>)", opt.DefValue)
			}
		}

		_, _ = fmt.Fprint(c.Output(), b.String(), "\n")
	})
}

// WrapStdFlag wrap the go flag.CommandLine instance
func WrapStdFlag() *CFlags {
	return NewEmptyCFlags(func(c *CFlags) {
		c.FlagSet = flag.CommandLine
		c.Usage = c.ShowHelp
	})
}
