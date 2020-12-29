package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	ss := []string{"a", "b", "c"}

	arrutil.Reverse(ss)

	assert.Equal(t, []string{"c", "b", "a"}, ss)
}

func TestStringsRemove(t *testing.T) {
	ss := []string{"a", "b", "c"}

	ns := arrutil.StringsRemove(ss, "b")
	assert.Contains(t, ns, "a")
	assert.NotContains(t, ns, "b")
}

func TestTrimStrings(t *testing.T) {
	is := assert.New(t)

	// TrimStrings
	ss := arrutil.TrimStrings([]string{" a", "b ", " c "})
	is.Equal("[a b c]", fmt.Sprint(ss))
	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
	is.Equal("[a b c]", fmt.Sprint(ss))
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.StringsToInts([]string{"a", "b"})
	is.Error(err)
}

func TestIntsHas(t *testing.T) {
	ints := []int{2, 4, 5}
	assert.True(t, arrutil.IntsHas(ints, 2))
	assert.True(t, arrutil.IntsHas(ints, 5))
	assert.False(t, arrutil.IntsHas(ints, 3))
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
	assert.False(t, arrutil.StringsHas(ss, "c"))
}

func TestContains(t *testing.T) {
	is := assert.New(t)
	tests := map[interface{}]interface{}{
		1:   []int{1, 2, 3},
		2:   []int8{1, 2, 3},
		3:   []int16{1, 2, 3},
		4:   []int32{4, 2, 3},
		5:   []int64{5, 2, 3},
		6:   []uint{6, 2, 3},
		7:   []uint8{7, 2, 3},
		8:   []uint16{8, 2, 3},
		9:   []uint32{9, 2, 3},
		10:  []uint64{10, 3},
		11:  []string{"11", "3"},
		'a': []int64{97},
		'b': []rune{'a', 'b'},
		'c': []byte{'a', 'b', 'c'}, // byte -> uint8
		"a": []string{"a", "b", "c"},
	}

	for val, list := range tests {
		is.True(arrutil.Contains(list, val))
		is.False(arrutil.NotContains(list, val))
	}

	is.False(arrutil.Contains(nil, []int{}))
	is.False(arrutil.Contains('a', []int{}))
	//
	is.False(arrutil.Contains([]int{2, 3}, []int{2}))
	is.False(arrutil.Contains([]string{"a", "b"}, 12))
	is.False(arrutil.Contains(nil, 12))
	is.False(arrutil.Contains(map[int]int{2: 3}, 12))

	tests1 := map[interface{}]interface{}{
		2:   []int{1, 3},
		"a": []string{"b", "c"},
	}

	for val, list := range tests1 {
		is.True(arrutil.NotContains(list, val))
		is.False(arrutil.Contains(list, val))
	}
}
