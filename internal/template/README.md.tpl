# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/goutil?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/goutil)](https://github.com/gookit/goutil)
[![GoDoc](https://godoc.org/github.com/gookit/goutil?status.svg)](https://pkg.go.dev/github.com/gookit/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/goutil)](https://goreportcard.com/report/github.com/gookit/goutil)
[![Unit-Tests](https://github.com/gookit/goutil/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/goutil/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)

💪 Useful utils package for the Go: int, string, array/slice, map, error, time, format, CLI, ENV, filesystem, system, testing and more.

- [`arrutil`](./arrutil): Array/Slice util functions. eg: check, convert, formatting
- [`cflag`](./cflag):  Wraps and extends go `flag.FlagSet` to build simple command line applications
- [`cliutil`](./cliutil) Command-line util functions. eg: read input, exec command, cmdline parse/build
- [`dump`](./dump):  Simple variable printing tool, printing slice, map will automatically wrap each element and display the call location
- [`errorx`](./errorx) Provide an enhanced error implements for go, allow with stacktraces and wrap another error.
- [`envutil`](./envutil) ENV util for current runtime env information. eg: get one, get info, parse var
- [`fmtutil`](./fmtutil) Format data util functions. eg: data, size
- [`fsutil`](./fsutil) Filesystem util functions, quick create, read and write file. eg: file and dir check, operate
- [`jsonutil`](./jsonutil) some util functions for quick read, write, encode, decode JSON data.
- [`maputil`](./maputil) Map data util functions. eg: convert, sub-value get, simple merge
- [`mathutil`](./mathutil) Math(int, number) util functions. eg: convert, math calc, random
- `netutil` Network util functions
  - `netutil/httpreq` An easier-to-use HTTP client that wraps http.Client
- [`stdutil`](./stdutil) Provide some commonly std util functions.
- [`structs`](./structs) Provide some extends util functions for struct. eg: tag parse, struct data
- [`strutil`](./strutil) String util functions. eg: bytes, check, convert, encode, format and more
- [`sysutil`](./sysutil) System util functions. eg: sysenv, exec, user, process
- [`testutil`](./testutil) Test help util functions. eg: http test, mock ENV value
- [`timex`](./timex) Provides an enhanced time.Time implementation. Add more commonly used functional methods
  - such as: DayStart(), DayAfter(), DayAgo(), DateFormat() and more.

> **[中文说明](README.zh-CN.md)**

## GoDoc

- [Godoc for github](https://pkg.go.dev/github.com/gookit/goutil)

## Install

```shell
go get github.com/gookit/goutil
```

## Packages
{{pgkFuncs}}
## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...
go test ./...
```

Testing in docker:

```shell
cd goutil
docker run -ti -v $(pwd):/go/work golang:1.18
root@xx:/go/work# go test ./...
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

## License

[MIT](LICENSE)
