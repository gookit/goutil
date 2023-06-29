package arrutil

import (
	"sort"
	"strings"

	"github.com/gookit/goutil/comdef"
)

// Ints type
type Ints []int

// String to string
func (is Ints) String() string {
	return ToString(is)
}

// Has given element
func (is Ints) Has(i int) bool {
	for _, iv := range is {
		if i == iv {
			return true
		}
	}
	return false
}

// First element value.
func (is Ints) First() int {
	if len(is) > 0 {
		return is[0]
	}
	panic("empty int slice")
}

// Last element value.
func (is Ints) Last() int {
	if len(is) > 0 {
		return is[len(is)-1]
	}
	panic("empty int slice")
}

// Sort the int slice
func (is Ints) Sort() {
	sort.Ints(is)
}

// Strings type
type Strings []string

// String to string
func (ss Strings) String() string {
	return strings.Join(ss, ",")
}

// Join to string
func (ss Strings) Join(sep string) string {
	return strings.Join(ss, sep)
}

// Has given element
func (ss Strings) Has(sub string) bool {
	return ss.Contains(sub)
}

// Contains given element
func (ss Strings) Contains(sub string) bool {
	for _, s := range ss {
		if s == sub {
			return true
		}
	}
	return false
}

// First element value.
func (ss Strings) First() string {
	if len(ss) > 0 {
		return ss[0]
	}
	panic("empty string list")
}

// Last element value.
func (ss Strings) Last() string {
	if len(ss) > 0 {
		return ss[len(ss)-1]
	}
	panic("empty string list")
}

// ScalarList definition for any type
type ScalarList[T comdef.ScalarType] []T

// IsEmpty check
func (ls ScalarList[T]) IsEmpty() bool {
	return len(ls) == 0
}

// String to string
func (ls ScalarList[T]) String() string {
	return ToString(ls)
}

// Has given element
func (ls ScalarList[T]) Has(el T) bool {
	return ls.Contains(el)
}

// Contains given element
func (ls ScalarList[T]) Contains(el T) bool {
	for _, v := range ls {
		if v == el {
			return true
		}
	}
	return false
}

// First element value.
func (ls ScalarList[T]) First() T {
	if len(ls) > 0 {
		return ls[0]
	}
	panic("empty list")
}

// Last element value.
func (ls ScalarList[T]) Last() T {
	if ln := len(ls); ln > 0 {
		return ls[ln-1]
	}
	panic("empty list")
}

// Remove given element
func (ls ScalarList[T]) Remove(el T) ScalarList[T] {
	return Filter(ls, func(v T) bool {
		return v != el
	})
}

// Filter the slice, default will filter zero value.
func (ls ScalarList[T]) Filter(filter ...comdef.MatchFunc[T]) ScalarList[T] {
	return Filter(ls, filter...)
}

// Map the slice to new slice. TODO syntax ERROR: Method cannot have type parameters
// func (ls ScalarList[T]) Map[V any](mapFn MapFn[T, V]) ScalarList[V] {
// 	return Map(ls, mapFn)
// }

// Sort the slice
// func (ls ScalarList[T]) Sort() {
// sort.Sort(ls)
// }
