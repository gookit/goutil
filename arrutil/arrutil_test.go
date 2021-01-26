package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	is := assert.New(t)
	tests := map[interface{}]interface{}{
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

func TestGetRandomOne(t *testing.T) {
	is := assert.New(t)
	// int slice
	intSlice := []int{1, 2, 3, 4, 5, 6}
	intVal := arrutil.GetRandomOne(intSlice)
	intVal1 := arrutil.GetRandomOne(intSlice)
	for intVal == intVal1 {
		intVal1 = arrutil.GetRandomOne(intSlice)
	}
	assert.IsType(t, 0, intVal)
	is.True(arrutil.Contains(intSlice, intVal))
	assert.IsType(t, 0, intVal1)
	is.True(arrutil.Contains(intSlice, intVal1))
	assert.NotEqual(t, intVal, intVal1)

	// int array
	intArray := [6]int{1, 2, 3, 4, 5, 6}
	intReturned := arrutil.GetRandomOne(intArray)
	intReturned1 := arrutil.GetRandomOne(intArray)
	for intReturned == intReturned1 {
		intReturned1 = arrutil.GetRandomOne(intArray)
	}
	assert.IsType(t, 0, intReturned)
	is.True(arrutil.Contains(intArray, intReturned))
	assert.IsType(t, 0, intReturned1)
	is.True(arrutil.Contains(intArray, intReturned1))
	assert.NotEqual(t, intReturned, intReturned1)

	// string slice
	strSlice := []string{"aa", "bb", "cc", "dd"}
	strVal := arrutil.GetRandomOne(strSlice)
	strVal1 := arrutil.GetRandomOne(strSlice)
	for strVal == strVal1 {
		strVal1 = arrutil.GetRandomOne(strSlice)
	}
	assert.IsType(t, string(""), strVal)
	is.True(arrutil.Contains(strSlice, strVal))
	assert.IsType(t, string(""), strVal1)
	is.True(arrutil.Contains(strSlice, strVal1))
	assert.NotEqual(t, strVal, strVal1)

	// string array
	strArray := [4]string{"aa", "bb", "cc", "dd"}
	strReturned := arrutil.GetRandomOne(strArray)
	strReturned1 := arrutil.GetRandomOne(strArray)
	for strReturned == strReturned1 {
		strReturned1 = arrutil.GetRandomOne(strArray)
	}
	assert.IsType(t, "", strReturned)
	is.True(arrutil.Contains(strArray, strReturned))
	assert.IsType(t, "", strReturned1)
	is.True(arrutil.Contains(strArray, strReturned1))
	assert.NotEqual(t, strReturned, strReturned1)

	// byte slice
	byteSlice := []byte("abcdefg")
	byteVal := arrutil.GetRandomOne(byteSlice)
	byteVal1 := arrutil.GetRandomOne(byteSlice)
	for byteVal == byteVal1 {
		byteVal1 = arrutil.GetRandomOne(byteSlice)
	}
	assert.IsType(t, byte('a'), byteVal)
	is.True(arrutil.Contains(byteSlice, byteVal))
	assert.IsType(t, byte('a'), byteVal1)
	is.True(arrutil.Contains(byteSlice, byteVal1))
	assert.NotEqual(t, byteVal, byteVal1)

	// int
	invalidIntData := int(404)
	invalidIntReturned := arrutil.GetRandomOne(invalidIntData)
	assert.IsType(t, int(0), invalidIntReturned)

	// float
	invalidDataFloat := float32(3.14)
	invalidFloatReturned := arrutil.GetRandomOne(invalidDataFloat)
	assert.IsType(t, float32(3.1), invalidFloatReturned)
}
