# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/goutil?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/goutil)](https://github.com/gookit/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/goutil)](https://goreportcard.com/report/github.com/gookit/goutil)
[![Unit-Tests](https://github.com/gookit/goutil/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/goutil/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/gookit/goutil.svg)](https://pkg.go.dev/github.com/gookit/goutil)

💪 Useful utils(**800+**) package for the Go: int, string, array/slice, map, error, time, format, CLI, ENV, filesystem, system, testing and more.

> **[中文说明](README.zh-CN.md)**

## Packages

### Basic packages

- [`arrutil`](arrutil): Array/Slice util functions. eg: check, convert, formatting, enum, collections
- [`byteutil`](byteutil): Provide some common bytes util functions. eg: convert, check and more
- [`maputil`](maputil) Map data util functions. eg: convert, sub-value get, simple merge
- [`mathutil`](mathutil) Math(int, number) util functions. eg: convert, math calc, random
- [`reflects`](reflects) Provide extends reflect util functions.
- [`structs`](structs) Provide some extends util functions for struct. eg: tag parse, struct data init
- [`strutil`](strutil) String util functions. eg: bytes, check, convert, encode, format and more
- [`sysutil`](sysutil) System util functions. eg: sysenv, exec, user, process
- [`cliutil`](cliutil) Command-line util functions. eg: colored print, read input, exec command
- [`envutil`](envutil) ENV util for current runtime env information. eg: get one, get info, parse var
- [`fsutil`](fsutil) Filesystem util functions, quick create, read and write file. eg: file and dir check, operate
- [`jsonutil`](jsonutil) Provide some util functions for quick read, write, encode, decode JSON data.

### Debug & Test & Errors

- [`dump`](dump): GO value printing tool. print slice, map will auto wrap each element and display the call location
- [`errorx`](errorx) Provide an enhanced error implements for go, allow with stacktrace and wrap another error.
- [`assert`](testutil/assert) Provides commonly asserts functions for help testing
- [`testutil`](testutil) Test help util functions. eg: http test, mock ENV value
- [`fakeobj`](x/fakeobj) provides a fake object for testing. such as fake fs.File, fs.FileInfo, fs.DirEntry etc.

### Extra Tools packages

- [`cflag`](cflag):  Wraps and extends go `flag.FlagSet` to build simple command line applications
- [`ccolor`](x/ccolor): Simple command-line color output library that uses ANSI color codes to output text with colors.
- [`timex`](timex) Provides an enhanced time.Time implementation. Add more commonly used functional methods
  - Provides datetime format parsing like `Y-m-d H:i:s`
  - such as: DayStart(), DayAfter(), DayAgo(), DateFormat() and more.
- [httpreq](netutil/httpreq) An easier-to-use HTTP client that wraps http.Client, and with some http utils.
- [syncs](syncs) Provides synchronization primitives util functions.

**More ...**

- [`cmdline`](cliutil/cmdline) Provide cmdline parse, args build to cmdline
- [`encodes`](encodes): Provide some encoding/decoding, hash, crypto util functions. eg: base64, hex, etc.
- [`finder`](x/finder) Provides a simple and convenient file/dir lookup function, supports filtering, excluding, matching, ignoring, etc.
- [`netutil`](netutil) Network util functions. eg: Ip, IpV4, IpV6, Mac, Port, Hostname, etc.
- [textutil](strutil/textutil) Provide some extensions text handle util functions. eg: text replace, etc.
- [textscan](strutil/textscan) Implemented a parser that quickly scans and analyzes text content. It can be used to parse INI, Properties and other formats
- [`cmdr`](sysutil/cmdr) Provide for quick build and run a cmd, batch run multi cmd tasks
- [`clipboard`](x/clipboard) Provide a simple clipboard read and write operations.
- [`process`](sysutil/process) Provide some process handle util functions.
- [`fmtutil`](x/fmtutil) Format data util functions. eg: data, size, time
- [`goinfo`](x/goinfo) provide some standard util functions for go.

## Go Doc

Please see [Go doc](https://pkg.go.dev/github.com/gookit/goutil).
Wiki docs on [DeepWiki - gookit/goutil](https://deepwiki.com/gookit/goutil)

## Install

```shell
go get github.com/gookit/goutil
```

## Usage

```go
// github.com/gookit/goutil
is.True(goutil.IsEmpty(nil))
is.False(goutil.IsEmpty("abc"))

is.True(goutil.IsEqual("a", "a"))
is.True(goutil.IsEqual([]string{"a"}, []string{"a"}))
is.True(goutil.IsEqual(23, 23))

is.True(goutil.Contains("abc", "a"))
is.True(goutil.Contains([]string{"abc", "def"}, "abc"))
is.True(goutil.Contains(map[int]string{2: "abc", 4: "def"}, 4))

// convert type
str = goutil.String(23) // "23"
iVal = goutil.Int("-2") // 2
i64Val = goutil.Int64("-2") // -2
u64Val = goutil.Uint("2") // 2
```

### Dump go variable

```go
dump.Print(somevar, somevar2, ...)
```

**dump nested struct**

![preview-nested-struct](dump/_examples/preview-nested-struct.png)

## Packages
{{pgkFuncs}}
## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...

# testing
go test -v ./...
go test -v -run ^TestErr$
go test -v -run ^TestErr$ ./testutil/assert/...
```

Testing in docker:

```shell
cd goutil

docker run -ti -v $(pwd):/go/goutil -e GOPROXY=https://goproxy.cn,direct golang:1.23
# on Windows
docker run -ti -v "${PWD}:/go/goutil" -e GOPROXY=https://goproxy.cn,direct golang:1.23

root@xx:/go/goutil# go test ./...
```

## Gookit packages

- [gookit/ini](https://github.com/gookit/ini) Go config management, use INI files
- [gookit/rux](https://github.com/gookit/rux) Simple and fast request router for golang HTTP
- [gookit/gcli](https://github.com/gookit/gcli) Build CLI application, tool library, running CLI commands
- [gookit/slog](https://github.com/gookit/slog) Lightweight, easy to extend, configurable logging library written in Go
- [gookit/color](https://github.com/gookit/color) A command-line color library with true color support, universal API methods and Windows support
- [gookit/event](https://github.com/gookit/event) Lightweight event manager and dispatcher implements by Go
- [gookit/cache](https://github.com/gookit/cache) Generic cache use and cache manager for golang. support File, Memory, Redis, Memcached.
- [gookit/config](https://github.com/gookit/config) Go config management. support JSON, YAML, TOML, INI, HCL, ENV and Flags
- [gookit/filter](https://github.com/gookit/filter) Provide filtering, sanitizing, and conversion of golang data
- [gookit/validate](https://github.com/gookit/validate) Use for data validation and filtering. support Map, Struct, Form data
- [gookit/goutil](https://github.com/gookit/goutil) Some utils for the Go: string, array/slice, map, format, cli, env, filesystem, test and more
- More, please see https://github.com/gookit

## Related

- https://github.com/duke-git/lancet
- https://github.com/samber/lo
- https://github.com/zyedidia/generic
- https://github.com/thoas/go-funk

## License

[MIT](LICENSE)
