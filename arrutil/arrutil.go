// Package arrutil provides some util functions for array, slice
package arrutil

import (
	"reflect"
	"strings"

	"github.com/gookit/goutil/mathutil"
)

// Contains assert array(strings, intXs, uintXs) should be contains the given value(int(X),string).
func Contains(arr, val interface{}) bool {
	if val == nil || arr == nil {
		return false
	}

	// if is string value
	if strVal, ok := val.(string); ok {
		if ss, ok := arr.([]string); ok {
			return StringsHas(ss, strVal)
		}

		rv := reflect.ValueOf(arr)
		if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
			for i := 0; i < rv.Len(); i++ {
				if v, ok := rv.Index(i).Interface().(string); ok && strings.EqualFold(v, strVal) {
					return true
				}
			}
		}

		return false
	}

	// as int value
	intVal, err := mathutil.Int64(val)
	if err != nil {
		return false
	}

	if int64s, ok := toInt64Slice(arr); ok {
		return Int64sHas(int64s, intVal)
	}
	return false
}

// NotContains array(strings, ints, uints) should be not contains the given value.
func NotContains(arr, val interface{}) bool {
	return false == Contains(arr, val)
}

// GetRandomOne get random element from an array/slice
func GetRandomOne(arr interface{}) interface{} {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return arr
	}

	i := mathutil.RandomInt(0, rv.Len())
	r := rv.Index(i).Interface()

	return r
}

func toInt64Slice(arr interface{}) (ret []int64, ok bool) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return
	}

	for i := 0; i < rv.Len(); i++ {
		i64, err := mathutil.Int64(rv.Index(i).Interface())
		if err != nil {
			return []int64{}, false
		}

		ret = append(ret, i64)
	}

	ok = true
	return
}
