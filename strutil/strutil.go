// Package strutil provide some string,char,byte util functions
package strutil

import (
	"errors"
	"strings"

	"github.com/gookit/goutil/comdef"
)

// OrCond return s1 on cond is True, OR return s2.
// Like: cond ? s1 : s2
func OrCond(cond bool, s1, s2 string) string {
	if cond {
		return s1
	}
	return s2
}

// BlankOr return default value on val is blank, otherwise return val
func BlankOr(val, defVal string) string {
	val = strings.TrimSpace(val)
	if val != "" {
		return val
	}
	return defVal
}

// ZeroOr return default value on val is zero, otherwise return val. same of OrElse()
func ZeroOr[T ~string](val, defVal T) T {
	if val != "" {
		return val
	}
	return defVal
}

// ErrorOr return default value on err is not nil, otherwise return s
// func ErrorOr(s string, err error, defVal string) string {
// 	if err != nil {
// 		return defVal
// 	}
// 	return s
// }

// OrElse return default value on s is empty, otherwise return s
func OrElse(s, orVal string) string {
	if s != "" {
		return s
	}
	return orVal
}

// OrElseNilSafe return default value on s is nil, otherwise return s
func OrElseNilSafe(s *string, orVal string) string {
	if s == nil || *s == "" {
		return orVal
	}
	return *s
}

// OrHandle return fn(s) on s is not empty.
func OrHandle(s string, fn comdef.StringHandleFunc) string {
	if s != "" {
		return fn(s)
	}
	return s
}

// Valid return first not empty element.
func Valid(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}

// SubstrCount returns the number of times the substr substring occurs in the s string.
// Actually, it comes from strings.Count().
//
//   - s The string to search in
//   - substr The substring to search for
//   - params[0] The rune offset where to start counting.
//   - params[1] The maximum number of runes after the specified offset to search.
//
// NOTE: the offset is a rune index, not a byte index. Like Substr(), an offset at or
// after the end of the string gives an empty range, so the count is 0.
func SubstrCount(s, substr string, params ...uint64) (int, error) {
	larg := len(params)
	hasArgs := larg != 0
	if hasArgs && larg > 2 {
		return 0, errors.New("too many parameters")
	}
	if !hasArgs {
		return strings.Count(s, substr), nil
	}

	runes := []rune(s)
	runeCount := len(runes)
	// same boundary convention as Substr(). it also keeps the uint64 -> int
	// conversion below in range, on all platforms.
	if params[0] >= uint64(runeCount) {
		return strings.Count("", substr), nil
	}

	offset := int(params[0])
	end := runeCount
	if larg == 2 && params[1] < uint64(runeCount-offset) {
		end = offset + int(params[1])
	}

	return strings.Count(string(runes[offset:end]), substr), nil
}
