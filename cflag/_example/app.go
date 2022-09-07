package main

import (
	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/dump"
)

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
	app := cflag.NewApp()
	app.Desc = "this is my cli application"
	app.Version = "1.0.2"

	// go run ./cflag/_example/app.go demo --name inhere --age 333 val0 val1
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
