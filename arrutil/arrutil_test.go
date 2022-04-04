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
	assert.Len(t, ns, 2)
}

func TestTrimStrings(t *testing.T) {
	is := assert.New(t)

	// TrimStrings
	ss := arrutil.TrimStrings([]string{" a", "b ", " c "})
	is.Equal("[a b c]", fmt.Sprint(ss))
	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
	is.Equal("[a b c]", fmt.Sprint(ss))
	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",", ".")
	is.Equal("[a b c]", fmt.Sprint(ss))
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
	is.True(arrutil.HasValue(intSlice, intVal))
	assert.IsType(t, 0, intVal1)
	is.True(arrutil.HasValue(intSlice, intVal1))
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
