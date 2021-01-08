// Package arrutil provides some util functions for array, slice
package arrutil

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/goutil/mathutil"
)

// Reverse string slice [site user info 0] -> [0 info user site]
func Reverse(ss []string) {
	ln := len(ss)

	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		// fmt.Println(i, "<=>", li)
		ss[i], ss[li] = ss[li], ss[i]
	}
}

// StringsRemove an value form an string slice
func StringsRemove(ss []string, s string) []string {
	var ns []string
	for _, v := range ss {
		if v != s {
			ns = append(ns, v)
		}
	}

	return ns
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

// TrimStrings trim string slice item.
func TrimStrings(ss []string, cutSet ...string) (ns []string) {
	hasCutSet := len(cutSet) > 0 && cutSet[0] != ""

	for _, str := range ss {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, cutSet[0]))
		} else {
			ns = append(ns, strings.TrimSpace(str))
		}
	}
	return
}

// IntsHas check the []int contains the given value
func IntsHas(ints []int, val int) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// Int64sHas check the []int64 contains the given value
func Int64sHas(ints []int64, val int64) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// StringsHas check the []string contains the given element
func StringsHas(ss []string, val string) bool {
	for _, ele := range ss {
		if ele == val {
			return true
		}
	}
	return false
}

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
