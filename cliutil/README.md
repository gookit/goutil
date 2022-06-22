# Cli Util

- Helper util functions in cli
- Color print in console terminal
- Read terminal message input
- Build command line application
- Command line args string parse
- Build command line string from []string

## Install

```shell
go get github.com/gookit/goutil/cliutil
```

## Helper functions

```go
cliutil.Workdir() // current workdir
cliutil.BinDir() // the program exe file dir
cliutil.QuickExec("echo $SHELL")
```

## Create CLI app

```go
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

	c.MustParse(os.Args[1:])
}
```

Show help:

![cmd-help](_demo/cmd-help.png)

## Color print

Quick color print in console:

```go
cliutil.Redln("ln:red color message print in cli.")
cliutil.Blueln("ln:blue color message print in cli.")
cliutil.Cyanln("ln:cyan color message print in cli.")
```

![color-print](_demo/color-print.png)

## Read input

```go
name := cliutil.ReadInput("Your name: ")
name := cliutil.ReadLine("Your name: ")

ans, _ := cliutil.ReadFirst("continue?[y/n] ")

pwd := cliutil.ReadPassword("Input password:")
```

## Parse command line as args

parse input command line to `[]string`, such as cli `os.Args`

```go
package main

import (
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
)

func main() {
	args := cliutil.ParseLine(`./app top sub --msg "has multi words"`)
	dump.P(args)
}
```

**output**:

```text
PRINT AT github.com/gookit/goutil/cliutil_test.TestParseLine(line_parser_test.go:30)
[]string [ #len=5
  string("./app"), #len=5
  string("top"), #len=3
  string("sub"), #len=3
  string("--msg"), #len=5
  string("has multi words"), #len=15
]
```

## Build command line from args

```go
	s := cliutil.BuildLine("./myapp", []string{
		"-a", "val0",
		"-m", "this is message",
		"arg0",
	})
	fmt.Println("Build line:", s)
```

**output**:

```text
Build line: ./myapp -a val0 -m "this is message" arg0
```
