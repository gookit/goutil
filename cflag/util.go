package cflag

import (
	"flag"
	"reflect"
	"strings"
)

// IsZeroValue determines whether the string represents the zero
// value for a flag.
//
// from flag.isZeroValue() and more return the second arg for check is string.
func IsZeroValue(opt *flag.Flag, value string) (bool, bool) {
	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(opt.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}

	return value == z.Interface().(flag.Value).String(), z.Kind() == reflect.String
}

// AddPrefix for render flag options help
func AddPrefix(name string) string {
	if len(name) > 1 {
		return "--" + name
	}
	return "-" + name
}

// AddPrefixes for render flag options help
func AddPrefixes(name string, shorts []string) string {
	shortLn := len(shorts)
	if shortLn == 0 {
		return AddPrefix(name)
	}

	withPfx := make([]string, 0, shortLn+1)
	withPfx = append(withPfx, AddPrefix(name))
	for _, short := range shorts {
		withPfx = append(withPfx, AddPrefix(short))
	}

	return strings.Join(withPfx, ", ")
}

// SplitShortcut string to []string
func SplitShortcut(shortcut string) (shorts []string) {
	for _, sub := range strings.Split(shortcut, ",") {
		sub = strings.TrimSpace(sub)
		if sub != "" {
			sub = strings.Trim(sub, "- ")
			if sub != "" {
				shorts = append(shorts, sub)
			}
		}
	}

	return
}

// ParseStopMark string
const ParseStopMark = "--"

// ReplaceShorts replace shorts to full option. will stop on ParseStopMark
// eg: '-f' -> '--file'.
func ReplaceShorts(args []string, shortsMap map[string]string) []string {
	if len(args) == 0 {
		return args
	}

	fmtArgs := make([]string, 0, len(args))

	for i, arg := range args {
		if arg == "" || arg[0] != '-' || len(arg) > 48 {
			fmtArgs = append(fmtArgs, arg)
			continue
		}

		if arg == ParseStopMark {
			fmtArgs = append(fmtArgs, args[i:]...)
			break
		}

		var handled bool
		for short, name := range shortsMap {
			// is short name, replace to full opt
			if arg == AddPrefix(short) {
				handled = true
				fmtArgs = append(fmtArgs, AddPrefix(name))
				break
			}
		}

		if !handled {
			fmtArgs = append(fmtArgs, arg)
		}
	}

	return fmtArgs
}
