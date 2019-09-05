// Package arrutil provides some util functions for array, slice
package arrutil

// Reverse string slice [site user info 0] -> [0 info user site]
func Reverse(ss []string) {
	ln := len(ss)

	for i := 0; i < int(ln/2); i++ {
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
