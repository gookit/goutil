# Go Util

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gookit/goutil?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gookit/goutil)](https://github.com/gookit/goutil)
[![GoDoc](https://godoc.org/github.com/gookit/goutil?status.svg)](https://pkg.go.dev/github.com/gookit/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/goutil)](https://goreportcard.com/report/github.com/gookit/goutil)
[![Unit-Tests](https://github.com/gookit/goutil/workflows/Unit-Tests/badge.svg)](https://github.com/gookit/goutil/actions)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)

💪 Useful utils for the Go: string, array/slice, map, format, CLI, ENV, filesystem, testing and more.

- `arrutil` array/slice util functions
- `dump`  Simple variable printing tool, printing slice, map will automatically wrap each element and display the call location
- `cliutil` CLI util functions
- `envutil` ENV util for check current runtime env information
- `fmtutil` format data util functions
- `fsutil` filesystem util functions
- `jsonutil` JSON util functions
- `maputil` map util functions
- `mathutil` math util functions
- `netutil` network util functions
- `strutil` string util functions
- `testutil` test help util functions

> **[中文说明](README.zh-CN.md)**

## GoDoc

- [Godoc for github](https://pkg.go.dev/github.com/gookit/goutil)

## Packages

### Array/Slice

> Package `github.com/gookit/goutil/arrutil`

```go
// source at arrutil/arrutil.go
func Reverse(ss []string)
func StringsRemove(ss []string, s string) []string
func TrimStrings(ss []string, cutSet ...string) (ns []string)
func GetRandomOne(arr interface{}) interface{}
// source at arrutil/check.go
func IntsHas(ints []int, val int) bool
func Int64sHas(ints []int64, val int64) bool
func StringsHas(ss []string, val string) bool
func HasValue(arr, val interface{}) bool
func Contains(arr, val interface{}) bool
func NotContains(arr, val interface{}) bool
// source at arrutil/convert.go
func ToInt64s(arr interface{})(ret []int64, err error)
func MustToInt64s(arr interface{}) []int64
func SliceToInt64s(arr []interface{}) []int64
func ToStrings(arr interface{})(ret []string, err error)
func MustToStrings(arr interface{}) []string
func SliceToStrings(arr []interface{}) []string
func StringsToInts(ss []string) (ints []int, err error)
```

### CLI

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
// source at cliutil/read.go
func ReadInput(question string) (string, error)
func ReadLine(question string) (string, error)
func ReadFirst(question string) (string, error)
func ReadFirstByte(question string) (byte, error)
func ReadFirstRune(question string) (rune, error)
// source at cliutil/read_nonwin.go
func ReadPassword(question ...string) string
```

#### Examples

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

output:

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


### Dump

> Package `github.com/gookit/goutil/dump`

```go
// source at dump/dump.go
func Std() *Dumper
func Reset()
func Config(fn func(opts *Options))
func Print(vs ...interface{})
func Println(vs ...interface{})
func Fprint(w io.Writer, vs ...interface{})
// source at dump/dumper.go
func NewDumper(out io.Writer, skip int) *Dumper
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


### ENV

> Package `github.com/gookit/goutil/envutil`

```go
// source at envutil/envutil.go
func VarParse(str string) string
func ParseEnvValue(val string) (newVal string)
// source at envutil/get.go
func Getenv(name string, def ...string) string
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

### Formatting

> Package `github.com/gookit/goutil/fmtutil`

```go
// source at fmtutil/format.go
func DataSize(bytes uint64) string
func PrettyJSON(v interface{}) (string, error)
func StringsToInts(ss []string) (ints []int, err error)
func ArgsWithSpaces(args []interface{}) (message string)
// source at fmtutil/time.go
func HowLongAgo(sec int64) string
```

### FileSystem

> Package `github.com/gookit/goutil/fsutil`

```go
// source at fsutil/check.go
func Dir(fpath string) string
func Name(fpath string) string
func FileExt(fpath string) string
func Suffix(fpath string) string
func PathExists(path string) bool
func IsDir(path string) bool
func FileExists(path string) bool
func IsFile(path string) bool
func IsAbsPath(aPath string) bool
func IsImageFile(path string) bool
func IsZipFile(filepath string) bool
// source at fsutil/finder.go
func EmptyFinder() *FileFinder
func NewFinder(dirPaths []string, filePaths ...string) *FileFinder
func ExtFilterFunc(exts []string, include bool) FileFilterFunc
func SuffixFilterFunc(suffixes []string, include bool) FileFilterFunc
func PathNameFilterFunc(names []string, include bool) FileFilterFunc
func DotFileFilterFunc(include bool) FileFilterFunc
func ModTimeFilterFunc(limitSec int, op rune, include bool) FileFilterFunc
func GlobFilterFunc(patterns []string, include bool) FileFilterFunc
func RegexFilterFunc(pattern string, include bool) FileFilterFunc
func DotDirFilterFunc(include bool) DirFilterFunc
func DirNameFilterFunc(names []string, include bool) DirFilterFunc
// source at fsutil/fsutil.go
func OSTempFile(pattern string) (*os.File, error)
func TempFile(dir, pattern string) (*os.File, error)
func OSTempDir(pattern string) (string, error)
func TempDir(dir, pattern string) (string, error)
func ExpandPath(path string) string
func MimeType(path string) (mime string)
func ReaderMimeType(r io.Reader) (mime string)
// source at fsutil/operate.go
func Mkdir(dirPath string, perm os.FileMode) error
func MkParentDir(fpath string) error
func MustReadFile(filePath string) []byte
func ReadExistFile(filePath string) []byte
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error)
func QuickOpenFile(filepath string) (*os.File, error)
func CreateFile(fpath string, filePerm, dirPerm os.FileMode) (*os.File, error)
func MustCreateFile(filePath string, filePerm, dirPerm os.FileMode) *os.File
func CopyFile(src string, dst string) error
func MustCopyFile(src string, dst string)
func MustRemove(fpath string)
func QuietRemove(fpath string)
func DeleteIfExist(fpath string) error
func DeleteIfFileExist(fpath string) error
func Unzip(archive, targetDir string) (err error)
```

#### Examples

**files finder:**

```go
package main

import (
	"fmt"
	"os"

	"github.com/gookit/goutil/fsutil"
)

func main() {
	f := fsutil.EmptyFinder()

	f.
		AddDir("./testdata").
		AddFile("finder.go").
		NoDotFile().
		// NoDotDir().
		Find().
		Each(func(filePath string) {
			fmt.Println(filePath)
		})

	fsutil.NewFinder([]string{"./testdata"}).
		AddFile("finder.go").
		NoDotDir().
		EachStat(func(fi os.FileInfo, filePath string) {
			fmt.Println(filePath, "=>", fi.ModTime())
		})
}
```


### JSON

> Package `github.com/gookit/goutil/jsonutil`

```go
// source at jsonutil/jsonutil.go
func WriteFile(filePath string, data interface{}) error
func ReadFile(filePath string, v interface{}) error
func Encode(v interface{}) ([]byte, error)
func Decode(json []byte, v interface{}) error
func Pretty(v interface{}) (string, error)
func StripComments(src string) string
```

### Map

> Package `github.com/gookit/goutil/maputil`

```go
// source at maputil/convert.go
func KeyToLower(src map[string]string) map[string]string
func ToStringMap(src map[string]interface{}) map[string]string
func HttpQueryString(data map[string]interface{}) string
// source at maputil/maputil.go
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string
func GetByPath(key string, mp map[string]interface{}) (val interface{}, ok bool)
func Keys(mp interface{}) (keys []string)
func Values(mp interface{}) (values []interface{})
```

### Math/Number

> Package `github.com/gookit/goutil/mathutil`

```go
// source at mathutil/convert.go
func Int(in interface{}) (int, error)
func MustInt(in interface{}) int
func ToInt(in interface{}) (iVal int, err error)
func Uint(in interface{}) (uint64, error)
func MustUint(in interface{}) uint64
func ToUint(in interface{}) (u64 uint64, err error)
func Int64(in interface{}) (int64, error)
func MustInt64(in interface{}) int64
func ToInt64(in interface{}) (i64 int64, err error)
func Float(in interface{}) (float64, error)
func ToFloat(in interface{}) (f64 float64, err error)
func MustFloat(in interface{}) float64
// source at mathutil/number.go
func IsNumeric(c byte) bool
func Percent(val, total int) float64
func ElapsedTime(startTime time.Time) string
func DataSize(size uint64) string
func HowLongAgo(sec int64) string
// source at mathutil/random.go
func RandomInt(min, max int) int
```

### Struct

> Package `github.com/gookit/goutil/structs`

```go
// source at structs/alias.go
func NewAliases(checker func(alias string)) *Aliases
// source at structs/tags.go
func ParseTags(v interface{}) error
func ParseReflectTags(v reflect.Value) error
```

### String

> Package `github.com/gookit/goutil/strutil`

```go
// source at strutil/check.go
func IsNumeric(c byte) bool
func IsAlphabet(char uint8) bool
func IsAlphaNum(c uint8) bool
func StrPos(s, sub string) int
func BytePos(s string, bt byte) int
func RunePos(s string, ru rune) int
func IsStartOf(s, sub string) bool
func IsEndOf(s, sub string) bool
func Len(s string) int
func Utf8len(s string) int
func ValidUtf8String(s string) bool
func IsSpace(c byte) bool
func IsSpaceRune(r rune) bool
func IsBlank(s string) bool
func IsBlankBytes(bs []byte) bool
func IsSymbol(r rune) bool
// source at strutil/convert.go
func String(val interface{}) (string, error)
func MustString(in interface{}) string
func ToString(val interface{}) (str string, err error)
func AnyToString(val interface{}, defaultAsErr bool) (str string, err error)
func ToBool(s string) (bool, error)
func MustBool(s string) bool
func Bool(s string) (bool, error)
func Int(s string) (int, error)
func ToInt(s string) (int, error)
func MustInt(s string) int
func ToInts(s string, sep ...string) ([]int, error)
func ToIntSlice(s string, sep ...string) (ints []int, err error)
func ToArray(s string, sep ...string) []string
func ToSlice(s string, sep ...string) []string
func ToOSArgs(s string) []string
func ToTime(s string, layouts ...string) (t time.Time, err error)
// source at strutil/encode.go
func Base64(str string) string
func B64Encode(str string) string
func URLEncode(s string) string
func URLDecode(s string) string
// source at strutil/find_similar.go
func NewComparator(src, dst string) *SimilarComparator
func Similarity(s, t string, rate float32) (float32, bool)
// source at strutil/format.go
func Lowercase(s string) string
func Uppercase(s string) string
func UpperWord(s string) string
func LowerFirst(s string) string
func UpperFirst(s string) string
func Snake(s string, sep ...string) string
func SnakeCase(s string, sep ...string) string
func Camel(s string, sep ...string) string
func CamelCase(s string, sep ...string) string
// source at strutil/id.go
func MicroTimeID() string
func MicroTimeHexID() string
// source at strutil/random.go
func Md5(src interface{}) string
func GenMd5(src interface{}) string
func RandomChars(ln int) string
func RandomCharsV2(ln int) string
func RandomCharsV3(ln int) string
func RandomBytes(length int) ([]byte, error)
func RandomString(length int) (string, error)
// source at strutil/strutil.go
func Trim(s string, cutSet ...string) string
func TrimLeft(s string, cutSet ...string) string
func TrimRight(s string, cutSet ...string) string
func FilterEmail(s string) string
func Split(s, sep string) (ss []string)
func Substr(s string, pos, length int) string
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
```

### System

> Package `github.com/gookit/goutil/sysutil`

```go
// source at sysutil/exec.go
func QuickExec(cmdLine string, workDir ...string) (string, error)
func ExecLine(cmdLine string, workDir ...string) (string, error)
func ExecCmd(binName string, args []string, workDir ...string) (string, error)
func ShellExec(cmdLine string, shells ...string) (string, error)
func FindExecutable(binName string) (string, error)
func Executable(binName string) (string, error)
func HasExecutable(binName string) bool
// source at sysutil/sysenv.go
func UserHomeDir() string
func HomeDir() string
func ExpandPath(path string) string
func Hostname() string
func IsWin() bool
func IsWindows() bool
func IsMac() bool
func IsLinux() bool
func IsMSys() bool
func IsConsole(out io.Writer) bool
func IsTerminal(fd uintptr) bool
func StdIsTerminal() bool
func CurrentShell(onlyName bool) (path string)
func HasShellEnv(shell string) bool
func IsShellSpecialVar(c uint8) bool
// source at sysutil/sysutil_nonwin.go
func Kill(pid int, signal syscall.Signal) error
func ProcessExists(pid int) bool
```

### Testing

> Package `github.com/gookit/goutil/testutil`

```go
// source at testutil/httpmock.go
func NewHttpRequest(method, path string, data *MD) *http.Request
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder
// source at testutil/testutil.go
func DiscardStdout() error
func ReadOutput() (s string)
func RewriteStdout()
func RestoreStdout() (s string)
func RewriteStderr()
func RestoreStderr() (s string)
func MockEnvValue(key, val string, fn func(nv string))
func MockEnvValues(kvMap map[string]string, fn func())
```

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
