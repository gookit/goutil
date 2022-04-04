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

// StringsRemove a value form a string slice
func StringsRemove(ss []string, s string) []string {
	ns := make([]string, 0, len(ss))
	for _, v := range ss {
		if v != s {
			ns = append(ns, v)
		}
	}

	return ns
}

// TrimStrings trim string slice item.
//
// Usage:
//	// output: [a, b, c]
// 	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
func TrimStrings(ss []string, cutSet ...string) (ns []string) {
	cutSetLn := len(cutSet)
	hasCutSet := cutSetLn > 0 && cutSet[0] != ""

	var trimSet string
	if hasCutSet {
		trimSet = cutSet[0]
	}
	if cutSetLn > 1 {
		trimSet = strings.Join(cutSet, "")
	}

	for _, str := range ss {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, trimSet))
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
