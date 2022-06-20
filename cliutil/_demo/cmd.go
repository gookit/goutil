package main

import (
	"os"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/cliutil/cflag"
)

// go run ./_demo/cmd.go
// go run ./cliutil/_demo/cmd.go -h
// go run ./cliutil/_demo/cmd.go --name inhere ab cd
func main() {
	opts := struct {
		age  int
		name string
		str1 string
		bol  bool
	}{}

	c := cflag.NewCFlags()
	c.WithDesc("this is a demo command")
	c.IntVar(&opts.age, "age", 0, "this is a int option;;")
	c.StringVar(&opts.name, "name", "", "this is a string option and required;true")
	c.StringVar(&opts.str1, "str1", "def-val", "this is a string option with default value")

	c.AddArg("arg1", "this is arg1", true, nil)
	c.AddArg("arg2", "this is arg2", true, nil)
	// c.AddArg("arg2", "this is arg2", false, "def-val")

	c.Func = func(c *cflag.CFlags) error {
		cliutil.Infoln("hello, this is", c.Name())
		cliutil.Infoln("option.name =", opts.name)
		cliutil.Infoln("option.str1 =", opts.str1)
		cliutil.Infoln("arg1 =", c.Arg("arg1").String())
		cliutil.Infoln("arg2 =", c.Arg("arg2").String())

		return nil
	}

	c.MustParse(os.Args[1:])
}
