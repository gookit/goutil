# cflag

`cflag` wraps and extends Go's standard `flag.FlagSet` for small command-line applications.

- Same basic usage model as Go `flag`
- Pretty help output
- Short option aliases, including multiple aliases
- Required options and required positional arguments
- Named positional arguments
- Argument validators and option validators
- Options can appear before or after positional arguments
- Multi-command applications via `cflag/capp`

> **[中文说明](README.zh-CN.md)**

## Install

```shell
go get github.com/gookit/goutil/cflag
```

## Go Docs

- [github.com/gookit/goutil/cflag](https://pkg.go.dev/github.com/gookit/goutil/cflag)
- [github.com/gookit/goutil/cflag/capp](https://pkg.go.dev/github.com/gookit/goutil/cflag/capp)

## Quick Start

Examples are available in [_example/cmd.go](_example/cmd.go).

```go
package main

import (
	"fmt"

	"github.com/gookit/goutil/cflag"
)

var opts = struct {
	age  int
	name string
	str1 string
	lOpt string
}{}

func main() {
	c := cflag.New(func(c *cflag.CFlags) {
		c.Desc = "this is a demo command"
		c.Version = "0.5.1"
	})

	c.IntVar(&opts.age, "age", 0, "this is an int option;;a")
	c.StringVar(&opts.name, "name", "", "this is a string option and required;true")
	c.StringVar(&opts.str1, "str1", "def-val", "this is a string option with default value;;s")
	c.StringVar(&opts.lOpt, "long-opt", "", "this is a string option with shortcuts;;lo")

	c.AddArg("arg1", "this is arg1", true, nil)
	c.AddArg("arg2", "this is arg2", false, "default value")

	c.Func = func(c *cflag.CFlags) error {
		fmt.Println("command:", c.Name())
		fmt.Println("option.age:", opts.age)
		fmt.Println("option.name:", opts.name)
		fmt.Println("option.str1:", opts.str1)
		fmt.Println("option.lOpt:", opts.lOpt)
		fmt.Println("arg1:", c.Arg("arg1").String())
		fmt.Println("arg2:", c.Arg("arg2").String())
		fmt.Println("remain args:", c.RemainArgs())
		return nil
	}

	c.MustParse(nil)
}
```

Run it:

```shell
go run ./cflag/_example/cmd.go --name inhere -a 12 --lo val ab cd
go run ./cflag/_example/cmd.go ab --name inhere -a 12 --lo val cd
```

## Options

`cflag` extends the standard flag usage string to configure required options and shortcuts.

Usage format:

```text
desc
desc;required
desc;required;shorts
```

- `desc`: option description
- `required`: bool string, such as `true`, `on`, `yes`, `false`, `off`, `no`
- `shorts`: comma-separated aliases, such as `s` or `s,short`

Examples:

```go
// Required option.
c.StringVar(&opts.name, "name", "", "user name;true")

// Optional option with short alias "-s".
c.StringVar(&opts.str1, "str1", "def-val", "string value;;s")

// Optional option with aliases "-lo" and "-l".
c.StringVar(&opts.lOpt, "long-opt", "", "long option;;lo,l")
```

## Positional Arguments

Bind positional arguments with `AddArg`.

```go
c.AddArg("arg1", "this is arg1")
c.AddArg("arg2", "this is required arg2", true)
c.AddArg("arg3", "this arg has default value", false, "default value")
c.AddArg("extras", "array argument, must be the last one", false, nil, true)
```

Read values by name:

```go
fmt.Println(c.Arg("arg1").String())
fmt.Println(c.Arg("extras").Strings())
fmt.Println(c.RemainArgs())
```

`FlagArg` fields:

```go
type FlagArg struct {
	Name     string
	Desc     string
	Required bool
	Arrayed  bool // Arrayed argument must be the last bound argument.
}
```

## Option and Argument Order

For a single command, options can be placed before or after positional arguments.

These two commands are equivalent:

```shell
mycmd --name inhere -a 12 keyword more
mycmd keyword --name inhere more -a 12
```

`--` stops option parsing. Anything after it is treated as an argument:

```shell
mycmd keyword -- --name inhere
```

Result:

```text
arg keyword = keyword
remain args = [--name inhere]
```

Unknown options before the first positional argument still return a flag error. Unknown options after a positional argument are kept as arguments unless they match a registered option.

## Help Output

```shell
go run ./cflag/_example/cmd.go -h
```

![cmd-help](_example/cmd-help.png)

## Required Checks

```shell
go run ./cflag/_example/cmd.go -a 22
go run ./cflag/_example/cmd.go --name inhere
```

![cmd-required.png](_example/cmd-required.png)

## Multi-Command Apps

Use `cflag/capp` to build a small multi-command application.

Examples are available in [_example/app.go](_example/app.go).

```go
package main

import (
	"fmt"

	"github.com/gookit/goutil/cflag/capp"
)

var demoOpts = struct {
	age  int
	name string
}{}

func main() {
	app := capp.NewApp()
	app.Name = "myapp"
	app.Desc = "this is my cli application"
	app.Version = "1.0.2"

	cmd := capp.NewCmd("demo", "this is a demo command")
	cmd.IntVar(&demoOpts.age, "age", 0, "this is an int option;;a")
	cmd.StringVar(&demoOpts.name, "name", "", "this is a string option and required;true")
	cmd.AddArg("arg1", "this is arg1", true, nil)
	cmd.AddArg("arg2", "this is arg2", false, nil)
	cmd.Func = func(c *capp.Cmd) error {
		fmt.Println("age:", demoOpts.age)
		fmt.Println("name:", demoOpts.name)
		fmt.Println("arg1:", c.Arg("arg1").String())
		fmt.Println("arg2:", c.Arg("arg2").String())
		return nil
	}

	app.Add(cmd)
	app.Run()
}
```

Show commands:

```shell
go run ./cflag/_example/app.go -h
```

![app-help](_example/app-help.png)

Run a command:

```shell
go run ./cflag/_example/app.go demo --name inhere --age 333 val0 val1
go run ./cflag/_example/app.go demo val0 --name inhere val1 --age 333
```

![app-run](_example/app-run.png)

### Global Options and Command Options

`capp` parses global options only before the command name. Command options can appear anywhere inside the command arguments.

```shell
myapp --workdir /tmp demo val0 --name inhere -a 333 val1
```

Parse result:

```text
global option: --workdir /tmp
command: demo
command options: --name inhere, -a 333
command args: val0 val1
```

This avoids treating command options as global options.

Use `--` inside command arguments to stop command option parsing:

```shell
myapp demo val0 -- --name inhere
```

Here `--name inhere` remains in command arguments.
