package cflag_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func ExampleNewApp() {
	app := cflag.NewApp()
	app.Desc = "this is my cli application"
	app.Version = "1.0.2"

	var c1Opts = struct {
		age  int
		name string
	}{}

	c1 := cflag.NewCmd("demo", "this is a demo command")
	c1.OnAdd = func(c *cflag.Cmd) {
		c.IntVar(&c1Opts.age, "age", 0, "this is a int option;;a")
		c.StringVar(&c1Opts.name, "name", "", "this is a string option and required;true")

		c.AddArg("arg1", "this is arg1", true, nil)
		c.AddArg("arg2", "this is arg2", false, nil)
	}

	c1.Func = func(c *cflag.Cmd) error {
		dump.P(c1Opts, c.Args())
		return nil
	}

	var c2Opts = struct {
		str1 string
		lOpt string
		bol  bool
	}{}

	c2 := cflag.NewCmd("other", "this is another demo command")
	{
		c2.StringVar(&c2Opts.str1, "str1", "def-val", "this is a string option with default value;;s")
		c2.StringVar(&c2Opts.lOpt, "long-opt", "", "this is a string option with shorts;;lo")

		c2.Func = func(c *cflag.Cmd) error {
			dump.P(c2Opts)
			return nil
		}
	}

	app.Add(c1, c2)
	app.Run()
}

func TestApp_Run(t *testing.T) {
	app := cflag.NewApp(func(app *cflag.App) {
		app.Name = "myapp"
		app.Desc = "this is my cli application"
		app.Version = "1.0.2"
	})

	// test invalid
	assert.Panics(t, func() {
		app.Add(cflag.NewCmd("", "empty name"))
	})

	// add command
	var c1Opts = struct {
		age  int
		name string
	}{}

	c1 := cflag.NewCmd("demo", "this is a demo command")
	c1.OnAdd = func(c *cflag.Cmd) {
		c.IntVar(&c1Opts.age, "age", 0, "this is a int option;;a")
		c.StringVar(&c1Opts.name, "name", "", "this is a string option and required;true")

		c.AddArg("arg1", "this is arg1", true, nil)
		c.AddArg("arg2", "this is arg2", false, nil)
	}
	c1.Config(func(c *cflag.Cmd) {
		c.Func = func(c *cflag.Cmd) error {
			dump.P(c1Opts, c.Args())
			return nil
		}
	})

	app.Add(c1)

	// repeat name
	assert.Panics(t, func() {
		app.Add(c1)
	})

	// add cmd by struct
	app.Add(&cflag.Cmd{
		Name: "demo2",
		Desc: "this is demo2 command",
		Func: func(c *cflag.Cmd) error {
			dump.P("hi, on demo2 command")
			return nil
		},
	})

	// show help1
	osArgs := os.Args
	os.Args = []string{"./myapp"}
	app.AfterHelpBuild = func(buf *strutil.Buffer) {
		help := buf.String()
		assert.StrContains(t, help, "-h, --help")
		assert.StrContains(t, help, "demo")
		assert.StrContains(t, help, "This is a demo command")
	}
	app.Run()
	os.Args = osArgs

	// show help2
	buf := new(bytes.Buffer)
	app.HelpWriter = buf
	err := app.RunWithArgs([]string{"--help"})
	assert.NoErr(t, err)

	help := buf.String()
	assert.StrContains(t, help, "-h, --help")
	assert.StrContains(t, help, "demo")
	assert.StrContains(t, help, "This is a demo command")

	// run ... error
	err = app.RunWithArgs([]string{"notExists"})
	assert.ErrMsg(t, err, `input not exists command "notExists"`)

	err = app.RunWithArgs([]string{"--invalid"})
	assert.ErrMsg(t, err, `provide undefined flag option "--invalid"`)

	// run
	err = app.RunWithArgs([]string{"demo", "-a", "230", "--name", "inhere", "val1"})
	assert.NoErr(t, err)
	assert.Eq(t, 230, c1Opts.age)
	assert.Eq(t, "inhere", c1Opts.name)
	assert.Eq(t, "val1", c1.Arg("arg1").String())
}
