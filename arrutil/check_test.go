package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSliceHas(t *testing.T) {
	ints := []int{2, 4, 5}
	assert.True(t, arrutil.SliceHas(ints, 2))
	assert.False(t, arrutil.SliceHas(ints, 3))

	int64s := []int64{2, 4, 5}
	assert.True(t, arrutil.SliceHas(int64s, 2))
	assert.False(t, arrutil.SliceHas(int64s, 3))

	strs := []string{"2", "4", "5"}
	assert.True(t, arrutil.SliceHas(strs, "2"))
	assert.False(t, arrutil.SliceHas(strs, "3"))
}

func TestIntsHas(t *testing.T) {
	ints := []int{2, 4, 5}
	assert.True(t, arrutil.IntsHas(ints, 2))
	assert.True(t, arrutil.IntsHas(ints, 5))
	assert.False(t, arrutil.IntsHas(ints, 3))

	uints := []uint{2, 4, 5}
	assert.True(t, arrutil.IntsHas(uints, 2))
	assert.False(t, arrutil.IntsHas(uints, 3))
}

func TestInt64sHas(t *testing.T) {
	ints := []int64{2, 4, 5}
	assert.True(t, arrutil.Int64sHas(ints, 2))
	assert.True(t, arrutil.Int64sHas(ints, 5))
	assert.False(t, arrutil.Int64sHas(ints, 3))
}

func TestStringsHas(t *testing.T) {
	ss := []string{"a", "b"}
	assert.True(t, arrutil.StringsHas(ss, "a"))
	assert.True(t, arrutil.StringsHas(ss, "b"))
	assert.True(t, arrutil.InStrings("b", ss))

	assert.False(t, arrutil.StringsHas(ss, "c"))
	assert.False(t, arrutil.InStrings("c", ss))
}

func TestInAndNotIn(t *testing.T) {
	is := assert.New(t)

	arr := []int{1, 2, 3}
	is.True(arrutil.In(2, arr))
	is.False(arrutil.NotIn(2, arr))

	arr1 := []rune{'a', 'b'}
	is.True(arrutil.In('b', arr1))
	is.False(arrutil.NotIn('b', arr1))

	arr2 := []string{"a", "b", "c"}
	is.True(arrutil.In("b", arr2))
	is.False(arrutil.NotIn("b", arr2))
}

func TestContainsAll(t *testing.T) {
	is := assert.New(t)

	arr := []int{1, 2, 3}
	is.True(arrutil.ContainsAll(arr, []int{2}))
	is.False(arrutil.ContainsAll(arr, []int{2, 45}))
	is.True(arrutil.IsParent(arr, []int{2}))

	arr2 := []string{"a", "b", "c"}
	is.True(arrutil.ContainsAll(arr2, []string{"b"}))
	is.False(arrutil.ContainsAll(arr2, []string{"b", "e"}))
	is.True(arrutil.IsParent(arr2, []string{"b"}))
}

func TestContains(t *testing.T) {
	is := assert.New(t)
	tests := map[any]any{
		1:    []int{1, 2, 3},
		2:    []int8{1, 2, 3},
		3:    []int16{1, 2, 3},
		4:    []int32{4, 2, 3},
		5:    []int64{5, 2, 3},
		6:    []uint{6, 2, 3},
		7:    []uint8{7, 2, 3},
		8:    []uint16{8, 2, 3},
		9:    []uint32{9, 2, 3},
		10:   []uint64{10, 3},
		11:   []string{"11", "3"},
		'a':  []int64{97},
		'b':  []rune{'a', 'b'},
		'c':  []byte{'a', 'b', 'c'}, // byte -> uint8
		"a":  []string{"a", "b", "c"},
		12:   [5]uint{12, 1, 2, 3, 4},
		'A':  [3]rune{'A', 'B', 'C'},
		'd':  [4]byte{'a', 'b', 'c', 'd'},
		"aa": [3]string{"aa", "bb", "cc"},
	}

	for val, list := range tests {
		is.True(arrutil.Contains(list, val))
		is.False(arrutil.NotContains(list, val))
	}

	is.False(arrutil.Contains(nil, []int{}))
	is.False(arrutil.Contains('a', []int{}))
	//
	is.False(arrutil.Contains([]int{2, 3}, []int{2}))
	is.False(arrutil.Contains([]int{2, 3}, "a"))
	is.False(arrutil.Contains([]string{"a", "b"}, 12))
	is.False(arrutil.Contains(nil, 12))
	is.False(arrutil.Contains(map[int]int{2: 3}, 12))

	tests1 := map[any]any{
		2:   []int{1, 3},
		"a": []string{"b", "c"},
	}

	for val, list := range tests1 {
		is.True(arrutil.NotContains(list, val))
		is.False(arrutil.Contains(list, val))
	}
}
