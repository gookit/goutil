# Cli Util

`cliutil` provides some extra util functions for CLI.

- Helper util functions in cli
- Color print in console terminal
- Read terminal message input
- Command line args string parse
- Build command line string from `[]string`

## Install

```shell
go get github.com/gookit/goutil/cliutil
```

## Go docs

- [Go docs](https://pkg.go.dev/github.com/gookit/goutil/cliutil)

## Helper functions

```go
cliutil.Workdir() // current workdir
cliutil.BinDir() // the program exe file dir
cliutil.QuickExec("echo $SHELL")
```

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
if cliutil.InputIsYes(ans) {
	// do something ...
}

ans, _ := cliutil.ReadFirstByte("continue?[y/n] ")
if cliutil.ByteIsYes(ans) {
	// do something ...
}
```

### Read Password

```go
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

## Functions API

> **Note**: doc by run `go doc ./fsutil`

```go
func BinDir() string
func BinFile() string
func BinName() string
func Bluef(format string, a ...interface{})
func Blueln(a ...interface{})
func Bluep(a ...interface{})
func BuildLine(binFile string, args []string) string
func BuildOptionHelpName(names []string) string
func CurrentShell(onlyName bool) (path string)
func Cyanf(format string, a ...interface{})
func Cyanln(a ...interface{})
func Cyanp(a ...interface{})
func Errorf(format string, a ...interface{})
func Errorln(a ...interface{})
func Errorp(a ...interface{})
func ExecCmd(binName string, args []string, workDir ...string) (string, error)
func ExecCommand(binName string, args []string, workDir ...string) (string, error)
func ExecLine(cmdLine string, workDir ...string) (string, error)
func FirstLine(output string) string
func Grayf(format string, a ...interface{})
func Grayln(a ...interface{})
func Grayp(a ...interface{})
func Greenf(format string, a ...interface{})
func Greenln(a ...interface{})
func Greenp(a ...interface{})
func HasShellEnv(shell string) bool
func Infof(format string, a ...interface{})
func Infoln(a ...interface{})
func Infop(a ...interface{})
func LineBuild(binFile string, args []string) string
func Magentaf(format string, a ...interface{})
func Magentaln(a ...interface{})
func Magentap(a ...interface{})
func OutputLines(output string) []string
func ParseLine(line string) []string
func QuickExec(cmdLine string, workDir ...string) (string, error)
func ReadFirst(question string) (string, error)
func ReadFirstByte(question string) (byte, error)
func ReadFirstRune(question string) (rune, error)
func ReadInput(question string) (string, error)
func ReadLine(question string) (string, error)
func ReadPassword(question ...string) string
func Redf(format string, a ...interface{})
func Redln(a ...interface{})
func Redp(a ...interface{})
func ShellExec(cmdLine string, shells ...string) (string, error)
func ShellQuote(s string) string
func String2OSArgs(line string) []string
func StringToOSArgs(line string) []string
func Successf(format string, a ...interface{})
func Successln(a ...interface{})
func Successp(a ...interface{})
func Warnf(format string, a ...interface{})
func Warnln(a ...interface{})
func Warnp(a ...interface{})
func Workdir() string
func Yellowf(format string, a ...interface{})
func Yellowln(a ...interface{})
func Yellowp(a ...interface{})
```

## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...
```

**Testing**:

```shell
go test -v ./cliutil/...
```

**Test limit by regexp**:

```shell
go test -v -run ^TestSetByKeys ./cliutil/...
```

## Projects using `cliutil`

`cliutil` is used in these projects:

- https://github.com/gookit/gcli
