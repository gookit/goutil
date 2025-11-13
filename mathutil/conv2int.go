package mathutil

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/internal/checkfn"
)

/*************************************************************
 * region convert to int
 *************************************************************/

// Int convert value to int
func Int(in any) (int, error) { return ToInt(in) }

// SafeInt convert value to int, will ignore error
func SafeInt(in any) int {
	val, _ := ToInt(in)
	return val
}

// QuietInt convert value to int, will ignore error
func QuietInt(in any) int { return SafeInt(in) }

// IntOrPanic convert value to int, will panic on error
func IntOrPanic(in any) int {
	val, err := ToInt(in)
	if err != nil {
		panic(err)
	}
	return val
}

// MustInt convert value to int, will panic on error
func MustInt(in any) int { return IntOrPanic(in) }

// IntOrDefault convert value to int, return defaultVal on failed
func IntOrDefault(in any, defVal int) int { return IntOr(in, defVal) }

// IntOr convert value to int, return defaultVal on failed
func IntOr(in any, defVal int) int {
	val, err := ToIntWith(in)
	if err != nil {
		return defVal
	}
	return val
}

// IntOrErr convert value to int, return error on failed
func IntOrErr(in any) (int, error) { return ToIntWith(in) }

// ToInt convert value to int, return error on failed
func ToInt(in any) (int, error) { return ToIntWith(in) }

// ToIntWith convert value to int, can with some option func.
//
// Example:
//
//	ToIntWithFunc(val, mathutil.WithNilAsFail, mathutil.WithUserConvFn(func(in any) (int, error) {
//	})
func ToIntWith(in any, optFns ...ConvOptionFn[int]) (iVal int, err error) {
	if len(optFns) == 0 && in == nil {
		return 0, nil
	}

	var opt *ConvOption[int]
	if len(optFns) > 0 {
		opt = NewConvOption(optFns...)
		if in == nil && opt.NilAsFail {
			err = comdef.ErrConvType
			return
		}
	}

	if tVal, ok := in.(string); ok {
		// in strict mode, cannot convert string to int
		if opt != nil && opt.StrictMode {
			err = comdef.ErrConvType
			return
		}

		// try convert to int
		iVal, err = TryStrInt(tVal)
		return
	}

	switch tVal := in.(type) {
	case int:
		iVal = tVal
	case *int: // default support int ptr type
		iVal = *tVal
	case int8:
		iVal = int(tVal)
	case int16:
		iVal = int(tVal)
	case int32:
		iVal = int(tVal)
	case int64:
		if tVal > math.MaxInt32 {
			err = fmt.Errorf("value overflow int32. input: %v", tVal)
		} else {
			iVal = int(tVal)
		}
	case uint:
		if tVal > math.MaxInt32 {
			err = fmt.Errorf("value overflow int32. input: %v", tVal)
		} else {
			iVal = int(tVal)
		}
	case uint8:
		iVal = int(tVal)
	case uint16:
		iVal = int(tVal)
	case uint32:
		if tVal > math.MaxInt32 {
			err = fmt.Errorf("value overflow int32. input: %v", tVal)
		} else {
			iVal = int(tVal)
		}
	case uint64:
		if tVal > math.MaxInt32 {
			err = fmt.Errorf("value overflow int32. input: %v", tVal)
		} else {
			iVal = int(tVal)
		}
	case float32:
		iVal = int(tVal)
	case float64:
		iVal = int(tVal)
	case time.Duration:
		if tVal > math.MaxInt32 {
			err = fmt.Errorf("value overflow int32. input: %v", tVal)
		} else {
			iVal = int(tVal)
		}
	case comdef.Int64able: // eg: json.Number
		var i64 int64
		if i64, err = tVal.Int64(); err == nil {
			if i64 > math.MaxInt32 {
				err = fmt.Errorf("value overflow int32. input: %v", tVal)
			} else {
				iVal = int(i64)
			}
		}
	default:
		if opt == nil {
			err = comdef.ErrConvType
			return
		}

		if opt.UserConvFn != nil {
			iVal, err = opt.UserConvFn(in)
		} else if opt.HandlePtr {
			if rv := reflect.ValueOf(in); rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
				if checkfn.IsSimpleKind(rv.Kind()) {
					return ToIntWith(rv.Interface(), optFns...)
				}
			}
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * region convert to int64
 *************************************************************/

// Int64 convert value to int64, return error on failed
func Int64(in any) (int64, error) { return ToInt64(in) }

// SafeInt64 convert value to int64, will ignore error
func SafeInt64(in any) int64 {
	i64, _ := ToInt64With(in)
	return i64
}

// QuietInt64 convert value to int64, will ignore error
func QuietInt64(in any) int64 { return SafeInt64(in) }

// MustInt64 convert value to int64, will panic on error
func MustInt64(in any) int64 {
	i64, err := ToInt64With(in)
	if err != nil {
		panic(err)
	}
	return i64
}

// Int64OrDefault convert value to int64, return default val on failed
func Int64OrDefault(in any, defVal int64) int64 { return Int64Or(in, defVal) }

// Int64Or convert value to int64, return default val on failed
func Int64Or(in any, defVal int64) int64 {
	i64, err := ToInt64With(in)
	if err != nil {
		return defVal
	}
	return i64
}

// ToInt64 convert value to int64, return error on failed
func ToInt64(in any) (int64, error) { return ToInt64With(in) }

// Int64OrErr convert value to int64, return error on failed
func Int64OrErr(in any) (int64, error) { return ToInt64With(in) }

// ToInt64With try to convert value to int64. can with some option func, more see ConvOption.
func ToInt64With(in any, optFns ...ConvOptionFn[int64]) (i64 int64, err error) {
	if len(optFns) == 0 && in == nil {
		return 0, nil
	}

	var opt *ConvOption[int64]
	if len(optFns) > 0 {
		opt = NewConvOption(optFns...)
		if in == nil && opt.NilAsFail {
			err = comdef.ErrConvType
			return
		}
	}

	if tVal, ok := in.(string); ok {
		// in strict mode, cannot convert string to int
		if opt != nil && opt.StrictMode {
			err = comdef.ErrConvType
			return
		}
		// try convert to int64
		i64, err = TryStrInt64(tVal)
		return
	}

	switch tVal := in.(type) {
	case int:
		i64 = int64(tVal)
	case int8:
		i64 = int64(tVal)
	case int16:
		i64 = int64(tVal)
	case int32:
		i64 = int64(tVal)
	case int64:
		i64 = tVal
	case *int64: // default support int64 ptr type
		i64 = *tVal
	case uint:
		i64 = int64(tVal)
	case uint8:
		i64 = int64(tVal)
	case uint16:
		i64 = int64(tVal)
	case uint32:
		i64 = int64(tVal)
	case uint64:
		i64 = int64(tVal)
	case float32:
		// in strict mode, cannot convert float to int
		if opt != nil && opt.StrictMode {
			err = comdef.ErrConvType
			return
		}
		i64 = int64(tVal)
	case float64:
		// in strict mode, cannot convert float to int
		if opt != nil && opt.StrictMode {
			err = comdef.ErrConvType
			return
		}
		if tVal > math.MaxInt64 {
			err = comdef.ErrConvType
			return
		}
		i64 = int64(tVal)
	case time.Duration:
		i64 = int64(tVal)
	case comdef.Int64able: // eg: json.Number
		i64, err = tVal.Int64()
	default:
		if opt == nil {
			err = comdef.ErrConvType
			return
		}

		if opt.UserConvFn != nil {
			i64, err = opt.UserConvFn(in)
		} else if opt.HandlePtr {
			if rv := reflect.ValueOf(in); rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
				if checkfn.IsSimpleKind(rv.Kind()) {
					return ToInt64With(rv.Interface(), optFns...)
				}
			}
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * region convert to uint
 *************************************************************/

// Uint convert any to uint, return error on failed
func Uint(in any) (uint, error) { return ToUint(in) }

// SafeUint convert any to uint, will ignore error
func SafeUint(in any) uint {
	val, _ := ToUint(in)
	return val
}

// QuietUint convert any to uint, will ignore error
func QuietUint(in any) uint { return SafeUint(in) }

// MustUint convert any to uint, will panic on error
func MustUint(in any) uint {
	val, err := ToUintWith(in)
	if err != nil {
		panic(err)
	}
	return val
}

// UintOrDefault convert any to uint, return default val on failed
func UintOrDefault(in any, defVal uint) uint { return UintOr(in, defVal) }

// UintOr convert any to uint, return default val on failed
func UintOr(in any, defVal uint) uint {
	val, err := ToUintWith(in)
	if err != nil {
		return defVal
	}
	return val
}

// UintOrErr convert value to uint, return error on failed
func UintOrErr(in any) (uint, error) { return ToUintWith(in) }

// ToUint convert value to uint, return error on failed
func ToUint(in any) (u64 uint, err error) { return ToUintWith(in) }

// ToUintWith try to convert value to uint. can with some option func, more see ConvOption.
func ToUintWith(in any, optFns ...ConvOptionFn[uint]) (uVal uint, err error) {
	if len(optFns) == 0 && in == nil {
		return 0, nil
	}

	var opt *ConvOption[uint]
	if len(optFns) > 0 {
		opt = NewConvOption(optFns...)
		if in == nil && opt.NilAsFail {
			err = comdef.ErrConvType
			return
		}
	}

	if tVal, ok := in.(string); ok {
		// in strict mode, cannot convert string to int
		if opt != nil && opt.StrictMode {
			err = comdef.ErrConvType
			return
		}

		// try convert to uint64
		var u64 uint64
		if u64, err = TryStrUint64(tVal); err == nil {
			uVal = uint(u64)
		}
		return
	}

	switch tVal := in.(type) {
	case int:
		uVal = uint(tVal)
	case int8:
		uVal = uint(tVal)
	case int16:
		uVal = uint(tVal)
	case int32:
		uVal = uint(tVal)
	case int64:
		uVal = uint(tVal)
	case uint:
		uVal = tVal
	case *uint: // default support uint ptr type
		uVal = *tVal
	case uint8:
		uVal = uint(tVal)
	case uint16:
		uVal = uint(tVal)
	case uint32:
		uVal = uint(tVal)
	case uint64:
		uVal = uint(tVal)
	case float32:
		uVal = uint(tVal)
	case float64:
		uVal = uint(tVal)
	case time.Duration:
		uVal = uint(tVal)
	case comdef.Int64able: // eg: json.Number
		var i64 int64
		i64, err = tVal.Int64()
		uVal = uint(i64)
	default:
		if opt == nil {
			err = comdef.ErrConvType
			return
		}

		if opt.UserConvFn != nil {
			uVal, err = opt.UserConvFn(in)
		} else if opt.HandlePtr {
			if rv := reflect.ValueOf(in); rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
				if checkfn.IsSimpleKind(rv.Kind()) {
					return ToUintWith(rv.Interface(), optFns...)
				}
			}
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * region convert to uint64
 *************************************************************/

// Uint64 convert any to uint64, return error on failed
func Uint64(in any) (uint64, error) { return ToUint64(in) }

// QuietUint64 convert any to uint64, will ignore error
func QuietUint64(in any) uint64 { return SafeUint64(in) }

// SafeUint64 convert any to uint64, will ignore error
func SafeUint64(in any) uint64 {
	val, _ := ToUint64(in)
	return val
}

// MustUint64 convert any to uint64, will panic on error
func MustUint64(in any) uint64 {
	val, err := ToUint64With(in)
	if err != nil {
		panic(err)
	}
	return val
}

// Uint64OrDefault convert any to uint64, return default val on failed
func Uint64OrDefault(in any, defVal uint64) uint64 { return Uint64Or(in, defVal) }

// Uint64Or convert any to uint64, return default val on failed
func Uint64Or(in any, defVal uint64) uint64 {
	val, err := ToUint64With(in)
	if err != nil {
		return defVal
	}
	return val
}

// Uint64OrErr convert value to uint64, return error on failed
func Uint64OrErr(in any) (uint64, error) { return ToUint64With(in) }

// ToUint64 convert value to uint64, return error on failed
func ToUint64(in any) (uint64, error) { return ToUint64With(in) }

// ToUint64With try to convert value to uint64. can with some option func, more see ConvOption.
func ToUint64With(in any, optFns ...ConvOptionFn[uint64]) (u64 uint64, err error) {
	if len(optFns) == 0 && in == nil {
		return 0, nil
	}

	var opt *ConvOption[uint64]
	if len(optFns) > 0 {
		opt = NewConvOption(optFns...)
		if in == nil && opt.NilAsFail {
			err = comdef.ErrConvType
			return
		}
	}

	if tVal, ok := in.(string); ok {
		// in strict mode, cannot convert string to int
		if opt != nil && opt.StrictMode {
			err = comdef.ErrConvType
			return
		}
		// try convert to uint64
		u64, err = TryStrUint64(tVal)
		return
	}

	switch tVal := in.(type) {
	case int:
		u64 = uint64(tVal)
	case int8:
		u64 = uint64(tVal)
	case int16:
		u64 = uint64(tVal)
	case int32:
		u64 = uint64(tVal)
	case int64:
		u64 = uint64(tVal)
	case uint:
		u64 = uint64(tVal)
	case uint8:
		u64 = uint64(tVal)
	case uint16:
		u64 = uint64(tVal)
	case uint32:
		u64 = uint64(tVal)
	case uint64:
		u64 = tVal
	case *uint64: // default support uint64 ptr type
		u64 = *tVal
	case float32:
		u64 = uint64(tVal)
	case float64:
		u64 = uint64(tVal)
	case time.Duration:
		u64 = uint64(tVal)
	case comdef.Int64able: // eg: json.Number
		var i64 int64
		i64, err = tVal.Int64()
		u64 = uint64(i64)
	default:
		if opt == nil {
			err = comdef.ErrConvType
			return
		}

		if opt.UserConvFn != nil {
			u64, err = opt.UserConvFn(in)
		} else if opt.HandlePtr {
			if rv := reflect.ValueOf(in); rv.Kind() == reflect.Pointer {
				rv = rv.Elem()
				if checkfn.IsSimpleKind(rv.Kind()) {
					return ToUint64With(rv.Interface(), optFns...)
				}
			}
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * region string to intX/uintX
 *************************************************************/

// StrInt convert string to int, ignore error
func StrInt(s string) int {
	iVal, _ := strconv.Atoi(strings.TrimSpace(s))
	return iVal
}

// StrIntOr convert string to int, return default val on failed
func StrIntOr(s string, defVal int) int {
	iVal, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return defVal
	}
	return iVal
}

// TryStrInt convert string to int, return error on failed.
//
//  - empty string will return 0.
//  - allow float string.
func TryStrInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}

	// try convert to int
	iVal, err := strconv.Atoi(s)

	// handle the case where the string might be a float
	if err != nil && checkfn.IsNumeric(s) {
		var floatVal float64
		if floatVal, err = strconv.ParseFloat(s, 64); err == nil {
			iVal = int(math.Round(floatVal))
			err = nil
		}
	}
	return iVal, err
}

// TryStrInt64 convert string to int64, return error on failed.
//
//  - empty string will return 0.
//  - allow float string.
func TryStrInt64(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	i64, err := strconv.ParseInt(s, 10, 0)

	// handle the case where the string might be a float
	if err != nil && checkfn.IsNumeric(s) {
		var floatVal float64
		if floatVal, err = strconv.ParseFloat(s, 64); err == nil {
			i64 = int64(math.Round(floatVal))
			err = nil
		}
	}
	return i64, err
}

// TryStrUint64 try to convert string to uint64, return error on failed
//
//  - empty string will return 0.
//  - allow float string.
func TryStrUint64(s string) (uint64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}

	// try convert to int64
	u64, err := strconv.ParseUint(s, 10, 0)

	// handle the case where the string might be a float
	if err != nil && checkfn.IsPositiveNum(s) {
		var floatVal float64
		if floatVal, err = strconv.ParseFloat(s, 64); err == nil {
			u64 = uint64(math.Round(floatVal))
			err = nil
		}
	}
	return u64, err
}
