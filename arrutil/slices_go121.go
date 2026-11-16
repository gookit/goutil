//go:build go1.21

package arrutil

import (
	"slices"

	"github.com/gookit/goutil/comdef"
)

// SliceHas check the slice contains the given value. alias of slices.Contains
func SliceHas[T comdef.ScalarType](slice []T, val T) bool {
	return slices.Contains(slice, val)
}

// IntsHas check the []comdef.Integer contains the given value
func IntsHas[T comdef.Integer](ints []T, val T) bool {
	return slices.Contains(ints, val)
}

// Int64sHas check the []int64 contains the given value
func Int64sHas(ints []int64, val int64) bool {
	return slices.Contains(ints, val)
}

// StringsHas check the []string contains the given element
func StringsHas[T ~string](ss []T, val T) bool {
	return slices.Contains(ss, val)
}

// In check the given value whether in the list
func In[T comdef.ScalarType](value T, list []T) bool {
	return slices.Contains(list, value)
}

// IndexOf value in given slice. returns -1 if not found
func IndexOf[T comdef.NumberOrString](val T, list []T) int {
	return slices.Index(list, val)
}

// Reverse any T slice in place.
//
// eg: []string{"site", "user", "info", "0"} -> []string{"0", "info", "user", "site"}
func Reverse[T any](ls []T) {
	slices.Reverse(ls)
}

// CloneSlice Clone a slice.
//
// NOTE: unlike slices.Clone a nil slice gives back an empty non-nil slice.
func CloneSlice[T any](data []T) []T {
	if data == nil {
		return make([]T, 0)
	}
	return slices.Clone(data)
}

// Sort the given slice in place. It panics if the element type is not ordered.
func Sort[T comdef.Compared](ls []T) {
	slices.Sort(ls)
}

// SortFunc sorts the slice in place by the given comparison function.
//
// It panics if cmp is not a valid comparison function for the element type.
func SortFunc[T any](ls []T, cmp Comparer[T]) {
	slices.SortFunc(ls, cmp)
}

// SortStableFunc sorts the slice in place by the given comparison function,
// keeping the original order of equal elements.
func SortStableFunc[T any](ls []T, cmp Comparer[T]) {
	slices.SortStableFunc(ls, cmp)
}

// IsSorted reports whether the given slice is sorted in increasing order.
func IsSorted[T comdef.Compared](ls []T) bool {
	return slices.IsSorted(ls)
}

// IsSortedFunc reports whether the given slice is sorted by the given comparison function.
func IsSortedFunc[T any](ls []T, cmp Comparer[T]) bool {
	return slices.IsSortedFunc(ls, cmp)
}

// ContainsFunc reports whether any element of the slice satisfies the given function.
func ContainsFunc[T any](ls []T, fn comdef.MatchFunc[T]) bool {
	return slices.ContainsFunc(ls, fn)
}

// IndexFunc returns the index of the first element that satisfies the given function,
// or -1 if there is none.
func IndexFunc[T any](ls []T, fn comdef.MatchFunc[T]) int {
	return slices.IndexFunc(ls, fn)
}

// Compact replaces consecutive runs of equal elements with a single copy, in place.
//
// The slice is modified and the elements after the returned length are zeroed.
func Compact[T comparable](ls []T) []T {
	return compactTail(ls, slices.Compact(ls))
}

// CompactFunc is like Compact, but uses eq to compare the elements.
func CompactFunc[T any](ls []T, eq func(a, b T) bool) []T {
	return compactTail(ls, slices.CompactFunc(ls, eq))
}

// compactTail zeroes the elements between the compacted length and the original length.
//
// slices.Compact only zeroes them from go1.22, so do it here to keep the result the same
// on every Go version.
func compactTail[T any](ls, compacted []T) []T {
	var zero T
	for i := len(compacted); i < len(ls); i++ {
		ls[i] = zero
	}
	return compacted
}

// Max returns the maximum value of the slice. It panics on an empty slice.
func Max[T comdef.Compared](ls []T) T {
	return slices.Max(ls)
}

// Min returns the minimum value of the slice. It panics on an empty slice.
func Min[T comdef.Compared](ls []T) T {
	return slices.Min(ls)
}

// MaxFunc returns the maximum value of the slice by the given comparison function.
// It panics on an empty slice.
func MaxFunc[T any](ls []T, cmp Comparer[T]) T {
	return slices.MaxFunc(ls, cmp)
}

// MinFunc returns the minimum value of the slice by the given comparison function.
// It panics on an empty slice.
func MinFunc[T any](ls []T, cmp Comparer[T]) T {
	return slices.MinFunc(ls, cmp)
}

// Equal reports whether the two slices have the same elements in the same order.
//
// A nil slice and an empty slice are considered equal.
func Equal[T comparable](a, b []T) bool {
	return slices.Equal(a, b)
}

// EqualFunc is like Equal, but compares the elements with the given function.
func EqualFunc[T any](a, b []T, eq func(a, b T) bool) bool {
	return slices.EqualFunc(a, b, eq)
}

// Grow increases the slice capacity to guarantee space for n more elements.
func Grow[T any](ls []T, n int) []T {
	return slices.Grow(ls, n)
}
