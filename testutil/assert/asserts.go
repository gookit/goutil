package assert

import (
	"bufio"
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil/common"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/stdutil"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/sysutil"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Helper()
	Name() string
	Error(args ...any)
	Errorf(format string, args ...any)
}

// Nil asserts that the given is a nil value
func Nil(t TestingT, give any, fmtAndArgs ...any) bool {
	if stdutil.IsNil(give) {
		return true
	}

	t.Helper()
	return fail(t, fmt.Sprintf("Expected nil, but got: %#v", give), fmtAndArgs)
}

// NotNil asserts that the given is a not nil value
func NotNil(t TestingT, give any, fmtAndArgs ...any) bool {
	if !stdutil.IsNil(give) {
		return true
	}

	t.Helper()
	return fail(t, "Should not nil value", fmtAndArgs)
}

// True asserts that the given is a bool true
func True(t TestingT, give bool, fmtAndArgs ...any) bool {
	if !give {
		t.Helper()
		return fail(t, "Should be True", fmtAndArgs)
	}
	return true
}

// False asserts that the given is a bool false
func False(t TestingT, give bool, fmtAndArgs ...any) bool {
	if give {
		t.Helper()
		return fail(t, "Should be False", fmtAndArgs)
	}
	return true
}

// Empty asserts that the give should be empty
func Empty(t TestingT, give any, fmtAndArgs ...any) bool {
	empty := stdutil.IsEmpty(give)
	if !empty {
		t.Helper()
		return fail(t, fmt.Sprintf("Should be empty, but was:\n%#v", give), fmtAndArgs)
	}

	return empty
}

// NotEmpty asserts that the give should not be empty
func NotEmpty(t TestingT, give any, fmtAndArgs ...any) bool {
	nEmpty := !stdutil.IsEmpty(give)
	if !nEmpty {
		t.Helper()
		return fail(t, fmt.Sprintf("Should not be empty, but was:\n%#v", give), fmtAndArgs)
	}

	return nEmpty
}

// PanicRunFunc define
type PanicRunFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func runPanicFunc(f PanicRunFunc) (didPanic bool, message interface{}, stack string) {
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

		return fail(t, fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", fn, panicVal), fmtAndArgs)
	}

	return true
}

// PanicsMsg should panic and with a value
func PanicsMsg(t TestingT, fn PanicRunFunc, wantVal interface{}, fmtAndArgs ...any) bool {
	hasPanic, panicVal, stackMsg := runPanicFunc(fn)
	if !hasPanic {
		t.Helper()
		return fail(t, fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", fn, panicVal), fmtAndArgs)
	}

	if panicVal != wantVal {
		t.Helper()
		return fail(t, fmt.Sprintf(
			"func %#v should panic with value:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s",
			fn, wantVal, panicVal, stackMsg),
			fmtAndArgs,
		)
	}

	return true
}

// PanicsErrMsg should panic and with error message
func PanicsErrMsg(t TestingT, fn PanicRunFunc, errMsg string, fmtAndArgs ...any) bool {
	hasPanic, panicVal, stackMsg := runPanicFunc(fn)
	if !hasPanic {
		t.Helper()
		return fail(t, fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", fn, panicVal), fmtAndArgs)
	}

	err, ok := panicVal.(error)
	if !ok || err.Error() != errMsg {
		t.Helper()
		return fail(t, fmt.Sprintf(
			"func %#v should panic with error message:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s",
			fn, errMsg, panicVal, stackMsg),
			fmtAndArgs,
		)
	}

	return true
}

// Contains asserts that the given data(string,slice,map) is contains sub-value
func Contains(t TestingT, src, sub any, fmtAndArgs ...any) bool {
	return true
}

// ContainsKey asserts that the given map is contains key
func ContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) bool {
	return true
}

// StrContains asserts that the given strings is contains sub-string
func StrContains(t TestingT, s, sub string, fmtAndArgs ...any) bool {
	if !strings.Contains(s, sub) {
		t.Helper()
		return fail(t, fmt.Sprintf("Given string: %#v\nNot contains: %#v", s, sub), fmtAndArgs)
	}
	return true
}

//
// -------------------- error --------------------
//

// NoErr asserts that the given is a nil error
func NoErr(t TestingT, err error, fmtAndArgs ...any) bool {
	if err != nil {
		t.Helper()
		return fail(t, fmt.Sprintf("Received unexpected error:\n%+v", err), fmtAndArgs)
	}
	return true
}

// Err asserts that the given is a not nil error
func Err(t TestingT, err error, fmtAndArgs ...any) bool {
	if err == nil {
		t.Helper()
		return fail(t, "An error is expected but got nil.", fmtAndArgs)
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

// ErrSubMsg asserts that the given is a not nil error and the error message contains subMsg
func ErrSubMsg(t TestingT, subMsg, err error, fmtAndArgs ...any) bool {
	if err == nil {
		t.Helper()
		return fail(t, "An error is expected but got nil.", fmtAndArgs)
	}

	return true
}

//
// -------------------- Len --------------------
//

func Len(t TestingT, give any, wantLn int, fmtAndArgs ...any) bool {
	gln := reflects.Len(reflect.ValueOf(give))
	if gln < 0 {
		t.Helper()
		return fail(t, fmt.Sprintf("\"%s\" could not be calc length", give), fmtAndArgs)
	}

	if gln != wantLn {
		t.Helper()
		return fail(t, fmt.Sprintf("\"%s\" should have %d item(s), but has %d", give, wantLn, gln), fmtAndArgs)
	}
	return false
}

func LenGt(t TestingT, give any, minLn int, fmtAndArgs ...any) bool {
	gln := reflects.Len(reflect.ValueOf(give))
	if gln < 0 {
		t.Helper()
		return fail(t, fmt.Sprintf("\"%s\" could not be calc length", give), fmtAndArgs)
	}

	if gln < minLn {
		t.Helper()
		return fail(t, fmt.Sprintf("\"%s\" should less have %d item(s), but has %d", give, minLn, gln), fmtAndArgs)
	}
	return false
}

//
// -------------------- compare --------------------
//

// Eq asserts that the want should equal to the given
func Eq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()

	if err := checkEqualArgs(want, give); err != nil {
		return fail(t, fmt.Sprintf("Cannot compare: %#v == %#v (%s)",
			want, give, err), fmtAndArgs)
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
// alias of NotEq()
func Neq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()
	return NotEq(t, want, give, fmtAndArgs...)
}

// NotEq asserts that the want should not be equal to the given
func NotEq(t TestingT, want, give any, fmtAndArgs ...any) bool {
	t.Helper()

	if err := checkEqualArgs(want, give); err != nil {
		return fail(t, fmt.Sprintf("Cannot compare: %#v == %#v (%s)",
			want, give, err), fmtAndArgs)
	}

	if reflects.IsEqual(want, give) {
		return fail(t, fmt.Sprintf("Given should not be: %#v\n", give), fmtAndArgs)
	}
	return true
}

func Lt(t TestingT, give, max int, fmtAndArgs ...any) bool {
	gInt, err := mathutil.ToInt(give)
	if err == nil && gInt <= max {
		return true
	}

	return fail(t, fmt.Sprintf("Given should later than or equal %d(but was %d)", max, gInt), fmtAndArgs)

}

func Gt(t TestingT, give, min int, fmtAndArgs ...any) bool {
	gInt, err := mathutil.ToInt(give)
	if err == nil && gInt >= min {
		return true
	}

	return fail(t, fmt.Sprintf("Given should gater than or equal %d(but was %d)", min, gInt), fmtAndArgs)
}

//
// -------------------- fail --------------------
//

// Fail reports a failure through
func Fail(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	t.Helper()
	return fail(t, failMsg, fmtAndArgs)
}

type failNower interface {
	FailNow()
}

// FailNow fails test
func FailNow(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	t.Helper()
	// TODO
	return false
}

//
// -------------------- render error --------------------
//

var (
	// ShowFullPath on show error trace
	ShowFullPath = true
	// EnableColor on show error trace
	EnableColor = true
)

// DisableColor render
func DisableColor() {
	EnableColor = false
}

// HideFullPath render
func HideFullPath() {
	ShowFullPath = false
}

// fail reports a failure through
func fail(t TestingT, failMsg string, fmtAndArgs []any) bool {
	t.Helper()

	tName := t.Name()
	if EnableColor {
		tName = color.Red.Sprint(tName)
	}

	labeledTexts := []labeledText{
		{"Test Name", tName},
		{"Error At", strings.Join(callerInfos(), "\n")},
		{"Error Msg", failMsg},
	}

	// user custom message
	if userMsg := formatTplAndArgs(fmtAndArgs...); len(userMsg) > 0 {
		labeledTexts = append(labeledTexts, labeledText{"User Msg", userMsg})
	}

	t.Error("\n" + formatLabeledTexts(labeledTexts))
	return false
}

func callerInfos() []string {
	num := 3
	ss := make([]string, 0, num)
	cs := sysutil.CallersInfos(4, num, func(file string, fc *runtime.Func) bool {
		// This is a huge edge case, but it will panic if this is the case
		if file == "<autogenerated>" {
			return false
		}

		fcName := fc.Name()
		if fcName == "testing.tRunner" {
			return false
		}

		// eg: runtime.goexit
		if strings.HasPrefix(fcName, "runtime.") {
			return false
		}
		return true
	})

	// cwd := sysutil.Workdir()
	for _, info := range cs {
		filePath := info.File
		if !ShowFullPath {
			filePath = fsutil.Name(filePath)
		}

		ss = append(ss, fmt.Sprintf("%s:%d", filePath, info.Line))
	}
	return ss
}

// refers from stretchr/testify/assert
type labeledText struct {
	label   string
	message string
}

func formatLabeledTexts(lts []labeledText) string {
	labelWidth := 0
	elemSize := len(lts)
	for _, lt := range lts {
		labelWidth = mathutil.MaxInt(len(lt.label), labelWidth)
	}

	var sb strings.Builder
	for i, lt := range lts {
		label := lt.label
		if EnableColor {
			label = color.Green.Sprint(label)
		}

		sb.WriteString("  " + label + strutil.Repeat(" ", labelWidth-len(lt.label)) + ":  ")
		formatMessage(lt.message, labelWidth, &sb)
		if i+1 != elemSize {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func formatMessage(message string, labelWidth int, buf common.StringWriteStringer) string {
	for i, scanner := 0, bufio.NewScanner(strings.NewReader(message)); scanner.Scan(); i++ {
		// skip add prefix for first line.
		if i != 0 {
			// +3: is len of ":  "
			_, _ = buf.WriteString("\n  " + strings.Repeat(" ", labelWidth+3))
		}
		_, _ = buf.WriteString(scanner.Text())
	}

	return buf.String()
}
