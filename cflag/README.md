# cflag

`cflag` - Wraps and extends go `flag.FlagSet` to build simple command line applications

- Support auto render a pretty help panel
- Allow to add shortcuts for flag option
- Allow binding named arguments
- Allow set required for argument or option
- Allow set validator for argument or option

## Install

```shell
go get github.com/gookit/goutil/cflag
```

## Usage

```go
package main

import (
	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/cliutil"
)

// go run ./_example/cmd.go
// go run ./cflag/_example/cmd.go -h
// go run ./cflag/_example/cmd.go --name inhere ab cd
func main() {
	opts := struct {
		age  int
		name string
		str1 string
		lOpt string
		bol  bool
	}{}

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

	c.Func = func(c *cflag.CFlags) error {
		cliutil.Infoln("hello, this is command:", c.Name())
		cliutil.Infoln("option.age =", opts.age)
		cliutil.Infoln("option.name =", opts.name)
		cliutil.Infoln("option.str1 =", opts.str1)
		cliutil.Infoln("option.lOpt =", opts.lOpt)
		cliutil.Infoln("arg1 =", c.Arg("arg1").String())
		cliutil.Infoln("arg2 =", c.Arg("arg2").String())

		return nil
	}

	// c.MustParse(os.Args[1:])
	c.MustParse(nil)
}
```

Show help:

![cmd-help](_example/cmd-help.png)
