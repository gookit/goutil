package main

import (
	"os"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/cliutil"
)

var opts = struct {
	age  int
	name string
	str1 string
	lOpt string
	bol  bool
}{}

// go run ./_example/cmd.go
// go run ./cflag/_example/cmd.go -h
// go run ./cflag/_example/cmd.go --name inhere -a 12 --lo val ab cd
func main() {
	c := cflag.New(func(c *cflag.CFlags) {
		c.Desc = "this is a demo command"
		c.Version = "0.5.1"
	})
	c.IntVar(&opts.age, "age", 0, "this is a int option;;a")
	c.StringVar(&opts.name, "name", "", "this is a string option and required;true")
	c.StringVar(&opts.str1, "str1", "def-val", "this is a string option with default value;;s")
	c.StringVar(&opts.lOpt, "long-opt", "", "this is a string option with shorts;;lo")

	c.AddArg("arg1", "this is arg1", true, nil)
	c.AddArg("arg2", "this is arg2", true, nil)
	// c.AddArg("arg2", "this is arg2", false, "def-val")

	c.Func = func(c *cflag.CFlags) error {
		cliutil.Magentaln("hello, this is command:", c.Name())
		cliutil.Yellowln("option values:")
		cliutil.Infoln("opts.age =", opts.age)
		cliutil.Infoln("opts.name =", opts.name)
		cliutil.Infoln("opts.str1 =", opts.str1)
		cliutil.Infoln("opts.lOpt =", opts.lOpt)
		cliutil.Yellowln("argument values:")
		cliutil.Infoln("arg1 =", c.Arg("arg1").String())
		cliutil.Infoln("arg2 =", c.Arg("arg2").String())

		cliutil.Infoln("\nremain args =", c.RemainArgs())

		return nil
	}

	c.MustParse(os.Args[1:])
}
