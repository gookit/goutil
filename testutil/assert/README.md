# Assert Utils

Package assert provides some tool functions for use with the Go testing.

## Install

```bash
go get github.com/gookit/goutil/sysutil
```

## GoDocs

Please see [Go docs](https://pkg.go.dev/github.com/gookit/goutil/testutil/assert)

## Function API

```go
func Contains(t TestingT, src, elem any, fmtAndArgs ...any) bool
func ContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool
func ContainsKeys(t TestingT, mp any, keys any, fmtAndArgs ...any) bool
func DisableColor()
func Empty(t TestingT, give any, fmtAndArgs ...any) bool
func Eq(t TestingT, want, give any, fmtAndArgs ...any) bool
func Equal(t TestingT, want, give any, fmtAndArgs ...any) bool
func Err(t TestingT, err error, fmtAndArgs ...any) bool
func ErrIs(t TestingT, err, wantErr error, fmtAndArgs ...any) bool
func ErrMsg(t TestingT, err error, wantMsg string, fmtAndArgs ...any) bool
func ErrSubMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool
func Fail(t TestingT, failMsg string, fmtAndArgs ...any) bool
func FailNow(t TestingT, failMsg string, fmtAndArgs ...any) bool
func False(t TestingT, give bool, fmtAndArgs ...any) bool
func Gt(t TestingT, give, min int, fmtAndArgs ...any) bool
func HideFullPath()
func IsKind(t TestingT, wantKind reflect.Kind, give any, fmtAndArgs ...any) bool
func IsType(t TestingT, wantType, give any, fmtAndArgs ...any) bool
func Len(t TestingT, give any, wantLn int, fmtAndArgs ...any) bool
func LenGt(t TestingT, give any, minLn int, fmtAndArgs ...any) bool
func Lt(t TestingT, give, max int, fmtAndArgs ...any) bool
func Neq(t TestingT, want, give any, fmtAndArgs ...any) bool
func Nil(t TestingT, give any, fmtAndArgs ...any) bool
func NoErr(t TestingT, err error, fmtAndArgs ...any) bool
func NotContains(t TestingT, src, elem any, fmtAndArgs ...any) bool
func NotEmpty(t TestingT, give any, fmtAndArgs ...any) bool
func NotEq(t TestingT, want, give any, fmtAndArgs ...any) bool
func NotEqual(t TestingT, want, give any, fmtAndArgs ...any) bool
func NotNil(t TestingT, give any, fmtAndArgs ...any) bool
func NotPanics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool
func NotSame(t TestingT, want, actual any, fmtAndArgs ...any) bool
func Panics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool
func PanicsErrMsg(t TestingT, fn PanicRunFunc, errMsg string, fmtAndArgs ...any) bool
func PanicsMsg(t TestingT, fn PanicRunFunc, wantVal interface{}, fmtAndArgs ...any) bool
func Same(t TestingT, wanted, actual any, fmtAndArgs ...any) bool
func StrContains(t TestingT, s, sub string, fmtAndArgs ...any) bool
func True(t TestingT, give bool, fmtAndArgs ...any) bool
type Assertions struct{ ... }
    func New(t TestingT) *Assertions
```

## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...
go test ./...
```
