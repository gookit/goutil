package assert

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/internal/checkfn"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/strutil"
)

// Nil asserts that the given is a nil value
func Nil(t TestingT, give any, fmtAndArgs ...any) bool {
	if checkfn.IsNil(give) {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Expected nil, but got:\n %+v", give), fmtAndArgs)
}

// NotNil asserts that the given is a not nil value
func NotNil(t TestingT, give any, fmtAndArgs ...any) bool {
	if !checkfn.IsNil(give) {
		return true
	}

	t.Helper()
	return fail(t, "Should not nil value", fmtAndArgs)
}

// True asserts that the given is a bool true
func True(t TestingT, give bool, fmtAndArgs ...any) bool {
	if !give {
		t.Helper()
		return fail(t, "Result should be True", fmtAndArgs)
	}
	return true
}

// False asserts that the given is a bool false
func False(t TestingT, give bool, fmtAndArgs ...any) bool {
	if give {
		t.Helper()
		return fail(t, "Result should be False", fmtAndArgs)
	}
	return true
}

// Empty asserts that the give should be empty
func Empty(t TestingT, give any, fmtAndArgs ...any) bool {
	empty := isEmpty(give)
	if !empty {
		t.Helper()
		return fail(t, fmt.Sprintf("Should be empty, but was:\n%#v", give), fmtAndArgs)
	}

	return empty
}

// NotEmpty asserts that the give should not be empty
func NotEmpty(t TestingT, give any, fmtAndArgs ...any) bool {
	nEmpty := !isEmpty(give)
	if !nEmpty {
		t.Helper()
		return fail(t, fmt.Sprintf("Should not be empty, but was:\n%#v", give), fmtAndArgs)
	}

	return nEmpty
}

// PanicRunFunc define
type PanicRunFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func runPanicFunc(f PanicRunFunc) (didPanic bool, message any, stack string) {
	didPanic = true
	defer func() {
		message = recover()
		if didPanic {
			stack = string(debug.Stack())
		}
	}()

	// call the target function
	f()
	didPanic = false

	return
}

// Panics asserts that the code inside the specified func panics.
func Panics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool {
	if hasPanic, panicVal, _ := runPanicFunc(fn); !hasPanic {
		t.Helper()
		funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		return fail(t, fmt.Sprintf("func '%s' should panic\n\tPanic value:\t%#v", funcName, panicVal), fmtAndArgs)
	}
	return true
}

// NotPanics asserts that the code inside the specified func NOT panics.
func NotPanics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) bool {
	if hasPanic, panicVal, stackMsg := runPanicFunc(fn); hasPanic {
		t.Helper()
		funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()

		return fail(t, fmt.Sprintf(
			"func %s should not panic\n\tPanic value:\t%#v\n\tPanic stack:\t%s",
			funcName, panicVal, stackMsg,
		), fmtAndArgs)
	}

	return true
}

// PanicsMsg should panic and with a value
func PanicsMsg(t TestingT, fn PanicRunFunc, wantVal any, fmtAndArgs ...any) bool {
	funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	hasPanic, panicVal, stackMsg := runPanicFunc(fn)
	if !hasPanic {
		t.Helper()
		return fail(t, fmt.Sprintf("func %s should panic\n\tPanic value:\t%#v", funcName, panicVal), fmtAndArgs)
	}

	if panicVal != wantVal {
		t.Helper()
		return fail(t, fmt.Sprintf(
			"func %s should panic.\n\tWant  value:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s",
			funcName, wantVal, panicVal, stackMsg),
			fmtAndArgs)
	}

	return true
}

// PanicsErrMsg should panic and with error message
func PanicsErrMsg(t TestingT, fn PanicRunFunc, errMsg string, fmtAndArgs ...any) bool {
	funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	hasPanic, panicVal, stackMsg := runPanicFunc(fn)
	if !hasPanic {
		t.Helper()
		return fail(t, fmt.Sprintf("func %s should panic\n\tPanic value:\t%#v", funcName, panicVal), fmtAndArgs)
	}

	err, ok := panicVal.(error)
	if !ok {
		t.Helper()
		return fail(t, fmt.Sprintf("func %s should panic and is error type,\nbut type was: %T", funcName, panicVal), fmtAndArgs)
	}

	if err.Error() != errMsg {
		t.Helper()
		return fail(t, fmt.Sprintf(
			"func %s should panic.\n\tWant  error:\t%#v\n\tPanic error:\t%#v\n\tPanic stack:\t%s",
			funcName, errMsg, err.Error(), stackMsg),
			fmtAndArgs,
		)
	}

	return true
}

// Contains asserting that the given data(string,slice,map) should contain element
//
// TIP: only support types: string, map, array, slice
//
//	map         - check key exists
//	string      - check substring exists
//	array,slice - check sub-element exists
func Contains(t TestingT, src, elem any, fmtAndArgs ...any) bool {
	valid, found := checkfn.Contains(src, elem)
	if valid && found {
		return true
	}

	t.Helper()

	// src invalid
	if !valid {
		return fail(t, fmt.Sprintf("%#v could not be applied builtin len()", src), fmtAndArgs)
	}

	// not found
	return fail(t, fmt.Sprintf("%#v\nShould contain: %#v", src, elem), fmtAndArgs)
}

// NotContains asserts that the given data(string,slice,map) should not contain element
//
// TIP: only support types: string, map, array, slice
//
//	map         - check key exists
//	string      - check substring exists
//	array,slice - check sub-element exists
func NotContains(t TestingT, src, elem any, fmtAndArgs ...any) bool {
	valid, found := checkfn.Contains(src, elem)
	if valid && !found {
		return true
	}

	t.Helper()

	// src invalid
	if !valid {
		return fail(t, fmt.Sprintf("%#v could not be applied builtin len()", src), fmtAndArgs)
	}

	// found
	return fail(t, fmt.Sprintf("%#v\nShould not contain: %#v", src, elem), fmtAndArgs)
}

// ContainsKey asserts that the given map is containing key
func ContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool {
	if !maputil.HasKey(mp, key) {
		t.Helper()
		return fail(t, fmt.Sprintf(
			"Map should contains the key: %#v\nMap data:\n%v",
			key, maputil.FormatIndent(mp, "  "),
		), fmtAndArgs)
	}

	return true
}

// NotContainsKey asserts that the given map is not contains key
func NotContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool {
	if maputil.HasKey(mp, key) {
		t.Helper()
		return fail(t,
			fmt.Sprintf(
				"Map should not contains the key: %#v\nMap data:\n%v",
				key,
				maputil.FormatIndent(mp, "  "),
			),
			fmtAndArgs,
		)
	}

	return true
}

// ContainsKeys asserts that the map is contains all given keys
//
// Usage:
//
//	ContainsKeys(t, map[string]any{...}, []string{"key1", "key2"})
func ContainsKeys(t TestingT, mp any, keys any, fmtAndArgs ...any) bool {
	anyKeys, err := arrutil.AnyToSlice(keys)
	if err != nil {
		t.Helper()
		return fail(t, err.Error(), fmtAndArgs)
	}

	ok, noKey := maputil.HasAllKeys(mp, anyKeys...)
	if !ok {
		t.Helper()
		return fail(t, fmt.Sprintf(
			"Map should contains the key: %#v\nMap data:\n%v",
			noKey, maputil.FormatIndent(mp, "  "),
		), fmtAndArgs)
	}

	return true
}

// NotContainsKeys asserts that the map is not contains all given keys
//
// Usage:
//
//	NotContainsKeys(t, map[string]any{...}, []string{"key1", "key2"})
func NotContainsKeys(t TestingT, mp any, keys any, fmtAndArgs ...any) bool {
	anyKeys, err := arrutil.AnyToSlice(keys)
	if err != nil {
		t.Helper()
		return fail(t, err.Error(), fmtAndArgs)
	}

	ok, hasKey := maputil.HasOneKey(mp, anyKeys...)
	if ok {
		t.Helper()
		return fail(t, fmt.Sprintf("Map should not contains the key: %#v\nMap data:\n%v",
			hasKey, maputil.FormatIndent(mp, "  "),
		), fmtAndArgs)
	}

	return true
}

// ContainsElems asserts that the given list should contain sub elements.
func ContainsElems[T comdef.ScalarType](t TestingT, list, sub []T, fmtAndArgs ...any) bool {
	if arrutil.ContainsAll(list, sub) {
		return true
	}

	t.Helper()

	// not contains all
	return fail(t, fmt.Sprintf("%#v\nShould contain: %#v", list, sub), fmtAndArgs)
}

// StrContains asserts that the given string should contain substring
func StrContains(t TestingT, s, sub string, fmtAndArgs ...any) bool {
	if strings.Contains(s, sub) {
		return true
	}

	t.Helper()
	return fail(t,
		fmt.Sprintf("String check fail:\nGiven string: %#v\nNot contains: %#v", s, sub),
		fmtAndArgs,
	)
}

// StrNotContains asserts that the given string should not contain substring
func StrNotContains(t TestingT, s, sub string, fmtAndArgs ...any) bool {
	if !strings.Contains(s, sub) {
		return true
	}

	t.Helper()
	return fail(t,
		fmt.Sprintf("String check fail:\nGiven string: %#v\nShould not contains: %#v", s, sub),
		fmtAndArgs,
	)
}

// StrContainsAll asserts that the given string should contain all substrings
func StrContainsAll(t TestingT, s string, subs []string, fmtAndArgs ...any) bool {
	if strutil.ContainsAll(s, subs) {
		return true
	}

	t.Helper()
	return fail(t,
		fmt.Sprintf("String check fail:\nGiven string: %#v\nShould contains all: %#v", s, subs),
		fmtAndArgs,
	)
}

// StrCount asserts that the given string should contain substring and count
func StrCount(t TestingT, s, sub string, count int, fmtAndArgs ...any) bool {
	if strings.Count(s, sub) == count {
		return true
	}

	t.Helper()
	return fail(t,
		fmt.Sprintf("String check fail:\nGiven string: %s\nNot contains %q count: %d", s, sub, count),
		fmtAndArgs,
	)
}

//
// -------------------- filesystem --------------------
//

// FileExists asserts that the given file exists
func FileExists(t TestingT, filePath string, fmtAndArgs ...any) bool {
	t.Helper()
	if !fsutil.IsFile(filePath) {
		return fail(t, fmt.Sprintf("File should exists: %s", filePath), fmtAndArgs)
	}
	return true
}

// FileNotExists asserts that the given file not exists
func FileNotExists(t TestingT, filePath string, fmtAndArgs ...any) bool {
	t.Helper()
	if fsutil.IsFile(filePath) {
		return fail(t, fmt.Sprintf("File should not exists: %s", filePath), fmtAndArgs)
	}
	return true
}

// DirExists asserts that the given dir exists
func DirExists(t TestingT, dirPath string, fmtAndArgs ...any) bool {
	t.Helper()
	if !fsutil.IsDir(dirPath) {
		return fail(t, fmt.Sprintf("Directory should exists: %s", dirPath), fmtAndArgs)
	}
	return true
}

// DirNotExists asserts that the given dir not exists
func DirNotExists(t TestingT, dirPath string, fmtAndArgs ...any) bool {
	t.Helper()
	if fsutil.IsDir(dirPath) {
		return fail(t, fmt.Sprintf("Directory should not exists: %s", dirPath), fmtAndArgs)
	}
	return true
}

//
// -------------------- error --------------------
//

// NoError asserts that the given is a nil error. alias of NoError()
func NoError(t TestingT, err error, fmtAndArgs ...any) bool {
	t.Helper()
	return NoErr(t, err, fmtAndArgs...)
}

// NoErr asserts that the given is a nil error
func NoErr(t TestingT, err error, fmtAndArgs ...any) bool {
	if err != nil {
		t.Helper()
		return fail(t, fmt.Sprintf("Received unexpected error:\n%+v", err), fmtAndArgs)
	}
	return true
}

// Error asserts that the given is a not nil error. alias of Error()
func Error(t TestingT, err error, fmtAndArgs ...any) bool {
	t.Helper()
	return Err(t, err, fmtAndArgs...)
}

// Err asserts that the given is a not nil error
func Err(t TestingT, err error, fmtAndArgs ...any) bool {
	if err == nil {
		t.Helper()
		return fail(t, "An error is expected but got nil.", fmtAndArgs)
	}
	return true
}

// ErrIs asserts that the given error is equals wantErr
func ErrIs(t TestingT, err, wantErr error, fmtAndArgs ...any) bool {
	if err == nil {
		t.Helper()
		return fail(t, "An error is expected but got nil.", fmtAndArgs)
	}

	if !errors.Is(err, wantErr) {
		t.Helper()
		return fail(t, fmt.Sprintf("Expect given err is equals %#v.", wantErr), fmtAndArgs)
	}

	return true
}

// ErrMsg asserts that the given is a not nil error and error message equals wantMsg
func ErrMsg(t TestingT, err error, wantMsg string, fmtAndArgs ...any) bool {
	if err == nil {
		t.Helper()
		return fail(t, "An error is expected but got nil.", fmtAndArgs)
	}

	errMsg := err.Error()
	if errMsg != wantMsg {
		t.Helper()
		return fail(t, fmt.Sprintf("Error message not equal:\n"+
			"expect: %q\n"+
			"actual: %q", wantMsg, errMsg), fmtAndArgs)
	}

	return true
}

// ErrMsgContains asserts that the given is a not nil error and error message contains subMsg
func ErrMsgContains(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	if err == nil {
		t.Helper()
		return fail(t, "An error is expected but got nil.", fmtAndArgs)
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, subMsg) {
		t.Helper()
		return fail(t, fmt.Sprintf("Error message check fail:\n"+
			"error  message : %q\n"+
			"should contains: %q", errMsg, subMsg), fmtAndArgs)
	}

	return true
}

// ErrSubMsg asserts that the given is a not nil error and the error message contains subMsg
func ErrSubMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	t.Helper()
	return ErrMsgContains(t, err, subMsg, fmtAndArgs...)
}

// ErrHasMsg asserts that the given is a not nil error and the error message contains subMsg
func ErrHasMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) bool {
	t.Helper()
	return ErrMsgContains(t, err, subMsg, fmtAndArgs...)
}

//
// -------------------- Len --------------------
//

// Len assert given length is equals to wantLn
func Len(t TestingT, give any, wantLn int, fmtAndArgs ...any) bool {
	gln := reflects.Len(reflect.ValueOf(give))
	if gln < 0 {
		t.Helper()
		return fail(t, fmt.Sprintf("type '%T' could not be calc length", give), fmtAndArgs)
	}

	if gln != wantLn {
		t.Helper()
		return fail(t, fmt.Sprintf("\"%v\" should have %d item(s), but has %d", give, wantLn, gln), fmtAndArgs)
	}
	return true
}

// LenGt assert given length is greater than to minLn
func LenGt(t TestingT, give any, minLn int, fmtAndArgs ...any) bool {
	gln := reflects.Len(reflect.ValueOf(give))
	if gln < 0 {
		t.Helper()
		return fail(t, fmt.Sprintf("type '%T' could not be calc length", give), fmtAndArgs)
	}

	if gln <= minLn {
		t.Helper()
		return fail(t, fmt.Sprintf("\"%v\" should have more than %d item(s), but has %d", give, minLn, gln), fmtAndArgs)
	}
	return true
}

//
// -------------------- compare --------------------
//

// Equal asserts that the want should equal to the given.
//
// alias of Eq()
func Equal(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()
	return Eq(t, want, give, fmtAndArgs...)
}

// Eq asserts that the want should equal to the given
func Eq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()

	if err := checkEqualArgs(want, give); err != nil {
		return fail(t,
			fmt.Sprintf("Cannot compare: %#v == %#v (%s)", want, give, err),
			fmtAndArgs,
		)
	}

	// float 特殊处理 - 默认按 0.009 精度对比
	if mathutil.IsFloat(want) {
		if mathutil.InDeltaAny(want, give, 0.009) {
			return true
		}
		return fail(t, fmt.Sprintf("Not equal(inDelta=0.009): \n"+
			"expect: %s\n"+
			"actual: %s", want, give), fmtAndArgs)
	}

	if !reflects.IsEqual(want, give) {
		// TODO diff := diff(want, give)
		want, give = formatUnequalValues(want, give)
		return fail(t, fmt.Sprintf("Not equal: \n"+
			"expect: %s\n"+
			"actual: %s", want, give), fmtAndArgs)
	}

	return true
}

// Neq asserts that the want should not be equal to the given.
//
// alias of NotEq()
func Neq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()
	return NotEq(t, want, give, fmtAndArgs...)
}

// NotEqual asserts that the want should not be equal to the given.
//
// alias of NotEq()
func NotEqual(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()
	return NotEq(t, want, give, fmtAndArgs...)
}

// NotEq asserts that the want should not be equal to the given
func NotEq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()

	if err := checkEqualArgs(want, give); err != nil {
		return fail(t, fmt.Sprintf("Cannot compare: %#v == %#v (%s)", want, give, err), fmtAndArgs)
	}

	if reflects.IsEqual(want, give) {
		return fail(t, fmt.Sprintf("Given should not be: %#v, but give: %+v\n", want, give), fmtAndArgs)
	}
	return true
}

// Lt asserts that the give(intX,uintX,floatX) should not be less than max
func Lt(t TestingT, give, max any, fmtAndArgs ...any) bool {
	if mathutil.Compare(give, max, "lt") {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Given %v should less than %v", give, max), fmtAndArgs)
}

// Lte asserts that the give(intX,uintX,floatX) should not be less than or equals to max
func Lte(t TestingT, give, max any, fmtAndArgs ...any) bool {
	if mathutil.Compare(give, max, "lte") {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Given %v should less than or equal %v", give, max), fmtAndArgs)
}

// Gt asserts that the give(intX,uintX,floatX) should not be greater than min
func Gt(t TestingT, give, min any, fmtAndArgs ...any) bool {
	if mathutil.Compare(give, min, "gt") {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Given %v should greater than %v", give, min), fmtAndArgs)
}

// Gte asserts that the give(intX,uintX,floatX) should not be greater than or equals to min
func Gte(t TestingT, give, min any, fmtAndArgs ...any) bool {
	if mathutil.Compare(give, min, "gte") {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Given %v should greater than or equal %v", give, min), fmtAndArgs)
}

// EqInt asserts that the want should equal to the given intX.
//
// NOTE: Will always convert to int64 to compare.
//
// Example:
// 	assert.Eq(t, uint(1), int(1)) // false
// 	assert.EqInt(t, uint(1), int(1)) // true
func EqInt(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()

	wantI64, err := mathutil.Int64(want)
	if err != nil {
		return fail(t, fmt.Sprintf("Want value cannot convert to int64: %#v, %#v. err=%v", want, give, err), fmtAndArgs)
	}
	giveI64, err1 := mathutil.Int64(give)
	if err1 != nil {
		return fail(t, fmt.Sprintf("Give value cannot convert to int64: %#v, %#v. err=%v", want, give, err1), fmtAndArgs)
	}

	if wantI64 == giveI64 {
		return true
	}
	return fail(t, fmt.Sprintf("Given %v(i64=%d) should equal to %v(i64=%d)", give, giveI64, want, wantI64), fmtAndArgs)
}

// EqFloat asserts that the want should equal to the given with delta. alias of InDelta()
func EqFloat(t TestingT, want, give any, delta float64, fmtAndArgs ...any) bool {
	t.Helper()
	return InDelta(t, want, give, delta, fmtAndArgs...)
}

// InDelta assert that two floating-point values differ from each other within a certain range.
func InDelta(t TestingT, want, give any, delta float64, fmtAndArgs ...any) bool {
	if mathutil.InDeltaAny(want, give, delta) {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Given %v should in delta %v with %v", give, delta, want), fmtAndArgs)
}

// IsType assert data type equals
//
// Usage:
//
//	assert.IsType(t, 0, val) // assert type is int
func IsType(t TestingT, wantType, give any, fmtAndArgs ...any) bool {
	if reflects.IsEqual(reflect.TypeOf(wantType), reflect.TypeOf(give)) {
		return true
	}

	t.Helper()
	return fail(t,
		fmt.Sprintf("Expected to be of type %v, but was %v", reflect.TypeOf(wantType), reflect.TypeOf(give)),
		fmtAndArgs,
	)
}

// IsKind assert data reflect.Kind equals.
// If `give` is ptr or interface, will get real kind.
//
// Usage:
//
//	assert.IsKind(t, reflect.Int, val) // assert type is int kind.
func IsKind(t TestingT, wantKind reflect.Kind, give any, fmtAndArgs ...any) bool {
	giveKind := reflects.Elem(reflect.ValueOf(give)).Kind()
	if wantKind == giveKind {
		return true
	}

	t.Helper()
	return fail(t,
		fmt.Sprintf("Expected to be of kind %v, but was %v", wantKind, giveKind),
		fmtAndArgs,
	)
}

// Same asserts that two pointers reference the same object.
//
//	assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Same(t TestingT, wanted, actual any, fmtAndArgs ...any) bool {
	if samePointers(wanted, actual) {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Not same: \n"+
		"wanted: %p %#v\n"+
		"actual: %p %#v", wanted, wanted, actual, actual), fmtAndArgs)
}

// NotSame asserts that two pointers do not reference the same object.
//
//	assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func NotSame(t TestingT, want, actual any, fmtAndArgs ...any) bool {
	if !samePointers(want, actual) {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Expect and actual is same object: %p %#v", want, want), fmtAndArgs)
}

// samePointers compares two generic interface objects and returns whether
// they point to the same object
func samePointers(first, second any) bool {
	firstPtr, secondPtr := reflect.ValueOf(first), reflect.ValueOf(second)
	if firstPtr.Kind() != reflect.Ptr || secondPtr.Kind() != reflect.Ptr {
		return false
	}

	firstType, secondType := reflect.TypeOf(first), reflect.TypeOf(second)
	if firstType != secondType {
		return false
	}

	// compare pointer addresses
	return first == second
}
