package mathutil

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/internal/checkfn"
	"github.com/gookit/goutil/internal/comfunc"
)

// ToIntFunc convert value to int
type ToIntFunc func(any) (int, error)

// ToInt64Func convert value to int64
type ToInt64Func func(any) (int64, error)

// ToUintFunc convert value to uint
type ToUintFunc func(any) (uint, error)

// ToUint64Func convert value to uint
type ToUint64Func func(any) (uint64, error)

// ToFloatFunc convert value to float
type ToFloatFunc func(any) (float64, error)

// ToTypeFunc convert value to defined type
type ToTypeFunc[T any] func(any) (T, error)

// ConvOption convert options
type ConvOption[T any] struct {
	// if ture: value is nil, will return convert error;
	// if false(default): value is nil, will convert to zero value
	NilAsFail bool
	// HandlePtr auto convert ptr type(int,float,string) value. eg: *int to int
	// 	- if true: will use real type try convert. default is false
	//	- NOTE: current T type's ptr is default support.
	HandlePtr bool
	// StrictMode for convert value. default is false
	//
	// TRUE:
	//  - to int: string, float will return error
	StrictMode bool
	// set custom fallback convert func for not supported type.
	UserConvFn ToTypeFunc[T]
}

// NewConvOption create a new ConvOption
func NewConvOption[T any](optFns ...ConvOptionFn[T]) *ConvOption[T] {
	opt := &ConvOption[T]{}
	opt.WithOption(optFns...)
	return opt
}

// WithOption set convert option
func (opt *ConvOption[T]) WithOption(optFns ...ConvOptionFn[T]) {
	for _, fn := range optFns {
		if fn != nil {
			fn(opt)
		}
	}
}

// ConvOptionFn convert option func
type ConvOptionFn[T any] func(opt *ConvOption[T])

// WithNilAsFail set ConvOption.NilAsFail option
//
// Example:
//
//	ToIntWithFunc(val, mathutil.WithNilAsFail[int])
func WithNilAsFail[T any](opt *ConvOption[T]) {
	opt.NilAsFail = true
}

// WithHandlePtr set ConvOption.HandlePtr option
func WithHandlePtr[T any](opt *ConvOption[T]) {
	opt.HandlePtr = true
}

// WithUserConvFn set ConvOption.UserConvFn option
func WithUserConvFn[T any](fn ToTypeFunc[T]) ConvOptionFn[T] {
	return func(opt *ConvOption[T]) {
		opt.UserConvFn = fn
	}
}

/*************************************************************
 * region Strict to int/uint
 *************************************************************/

// StrictInt check the given value is an integer(intX,uintX), return the int64 value and true if success
func StrictInt(val any) (int64, bool) {
	switch tVal := val.(type) {
	case int:
		return int64(tVal), true
	case int8:
		return int64(tVal), true
	case int16:
		return int64(tVal), true
	case int32:
		return int64(tVal), true
	case int64:
		return tVal, true
	case uint:
		return int64(tVal), true
	case uint8:
		return int64(tVal), true
	case uint16:
		return int64(tVal), true
	case uint32:
		return int64(tVal), true
	case uint64:
		return int64(tVal), true
	case uintptr:
		return int64(tVal), true
	default:
		return 0, false
	}
}

// StrictUint strict check value is integer(intX,uintX) and convert to uint64.
func StrictUint(val any) (uint64, bool) {
	switch tVal := val.(type) {
	case int:
		return uint64(tVal), true
	case int8:
		return uint64(tVal), true
	case int16:
		return uint64(tVal), true
	case int32:
		return uint64(tVal), true
	case int64:
		return uint64(tVal), true
	case uint:
		return uint64(tVal), true
	case uint8:
		return uint64(tVal), true
	case uint16:
		return uint64(tVal), true
	case uint32:
		return uint64(tVal), true
	case uint64:
		return tVal, true
	case uintptr:
		return uint64(tVal), true
	default:
		return 0, false
	}
}

/*************************************************************
 * region convert to float64
 *************************************************************/

// QuietFloat convert value to float64, will ignore error. alias of SafeFloat
func QuietFloat(in any) float64 { return SafeFloat(in) }

// SafeFloat convert value to float64, will ignore error
func SafeFloat(in any) float64 {
	val, _ := ToFloatWith(in)
	return val
}

// FloatOrPanic convert value to float64, will panic on error
func FloatOrPanic(in any) float64 { return MustFloat(in) }

// MustFloat convert value to float64, will panic on error
func MustFloat(in any) float64 {
	val, err := ToFloatWith(in)
	if err != nil {
		panic(err)
	}
	return val
}

// FloatOrDefault convert value to float64, will return default value on error
func FloatOrDefault(in any, defVal float64) float64 { return FloatOr(in, defVal) }

// FloatOr convert value to float64, will return default value on error
func FloatOr(in any, defVal float64) float64 {
	val, err := ToFloatWith(in)
	if err != nil {
		return defVal
	}
	return val
}

// Float convert value to float64, return error on failed
func Float(in any) (float64, error) { return ToFloatWith(in) }

// FloatOrErr convert value to float64, return error on failed
func FloatOrErr(in any) (float64, error) { return ToFloatWith(in) }

// ToFloat convert value to float64, return error on failed
func ToFloat(in any) (float64, error) { return ToFloatWith(in) }

// ToFloatWith try to convert value to float64. can with some option func, more see ConvOption.
func ToFloatWith(in any, optFns ...ConvOptionFn[float64]) (f64 float64, err error) {
	opt := NewConvOption(optFns...)
	if !opt.NilAsFail && in == nil {
		return 0, nil
	}

	switch tVal := in.(type) {
	case string:
		f64, err = strconv.ParseFloat(strings.TrimSpace(tVal), 64)
	case int:
		f64 = float64(tVal)
	case int8:
		f64 = float64(tVal)
	case int16:
		f64 = float64(tVal)
	case int32:
		f64 = float64(tVal)
	case int64:
		f64 = float64(tVal)
	case uint:
		f64 = float64(tVal)
	case uint8:
		f64 = float64(tVal)
	case uint16:
		f64 = float64(tVal)
	case uint32:
		f64 = float64(tVal)
	case uint64:
		f64 = float64(tVal)
	case float32:
		f64 = float64(tVal)
	case float64:
		f64 = tVal
	case *float64: // default support float64 ptr type
		f64 = *tVal
	case time.Duration:
		f64 = float64(tVal)
	case comdef.Float64able: // eg: json.Number
		f64, err = tVal.Float64()
	default:
		if opt.HandlePtr {
			if rv := reflect.ValueOf(in); rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
				if checkfn.IsSimpleKind(rv.Kind()) {
					return ToFloatWith(rv.Interface(), optFns...)
				}
			}
		}

		if opt.UserConvFn != nil {
			f64, err = opt.UserConvFn(in)
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * region intX/floatX to string
 *************************************************************/

// MustString convert intX/floatX value to string, will panic on error
func MustString(val any) string {
	str, err := ToStringWith(val)
	if err != nil {
		panic(err)
	}
	return str
}

// StringOrPanic convert intX/floatX value to string, will panic on error
func StringOrPanic(val any) string { return MustString(val) }

// StringOrDefault convert intX/floatX value to string, will return default value on error
func StringOrDefault(val any, defVal string) string { return StringOr(val, defVal) }

// StringOr convert intX/floatX value to string, will return default value on error
func StringOr(val any, defVal string) string {
	str, err := ToStringWith(val)
	if err != nil {
		return defVal
	}
	return str
}

// ToString convert intX/floatX value to string, return error on failed
func ToString(val any) (string, error) { return ToStringWith(val) }

// StringOrErr convert intX/floatX value to string, return error on failed
func StringOrErr(val any) (string, error) { return ToStringWith(val) }

// QuietString convert intX/floatX value to string, other type convert by fmt.Sprint
func QuietString(val any) string { return SafeString(val) }

// String convert intX/floatX value to string, other type convert by fmt.Sprint
func String(val any) string {
	str, _ := TryToString(val, false)
	return str
}

// SafeString convert intX/floatX value to string, other type convert by fmt.Sprint
func SafeString(val any) string {
	str, _ := TryToString(val, false)
	return str
}

// TryToString try convert intX/floatX value to string
//
// if defaultAsErr is False, will use fmt.Sprint convert other type
func TryToString(val any, defaultAsErr bool) (string, error) {
	var optFn comfunc.ConvOptionFn
	if !defaultAsErr {
		optFn = comfunc.WithUserConvFn(comfunc.StrBySprintFn)
	}
	return ToStringWith(val, optFn)
}

// ToStringWith try to convert value to string. can with some option func, more see comfunc.ConvOption.
func ToStringWith(in any, optFns ...comfunc.ConvOptionFn) (string, error) {
	return comfunc.ToStringWith(in, optFns...)
}
