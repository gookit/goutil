package mathutil

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errConvertFail = errors.New("convert data type is failure")
)

/*************************************************************
 * convert value to int
 *************************************************************/

// Int convert string to int
func Int(in interface{}) (int, error) {
	return ToInt(in)
}

// MustInt convert string to int
func MustInt(in interface{}) int {
	val, _ := ToInt(in)
	return val
}

// ToInt convert string to int
func ToInt(in interface{}) (iVal int, err error) {
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
		iVal = int(tVal)
	case uint:
		iVal = int(tVal)
	case uint8:
		iVal = int(tVal)
	case uint16:
		iVal = int(tVal)
	case uint32:
		iVal = int(tVal)
	case uint64:
		iVal = int(tVal)
	case float32:
		iVal = int(tVal)
	case float64:
		iVal = int(tVal)
	case string:
		iVal, err = strconv.Atoi(strings.TrimSpace(tVal))
	case nil:
		iVal = 0
	default:
		err = errConvertFail
	}
	return
}

/*************************************************************
 * convert value to uint
 *************************************************************/

// Uint convert string to uint
func Uint(in interface{}) (uint64, error) {
	return ToUint(in)
}

// MustUint convert string to uint
func MustUint(in interface{}) uint64 {
	val, _ := ToUint(in)
	return val
}

// ToUint convert string to uint
func ToUint(in interface{}) (u64 uint64, err error) {
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
	case string:
		u64, err = strconv.ParseUint(strings.TrimSpace(tVal), 10, 0)
	case nil:
		u64 = 0
	default:
		err = errConvertFail
	}
	return
}

/*************************************************************
 * convert value to int64
 *************************************************************/

// Int64 convert string to int64
func Int64(in interface{}) (int64, error) {
	return ToInt64(in)
}

// MustInt64 convert
func MustInt64(in interface{}) int64 {
	i64, _ := ToInt64(in)
	return i64
}

// ToInt64 convert string to int64
func ToInt64(in interface{}) (i64 int64, err error) {
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
	case nil:
		i64 = 0
	default:
		err = errConvertFail
	}
	return
}

/*************************************************************
 * convert value to float
 *************************************************************/

// Float convert string to float
func Float(s string) (float64, error) {
	return ToFloat(s)
}

// ToFloat convert string to float
func ToFloat(s string) (float64, error) {
	return strconv.ParseFloat(strings.TrimSpace(s), 0)
}

// MustFloat convert string to float
func MustFloat(s string) float64 {
	val, _ := strconv.ParseFloat(strings.TrimSpace(s), 0)
	return val
}
