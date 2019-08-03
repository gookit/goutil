# Go Util

[![GoDoc](https://godoc.org/github.com/gookit/goutil?status.svg)](https://godoc.org/github.com/gookit/goutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/goutil)](https://goreportcard.com/report/github.com/gookit/goutil)
[![Build Status](https://travis-ci.org/gookit/goutil.svg?branch=master)](https://travis-ci.org/gookit/goutil)

Go一些常用的工具函数收集、实现和整理

> **[EN README](README.md)**

- `arrutil` array/slice util
- `calc` number format calc
- `cliutil` CLI util
- `envutil` ENV util
- `fmtutil` format data tool
- `fsutil` filesystem util
- `jsonutil` JSON util
- `maputil` map util
- `netutil` network util
- `strutil` string util
- `testutil` test help util

## GoDoc

- [godoc for github](https://godoc.org/github.com/gookit/goutil)

## Util Packages

### Array/Slice

- package `github.com/gookit/goutil/arrutil`

```go
func Reverse(ss []string)
```

### Calc

- package `github.com/gookit/goutil/calc`

```go
func DataSize(size uint64) string
func ElapsedTime(startTime time.Time) string
func HowLongAgo(sec int64) string
func Percent(val, total int) float64
```

### CLI Util

- package `github.com/gookit/goutil/cliutil`

```go
func ExecCmd(binName string, args []string, workDir ...string) (string, error)
func ExecCommand(binName string, args []string, workDir ...string) (string, error)
func CurrentShell(onlyName bool) string
func ShellExec(cmdStr string, shells ...string) (string, error)
func HasShellEnv(shell string) bool
```

### ENV Util

- package `github.com/gookit/goutil/envutil`

```go
func HasShellEnv(shell string) bool
func IsConsole(out io.Writer) bool
func IsLinux() bool
func IsMSys() bool
func IsMac() bool
func IsSupport256Color() bool
func IsSupportColor() bool
func IsWin() bool
```

### Format Util

- package `github.com/gookit/goutil/fmtutil`

```go
func DataSize(bytes uint64) string
func HowLongAgo(sec int64) string
```

### Filesystem Util

- package `github.com/gookit/goutil/fsutil`

```go
func FileExists(path string) bool
func IsAbsPath(filename string) bool
func IsZipFile(filepath string) bool
func Unzip(archive, targetDir string) (err error)
```

### JSON Util

- package `github.com/gookit/goutil/jsonutil`

```go
func Decode(json []byte, v interface{}) error
func Encode(v interface{}) ([]byte, error)
func Pretty(v interface{}) (string, error)
func ReadFile(filePath string, v interface{}) error
func WriteFile(filePath string, data interface{}) error
func StripComments(src string) string
```

### Map Util

- package `github.com/gookit/goutil/maputil`

```go
func GetByPath(key string, mp map[string]interface{}) (val interface{}, ok bool)
func KeyToLower(src map[string]string) map[string]string
func Keys(mp interface{}) (keys []string)
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string
func Values(mp interface{}) (values []interface{})
```

### String Util

- package `github.com/gookit/goutil/strutil`

```go
func Base64Encode(src []byte) []byte
func GenMd5(s string) string
func LowerFirst(s string) string
func PadLeft(s, pad string, length int) string
func PadRight(s, pad string, length int) string
func Padding(s, pad string, length int, pos uint8) string
func PrettyJson(v interface{}) (string, error)
func RandomBytes(length int) ([]byte, error)
func RandomString(length int) (string, error)
func RenderTemplate(input string, data interface{}, isFile ...bool) string
func Repeat(s string, times int) string
func RepeatRune(char rune, times int) (chars []rune)
func Replaces(str string, pairs map[string]string) string
func Similarity(s, t string, rate float32) (float32, bool)
func Split(s, sep string) (ss []string)
func Substr(s string, pos, length int) string
func UpperFirst(s string) string
func UpperWord(s string) string
```

### Test Util

```go
func DiscardStdout() error
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder
func RestoreStdout() (s string)
func RewriteStdout()
```

## License

[MIT](LICENSE)
