package main

import (
	"github.com/gookit/goutil/cflag/capp"
	"github.com/gookit/goutil/dump"
)

var globalOpts = struct {
	test1      string
	workdir    string
	longOption string
}{}

var c1Opts = struct {
	age  int
	name string
}{}

var c2Opts = struct {
	str1 string
	lOpt string
	bol  bool
}{}

// go run ./_example/app.go
// go run ./cflag/_example/app.go -h
// go run ./cflag/_example/app.go demo -h
func main() {
	app := capp.NewApp()
	app.Desc = "this is my cli application"
	app.Version = "1.0.2"

	// global flags
	app.StringVar(&globalOpts.test1, "test1", "", "this is a global string option1")
	app.StringVar(&globalOpts.workdir, "workdir", "", "this is a global string option\nnew line text;off;w")
	app.StringVar(&globalOpts.longOption, "long-long-long-long-name", "", "this is a long long long ... option name")

	// go run ./cflag/_example/app.go demo --name inhere --age 333 val0 val1
	c1 := capp.NewCmd("demo", "this is a demo command")
	c1.OnAdd = func(c *capp.Cmd) {
		c.IntVar(&c1Opts.age, "age", 0, "this is a int option;;a")
		c.StringVar(&c1Opts.name, "name", "", "this is a string option and required;true")

		c.AddArg("arg1", "this is arg1", true, nil)
		c.AddArg("arg2", "this is arg2", false, nil)
	}
	c1.Func = func(c *capp.Cmd) error {
		dump.P(c1Opts, c.Args())
		return nil
	}

	c2 := capp.NewCmd("other", "this is another demo command", func(c *capp.Cmd) error {
		dump.P(c2Opts)
		return nil
	})
	{
		c2.StringVar(&c2Opts.str1, "str1", "def-val", "this is a string option with default value;;s")
		c2.StringVar(&c2Opts.lOpt, "long-opt", "", "this is a string option with shorts;;lo")
	}

	app.Add(c1, c2)
	app.Run()
}
