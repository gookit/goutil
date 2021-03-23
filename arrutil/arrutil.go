// Package arrutil provides some util functions for array, slice
package arrutil

import (
	"reflect"
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
