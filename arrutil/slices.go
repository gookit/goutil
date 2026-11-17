package arrutil

import "github.com/gookit/goutil/comdef"

// Clip removes the unused capacity of the given slice.
//
// The result shares the same backing array, use it to drop references to
// elements that are no longer reachable through the slice.
func Clip[T any](ls []T) []T {
	return ls[:len(ls):len(ls)]
}

// Insert inserts the values vs into the slice at index i and returns the result.
//
// It panics if i is out of the range [0, len(ls)]. Unlike slices.Insert it always
// returns a new slice, the given slice is never modified.
func Insert[T any](ls []T, i int, vs ...T) []T {
	if i < 0 || i > len(ls) {
		panic("arrutil.Insert: index out of range")
	}
	if len(vs) == 0 {
		return ls
	}

	newLs := make([]T, 0, len(ls)+len(vs))
	newLs = append(newLs, ls[:i]...)
	newLs = append(newLs, vs...)
	newLs = append(newLs, ls[i:]...)
	return newLs
}

// Concat returns a new slice with all the given slices concatenated.
//
// A nil result is returned when every given slice is empty.
func Concat[T any](ls ...[]T) []T {
	total := 0
	for _, item := range ls {
		total += len(item)
	}
	if total == 0 {
		return nil
	}

	newLs := make([]T, 0, total)
	for _, item := range ls {
		newLs = append(newLs, item...)
	}
	return newLs
}

// DeleteRange removes the elements ls[i:j] from the slice, in place.
//
// It panics if the indexes are out of the range [0, len(ls)] or i > j. The elements
// after the returned length are zeroed.
func DeleteRange[T any](ls []T, i, j int) []T {
	if i < 0 || j > len(ls) || i > j {
		panic("arrutil.DeleteRange: index out of range")
	}
	if i == j {
		return ls
	}

	n := copy(ls[i:], ls[j:])

	var zero T
	for k := i + n; k < len(ls); k++ {
		ls[k] = zero
	}
	return ls[:i+n]
}

// Replace replaces the elements ls[i:j] with the given values and returns the result.
//
// It panics if the indexes are out of the range [0, len(ls)] or i > j. Unlike
// slices.Replace it always returns a new slice, the given slice is never modified.
func Replace[T any](ls []T, i, j int, vs ...T) []T {
	if i < 0 || j > len(ls) || i > j {
		panic("arrutil.Replace: index out of range")
	}

	newLs := make([]T, 0, len(ls)-(j-i)+len(vs))
	newLs = append(newLs, ls[:i]...)
	newLs = append(newLs, vs...)
	newLs = append(newLs, ls[j:]...)
	return newLs
}

// Repeat returns a new slice with the given value repeated n times.
//
// It panics if n is negative.
func Repeat[T any](v T, n int) []T {
	if n < 0 {
		panic("arrutil.Repeat: negative count")
	}

	newLs := make([]T, n)
	for i := range newLs {
		newLs[i] = v
	}
	return newLs
}

// Count the occurrences of the given value in the slice.
func Count[T comparable](ls []T, v T) int {
	count := 0
	for _, item := range ls {
		if item == v {
			count++
		}
	}
	return count
}

// CountFunc counts the elements that satisfy the given function.
func CountFunc[T any](ls []T, fn comdef.MatchFunc[T]) int {
	count := 0
	for _, item := range ls {
		if fn(item) {
			count++
		}
	}
	return count
}
