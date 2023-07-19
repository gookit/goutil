package comfunc

import (
	"fmt"
	"strings"

	"github.com/gookit/goutil/comdef"
)

// Bool try to convert type to bool
func Bool(v any) bool {
	bl, _ := ToBool(v)
	return bl
}

// ToBool try to convert type to bool
func ToBool(v any) (bool, error) {
	if bl, ok := v.(bool); ok {
		return bl, nil
	}

	if str, ok := v.(string); ok {
		return StrToBool(str)
	}
	return false, comdef.ErrConvType
}

// StrToBool parse string to bool. like strconv.ParseBool()
func StrToBool(s string) (bool, error) {
	lower := strings.ToLower(s)
	switch lower {
	case "1", "on", "yes", "true":
		return true, nil
	case "0", "off", "no", "false":
		return false, nil
	}

	return false, fmt.Errorf("'%s' cannot convert to bool", s)
}

// FormatWithArgs format message with args
func FormatWithArgs(fmtAndArgs []any) string {
	ln := len(fmtAndArgs)
	if ln == 0 {
		return ""
	}

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
