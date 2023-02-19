package cflag

import (
	"flag"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/gookit/color"
)

const (
	// RegGoodName match a good option, argument name
	RegGoodName = `^[a-zA-Z][\w-]*$`
)

var (
	// GoodName good name for option and argument
	goodName = regexp.MustCompile(RegGoodName)
)

// IsGoodName check
func IsGoodName(name string) bool {
	return goodName.MatchString(name)
}

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

// AddPrefixes for render flag options help, name will first add.
func AddPrefixes(name string, shorts []string) string {
	return AddPrefixes2(name, shorts, false)
}

// AddPrefixes2 for render flag options help, can custom name add position.
func AddPrefixes2(name string, shorts []string, nameAtEnd bool) string {
	shortLn := len(shorts)
	if shortLn == 0 {
		return AddPrefix(name)
	}

	sort.Strings(shorts)
	withPfx := make([]string, 0, shortLn+1)
	if !nameAtEnd {
		withPfx = append(withPfx, AddPrefix(name))
	}

	// append shorts
	for _, short := range shorts {
		withPfx = append(withPfx, AddPrefix(short))
	}

	if nameAtEnd {
		withPfx = append(withPfx, AddPrefix(name))
	}
	return strings.Join(withPfx, ", ")
}

// SplitShortcut string to []string
func SplitShortcut(shortcut string) []string {
	return FilterNames(strings.Split(shortcut, ","))
}

// FilterNames for option names, will clear there are: "-+= "
func FilterNames(names []string) []string {
	filtered := make([]string, 0, len(names))
	for _, sub := range names {
		if sub = strings.TrimSpace(sub); sub != "" {
			sub = strings.Trim(sub, "-+= ")
			if sub != "" {
				filtered = append(filtered, sub)
			}
		}
	}
	return filtered
}

// IsFlagHelpErr check
func IsFlagHelpErr(err error) bool {
	if err == nil {
		return false
	}
	return err == flag.ErrHelp
}

// regex: "`[\w ]+`"
// regex: "`.+`"
var codeReg = regexp.MustCompile("`" + `.+` + "`")

// WrapColorForCode convert "hello `keywords`" to "hello <mga>keywords</>"
func WrapColorForCode(s string) string {
	if !strings.ContainsRune(s, '`') {
		return s
	}

	return codeReg.ReplaceAllStringFunc(s, func(code string) string {
		code = strings.Trim(code, "`")
		return color.WrapTag(code, "mga")
	})
}

// ParseStopMark string
const ParseStopMark = "--"

// ReplaceShorts replace shorts to full option. will stop on ParseStopMark
//
// For example:
//
//	eg: '-f' -> '--file'.
//	eg: '-n=tom' -> '--name=tom'.
func ReplaceShorts(args []string, shortsMap map[string]string) []string {
	if len(args) == 0 {
		return args
	}

	fmtArgs := make([]string, 0, len(args))

	for i, arg := range args {
		if arg == "" || arg[0] != '-' || len(arg) > 64 {
			fmtArgs = append(fmtArgs, arg)
			continue
		}

		if arg == ParseStopMark {
			fmtArgs = append(fmtArgs, args[i:]...)
			break
		}

		var handled bool
		for short, name := range shortsMap {
			sOpt := AddPrefix(short)
			// is short name, replace to full opt. eg: '-f' -> '--file'
			if arg == sOpt {
				handled = true
				fmtArgs = append(fmtArgs, AddPrefix(name))
				break
			}

			// special, use '=' split value. eg: '-n=tom' -> '--name=tom'
			if strings.HasPrefix(arg, sOpt+"=") {
				handled = true
				fullOpt := AddPrefix(name)
				fmtArgs = append(fmtArgs, fullOpt+arg[len(sOpt):])
				break
			}
		}

		if !handled {
			fmtArgs = append(fmtArgs, arg)
		}
	}

	return fmtArgs
}
