package arrutil

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// ErrInvalidType error
var ErrInvalidType = errors.New("the input param type is invalid")

/*************************************************************
 * helper func for strings
 *************************************************************/

// JoinStrings alias of strings.Join
func JoinStrings(sep string, ss ...string) string {
	return strings.Join(ss, sep)
}

// StringsJoin alias of strings.Join
func StringsJoin(sep string, ss ...string) string {
	return strings.Join(ss, sep)
}

// StringsToInts string slice to int slice
func StringsToInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}

// MustToStrings convert interface{}(allow: array,slice) to []string
func MustToStrings(arr interface{}) []string {
	ret, _ := ToStrings(arr)
	return ret
}

// StringsToSlice convert []string to []interface{}
func StringsToSlice(strings []string) []interface{} {
	args := make([]interface{}, len(strings))
	for i, s := range strings {
		args[i] = s
	}
	return args
}

/*************************************************************
 * helper func for slices
 *************************************************************/

// ToInt64s convert interface{}(allow: array,slice) to []int64
func ToInt64s(arr interface{}) (ret []int64, err error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		err = ErrInvalidType
		return
	}

	for i := 0; i < rv.Len(); i++ {
		i64, err := mathutil.Int64(rv.Index(i).Interface())
		if err != nil {
			return []int64{}, err
		}

		ret = append(ret, i64)
	}
	return
}

// MustToInt64s convert interface{}(allow: array,slice) to []int64
func MustToInt64s(arr interface{}) []int64 {
	ret, _ := ToInt64s(arr)
	return ret
}

// SliceToInt64s convert []interface{} to []int64
func SliceToInt64s(arr []interface{}) []int64 {
	i64s := make([]int64, len(arr))
	for i, v := range arr {
		i64s[i] = mathutil.MustInt64(v)
	}
	return i64s
}

// ToStrings convert interface{}(allow: array,slice) to []string
func ToStrings(arr interface{}) (ret []string, err error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		err = ErrInvalidType
		return
	}

	for i := 0; i < rv.Len(); i++ {
		str, err := strutil.ToString(rv.Index(i).Interface())
		if err != nil {
			return []string{}, err
		}

		ret = append(ret, str)
	}
	return
}

// SliceToStrings convert []interface{} to []string
func SliceToStrings(arr []interface{}) []string {
	ss := make([]string, len(arr))
	for i, v := range arr {
		ss[i] = strutil.MustString(v)
	}
	return ss
}

// AnyToString convert any array, slice to string
func AnyToString(arr interface{}) string {
	if arr == nil {
		return "[]"
	}

	rftVal := reflect.ValueOf(arr)
	if rftVal.Type().Kind() == reflect.Ptr {
		rftVal = rftVal.Elem()
	}

	if rftVal.Kind() != reflect.Slice && rftVal.Kind() != reflect.Array {
		return ""
	}

	if rftVal.Len() == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteByte('[')

	for i := 0; i < rftVal.Len(); i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strutil.MustString(rftVal.Index(i).Interface()))
	}

	sb.WriteByte(']')
	return sb.String()
}

// SliceToString convert []interface{} to string
func SliceToString(arr ...interface{}) string { return ToString(arr) }

// ToString convert []interface{} to string
func ToString(arr []interface{}) string {
	// like fmt.Println([]interface{}(nil))
	if arr == nil {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteByte('[')

	for i, v := range arr {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strutil.MustString(v))
	}

	sb.WriteByte(']')
	return sb.String()
}

// JoinSlice join []any slice to string.
func JoinSlice(sep string, arr ...interface{}) string {
	if arr == nil {
		return ""
	}

	var sb strings.Builder
	for i, v := range arr {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(strutil.MustString(v))
	}

	return sb.String()
}
