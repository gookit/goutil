# Go Util

[![GoDoc](https://godoc.org/github.com/gookit/goutil?status.svg)](https://godoc.org/github.com/gookit/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/goutil)](https://goreportcard.com/report/github.com/gookit/goutil)
[![Build Status](https://travis-ci.org/gookit/goutil.svg?branch=master)](https://travis-ci.org/gookit/goutil)
[![Coverage Status](https://coveralls.io/repos/github/gookit/goutil/badge.svg?branch=master)](https://coveralls.io/github/gookit/goutil?branch=master)

Go一些常用的工具函数收集、实现和整理

> **[EN README](README.md)**

- `arrutil` array/slice util
- `dump`  print debug vars
- `cliutil` CLI util
- `envutil` ENV util
- `fmtutil` format data tool
- `fsutil` filesystem util
- `jsonutil` JSON util
- `maputil` map util
- `mathutil` math util
- `netutil` network util
- `strutil` string util
- `testutil` test help util

## GoDoc

- [godoc for github](https://godoc.org/github.com/gookit/goutil)

## Util Packages

### Array/Slice

> package `github.com/gookit/goutil/arrutil`

```text
func Reverse(ss []string)
func StringsRemove(ss []string, s string) []string
func StringsToInts(ss []string) (ints []int, err error)
func TrimStrings(ss []string, cutSet ...string) (ns []string)
```

### CLI Util

> package `github.com/gookit/goutil/cliutil`

```text
func CurrentShell(onlyName bool) (path string)
func ExecCmd(binName string, args []string, workDir ...string) (string, error)
func ExecCommand(binName string, args []string, workDir ...string) (string, error)
func HasShellEnv(shell string) bool
func QuickExec(cmdLine string, workDir ...string) (string, error)
func ShellExec(cmdLine string, shells ...string) (string, error)
```

### Dump Util

> package `github.com/gookit/goutil/dump`

```text
func P(vs ...interface{})
func V(vs ...interface{})
func Print(vs ...interface{})
```

Usage please see [dump/README.md](dump/README.md)

### ENV Util

> package `github.com/gookit/goutil/envutil`

```text
func Getenv(name string, def ...string) string
func HasShellEnv(shell string) bool
func IsConsole(out io.Writer) bool
func IsLinux() bool
func IsMSys() bool
func IsMac() bool
func IsSupport256Color() bool
func IsSupportColor() bool
func IsWin() bool
func ParseEnvValue(val string) (newVal string)
```

### Format Util

> package `github.com/gookit/goutil/fmtutil`

```text
func DataSize(bytes uint64) string
func HowLongAgo(sec int64) string
func PrettyJSON(v interface{}) (string, error)
func StringsToInts(ss []string) (ints []int, err error)
```

### Filesystem Util

> package `github.com/gookit/goutil/fsutil`

```text
func FileExists(path string) bool
func IsAbsPath(filepath string) bool
func IsDir(path string) bool
func IsFile(path string) bool
func IsZipFile(filepath string) bool
func Mkdir(name string, perm os.FileMode) error
func PathExists(path string) bool
func Unzip(archive, targetDir string) (err error)
```

### JSON Util

> package `github.com/gookit/goutil/jsonutil`

```text
func Decode(json []byte, v interface{}) error
func Encode(v interface{}) ([]byte, error)
func Pretty(v interface{}) (string, error)
func ReadFile(filePath string, v interface{}) error
func WriteFile(filePath string, data interface{}) error
func StripComments(src string) string
```

### Map Util

- package `github.com/gookit/goutil/maputil`

```text
func GetByPath(key string, mp map[string]interface{}) (val interface{}, ok bool)
func KeyToLower(src map[string]string) map[string]string
func Keys(mp interface{}) (keys []string)
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string
func Values(mp interface{}) (values []interface{})
```

### Math Util

> package `github.com/gookit/goutil/mathutil`

```text
func DataSize(size uint64) string
func ElapsedTime(startTime time.Time) string
func Float(s string) (float64, error)
func HowLongAgo(sec int64) string
func Int(in interface{}) (int, error)
func Int64(in interface{}) (int64, error)
func MustFloat(s string) float64
func MustInt(in interface{}) int
func MustInt64(in interface{}) int64
func MustUint(in interface{}) uint64
func Percent(val, total int) float64
func ToFloat(s string) (float64, error)
func ToInt(in interface{}) (iVal int, err error)
func ToInt64(in interface{}) (i64 int64, err error)
func ToUint(in interface{}) (u64 uint64, err error)
func Uint(in interface{}) (uint64, error)
```

### String Util

> package `github.com/gookit/goutil/strutil`

```text
func B64Encode(str string) string
func Bool(s string) (bool, error)
func Camel(s string, sep ...string) string
func CamelCase(s string, sep ...string) string
func FilterEmail(s string) string
func GenMd5(src interface{}) string
func LowerFirst(s string) string
func Lowercase(s string) string
func Md5(src interface{}) string
func MustBool(s string) bool
func MustString(in interface{}) string
func PadLeft(s, pad string, length int) string
func PadRight(s, pad string, length int) string
func Padding(s, pad string, length int, pos uint8) string
func PrettyJSON(v interface{}) (string, error)
func RandomBytes(length int) ([]byte, error)
func RandomString(length int) (string, error)
func RenderTemplate(input string, data interface{}, fns template.FuncMap, isFile ...bool) string
func Repeat(s string, times int) string
func RepeatRune(char rune, times int) (chars []rune)
func Replaces(str string, pairs map[string]string) string
func Similarity(s, t string, rate float32) (float32, bool)
func Snake(s string, sep ...string) string
func SnakeCase(s string, sep ...string) string
func Split(s, sep string) (ss []string)
func String(val interface{}) (string, error)
func Substr(s string, pos, length int) string
func ToArray(s string, sep ...string) []string
func ToBool(s string) (bool, error)
func ToIntSlice(s string, sep ...string) (ints []int, err error)
func ToInts(s string, sep ...string) ([]int, error)
func ToSlice(s string, sep ...string) []string
func ToString(val interface{}) (str string, err error)
func ToTime(s string, layouts ...string) (t time.Time, err error)
func Trim(s string, cutSet ...string) string
func TrimLeft(s string, cutSet ...string) string
func TrimRight(s string, cutSet ...string) string
func URLDecode(s string) string
func URLEncode(s string) string
func UpperFirst(s string) string
func UpperWord(s string) string
func Uppercase(s string) string
```

### System Util

> package `github.com/gookit/goutil/sysutil`

```text
func CurrentShell(onlyName bool) (path string)
func ExecCmd(binName string, args []string, workDir ...string) (string, error)
func HasShellEnv(shell string) bool
func IsConsole(out io.Writer) bool
func IsLinux() bool
func IsMSys() bool
func IsMac() bool
func IsWin() bool
func Kill(pid int, signal syscall.Signal) error
func ProcessExists(pid int) bool
func QuickExec(cmdLine string, workDir ...string) (string, error)
func ShellExec(cmdStr string, shells ...string) (string, error)
```

### Test Util

> package `github.com/gookit/goutil/testutil`

```text
func DiscardStdout() error
func MockEnvValue(key, val string, fn func(nv string))
func MockEnvValues(kvMap map[string]string, fn func())
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder
func RestoreStdout() (s string)
func RewriteStdout()
```

## License

[MIT](LICENSE)
