package comfunc

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/internal/checkfn"
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
//
//  - only one element, format to string
//  - first is format: fmt.Sprintf(firstElem, fmtAndArgs[1:]...)
//  - all is args: return fmt.Sprint(fmtAndArgs...)
func FormatWithArgs(fmtAndArgs []any) string {
	ln := len(fmtAndArgs)
	if ln == 0 {
		return ""
	}

	first := fmtAndArgs[0]
	if ln == 1 {
		if str, ok := first.(string); ok {
			return str
		}
		return fmt.Sprintf("%+v", first)
	}

	// is template string.
	if tplStr, ok := first.(string); ok && strings.IndexByte(tplStr, '%') >= 0 {
		return fmt.Sprintf(tplStr, fmtAndArgs[1:]...)
	}
	return fmt.Sprint(fmtAndArgs...)
}

// ConvOption convert options
type ConvOption struct {
	// if ture: value is nil, will return convert error;
	// if false(default): value is nil, will convert to zero value
	NilAsFail bool
	// HandlePtr auto convert ptr type(int,float,string) value. eg: *int to int
	// 	- if true: will use real type try convert. default is false
	//	- NOTE: current T type's ptr is default support.
	HandlePtr bool
	// set custom fallback convert func for not supported type.
	UserConvFn comdef.ToStringFunc
}

// ConvOptionFn convert option func
type ConvOptionFn func(opt *ConvOption)

// StrBySprintFn convert any value to string by fmt.Sprint
var StrBySprintFn = func(v any) (string, error) {
	return fmt.Sprint(v), nil
}

// WithUserConvFn set ConvOption.UserConvFn option
func WithUserConvFn(fn comdef.ToStringFunc) ConvOptionFn {
	return func(opt *ConvOption) {
		opt.UserConvFn = fn
	}
}

// NewConvOption create a new ConvOption
func NewConvOption(optFns ...ConvOptionFn) *ConvOption {
	opt := &ConvOption{}
	opt.WithOption(optFns...)
	return opt
}

// WithOption set convert option
func (opt *ConvOption) WithOption(optFns ...ConvOptionFn) {
	for _, fn := range optFns {
		if fn != nil {
			fn(opt)
		}
	}
}

// ToStringWith try to convert value to string. can with some option func, more see ConvOption.
func ToStringWith(in any, optFns ...ConvOptionFn) (str string, err error) {
	switch value := in.(type) {
	case int:
		str = strconv.Itoa(value)
	case int8:
		str = strconv.Itoa(int(value))
	case int16:
		str = strconv.Itoa(int(value))
	case int32: // same as `rune`
		str = strconv.Itoa(int(value))
	case int64:
		str = strconv.FormatInt(value, 10)
	case uint:
		str = strconv.FormatUint(uint64(value), 10)
	case uint8:
		str = strconv.FormatUint(uint64(value), 10)
	case uint16:
		str = strconv.FormatUint(uint64(value), 10)
	case uint32:
		str = strconv.FormatUint(uint64(value), 10)
	case uint64:
		str = strconv.FormatUint(value, 10)
	case float32:
		str = strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		str = strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		str = strconv.FormatBool(value)
	case string:
		str = value
	case *string:
		str = *value
	case []byte:
		str = string(value)
	case time.Duration:
		str = strconv.FormatInt(int64(value), 10)
	case fmt.Stringer:
		str = value.String()
	case error:
		str = value.Error()
	default:
		if len(optFns) == 0 && in == nil {
			return "", nil
		}

		opt := NewConvOption(optFns...)
		if in == nil {
			if opt.NilAsFail {
				err = comdef.ErrConvType
			}
			return
		}

		if opt.HandlePtr {
			if rv := reflect.ValueOf(in); rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
				if checkfn.IsSimpleKind(rv.Kind()) {
					return ToStringWith(rv.Interface(), optFns...)
				}
			}
		}

		if opt.UserConvFn != nil {
			str, err = opt.UserConvFn(in)
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}
