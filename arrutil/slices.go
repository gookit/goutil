package arrutil

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
