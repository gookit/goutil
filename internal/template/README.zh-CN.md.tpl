# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/goutil?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/goutil)](https://github.com/gookit/goutil)
[![GoDoc](https://godoc.org/github.com/gookit/goutil?status.svg)](https://pkg.go.dev/github.com/gookit/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/goutil)](https://goreportcard.com/report/github.com/gookit/goutil)
[![Unit-Tests](https://github.com/gookit/goutil/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/goutil/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)

Go一些常用的工具函数实现、增强、收集和整理

- [`arrutil`](./arrutil) array/slice 相关操作的函数工具包 如：类型转换，元素检查等等
- [`cflag`](./cflag): 包装和扩展 go `flag.FlagSet` 以构建简单的命令行应用程序
- [`cliutil`](./cliutil) CLI 的一些工具函数包. eg: read input, exec command, cmdline parse/build
- [`dump`](./dump)  简单的变量打印工具，打印 slice, map 会自动换行显示每个元素，同时会显示打印调用位置
- [`errorx`](./errorx)  为 go 提供增强的错误实现，允许带有堆栈跟踪信息和包装另一个错误。
- [`envutil`](./envutil) ENV 信息获取判断工具包. eg: get one, get info, parse var
- [`fmtutil`](./fmtutil) 格式化数据工具函数 eg：数据size
- [`fsutil`](./fsutil) 文件系统操作相关的工具函数包. eg: file and dir check, operate
- [`jsonutil`](./jsonutil) 一些用于快速读取、写入、编码、解码 JSON 数据的实用函数。
- [`maputil`](./maputil) map 相关操作的函数工具包. eg: convert, sub-value get, simple merge
- [`mathutil`](./mathutil) int/number 相关操作的函数工具包. eg: convert, math calc, random
- `netutil/httpreq` 包装 http.Client 实现的更加易于使用的HTTP客户端
- [`stdutil`](./stdutil) 提供一些常用的 std util 函数。
- [`structs`](./structs) 为 struct 提供一些扩展 util 函数。 eg: tag parse, struct data
- [`strutil`](./strutil) string 相关操作的函数工具包. eg: bytes, check, convert, encode, format and more
- [`sysutil`](./sysutil) system 相关操作的函数工具包. eg: sysenv, exec, user, process
- [`testutil`](./testutil) test help 相关操作的函数工具包. eg: http test, mock ENV value
- [`timex`](./timex) 提供增强的 time.Time 实现。添加更多常用的功能方法
  - 例如: DayStart(), DayAfter(), DayAgo(), DateFormat() 等等

> **[EN README](README.md)**

## GoDoc

- [Godoc for github](https://pkg.go.dev/github.com/gookit/goutil)

## 获取

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
