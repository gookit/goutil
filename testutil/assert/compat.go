// Package assert provides compatibility for the old assert import path.
//
// Deprecated: use github.com/gookit/goutil/x/assert.
package assert

import (
	"reflect"

	"github.com/gookit/goutil/comdef"
	xassert "github.com/gookit/goutil/x/assert"
)

// TestingT is an interface wrapper around *testing.T.
//
// Deprecated: use github.com/gookit/goutil/x/assert.TestingT.
type TestingT = xassert.TestingT

// PanicRunFunc define.
//
// Deprecated: use github.com/gookit/goutil/x/assert.PanicRunFunc.
type PanicRunFunc = xassert.PanicRunFunc

var (
	// ShowFullPath on show error trace.
	//
	// Deprecated: use github.com/gookit/goutil/x/assert.ShowFullPath.
	ShowFullPath = xassert.ShowFullPath
	// EnableColor on show error trace.
	//
	// Deprecated: use github.com/gookit/goutil/x/assert.EnableColor.
	EnableColor = xassert.EnableColor
	// FailFast fail fast, stop test when first error(will call testing.T.FailNow()).
	//
	// Deprecated: use github.com/gookit/goutil/x/assert.FailFast.
	FailFast = xassert.FailFast
)

func syncConfig() {
	xassert.ShowFullPath = ShowFullPath
	xassert.EnableColor = EnableColor
	xassert.FailFast = FailFast
}

// Deprecated: use github.com/gookit/goutil/x/assert.DisableColor.
func DisableColor() {
	EnableColor = false
	syncConfig()
}

// Deprecated: use github.com/gookit/goutil/x/assert.HideFullPath.
func HideFullPath() {
	ShowFullPath = false
	syncConfig()
}

// Deprecated: use github.com/gookit/goutil/x/assert.SetFailFast.
func SetFailFast(enable bool) {
	FailFast = enable
	syncConfig()
}

// Deprecated: use github.com/gookit/goutil/x/assert.Must.
func Must(t TestingT, condition bool, fmtAndArgs ...any) {
	syncConfig()
	xassert.Must(t, condition, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Require.
func Require(t TestingT, condition bool, fmtAndArgs ...any) {
	syncConfig()
	xassert.Require(t, condition, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Fail.
func Fail(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Fail(t, failMsg, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.FailNow.
func FailNow(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.FailNow(t, failMsg, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Nil.
func Nil(t TestingT, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Nil(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotNil.
func NotNil(t TestingT, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotNil(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.True.
func True(t TestingT, give bool, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.True(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.False.
func False(t TestingT, give bool, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.False(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Empty.
func Empty(t TestingT, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Empty(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotEmpty.
func NotEmpty(t TestingT, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotEmpty(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Zero.
func Zero(t TestingT, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Zero(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotZero.
func NotZero(t TestingT, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotZero(t, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Panics.
func Panics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Panics(t, fn, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotPanics.
func NotPanics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotPanics(t, fn, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.PanicsMsg.
func PanicsMsg(t TestingT, fn PanicRunFunc, wantVal any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.PanicsMsg(t, fn, wantVal, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.PanicsErrMsg.
func PanicsErrMsg(t TestingT, fn PanicRunFunc, errMsg string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.PanicsErrMsg(t, fn, errMsg, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Contains.
func Contains(t TestingT, src, elem any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Contains(t, src, elem, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotContains.
func NotContains(t TestingT, src, elem any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotContains(t, src, elem, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ContainsKey.
func ContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ContainsKey(t, mp, key, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotContainsKey.
func NotContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotContainsKey(t, mp, key, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ContainsKeys.
func ContainsKeys(t TestingT, mp any, keys any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ContainsKeys(t, mp, keys, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotContainsKeys.
func NotContainsKeys(t TestingT, mp any, keys any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotContainsKeys(t, mp, keys, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ContainsElems.
func ContainsElems[T comdef.ScalarType](t TestingT, list, sub []T, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ContainsElems[T](t, list, sub, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ContainsElemsAny.
func ContainsElemsAny(t TestingT, list, sub any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ContainsElemsAny(t, list, sub, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.StrContains.
func StrContains(t TestingT, s, sub string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.StrContains(t, s, sub, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.StrNotContains.
func StrNotContains(t TestingT, s, sub string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.StrNotContains(t, s, sub, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.StrContainsAll.
func StrContainsAll(t TestingT, s string, subs []string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.StrContainsAll(t, s, subs, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.StrCount.
func StrCount(t TestingT, s, sub string, count int, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.StrCount(t, s, sub, count, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.FileExists.
func FileExists(t TestingT, filePath string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.FileExists(t, filePath, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.FileNotExists.
func FileNotExists(t TestingT, filePath string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.FileNotExists(t, filePath, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.DirExists.
func DirExists(t TestingT, dirPath string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.DirExists(t, dirPath, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.DirNotExists.
func DirNotExists(t TestingT, dirPath string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.DirNotExists(t, dirPath, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NoError.
func NoError(t TestingT, err error, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NoError(t, err, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NoErr.
func NoErr(t TestingT, err error, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NoErr(t, err, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Error.
func Error(t TestingT, err error, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Error(t, err, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Err.
func Err(t TestingT, err error, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Err(t, err, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ErrIs.
func ErrIs(t TestingT, err, wantErr error, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ErrIs(t, err, wantErr, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ErrMsg.
func ErrMsg(t TestingT, err error, wantMsg string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ErrMsg(t, err, wantMsg, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ErrMsgContains.
func ErrMsgContains(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ErrMsgContains(t, err, subMsg, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ErrSubMsg.
func ErrSubMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ErrSubMsg(t, err, subMsg, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.ErrHasMsg.
func ErrHasMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.ErrHasMsg(t, err, subMsg, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Len.
func Len(t TestingT, give any, wantLn int, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Len(t, give, wantLn, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.LenGt.
func LenGt(t TestingT, give any, minLn int, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.LenGt(t, give, minLn, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Equal.
func Equal(t TestingT, want, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Equal(t, want, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Eq.
func Eq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Eq(t, want, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Neq.
func Neq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Neq(t, want, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotEqual.
func NotEqual(t TestingT, want, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotEqual(t, want, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotEq.
func NotEq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotEq(t, want, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Lt.
func Lt(t TestingT, give, max any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Lt(t, give, max, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Lte.
func Lte(t TestingT, give, max any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Lte(t, give, max, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Gt.
func Gt(t TestingT, give, min any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Gt(t, give, min, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Gte.
func Gte(t TestingT, give, min any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Gte(t, give, min, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.EqInt.
func EqInt(t TestingT, want, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.EqInt(t, want, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.EqFloat.
func EqFloat(t TestingT, want, give any, delta float64, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.EqFloat(t, want, give, delta, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.InDelta.
func InDelta(t TestingT, want, give any, delta float64, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.InDelta(t, want, give, delta, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.IsType.
func IsType(t TestingT, wantType, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.IsType(t, wantType, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.IsKind.
func IsKind(t TestingT, wantKind reflect.Kind, give any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.IsKind(t, wantKind, give, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.Same.
func Same(t TestingT, wanted, actual any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.Same(t, wanted, actual, fmtAndArgs...)
}

// Deprecated: use github.com/gookit/goutil/x/assert.NotSame.
func NotSame(t TestingT, want, actual any, fmtAndArgs ...any) bool {
	syncConfig()
	return xassert.NotSame(t, want, actual, fmtAndArgs...)
}
