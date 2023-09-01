package mathutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/goutil/comdef"
)

// ToIntFunc convert value to int
type ToIntFunc func(any) (int, error)

// ToInt64Func convert value to int64
type ToInt64Func func(any) (int64, error)

// ToUintFunc convert value to uint
type ToUintFunc func(any) (uint64, error)

// ToFloatFunc convert value to float
type ToFloatFunc func(any) (float64, error)

/*************************************************************
 * convert value to int
 *************************************************************/

// Int convert value to int
func Int(in any) (int, error) {
	return ToInt(in)
}

// SafeInt convert value to int, will ignore error
func SafeInt(in any) int {
	val, _ := ToInt(in)
	return val
}

// QuietInt convert value to int, will ignore error
func QuietInt(in any) int {
	return SafeInt(in)
}

// MustInt convert value to int, will panic on error
func MustInt(in any) int {
	val, err := ToInt(in)
	if err != nil {
		panic(err)
	}
	return val
}

// IntOrPanic convert value to int, will panic on error
func IntOrPanic(in any) int {
	return MustInt(in)
}

// IntOrDefault convert value to int, return defaultVal on failed
func IntOrDefault(in any, defVal int) int {
	return IntOr(in, defVal)
}

// IntOr convert value to int, return defaultVal on failed
func IntOr(in any, defVal int) int {
	val, err := ToIntWithFunc(in, nil)
	if err != nil {
		return defVal
	}
	return val
}

// IntOrErr convert value to int, return error on failed
func IntOrErr(in any) (iVal int, err error) {
	return ToIntWithFunc(in, nil)
}

// ToInt convert value to int, return error on failed
func ToInt(in any) (iVal int, err error) {
	return ToIntWithFunc(in, nil)
}

// ToIntWithFunc convert value to int, will call usrFn on value type not supported.
func ToIntWithFunc(in any, usrFn ToIntFunc) (iVal int, err error) {
	switch tVal := in.(type) {
	case int:
		iVal = tVal
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
	case string:
		iVal, err = strconv.Atoi(strings.TrimSpace(tVal))
	case interface{ Int64() (int64, error) }: // eg: json.Number
		var i64 int64
		if i64, err = tVal.Int64(); err == nil {
			if i64 > math.MaxInt32 {
				err = fmt.Errorf("value overflow int32. input: %v", tVal)
			} else {
				iVal = int(i64)
			}
		}
	default:
		if usrFn != nil {
			return usrFn(in)
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

// StrInt convert.
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

/*************************************************************
 * convert value to uint
 *************************************************************/

// Uint convert any to uint, return error on failed
func Uint(in any) (uint64, error) {
	return ToUint(in)
}

// SafeUint convert any to uint, will ignore error
func SafeUint(in any) uint64 {
	val, _ := ToUint(in)
	return val
}

// QuietUint convert any to uint, will ignore error
func QuietUint(in any) uint64 {
	return SafeUint(in)
}

// MustUint convert any to uint, will panic on error
func MustUint(in any) uint64 {
	val, err := ToUintWithFunc(in, nil)
	if err != nil {
		panic(err)
	}
	return val
}

// UintOrDefault convert any to uint, return default val on failed
func UintOrDefault(in any, defVal uint64) uint64 {
	return UintOr(in, defVal)
}

// UintOr convert any to uint, return default val on failed
func UintOr(in any, defVal uint64) uint64 {
	val, err := ToUintWithFunc(in, nil)
	if err != nil {
		return defVal
	}
	return val
}

// UintOrErr convert value to uint, return error on failed
func UintOrErr(in any) (uint64, error) {
	return ToUintWithFunc(in, nil)
}

// ToUint convert value to uint, return error on failed
func ToUint(in any) (u64 uint64, err error) {
	return ToUintWithFunc(in, nil)
}

// ToUintWithFunc convert value to uint, will call usrFn on value type not supported.
func ToUintWithFunc(in any, usrFn ToUintFunc) (u64 uint64, err error) {
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
	case float32:
		u64 = uint64(tVal)
	case float64:
		u64 = uint64(tVal)
	case time.Duration:
		u64 = uint64(tVal)
	case interface{ Int64() (int64, error) }: // eg: json.Number
		var i64 int64
		i64, err = tVal.Int64()
		u64 = uint64(i64)
	case string:
		u64, err = strconv.ParseUint(strings.TrimSpace(tVal), 10, 0)
	default:
		if usrFn != nil {
			u64, err = usrFn(in)
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * convert value to int64
 *************************************************************/

// Int64 convert value to int64, return error on failed
func Int64(in any) (int64, error) {
	return ToInt64(in)
}

// SafeInt64 convert value to int64, will ignore error
func SafeInt64(in any) int64 {
	i64, _ := ToInt64WithFunc(in, nil)
	return i64
}

// QuietInt64 convert value to int64, will ignore error
func QuietInt64(in any) int64 {
	return SafeInt64(in)
}

// MustInt64 convert value to int64, will panic on error
func MustInt64(in any) int64 {
	i64, err := ToInt64WithFunc(in, nil)
	if err != nil {
		panic(err)
	}
	return i64
}

// Int64OrDefault convert value to int64, return default val on failed
func Int64OrDefault(in any, defVal int64) int64 {
	return Int64Or(in, defVal)
}

// Int64Or convert value to int64, return default val on failed
func Int64Or(in any, defVal int64) int64 {
	i64, err := ToInt64WithFunc(in, nil)
	if err != nil {
		return defVal
	}
	return i64
}

// Int64OrErr convert value to int64, return error on failed
func Int64OrErr(in any) (int64, error) {
	return ToInt64(in)
}

// ToInt64 convert value to int64, return error on failed
func ToInt64(in any) (i64 int64, err error) {
	return ToInt64WithFunc(in, nil)
}

// ToInt64WithFunc convert value to int64, will call usrFn on value type not supported.
func ToInt64WithFunc(in any, usrFn ToInt64Func) (i64 int64, err error) {
	switch tVal := in.(type) {
	case string:
		i64, err = strconv.ParseInt(strings.TrimSpace(tVal), 10, 0)
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
		i64 = int64(tVal)
	case float64:
		i64 = int64(tVal)
	case time.Duration:
		i64 = int64(tVal)
	case interface{ Int64() (int64, error) }: // eg: json.Number
		i64, err = tVal.Int64()
	default:
		if usrFn != nil {
			i64, err = usrFn(in)
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * convert value to float
 *************************************************************/

// QuietFloat convert value to float64, will ignore error. alias of SafeFloat
func QuietFloat(in any) float64 {
	return SafeFloat(in)
}

// SafeFloat convert value to float64, will ignore error
func SafeFloat(in any) float64 {
	val, _ := ToFloatWithFunc(in, nil)
	return val
}

// FloatOrPanic convert value to float64, will panic on error
func FloatOrPanic(in any) float64 {
	return MustFloat(in)
}

// MustFloat convert value to float64, will panic on error
func MustFloat(in any) float64 {
	val, err := ToFloatWithFunc(in, nil)
	if err != nil {
		panic(err)
	}
	return val
}

// FloatOrDefault convert value to float64, will return default value on error
func FloatOrDefault(in any, defVal float64) float64 {
	return FloatOr(in, defVal)
}

// FloatOr convert value to float64, will return default value on error
func FloatOr(in any, defVal float64) float64 {
	val, err := ToFloatWithFunc(in, nil)
	if err != nil {
		return defVal
	}
	return val
}

// Float convert value to float64, return error on failed
func Float(in any) (float64, error) {
	return ToFloatWithFunc(in, nil)
}

// FloatOrErr convert value to float64, return error on failed
func FloatOrErr(in any) (float64, error) {
	return ToFloatWithFunc(in, nil)
}

// ToFloat convert value to float64, return error on failed
func ToFloat(in any) (f64 float64, err error) {
	return ToFloatWithFunc(in, nil)
}

// ToFloatWithFunc convert value to float64, will call usrFn if value type not supported.
func ToFloatWithFunc(in any, usrFn ToFloatFunc) (f64 float64, err error) {
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
	case time.Duration:
		f64 = float64(tVal)
	case interface{ Float64() (float64, error) }: // eg: json.Number
		f64, err = tVal.Float64()
	default:
		if usrFn != nil {
			f64, err = usrFn(in)
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

/*************************************************************
 * convert intX/floatX to string
 *************************************************************/

// MustString convert intX/floatX value to string, will panic on error
func MustString(val any) string {
	str, err := ToStringWithFunc(val, nil)
	if err != nil {
		panic(err)
	}
	return str
}

// StringOrPanic convert intX/floatX value to string, will panic on error
func StringOrPanic(val any) string { return MustString(val) }

// StringOrDefault convert intX/floatX value to string, will return default value on error
func StringOrDefault(val any, defVal string) string {
	return StringOr(val, defVal)
}

// StringOr convert intX/floatX value to string, will return default value on error
func StringOr(val any, defVal string) string {
	str, err := ToStringWithFunc(val, nil)
	if err != nil {
		return defVal
	}
	return str
}

// ToString convert intX/floatX value to string, return error on failed
func ToString(val any) (string, error) {
	return ToStringWithFunc(val, nil)
}

// StringOrErr convert intX/floatX value to string, return error on failed
func StringOrErr(val any) (string, error) {
	return ToStringWithFunc(val, nil)
}

// QuietString convert intX/floatX value to string, other type convert by fmt.Sprint
func QuietString(val any) string {
	return SafeString(val)
}

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
func TryToString(val any, defaultAsErr bool) (str string, err error) {
	var usrFn comdef.ToStringFunc
	if !defaultAsErr {
		usrFn = func(v any) (string, error) {
			if val == nil {
				return "", nil
			}
			return fmt.Sprint(v), nil
		}
	}

	return ToStringWithFunc(val, usrFn)
}

// ToStringWithFunc try convert intX/floatX value to string, will call usrFn if value type not supported.
//
// if defaultAsErr is False, will use fmt.Sprint convert other type
func ToStringWithFunc(val any, usrFn comdef.ToStringFunc) (str string, err error) {
	switch value := val.(type) {
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
	case time.Duration:
		str = strconv.FormatInt(int64(value), 10)
	case string:
		str = value
	case fmt.Stringer:
		str = value.String()
	default:
		if usrFn != nil {
			str, err = usrFn(val)
		} else {
			err = comdef.ErrConvType
		}
	}
	return
}

// Percent returns a values percent of the total
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}
	return (float64(val) / float64(total)) * 100
}

// ElapsedTime calc elapsed time 计算运行时间消耗 单位 ms(毫秒)
//
// Deprecated: use timex.ElapsedTime()
func ElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3f", time.Since(startTime).Seconds()*1000)
}
