package arrutil_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/x/assert"
)

func TestSort(t *testing.T) {
	ints := []int{3, 1, 2}
	arrutil.Sort(ints)
	assert.Eq(t, []int{1, 2, 3}, ints)

	strs := []string{"b", "c", "a"}
	arrutil.Sort(strs)
	assert.Eq(t, []string{"a", "b", "c"}, strs)
}

func TestSortFunc(t *testing.T) {
	ints := []int{3, 1, 2}
	arrutil.SortFunc(ints, func(a, b int) int { return b - a })
	assert.Eq(t, []int{3, 2, 1}, ints)

	// stable: same first char keeps the original order
	strs := []string{"b1", "a1", "b2", "a2"}
	arrutil.SortStableFunc(strs, func(a, b string) int {
		switch {
		case a[0] < b[0]:
			return -1
		case a[0] > b[0]:
			return 1
		}
		return 0
	})
	assert.Eq(t, []string{"a1", "a2", "b1", "b2"}, strs)
}

func TestIsSorted(t *testing.T) {
	assert.True(t, arrutil.IsSorted([]int{1, 2, 3}))
	assert.True(t, arrutil.IsSorted([]int{1, 1, 2}))
	assert.False(t, arrutil.IsSorted([]int{2, 1, 3}))
	assert.True(t, arrutil.IsSorted([]int{}))

	assert.True(t, arrutil.IsSortedFunc([]string{"a", "bb"}, func(a, b string) int {
		return len(a) - len(b)
	}))
	assert.False(t, arrutil.IsSortedFunc([]string{"bb", "a"}, func(a, b string) int {
		return len(a) - len(b)
	}))
}

func TestContainsFunc(t *testing.T) {
	ss := []string{"a", "bb", "ccc"}
	assert.True(t, arrutil.ContainsFunc(ss, func(s string) bool { return len(s) == 2 }))
	assert.False(t, arrutil.ContainsFunc(ss, func(s string) bool { return len(s) == 4 }))

	assert.Eq(t, 1, arrutil.IndexFunc(ss, func(s string) bool { return len(s) == 2 }))
	assert.Eq(t, -1, arrutil.IndexFunc(ss, func(s string) bool { return len(s) == 4 }))
}

func TestCompact(t *testing.T) {
	ss := []string{"a", "a", "b", "b", "b", "c", "a"}
	got := arrutil.Compact(ss)
	assert.Eq(t, []string{"a", "b", "c", "a"}, got)
	// the elements after the new length are zeroed
	assert.Eq(t, []string{"a", "b", "c", "a", "", "", ""}, ss)

	assert.Eq(t, []int{1, 2}, arrutil.Compact([]int{1, 1, 2}))
	assert.Eq(t, 0, len(arrutil.Compact([]int{})))

	// ignore case
	got = arrutil.CompactFunc([]string{"A", "a", "b"}, func(a, b string) bool {
		return strings.EqualFold(a, b)
	})
	assert.Eq(t, []string{"A", "b"}, got)
}

func TestMaxMin(t *testing.T) {
	assert.Eq(t, 9, arrutil.Max([]int{3, 9, 5}))
	assert.Eq(t, 3, arrutil.Min([]int{3, 9, 5}))
	assert.Eq(t, "c", arrutil.Max([]string{"a", "c", "b"}))

	assert.Panics(t, func() {
		arrutil.Max([]int{})
	})
	assert.Panics(t, func() {
		arrutil.Min([]int{})
	})

	ss := []string{"a", "ccc", "bb"}
	assert.Eq(t, "ccc", arrutil.MaxFunc(ss, func(a, b string) int { return len(a) - len(b) }))
	assert.Eq(t, "a", arrutil.MinFunc(ss, func(a, b string) int { return len(a) - len(b) }))
	assert.Panics(t, func() {
		arrutil.MaxFunc([]string{}, func(a, b string) int { return 0 })
	})
}

func TestEqual(t *testing.T) {
	assert.True(t, arrutil.Equal([]int{1, 2}, []int{1, 2}))
	assert.False(t, arrutil.Equal([]int{1, 2}, []int{1, 2, 3}))
	assert.False(t, arrutil.Equal([]int{1, 2}, []int{2, 1}))
	// nil and empty are equal
	assert.True(t, arrutil.Equal([]int{}, []int(nil)))

	assert.True(t, arrutil.EqualFunc([]string{"A"}, []string{"a"}, func(a, b string) bool {
		return strings.EqualFold(a, b)
	}))
	assert.False(t, arrutil.EqualFunc([]string{"A", "b"}, []string{"a"}, func(a, b string) bool {
		return strings.EqualFold(a, b)
	}))
}

func TestInsertConcatClip(t *testing.T) {
	ls := []int{1, 2, 4}
	got := arrutil.Insert(ls, 2, 3)
	assert.Eq(t, []int{1, 2, 3, 4}, got)
	// the given slice is not modified
	assert.Eq(t, []int{1, 2, 4}, ls)

	assert.Eq(t, []int{1, 2, 4, 9}, arrutil.Insert(ls, 3, 9))
	assert.Eq(t, []int{9, 1, 2, 4}, arrutil.Insert(ls, 0, 9))
	assert.Panics(t, func() {
		arrutil.Insert(ls, 4, 9)
	})

	assert.Eq(t, []int{1, 2, 4}, arrutil.Concat([]int{1}, []int{2, 4}))
	assert.Eq(t, 0, len(arrutil.Concat([]int{}, []int{})))

	clipped := arrutil.Clip(append([]int{1, 2}, 3, 4)[:2])
	assert.Eq(t, 2, len(clipped))
	assert.Eq(t, 2, cap(clipped))
}

func TestGrow(t *testing.T) {
	ls := arrutil.Grow([]int{1, 2}, 3)
	assert.Eq(t, []int{1, 2}, ls)
	assert.True(t, cap(ls) >= 5)

	// enough capacity: the same slice is returned
	ls = ls[:0]
	assert.Eq(t, cap(ls), cap(arrutil.Grow(ls, 1)))

	assert.Panics(t, func() {
		arrutil.Grow([]int{1}, -1)
	})
}

func TestBinarySearch(t *testing.T) {
	ints := []int{1, 3, 5, 7}
	pos, found := arrutil.BinarySearch(ints, 5)
	assert.True(t, found)
	assert.Eq(t, 2, pos)

	// not found: the insert position
	pos, found = arrutil.BinarySearch(ints, 4)
	assert.False(t, found)
	assert.Eq(t, 2, pos)

	pos, found = arrutil.BinarySearch(ints, 0)
	assert.False(t, found)
	assert.Eq(t, 0, pos)

	pos, found = arrutil.BinarySearch(ints, 9)
	assert.False(t, found)
	assert.Eq(t, 4, pos)

	// by length
	ss := []string{"a", "bb", "ccc"}
	pos, found = arrutil.BinarySearchFunc(ss, "bb", func(a, b string) int { return len(a) - len(b) })
	assert.True(t, found)
	assert.Eq(t, 1, pos)
}

func TestDeleteRange(t *testing.T) {
	ls := []int{1, 2, 3, 4, 5}
	got := arrutil.DeleteRange(ls, 1, 3)
	assert.Eq(t, []int{1, 4, 5}, got)
	// the elements after the new length are zeroed
	assert.Eq(t, []int{1, 4, 5, 0, 0}, ls)

	// i == j: nothing removed
	assert.Eq(t, 3, len(arrutil.DeleteRange([]int{1, 2, 3}, 1, 1)))

	assert.Panics(t, func() {
		arrutil.DeleteRange([]int{1}, 2, 1)
	})
	assert.Panics(t, func() {
		arrutil.DeleteRange([]int{1}, 0, 2)
	})
}

func TestReplace(t *testing.T) {
	ls := []int{1, 2, 3, 4}
	got := arrutil.Replace(ls, 1, 3, 9, 8, 7)
	assert.Eq(t, []int{1, 9, 8, 7, 4}, got)
	// the given slice is not modified
	assert.Eq(t, []int{1, 2, 3, 4}, ls)

	// insert: i == j
	assert.Eq(t, []int{1, 2, 9, 3, 4}, arrutil.Replace(ls, 2, 2, 9))
	// remove: no value
	assert.Eq(t, []int{1, 4}, arrutil.Replace(ls, 1, 3))

	assert.Panics(t, func() {
		arrutil.Replace(ls, 3, 1)
	})
}

func TestRepeat(t *testing.T) {
	assert.Eq(t, []string{"a", "a", "a"}, arrutil.Repeat("a", 3))
	assert.Eq(t, 0, len(arrutil.Repeat(0, 0)))

	assert.Panics(t, func() {
		arrutil.Repeat(1, -1)
	})
}

func TestDeleteFunc(t *testing.T) {
	ls := []int{1, 2, 3, 4, 5, 6}
	got := arrutil.DeleteFunc(ls, func(v int) bool { return v%2 == 0 })
	assert.Eq(t, []int{1, 3, 5}, got)
	// the elements after the new length are zeroed
	assert.Eq(t, []int{1, 3, 5, 0, 0, 0}, ls)

	// nothing removed
	assert.Eq(t, 3, len(arrutil.DeleteFunc([]int{1, 2, 3}, func(v int) bool { return false })))
	// all removed
	assert.Eq(t, 0, len(arrutil.DeleteFunc([]int{1, 2}, func(v int) bool { return true })))
}

func TestCount(t *testing.T) {
	assert.Eq(t, 2, arrutil.Count([]int{1, 2, 1, 3}, 1))
	assert.Eq(t, 0, arrutil.Count([]int{1, 2}, 9))
	assert.Eq(t, 2, arrutil.CountFunc([]string{"a", "bb", "cc"}, func(s string) bool { return len(s) == 2 }))
}
