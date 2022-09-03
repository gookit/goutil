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
- [`cliutil`](./cliutil) CLI 的一些工具函数包. eg: read input, exec command
  - [cmdline](./cliutil/cmdline) 提供 cmdline 解析，args 构建到 cmdline
- [`dump`](./dump)  简单的变量打印工具，打印 slice, map 会自动换行显示每个元素，同时会显示打印调用位置
- [`errorx`](./errorx)  为 go 提供增强的错误实现，允许带有堆栈跟踪信息和包装另一个错误。
- [`envutil`](./envutil) ENV 信息获取判断工具包. eg: get one, get info, parse var
- [`fmtutil`](./fmtutil) 格式化数据工具函数 eg：数据size
- [`fsutil`](./fsutil) 文件系统操作相关的工具函数包. eg: file and dir check, operate
- [`jsonutil`](./jsonutil) 一些用于快速读取、写入、编码、解码 JSON 数据的实用函数。
- [`maputil`](./maputil) map 相关操作的函数工具包. eg: convert, sub-value get, simple merge
- [`mathutil`](./mathutil) int/number 相关操作的函数工具包. eg: convert, math calc, random
  - `netutil/httpreq` 包装 http.Client 实现的更加易于使用的HTTP客户端
- [`reflects`](./reflects) 提供一些扩展性的反射使用工具函数.
- [`stdutil`](./stdutil) 提供一些常用的 std util 函数。
- [`structs`](./structs) 为 struct 提供一些扩展 util 函数。 eg: tag parse, struct data
- [`strutil`](./strutil) string 相关操作的函数工具包. eg: bytes, check, convert, encode, format and more
- [`sysutil`](./sysutil) system 相关操作的函数工具包. eg: sysenv, exec, user, process
  - [clipboard](./sysutil/clipboard) 提供简单的剪贴板读写操作工具库
  - [cmdr](./sysutil/cmdr) 提供快速构建和运行一个cmd，批量运行多个cmd任务
  - [process](./sysutil/process) 提供一些进程操作相关的实用功能。
- [`testutil`](./testutil) test help 相关操作的函数工具包. eg: http test, mock ENV value
  - [assert](./testutil/assert) 用于帮助测试的断言函数工具包
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

### Array/Slice

> Package `github.com/gookit/goutil/arrutil`

```go
// source at arrutil/arrutil.go
func Reverse(ss []string)
func StringsRemove(ss []string, s string) []string
func StringsFilter(ss []string, filter ...func(s string) bool) []string
func StringsMap(ss []string, mapFn func(s string) string) []string
func TrimStrings(ss []string, cutSet ...string) (ns []string)
func GetRandomOne(arr interface{}) interface{} { return RandomOne(arr) }
func RandomOne(arr interface{}) interface{}
// source at arrutil/check.go
func IntsHas(ints []int, val int) bool
func Int64sHas(ints []int64, val int64) bool
func InStrings(elem string, ss []string) bool { return StringsHas(ss, elem) }
func StringsHas(ss []string, val string) bool
func HasValue(arr, val interface{}) bool
func Contains(arr, val interface{}) bool
func NotContains(arr, val interface{}) bool
// source at arrutil/collection.go
func TwowaySearch(data interface{}, item interface{}, fn Comparer) (int, error)
func MakeEmptySlice(itemType reflect.Type) interface{}
func CloneSlice(data interface{}) interface{}
func Excepts(first interface{}, second interface{}, fn Comparer) interface{}
func Intersects(first interface{}, second interface{}, fn Comparer) interface{}
func Union(first interface{}, second interface{}, fn Comparer) interface{}
func Find(source interface{}, fn Predicate) (interface{}, error)
func FindOrDefault(source interface{}, fn Predicate, defaultValue interface{}) interface{}
func TakeWhile(data interface{}, fn Predicate) interface{}
func ExceptWhile(data interface{}, fn Predicate) interface{}
// source at arrutil/convert.go
func JoinStrings(sep string, ss ...string) string
func StringsJoin(sep string, ss ...string) string
func StringsToInts(ss []string) (ints []int, err error)
func MustToStrings(arr interface{}) []string
func StringsToSlice(ss []string) []interface{}
func ToInt64s(arr interface{}) (ret []int64, err error)
func MustToInt64s(arr interface{}) []int64
func SliceToInt64s(arr []interface{}) []int64
func ToStrings(arr interface{}) (ret []string, err error)
func SliceToStrings(arr []interface{}) []string
func AnyToString(arr interface{}) string
func SliceToString(arr ...interface{}) string { return ToString(arr) }
func ToString(arr []interface{}) string
func JoinSlice(sep string, arr ...interface{}) string
// source at arrutil/format.go
func NewFormatter(arr interface{}) *ArrFormatter
func FormatIndent(arr interface{}, indent string) string
```
#### ArrUtil Usage

**check value**:

```go
arrutil.IntsHas([]int{2, 4, 5}, 2) // True
arrutil.Int64sHas([]int64{2, 4, 5}, 2) // True
arrutil.StringsHas([]string{"a", "b"}, "a") // True

// list and val interface{}
arrutil.Contains(list, val)
arrutil.Contains([]uint32{9, 2, 3}, 9) // True
```

**convert**:

```go
ints, err := arrutil.ToInt64s([]string{"1", "2"}) // ints: []int64{1, 2} 
ss, err := arrutil.ToStrings([]int{1, 2}) // ss: []string{"1", "2"}
```


### Cflag

> Package `github.com/gookit/goutil/cflag`

```go
// source at cflag/cflag.go
func SetDebug(open bool)
func New(fns ...func(c *CFlags)) *CFlags
func NewEmpty(fns ...func(c *CFlags)) *CFlags
func WithDesc(desc string) func(c *CFlags)
func WithVersion(version string) func(c *CFlags)
// source at cflag/optarg.go
func NewArg(name, desc string, required bool) *FlagArg
// source at cflag/util.go
func IsZeroValue(opt *flag.Flag, value string) (bool, bool)
func AddPrefix(name string) string
func AddPrefixes(name string, shorts []string) string
func AddPrefixes2(name string, shorts []string, nameAtEnd bool) string
func SplitShortcut(shortcut string) []string
func FilterNames(names []string) []string
func IsFlagHelpErr(err error) bool
func ReplaceShorts(args []string, shortsMap map[string]string) []string
```
#### `cflag` Usage

`cflag` 使用说明请看 [cflag/README.zh-CN.md](cflag/README.zh-CN.md)


### CLI/Console

> Package `github.com/gookit/goutil/cliutil`

```go
// source at cliutil/cliutil.go
func LineBuild(binFile string, args []string) string
func BuildLine(binFile string, args []string) string
func String2OSArgs(line string) []string
func StringToOSArgs(line string) []string
func ParseLine(line string) []string
func QuickExec(cmdLine string, workDir ...string) (string, error)
func ExecLine(cmdLine string, workDir ...string) (string, error)
func ExecCmd(binName string, args []string, workDir ...string) (string, error)
func ExecCommand(binName string, args []string, workDir ...string) (string, error)
func ShellExec(cmdLine string, shells ...string) (string, error)
func CurrentShell(onlyName bool) (path string)
func HasShellEnv(shell string) bool
func Workdir() string
func BinDir() string
func BinFile() string
func BinName() string
func BuildOptionHelpName(names []string) string
func ShellQuote(s string) string
func OutputLines(output string) []string
func FirstLine(output string) string
// source at cliutil/color_print.go
func Redp(a ...interface{}) { color.Red.Print(a...) }
func Redf(format string, a ...interface{}) { color.Red.Printf(format, a...) }
func Redln(a ...interface{}) { color.Red.Println(a...) }
func Bluep(a ...interface{}) { color.Blue.Print(a...) }
func Bluef(format string, a ...interface{}) { color.Blue.Printf(format, a...) }
func Blueln(a ...interface{}) { color.Blue.Println(a...) }
func Cyanp(a ...interface{}) { color.Cyan.Print(a...) }
func Cyanf(format string, a ...interface{}) { color.Cyan.Printf(format, a...) }
func Cyanln(a ...interface{}) { color.Cyan.Println(a...) }
func Grayp(a ...interface{}) { color.Gray.Print(a...) }
func Grayf(format string, a ...interface{}) { color.Gray.Printf(format, a...) }
func Grayln(a ...interface{}) { color.Gray.Println(a...) }
func Greenp(a ...interface{}) { color.Green.Print(a...) }
func Greenf(format string, a ...interface{}) { color.Green.Printf(format, a...) }
func Greenln(a ...interface{}) { color.Green.Println(a...) }
func Yellowp(a ...interface{}) { color.Yellow.Print(a...) }
func Yellowf(format string, a ...interface{}) { color.Yellow.Printf(format, a...) }
func Yellowln(a ...interface{}) { color.Yellow.Println(a...) }
func Magentap(a ...interface{}) { color.Magenta.Print(a...) }
func Magentaf(format string, a ...interface{}) { color.Magenta.Printf(format, a...) }
func Magentaln(a ...interface{}) { color.Magenta.Println(a...) }
func Infop(a ...interface{}) { color.Info.Print(a...) }
func Infof(format string, a ...interface{}) { color.Info.Printf(format, a...) }
func Infoln(a ...interface{}) { color.Info.Println(a...) }
func Successp(a ...interface{}) { color.Success.Print(a...) }
func Successf(format string, a ...interface{}) { color.Success.Printf(format, a...) }
func Successln(a ...interface{}) { color.Success.Println(a...) }
func Errorp(a ...interface{}) { color.Error.Print(a...) }
func Errorf(format string, a ...interface{}) { color.Error.Printf(format, a...) }
func Errorln(a ...interface{}) { color.Error.Println(a...) }
func Warnp(a ...interface{}) { color.Warn.Print(a...) }
func Warnf(format string, a ...interface{}) { color.Warn.Printf(format, a...) }
func Warnln(a ...interface{}) { color.Warn.Println(a...) }
// source at cliutil/read.go
func ReadInput(question string) (string, error)
func ReadLine(question string) (string, error)
func ReadFirst(question string) (string, error)
func ReadFirstByte(question string) (byte, error)
func ReadFirstRune(question string) (rune, error)
// source at cliutil/read_nonwin.go
func ReadPassword(question ...string) string
```

#### CLI Util Usage

**helper functions:**

```go
cliutil.Workdir() // current workdir
cliutil.BinDir() // the program exe file dir

cliutil.ReadInput("Your name?")
cliutil.ReadPassword("Input password:")
ans, _ := cliutil.ReadFirstByte("continue?[y/n] ")
```

**cmdline parse:**

```go
package main

import (
	"fmt"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
)

func main() {
	args := cliutil.ParseLine(`./app top sub --msg "has multi words"`)
	dump.P(args)

	s := cliutil.BuildLine("./myapp", []string{
		"-a", "val0",
		"-m", "this is message",
		"arg0",
	})
	fmt.Println("Build line:", s)
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

Build line: ./myapp -a val0 -m "this is message" arg0
```

> More, please see [./cliutil/README](cliutil/README.md)

### Dumper

> Package `github.com/gookit/goutil/dump`

```go
// source at dump/dump.go
func Std() *Dumper
func Reset()
func Config(fn func(opts *Options))
func Print(vs ...interface{})
func Println(vs ...interface{})
func Fprint(w io.Writer, vs ...interface{})
func Format(vs ...interface{}) string
func NoLoc(vs ...interface{})
func Clear(vs ...interface{})
// source at dump/dumper.go
func NewDumper(out io.Writer, skip int) *Dumper
func NewWithOptions(fn func(opts *Options)) *Dumper
func NewDefaultOptions(out io.Writer, skip int) *Options
```
#### Examples

example code:

```go
package main

import "github.com/gookit/goutil/dump"

// rum demo:
// 	go run ./dump/_examples/demo1.go
func main() {
	otherFunc1()
}

func otherFunc1() {
	dump.P(
		23,
		[]string{"ab", "cd"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, // len > 10
		map[string]interface{}{
			"key": "val", "sub": map[string]string{"k": "v"},
		},
		struct {
			ab string
			Cd int
		}{
			"ab", 23,
		},
	)
}
```

Preview:

![](dump/_examples/preview-demo1.png)

**nested struct**

> source code at `dump/dumper_test.TestStruct_WithNested`

![](dump/_examples/preview-nested-struct.png)


### ENV/Environment

> Package `github.com/gookit/goutil/envutil`

```go
// source at envutil/envutil.go
func VarReplace(s string) string { return os.ExpandEnv(s) }
func VarParse(val string) string
func ParseEnvValue(val string) string
func ParseValue(val string) (newVal string)
// source at envutil/get.go
func Getenv(name string, def ...string) string
func GetInt(name string, def ...int) int
func GetBool(name string, def ...bool) bool
func Environ() map[string]string
// source at envutil/info.go
func IsWin() bool
func IsWindows() bool
func IsMac() bool
func IsLinux() bool
func IsMSys() bool
func IsWSL() bool
func IsTerminal(fd uintptr) bool
func StdIsTerminal() bool
func IsConsole(out io.Writer) bool
func HasShellEnv(shell string) bool
func IsSupportColor() bool
func IsSupport256Color() bool
func IsSupportTrueColor() bool
```
#### ENV Util Usage

**helper functions:**

```go
envutil.IsWin()
envutil.IsMac()
envutil.IsLinux()

// get ENV value by key, can with default value
envutil.Getenv("APP_ENV", "dev")
envutil.GetInt("LOG_LEVEL", 1)
envutil.GetBool("APP_DEBUG", true)

// parse ENV var value from input string, support default value.
envutil.ParseValue("${ENV_NAME | defValue}")
```


### Errorx

> Package `github.com/gookit/goutil/errorx`

`errorx` 提供了增强的错误报告实现，包含调用堆栈信息并且可以包装上一级错误。

> 在打印 error 时会额外附带调用栈信息, 方便记录日志和查找问题。


```go
// source at errorx/errorx.go
func New(msg string) error
func Newf(tpl string, vars ...interface{}) error
func Errorf(tpl string, vars ...interface{}) error
func With(err error, msg string) error
func Withf(err error, tpl string, vars ...interface{}) error
func WithPrev(err error, msg string) error
func WithPrevf(err error, tpl string, vars ...interface{}) error
func WithStack(err error) error
func Traced(err error) error
func Stacked(err error) error
func WithOptions(msg string, fns ...func(opt *ErrStackOpt)) error
func Wrap(err error, msg string) error
func Wrapf(err error, tpl string, vars ...interface{}) error
// source at errorx/reply.go
func NewR(code int, msg string) ErrorR
func Fail(code int, msg string) ErrorR
func Suc(msg string) ErrorR
// source at errorx/stack.go
func FuncForPC(pc uintptr) *Func
func ResetStdOpt()
func Config(fns ...func(opt *ErrStackOpt))
func SkipDepth(skipDepth int) func(opt *ErrStackOpt)
func TraceDepth(traceDepth int) func(opt *ErrStackOpt)
// source at errorx/util.go
func Raw(msg string) error
func Rawf(tpl string, vars ...interface{}) error
func Cause(err error) error
func Unwrap(err error) error
func Previous(err error) error { return Unwrap(err) }
func ToErrorX(err error) (ex *ErrorX, ok bool)
func Has(err, target error) bool
func Is(err, target error) bool
func To(err error, target interface{}) bool
func As(err error, target interface{}) bool
```

#### Errorx 使用示例

**创建错误带有调用栈信息**

- 使用 `errorx.New` 替代 `errors.New`

```go
func doSomething() error {
    if false {
	    // return errors.New("a error happen")
	    return errorx.New("a error happen")
	}
}
```

- 使用 `errorx.Newf` 或者 `errorx.Errorf` 替代 `fmt.Errorf`

```go
func doSomething() error {
    if false {
	    // return fmt.Errorf("a error %s", "happen")
	    return errorx.Newf("a error %s", "happen")
	}
}
```

**包装上一级错误**

之前这样使用:

```go
    if err := SomeFunc(); err != nil {
	    return err
	}
```

可以替换成:

```go
    if err := SomeFunc(); err != nil {
	    return errors.Stacked(err)
	}
```

**使用效果示例**

更多关于 `errorx` 的使用请看 [./errorx/README](errorx/README.md)

```go
    err := errorx.New("the error message")

    fmt.Println(err)
    // fmt.Printf("%v\n", err)
    // fmt.Printf("%#v\n", err)
```

> from the test: `errorx/errorx_test.TestNew()`

**Output**:

```text
the error message
STACK:
github.com/gookit/goutil/errorx_test.returnXErr()
  /Users/inhere/Workspace/godev/gookit/goutil/errorx/errorx_test.go:21
github.com/gookit/goutil/errorx_test.returnXErrL2()
  /Users/inhere/Workspace/godev/gookit/goutil/errorx/errorx_test.go:25
github.com/gookit/goutil/errorx_test.TestNew()
  /Users/inhere/Workspace/godev/gookit/goutil/errorx/errorx_test.go:29
testing.tRunner()
  /usr/local/Cellar/go/1.18/libexec/src/testing/testing.go:1439
runtime.goexit()
  /usr/local/Cellar/go/1.18/libexec/src/runtime/asm_amd64.s:1571
```


### Formatting

> Package `github.com/gookit/goutil/fmtutil`

```go
// source at fmtutil/format.go
func DataSize(size uint64) string
func SizeToString(size uint64) string { return DataSize(size) }
func StringToByte(sizeStr string) uint64 { return ParseByte(sizeStr) }
func ParseByte(sizeStr string) uint64
func PrettyJSON(v interface{}) (string, error)
func StringsToInts(ss []string) (ints []int, err error)
func ArgsWithSpaces(args []interface{}) (message string)
// source at fmtutil/time.go
func HowLongAgo(sec int64) string
```

### File System

> Package `github.com/gookit/goutil/fsutil`

```go
// source at fsutil/check.go
func PathExists(path string) bool
func IsDir(path string) bool
func FileExists(path string) bool
func IsFile(path string) bool
func IsAbsPath(aPath string) bool
func IsImageFile(path string) bool
func IsZipFile(filepath string) bool
// source at fsutil/fsutil.go
func OSTempFile(pattern string) (*os.File, error)
func TempFile(dir, pattern string) (*os.File, error)
func OSTempDir(pattern string) (string, error)
func TempDir(dir, pattern string) (string, error)
func MimeType(path string) (mime string)
func ReaderMimeType(r io.Reader) (mime string)
// source at fsutil/info.go
func Dir(fpath string) string
func PathName(fpath string) string
func Name(fpath string) string
func FileExt(fpath string) string
func Suffix(fpath string) string
func Expand(pathStr string) string
func ExpandPath(pathStr string) string
func Realpath(pathStr string) string
func SplitPath(pathStr string) (dir, name string)
func GlobWithFunc(pattern string, fn func(filePath string) error) (err error)
func FindInDir(dir string, handleFn HandleFunc, filters ...FilterFunc) (e error)
// source at fsutil/operate.go
func Mkdir(dirPath string, perm os.FileMode) error
func MkParentDir(fpath string) error
func DiscardReader(src io.Reader)
func MustReadFile(filePath string) []byte
func MustReadReader(r io.Reader) []byte
func GetContents(in interface{}) []byte
func ReadExistFile(filePath string) []byte
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error)
func QuickOpenFile(filepath string, fileFlag ...int) (*os.File, error)
func OpenReadFile(filepath string) (*os.File, error)
func CreateFile(fpath string, filePerm, dirPerm os.FileMode, fileFlag ...int) (*os.File, error)
func MustCreateFile(filePath string, filePerm, dirPerm os.FileMode) *os.File
func PutContents(filePath string, data interface{}, fileFlag ...int) (int, error)
func WriteFile(filePath string, data interface{}, perm os.FileMode, fileFlag ...int) error
func CopyFile(srcPath, dstPath string) error
func MustCopyFile(srcPath, dstPath string)
func Remove(fPath string) error
func MustRemove(fPath string)
func QuietRemove(fPath string) { _ = os.Remove(fPath) }
func RmIfExist(fPath string) error { return DeleteIfExist(fPath) }
func DeleteIfExist(fPath string) error
func RmFileIfExist(fPath string) error { return DeleteIfFileExist(fPath) }
func DeleteIfFileExist(fPath string) error
func Unzip(archive, targetDir string) (err error)
```

#### FsUtil Usage

**files finder:**

```go
package main

import (
	"fmt"
	"os"

	"github.com/gookit/goutil/fsutil/finder"
)

func main() {
	f := finder.EmptyFinder()

	f.
		AddDir("./testdata").
		AddFile("finder.go").
		NoDotFile().
		// NoDotDir().
		Find().
		Each(func(filePath string) {
			fmt.Println(filePath)
		})

	finder.NewFinder([]string{"./testdata"}).
		AddFile("finder.go").
		NoDotDir().
		EachStat(func(fi os.FileInfo, filePath string) {
			fmt.Println(filePath, "=>", fi.ModTime())
		})
}
```


### JSON Utils

> Package `github.com/gookit/goutil/jsonutil`

```go
// source at jsonutil/jsonutil.go
func WriteFile(filePath string, data interface{}) error
func ReadFile(filePath string, v interface{}) error
func Pretty(v interface{}) (string, error)
func Encode(v interface{}) ([]byte, error)
func EncodePretty(v interface{}) ([]byte, error)
func EncodeToWriter(v interface{}, w io.Writer) error
func EncodeUnescapeHTML(v interface{}) ([]byte, error)
func Decode(bts []byte, ptr interface{}) error
func DecodeString(str string, ptr interface{}) error
func DecodeReader(r io.Reader, ptr interface{}) error
func Mapping(src, dst interface{}) error
func StripComments(src string) string
```

### Map

> Package `github.com/gookit/goutil/maputil`

```go
// source at maputil/check.go
func HasKey(mp, key any) (ok bool)
func HasAllKeys(mp any, keys ...any) (ok bool, noKey interface{})
// source at maputil/convert.go
func KeyToLower(src map[string]string) map[string]string
func ToStringMap(src map[string]any) map[string]string
func HttpQueryString(data map[string]any) string
func ToString(mp map[string]any) string
func ToString2(mp any) string
func FormatIndent(mp any, indent string) string
// source at maputil/format.go
func NewFormatter(mp any) *MapFormatter
// source at maputil/get.go
func DeepGet(mp map[string]any, path string) (val interface{})
func QuietGet(mp map[string]any, path string) (val interface{})
func GetByPath(path string, mp map[string]any) (val interface{}, ok bool)
func Keys(mp any) (keys []string)
func Values(mp any) (values []interface{})
// source at maputil/maputil.go
func MergeSMap(src, dst map[string]string, ignoreCase bool) map[string]string
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string
func MakeByPath(path string, val interface{}) (mp map[string]interface{})
func MakeByKeys(keys []string, val any) (mp map[string]interface{})
// source at maputil/setval.go
func SetByPath(mp *map[string]any, path string, val any) error
func SetByKeys(mp *map[string]any, keys []string, val any) (err error)
```

### Math/Number

> Package `github.com/gookit/goutil/mathutil`

```go
// source at mathutil/convert.go
func Int(in interface{}) (int, error)
func QuietInt(in interface{}) int
func MustInt(in interface{}) int
func IntOrPanic(in interface{}) int
func IntOrErr(in interface{}) (iVal int, err error)
func ToInt(in interface{}) (iVal int, err error)
func StrInt(s string) int
func Uint(in interface{}) (uint64, error)
func QuietUint(in interface{}) uint64
func MustUint(in interface{}) uint64
func UintOrErr(in interface{}) (uint64, error)
func ToUint(in interface{}) (u64 uint64, err error)
func Int64(in interface{}) (int64, error)
func QuietInt64(in interface{}) int64
func MustInt64(in interface{}) int64
func Int64OrErr(in interface{}) (int64, error)
func ToInt64(in interface{}) (i64 int64, err error)
func QuietFloat(in interface{}) float64
func FloatOrPanic(in interface{}) float64
func MustFloat(in interface{}) float64
func Float(in interface{}) (float64, error)
func FloatOrErr(in interface{}) (float64, error)
func ToFloat(in interface{}) (f64 float64, err error)
func TryToString(val interface{}, defaultAsErr bool) (str string, err error)
func StringOrPanic(val interface{}) string
func MustString(val interface{}) string
func ToString(val interface{}) (string, error)
func StringOrErr(val interface{}) (string, error)
func QuietString(val interface{}) string
func String(val interface{}) string
// source at mathutil/mathutil.go
func MaxFloat(x, y float64) float64
func MaxInt(x, y int) int
func SwapMaxInt(x, y int) (int, int)
func MaxI64(x, y int64) int64
func SwapMaxI64(x, y int64) (int64, int64)
func Compare(srcVal, dstVal interface{}, op string) (ok bool)
func CompInt64(srcI64, dstI64 int64, op string) (ok bool)
func CompFloat(srcF64, dstF64 float64, op string) (ok bool)
// source at mathutil/number.go
func IsNumeric(c byte) bool
func Percent(val, total int) float64
func ElapsedTime(startTime time.Time) string
func DataSize(size uint64) string
func HowLongAgo(sec int64) string
// source at mathutil/random.go
func RandomInt(min, max int) int
func RandInt(min, max int) int { return RandomInt(min, max) }
func RandIntWithSeed(min, max int, seed int64) int
func RandomIntWithSeed(min, max int, seed int64) int
```

### Reflects

> Package `github.com/gookit/goutil/reflects`

```go
// source at reflects/check.go
func HasChild(v reflect.Value) bool
func IsNil(v reflect.Value) bool
func IsFunc(val interface{}) bool
func IsEqual(src, dst interface{}) bool
func IsEmpty(v reflect.Value) bool
func IsEmptyValue(v reflect.Value) bool
// source at reflects/conv.go
func BaseTypeVal(v reflect.Value) (value interface{}, err error)
func ValueByKind(val interface{}, kind reflect.Kind) (rv reflect.Value, err error)
// source at reflects/type.go
func ToBaseKind(kind reflect.Kind) BKind
func ToBKind(kind reflect.Kind) BKind
func TypeOf(v interface{}) Type
// source at reflects/util.go
func Elem(v reflect.Value) reflect.Value
func Len(v reflect.Value) int
func SliceSubKind(typ reflect.Type) reflect.Kind
// source at reflects/value.go
func Wrap(rv reflect.Value) Value
func ValueOf(v interface{}) Value
```

### Stdio

> Package `github.com/gookit/goutil/stdio`

```go
// source at stdio/ioutil.go
func QuietFprint(w io.Writer, ss ...string)
func QuietFprintf(w io.Writer, tpl string, vs ...interface{})
func QuietFprintln(w io.Writer, ss ...string)
func QuietWriteString(w io.Writer, ss ...string)
func DiscardReader(src io.Reader)
func MustReadReader(r io.Reader) []byte
// source at stdio/writer.go
func NewWriteWrapper(w io.Writer) *WriteWrapper
```

### Standard

> Package `github.com/gookit/goutil/stdutil`

```go
// source at stdutil/chan.go
func WaitCloseSignals(closer io.Closer) error
func Go(f func() error) error
// source at stdutil/check.go
func IsNil(v interface{}) bool
func IsEmpty(v interface{}) bool
func IsFunc(val interface{}) bool
func IsEqual(src, dst interface{}) bool
func Contains(data, elem interface{}) bool
func IsContains(data, elem interface{}) bool
func CheckContains(data, elem interface{}) (valid, found bool)
func ValueIsEmpty(v reflect.Value) bool
func ValueLen(v reflect.Value) int
// source at stdutil/convert.go
func ToString(v interface{}) string
func MustString(v interface{}) string
func TryString(v interface{}) (string, error)
func BaseTypeVal(val interface{}) (value interface{}, err error)
func BaseTypeVal2(v reflect.Value) (value interface{}, err error)
// source at stdutil/gofunc.go
func FuncName(fn interface{}) string
func CutFuncName(fullFcName string) (pkgPath, shortFnName string)
func PkgName(fullFcName string) string
// source at stdutil/stack.go
func GetCallStacks(all bool) []byte
func GetCallerInfo(skip int) string
func SimpleCallersInfo(skip, num int) []string
func GetCallersInfo(skip, max int) []string
// source at stdutil/stdutil.go
func DiscardE(_ error) {}
func PanicIfErr(err error)
func PanicIf(err error)
func Panicf(format string, v ...interface{})
func GoVersion() string
```

### Structs

> Package `github.com/gookit/goutil/structs`

```go
// source at structs/alias.go
func NewAliases(checker func(alias string)) *Aliases
// source at structs/convert.go
func ToMap(st interface{}, optFns ...MapOptFunc) map[string]interface{}
func MustToMap(st interface{}, optFns ...MapOptFunc) map[string]interface{}
func TryToMap(st interface{}, optFns ...MapOptFunc) (map[string]interface{}, error)
func StructToMap(st interface{}, optFns ...MapOptFunc) (map[string]interface{}, error)
// source at structs/data.go
func NewMapData() *DataStore
// source at structs/structs.go
func MapStruct(srcSt, dstSt interface{})
func InitDefaults(ptr interface{}, optFns ...InitOptFunc) error
// source at structs/tags.go
func ParseTags(st interface{}, tagNames []string) (map[string]maputil.SMap, error)
func ParseReflectTags(rt reflect.Type, tagNames []string) (map[string]maputil.SMap, error)
func NewTagParser(tagNames ...string) *TagParser
func ParseTagValueDefault(field, tagVal string) (mp maputil.SMap, err error)
func ParseTagValueDefine(sep string, defines []string) TagValFunc
func ParseTagValueNamed(field, tagVal string) (mp maputil.SMap, err error)
// source at structs/value.go
func NewValue(val interface{}) *Value
```

### Strings

> Package `github.com/gookit/goutil/strutil`

```go
// source at strutil/bytes.go
func NewBuffer() *Buffer
func NewByteChanPool(maxSize int, width int, capWidth int) *ByteChanPool
// source at strutil/check.go
func NoCaseEq(s, t string) bool
func IsNumChar(c byte) bool
func IsNumeric(s string) bool
func IsAlphabet(char uint8) bool
func IsAlphaNum(c uint8) bool
func StrPos(s, sub string) int
func BytePos(s string, bt byte) int
func HasOneSub(s string, subs []string) bool
func HasAllSubs(s string, subs []string) bool
func IsStartsOf(s string, prefixes []string) bool
func HasOnePrefix(s string, prefixes []string) bool
func HasPrefix(s string, prefix string) bool { return strings.HasPrefix(s, prefix) }
func IsStartOf(s, prefix string) bool { return strings.HasPrefix(s, prefix) }
func HasSuffix(s string, suffix string) bool { return strings.HasSuffix(s, suffix) }
func IsEndOf(s, suffix string) bool { return strings.HasSuffix(s, suffix) }
func IsValidUtf8(s string) bool { return utf8.ValidString(s) }
func IsSpace(c byte) bool
func IsEmpty(s string) bool { return len(s) == 0 }
func IsBlank(s string) bool
func IsNotBlank(s string) bool
func IsBlankBytes(bs []byte) bool
func IsSymbol(r rune) bool
func IsVersion(s string) bool
func Compare(s1, s2, op string) bool
func VersionCompare(v1, v2, op string) bool
// source at strutil/convert.go
func Quote(s string) string { return strconv.Quote(s) }
func Unquote(s string) string
func Join(sep string, ss ...string) string { return strings.Join(ss, sep) }
func JoinSubs(sep string, ss []string) string { return strings.Join(ss, sep) }
func Implode(sep string, ss ...string) string { return strings.Join(ss, sep) }
func String(val interface{}) (string, error)
func QuietString(in interface{}) string
func MustString(in interface{}) string
func StringOrErr(val interface{}) (string, error)
func ToString(val interface{}) (string, error)
func AnyToString(val interface{}, defaultAsErr bool) (str string, err error)
func Byte2str(b []byte) string
func Byte2string(b []byte) string
func ToBytes(s string) (b []byte)
func ToBool(s string) (bool, error)
func QuietBool(s string) bool
func MustBool(s string) bool
func Bool(s string) (bool, error)
func Int(s string) (int, error)
func ToInt(s string) (int, error)
func QuietInt(s string) int
func MustInt(s string) int
func IntOrPanic(s string) int
func Int64(s string) int64
func QuietInt64(s string) int64
func ToInt64(s string) (int64, error)
func Int64OrErr(s string) (int64, error)
func MustInt64(s string) int64
func Int64OrPanic(s string) int64
func Ints(s string, sep ...string) []int
func ToInts(s string, sep ...string) ([]int, error) { return ToIntSlice(s, sep...) }
func ToIntSlice(s string, sep ...string) (ints []int, err error)
func ToArray(s string, sep ...string) []string { return ToSlice(s, sep...) }
func Strings(s string, sep ...string) []string { return ToSlice(s, sep...) }
func ToStrings(s string, sep ...string) []string { return ToSlice(s, sep...) }
func ToSlice(s string, sep ...string) []string
func ToOSArgs(s string) []string
func MustToTime(s string, layouts ...string) time.Time
func ToTime(s string, layouts ...string) (t time.Time, err error)
func ToDuration(s string) (time.Duration, error)
// source at strutil/encode.go
func EscapeJS(s string) string
func EscapeHTML(s string) string
func AddSlashes(s string) string
func StripSlashes(s string) string
func Md5(src interface{}) string { return GenMd5(src) }
func MD5(src interface{}) string { return GenMd5(src) }
func GenMd5(src interface{}) string
func URLEncode(s string) string
func URLDecode(s string) string
func B32Encode(str string) string
func B32Decode(str string) string
func Base64(str string) string
func B64Encode(str string) string
func B64Decode(str string) string
func NewBaseEncoder(base int) *BaseEncoder
// source at strutil/filter.go
func Trim(s string, cutSet ...string) string
func Ltrim(s string, cutSet ...string) string { return TrimLeft(s, cutSet...) }
func LTrim(s string, cutSet ...string) string { return TrimLeft(s, cutSet...) }
func TrimLeft(s string, cutSet ...string) string
func Rtrim(s string, cutSet ...string) string { return TrimRight(s, cutSet...) }
func RTrim(s string, cutSet ...string) string { return TrimRight(s, cutSet...) }
func TrimRight(s string, cutSet ...string) string
func FilterEmail(s string) string
// source at strutil/format.go
func Title(s string) string { return strings.ToTitle(s) }
func Lower(s string) string { return strings.ToLower(s) }
func Lowercase(s string) string { return strings.ToLower(s) }
func Upper(s string) string { return strings.ToUpper(s) }
func Uppercase(s string) string { return strings.ToUpper(s) }
func UpperWord(s string) string
func LowerFirst(s string) string
func UpperFirst(s string) string
func Snake(s string, sep ...string) string
func SnakeCase(s string, sep ...string) string
func Camel(s string, sep ...string) string
func CamelCase(s string, sep ...string) string
func Indent(s, prefix string) string
func IndentBytes(b, prefix []byte) []byte
// source at strutil/id.go
func MicroTimeID() string
func MicroTimeHexID() string
// source at strutil/random.go
func RandomChars(ln int) string
func RandomCharsV2(ln int) string
func RandomCharsV3(ln int) string
func RandomBytes(length int) ([]byte, error)
func RandomString(length int) (string, error)
// source at strutil/runes.go
func RuneIsWord(c rune) bool
func RuneIsLower(c rune) bool
func RuneIsUpper(c rune) bool
func RunePos(s string, ru rune) int { return strings.IndexRune(s, ru) }
func IsSpaceRune(r rune) bool
func Utf8Len(s string) int { return utf8.RuneCountInString(s) }
func Utf8len(s string) int { return utf8.RuneCountInString(s) }
func RuneCount(s string) int { return len([]rune(s)) }
func RuneWidth(r rune) int
func TextWidth(s string) int { return Utf8Width(s) }
func Utf8Width(s string) (size int)
func TextTruncate(s string, w int, tail string) string { return Utf8Truncate(s, w, tail) }
func Utf8Truncate(s string, w int, tail string) string
func TextSplit(s string, w int) []string { return Utf8Split(s, w) }
func Utf8Split(s string, w int) []string
func TextWrap(s string, w int) string { return WidthWrap(s, w) }
func WidthWrap(s string, w int) string
// source at strutil/similar_find.go
func NewComparator(src, dst string) *SimilarComparator
func Similarity(s, t string, rate float32) (float32, bool)
// source at strutil/split.go
func Cut(s, sep string) (before string, after string, found bool)
func MustCut(s, sep string) (before string, after string)
func TrimCut(s, sep string) (string, string)
func SplitValid(s, sep string) (ss []string) { return Split(s, sep) }
func Split(s, sep string) (ss []string)
func SplitNValid(s, sep string, n int) (ss []string) { return SplitN(s, sep, n) }
func SplitN(s, sep string, n int) (ss []string)
func SplitTrimmed(s, sep string) (ss []string)
func SplitNTrimmed(s, sep string, n int) (ss []string)
func Substr(s string, pos, length int) string
// source at strutil/strutil.go
func Padding(s, pad string, length int, pos uint8) string
func PadLeft(s, pad string, length int) string
func PadRight(s, pad string, length int) string
func Repeat(s string, times int) string
func RepeatRune(char rune, times int) (chars []rune)
func RepeatBytes(char byte, times int) (chars []byte)
func Replaces(str string, pairs map[string]string) string
func PrettyJSON(v interface{}) (string, error)
func RenderTemplate(input string, data interface{}, fns template.FuncMap, isFile ...bool) string
func RenderText(input string, data interface{}, fns template.FuncMap, isFile ...bool) string
func WrapTag(s, tag string) string
```

### System

> Package `github.com/gookit/goutil/sysutil`

```go
// source at sysutil/exec.go
func NewCmd(bin string, args ...string) *cmdr.Cmd
func FlushExec(bin string, args ...string) error
func QuickExec(cmdLine string, workDir ...string) (string, error)
func ExecLine(cmdLine string, workDir ...string) (string, error)
func ExecCmd(binName string, args []string, workDir ...string) (string, error)
func ShellExec(cmdLine string, shells ...string) (string, error)
// source at sysutil/stack.go
func CallersInfos(skip, num int, filters ...func(file string, fc *runtime.Func) bool) []*CallerInfo
// source at sysutil/sysenv.go
func IsMSys() bool
func IsConsole(out io.Writer) bool
func IsTerminal(fd uintptr) bool
func StdIsTerminal() bool
func Hostname() string
func CurrentShell(onlyName bool) (path string)
func HasShellEnv(shell string) bool
func IsShellSpecialVar(c uint8) bool
func EnvPaths() []string
func FindExecutable(binName string) (string, error)
func Executable(binName string) (string, error)
func HasExecutable(binName string) bool
func SearchPath(keywords string) []string
// source at sysutil/sysutil.go
func Workdir() string
func BinDir() string
func BinFile() string
func GoVersion() string
func ParseGoVersion(line string) (*GoInfo, error)
func OsGoInfo() (*GoInfo, error)
// source at sysutil/sysutil_nonwin.go
func IsWin() bool
func IsWindows() bool
func IsMac() bool
func IsDarwin() bool
func IsLinux() bool
func Kill(pid int, signal syscall.Signal) error
func ProcessExists(pid int) bool
func OpenBrowser(URL string) error
// source at sysutil/user.go
func MustFindUser(uname string) *user.User
func LoginUser() *user.User
func CurrentUser() *user.User
func UHomeDir() string
func UserHomeDir() string
func HomeDir() string
func UserDir(subPath string) string
func UserCacheDir(subPath string) string
func UserConfigDir(subPath string) string
func ExpandPath(path string) string
// source at sysutil/user_nonwin.go
func ChangeUserByName(newUname string) (err error)
func ChangeUserUidGid(newUid int, newGid int) (err error)
```

### Testing

> Package `github.com/gookit/goutil/testutil`

```go
// source at testutil/buffer.go
func NewBuffer() *Buffer
// source at testutil/envmock.go
func MockEnvValue(key, val string, fn func(nv string))
func MockEnvValues(kvMap map[string]string, fn func())
func MockOsEnvByText(envText string, fn func())
func MockOsEnv(mp map[string]string, fn func())
func MockCleanOsEnv(mp map[string]string, fn func())
// source at testutil/httpmock.go
func NewHttpRequest(method, path string, data *MD) *http.Request
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder
// source at testutil/testutil.go
func DiscardStdout() error
func ReadOutput() (s string)
func RewriteStdout()
func RestoreStdout(printData ...bool) (s string)
func RewriteStderr()
func RestoreStderr(printData ...bool) (s string)
// source at testutil/writer.go
func NewTestWriter() *TestWriter
```

### Timex

> Package `github.com/gookit/goutil/timex`

Provides an enhanced time.Time implementation, and add more commonly used functional methods.
```go
// source at timex/template.go
func ToLayout(template string) string
// source at timex/timex.go
func Now() *TimeX
func New(t time.Time) *TimeX
func Wrap(t time.Time) *TimeX
func FromTime(t time.Time) *TimeX
func Local() *TimeX
func FromUnix(sec int64) *TimeX
func FromDate(s string, template ...string) (*TimeX, error)
func FromString(s string, layouts ...string) (*TimeX, error)
func LocalByName(tzName string) *TimeX
// source at timex/util.go
func NowUnix() int64
func SetLocalByName(tzName string) error
func Format(t time.Time) string
func FormatBy(t time.Time, layout string) string
func Date(t time.Time, template string) string
func DateFormat(t time.Time, template string) string
func FormatByTpl(t time.Time, template string) string
func FormatUnix(sec int64) string
func FormatUnixBy(sec int64, layout string) string
func FormatUnixByTpl(sec int64, template string) string
func NowAddDay(day int) time.Time
func NowAddHour(hour int) time.Time
func NowAddMinutes(minutes int) time.Time
func NowAddSeconds(seconds int) time.Time
func NowHourStart() time.Time
func NowHourEnd() time.Time
func AddDay(t time.Time, day int) time.Time
func AddHour(t time.Time, hour int) time.Time
func AddMinutes(t time.Time, minutes int) time.Time
func AddSeconds(t time.Time, seconds int) time.Time
func HourStart(t time.Time) time.Time
func HourEnd(t time.Time) time.Time
func DayStart(t time.Time) time.Time
func DayEnd(t time.Time) time.Time
func TodayStart() time.Time
func TodayEnd() time.Time
func HowLongAgo(sec int64) string
func ToDuration(s string) (time.Duration, error)
```
#### Timex Usage

**Create timex instance**

```go
now := timex.Now()

// from time.Time
tx := timex.New(time.Now())
tx := timex.FromTime(time.Now())

// from time unix
tx := timex.FromUnix(1647411580)
```

Create from datetime string:

```go
// auto match layout by datetime
tx, err  := timex.FromString("2022-04-20 19:40:34")
// custom set the datetime layout
tx, err  := timex.FromString("2022-04-20 19:40:34", "2006-01-02 15:04:05")
// use date template as layout
tx, err  := timex.FromDate("2022-04-20 19:40:34", "Y-m-d H:I:S")
```

**Use timex instance**

```go
tx := timex.Now()
```

change time:

```go
tx.Yesterday()
tx.Tomorrow()

tx.DayStart() // get time at Y-m-d 00:00:00
tx.DayEnd() // get time at Y-m-d 23:59:59
tx.HourStart() // get time at Y-m-d H:00:00
tx.HourEnd() // get time at Y-m-d H:59:59

tx.AddDay(2)
tx.AddHour(1)
tx.AddMinutes(15)
tx.AddSeconds(120)
```

compare time:

```go
// before compare
tx.IsBefore(u time.Time)
tx.IsBeforeUnix(1647411580)
// after compare
tx.IsAfter(u time.Time)
tx.IsAfterUnix(1647411580)
```

**Helper functions**

```go
ts := timex.NowUnix() // current unix timestamp

t := NowAddDay(1) // from now add 1 day
t := NowAddHour(1) // from now add 1 hour
t := NowAddMinutes(3) // from now add 3 minutes
t := NowAddSeconds(180) // from now add 180 seconds
```

**Convert time to date by template**

Template Chars:

```text
 Y,y - year
  Y - year 2006
  y - year 06
 m - month 01-12
 d - day 01-31
 H,h - hour
  H - hour 00-23
  h - hour 01-12
 I,i - minute
  I - minute 00-59
  i - minute 0-59
 S,s - second
  S - second 00-59
  s - second 0-59
```

> More, please see [char map](./timex/template.go)

Examples, use timex:

```go
tx := timex.Now()
date := tx.DateFormat("Y-m-d H:I:S") // Output: 2022-04-20 19:09:03
date = tx.DateFormat("y-m-d h:i:s") // Output: 22-04-20 07:9:3
```

Format time.Time:

```go
tx := time.Now()
date := timex.DateFormat(tx, "Y-m-d H:I:S") // Output: 2022-04-20 19:40:34
```

More usage:

```go
ts := timex.NowUnix() // current unix timestamp

date := FormatUnix(ts, "2006-01-02 15:04:05") // Get: 2022-04-20 19:40:34
date := FormatUnixByTpl(ts, "Y-m-d H:I:S") // Get: 2022-04-20 19:40:34
```


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
