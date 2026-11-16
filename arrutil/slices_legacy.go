//go:build !go1.21

package arrutil

import (
	"sort"

	"github.com/gookit/goutil/comdef"
)

// NOTE: this file has the same functions as slices_go121.go, written by hand for
// the Go versions that do not have the slices package (go1.21+). Keep both in sync.

// SliceHas check the slice contains the given value
func SliceHas[T comdef.ScalarType](slice []T, val T) bool {
	for _, ele := range slice {
		if ele == val {
			return true
		}
	}
	return false
}

// IntsHas check the []comdef.Integer contains the given value
func IntsHas[T comdef.Integer](ints []T, val T) bool {
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
func StringsHas[T ~string](ss []T, val T) bool {
	for _, ele := range ss {
		if ele == val {
			return true
		}
	}
	return false
}

// In check the given value whether in the list
func In[T comdef.ScalarType](value T, list []T) bool {
	for _, elem := range list {
		if elem == value {
			return true
		}
	}
	return false
}

// IndexOf value in given slice. returns -1 if not found
func IndexOf[T comdef.NumberOrString](val T, list []T) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}

// Reverse any T slice in place.
//
// eg: []string{"site", "user", "info", "0"} -> []string{"0", "info", "user", "site"}
func Reverse[T any](ls []T) {
	ln := len(ls)
	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		ls[i], ls[li] = ls[li], ls[i]
	}
}

// CloneSlice Clone a slice.
//
// NOTE: unlike slices.Clone a nil slice gives back an empty non-nil slice.
func CloneSlice[T any](data []T) []T {
	nt := make([]T, 0, len(data))
	nt = append(nt, data...)
	return nt
}

// Sort the given slice in place. It panics if the element type is not ordered.
func Sort[T comdef.Compared](ls []T) {
	sort.Slice(ls, func(i, j int) bool {
		return ls[i] < ls[j]
	})
}

// SortFunc sorts the slice in place by the given comparison function.
func SortFunc[T any](ls []T, cmp Comparer[T]) {
	sort.Slice(ls, func(i, j int) bool {
		return cmp(ls[i], ls[j]) < 0
	})
}

// SortStableFunc sorts the slice in place by the given comparison function,
// keeping the original order of equal elements.
func SortStableFunc[T any](ls []T, cmp Comparer[T]) {
	sort.SliceStable(ls, func(i, j int) bool {
		return cmp(ls[i], ls[j]) < 0
	})
}

// IsSorted reports whether the given slice is sorted in increasing order.
func IsSorted[T comdef.Compared](ls []T) bool {
	for i := 1; i < len(ls); i++ {
		if ls[i] < ls[i-1] {
			return false
		}
	}
	return true
}

// IsSortedFunc reports whether the given slice is sorted by the given comparison function.
func IsSortedFunc[T any](ls []T, cmp Comparer[T]) bool {
	for i := 1; i < len(ls); i++ {
		if cmp(ls[i], ls[i-1]) < 0 {
			return false
		}
	}
	return true
}

// ContainsFunc reports whether any element of the slice satisfies the given function.
func ContainsFunc[T any](ls []T, fn comdef.MatchFunc[T]) bool {
	return IndexFunc(ls, fn) >= 0
}

// IndexFunc returns the index of the first element that satisfies the given function,
// or -1 if there is none.
func IndexFunc[T any](ls []T, fn comdef.MatchFunc[T]) int {
	for i, v := range ls {
		if fn(v) {
			return i
		}
	}
	return -1
}

// Compact replaces consecutive runs of equal elements with a single copy, in place.
//
// The slice is modified and the elements after the returned length are zeroed.
func Compact[T comparable](ls []T) []T {
	return CompactFunc(ls, func(a, b T) bool {
		return a == b
	})
}

// CompactFunc is like Compact, but uses eq to compare the elements.
func CompactFunc[T any](ls []T, eq func(a, b T) bool) []T {
	if len(ls) < 2 {
		return ls
	}

	i := 1
	for k := 1; k < len(ls); k++ {
		if !eq(ls[k], ls[k-1]) {
			if i != k {
				ls[i] = ls[k]
			}
			i++
		}
	}

	// zero the elements between the new length and the old length, like slices.Compact
	var zero T
	for k := i; k < len(ls); k++ {
		ls[k] = zero
	}
	return ls[:i]
}

// Max returns the maximum value of the slice. It panics on an empty slice.
func Max[T comdef.Compared](ls []T) T {
	if len(ls) == 0 {
		panic("arrutil.Max: empty list")
	}

	maxVal := ls[0]
	for _, v := range ls[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// Min returns the minimum value of the slice. It panics on an empty slice.
func Min[T comdef.Compared](ls []T) T {
	if len(ls) == 0 {
		panic("arrutil.Min: empty list")
	}

	minVal := ls[0]
	for _, v := range ls[1:] {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// MaxFunc returns the maximum value of the slice by the given comparison function.
// It panics on an empty slice.
func MaxFunc[T any](ls []T, cmp Comparer[T]) T {
	if len(ls) == 0 {
		panic("arrutil.MaxFunc: empty list")
	}

	maxVal := ls[0]
	for _, v := range ls[1:] {
		if cmp(v, maxVal) > 0 {
			maxVal = v
		}
	}
	return maxVal
}

// MinFunc returns the minimum value of the slice by the given comparison function.
// It panics on an empty slice.
func MinFunc[T any](ls []T, cmp Comparer[T]) T {
	if len(ls) == 0 {
		panic("arrutil.MinFunc: empty list")
	}

	minVal := ls[0]
	for _, v := range ls[1:] {
		if cmp(v, minVal) < 0 {
			minVal = v
		}
	}
	return minVal
}

// Equal reports whether the two slices have the same elements in the same order.
//
// A nil slice and an empty slice are considered equal.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// EqualFunc is like Equal, but compares the elements with the given function.
func EqualFunc[T any](a, b []T, eq func(a, b T) bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !eq(v, b[i]) {
			return false
		}
	}
	return true
}

// Grow increases the slice capacity to guarantee space for n more elements.
func Grow[T any](ls []T, n int) []T {
	if n < 0 {
		panic("arrutil.Grow: negative count")
	}
	if cap(ls)-len(ls) >= n {
		return ls
	}

	newCap := len(ls) + n
	if double := 2 * cap(ls); double > newCap {
		newCap = double
	}

	newLs := make([]T, len(ls), newCap)
	copy(newLs, ls)
	return newLs
}
