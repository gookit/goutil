package assert

import (
	"bufio"
	"fmt"
	"reflect"
	"time"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/reflects"
)

func checkEqualArgs(expected, actual any) error {
	if expected == nil && actual == nil {
		return nil
	}

	if reflects.IsFunc(expected) || reflects.IsFunc(actual) {
		return errorx.New("cannot take func type as argument")
	}
	return nil
}

// formatUnequalValues takes two values of arbitrary types and returns string
// representations appropriate to be presented to the user.
//
// If the values are not of like type, the returned strings will be prefixed
// with the type name, and the value will be enclosed in parenthesis similar
// to a type conversion in the Go grammar.
func formatUnequalValues(expected, actual any) (e string, a string) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return fmt.Sprintf("%T(%s)", expected, truncatingFormat(expected)),
			fmt.Sprintf("%T(%s)", actual, truncatingFormat(actual))
	}

	switch expected.(type) {
	case time.Duration:
		return fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)
	}

	return truncatingFormat(expected), truncatingFormat(actual)
}

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(data any) string {
	value := fmt.Sprintf("%T(%v)", data, data)
	// Give us some space the type info too if needed.
	max := bufio.MaxScanTokenSize - 100
	if len(value) > max {
		value = value[0:max] + "<... truncated>"
	}
	return value
}

func formatTplAndArgs(fmtAndArgs ...any) string {
	if len(fmtAndArgs) == 0 || fmtAndArgs == nil {
		return ""
	}

	ln := len(fmtAndArgs)
	first := fmtAndArgs[0]

	if ln == 1 {
		if msgAsStr, ok := first.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", first)
	}

	// is template string.
	if tplStr, ok := first.(string); ok {
		return fmt.Sprintf(tplStr, fmtAndArgs[1:]...)
	}
	return fmt.Sprint(fmtAndArgs...)
}
