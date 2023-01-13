# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/goutil?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/goutil)](https://github.com/gookit/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/goutil)](https://goreportcard.com/report/github.com/gookit/goutil)
[![Unit-Tests](https://github.com/gookit/goutil/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/goutil/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)
[![Go Reference](https://pkg.go.dev/badge/github.com/gookit/goutil.svg)](https://pkg.go.dev/github.com/gookit/goutil)

`goutil` Go 常用功能的扩展工具库。包含：数字，byte, 字符串，slice/数组，Map，结构体，反射，文本，文件，错误，时间日期，测试，特殊处理，格式化，常用信息获取等等。

> **[EN README](README.md)**

**基础工具包**

- [`arrutil`](./arrutil) array/slice 相关操作的函数工具包 如：类型转换，元素检查等等
- [`cliutil`](./cliutil) CLI 的一些工具函数包. eg: read input, exec command
  - [cmdline](./cliutil/cmdline) 提供 cmdline 解析，args 构建到 cmdline
- [`envutil`](./envutil) ENV 信息获取判断工具包. eg: get one, get info, parse var
- [`fmtutil`](./fmtutil) 格式化数据工具函数 eg：数据size
- [`fsutil`](./fsutil) 文件系统操作相关的工具函数包. eg: file and dir check, operate
- [`jsonutil`](./jsonutil) 一些用于快速读取、写入、编码、解码 JSON 数据的实用函数。
- [`maputil`](./maputil) map 相关操作的函数工具包. eg: convert, sub-value get, simple merge
- [`mathutil`](./mathutil) int/number 相关操作的函数工具包. eg: convert, math calc, random
- [`reflects`](./reflects) 提供一些扩展性的反射使用工具函数.
- [`stdutil`](./stdutil) 提供一些常用的 std util 函数。
- [`structs`](./structs) 为 struct 提供一些扩展 util 函数。 eg: tag parse, struct data
- [`strutil`](./strutil) string 相关操作的函数工具包. eg: bytes, check, convert, encode, format and more
- [`sysutil`](./sysutil) system 相关操作的函数工具包. eg: sysenv, exec, user, process
  - [process](./sysutil/process) 提供一些进程操作相关的实用功能。

**扩展工具包**

- [`cflag`](./cflag): 包装和扩展 go `flag.FlagSet` 以方便快速的构建简单的命令行应用程序
- [`dump`](./dump)  GO变量打印工具，打印 slice, map 会自动换行显示每个元素，同时会显示打印调用位置
- [`errorx`](./errorx)  为 go 提供增强的错误实现，允许带有堆栈跟踪信息和包装另一个错误。
- strutil:
  - `netutil/httpreq` 包装 http.Client 实现的更加易于使用的HTTP客户端
- strutil:
  - [textscan](strutil/textscan) 实现了一个快速扫描和分析文本内容的解析器. 可用于解析 INI, Properties 等格式内容
- sysutil:
  - [clipboard](sysutil/clipboard) 提供简单的剪贴板读写操作工具库
  - [cmdr](sysutil/cmdr) 提供快速构建和运行一个cmd，批量运行多个cmd任务
- [`testutil`](testutil) test help 相关操作的函数工具包. eg: http test, mock ENV value
  - [assert](testutil/assert) 用于帮助测试的断言函数工具包
- [`timex`](timex) 提供增强的 time.Time 实现。添加更多常用的功能方法
  - 提供类似 `Y-m-d H:i:s` 的日期时间格式解析处理
  - 例如: DayStart(), DayAfter(), DayAgo(), DateFormat() 等等

## GoDoc

- [Godoc for github](https://pkg.go.dev/github.com/gookit/goutil)

## 获取

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

## Related

- https://github.com/duke-git/lancet
- https://github.com/samber/lo
- https://github.com/zyedidia/generic
- https://github.com/thoas/go-funk

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
