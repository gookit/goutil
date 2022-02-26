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

// JoinStrings alias of strings.Join
func JoinStrings(ss []string, sep string) string {
	return strings.Join(ss, sep)
}

// StringsJoin alias of strings.Join
func StringsJoin(ss []string, sep string) string {
	return strings.Join(ss, sep)
}

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

// SliceToStrings convert []interface{} to []string
func SliceToStrings(arr []interface{}) []string {
	ss := make([]string, len(arr))
	for i, v := range arr {
		ss[i] = strutil.MustString(v)
	}
	return ss
}

// ToString convert []interface{} to string
func ToString(arr []interface{}) string {
	var b strings.Builder
	for i, v := range arr {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(strutil.MustString(v))
	}

	return b.String()
}

// SliceToString convert []interface{} to string
func SliceToString(arr ...interface{}) string {
	return ToString(arr)
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
