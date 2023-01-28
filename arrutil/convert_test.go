package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestToInt64s(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.ToInt64s([]string{"1", "2"})
	is.Nil(err)
	is.Eq("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = arrutil.MustToInt64s([]string{"1", "2"})
	is.Eq("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = arrutil.MustToInt64s([]any{"1", "2"})
	is.Eq("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = arrutil.SliceToInt64s([]any{"1", "2"})
	is.Eq("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.ToInt64s([]string{"a", "b"})
	is.Err(err)
}

func TestToStrings(t *testing.T) {
	is := assert.New(t)

	ss, err := arrutil.ToStrings([]int{1, 2})
	is.Nil(err)
	is.Eq(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = arrutil.MustToStrings([]int{1, 2})
	is.Eq(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = arrutil.MustToStrings([]any{1, 2})
	is.Eq(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = arrutil.SliceToStrings([]any{1, 2})
	is.Eq(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	as := arrutil.StringsToSlice([]string{"1", "2"})
	is.Eq(`[]interface {}{"1", "2"}`, fmt.Sprintf("%#v", as))

	ss, err = arrutil.ToStrings("b")
	is.Nil(err)
	is.Eq(`[]string{"b"}`, fmt.Sprintf("%#v", ss))

	_, err = arrutil.ToStrings([]any{[]int{1}, nil})
	is.Err(err)
}

func TestStringsToString(t *testing.T) {
	is := assert.New(t)

	is.Eq("a,b", arrutil.JoinStrings(",", []string{"a", "b"}...))
	is.Eq("a,b", arrutil.StringsJoin(",", []string{"a", "b"}...))
	is.Eq("a,b", arrutil.StringsJoin(",", "a", "b"))
}

func TestAnyToString(t *testing.T) {
	is := assert.New(t)
	arr := [2]string{"a", "b"}

	is.Eq("", arrutil.AnyToString(nil))
	is.Eq("[]", arrutil.AnyToString([]string{}))
	is.Eq("[a, b]", arrutil.AnyToString(arr))
	is.Eq("[a, b]", arrutil.AnyToString([]string{"a", "b"}))
	is.Eq("", arrutil.AnyToString("invalid"))
}

func TestSliceToString(t *testing.T) {
	is := assert.New(t)

	is.Eq("[]", arrutil.SliceToString(nil))
	is.Eq("[a,b]", arrutil.SliceToString("a", "b"))
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Eq("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.StringsToInts([]string{"a", "b"})
	is.Err(err)

	ints = arrutil.StringsAsInts([]string{"1", "2"})
	is.Eq("[]int{1, 2}", fmt.Sprintf("%#v", ints))
	is.Nil(arrutil.StringsAsInts([]string{"abc"}))
}

func TestConvType(t *testing.T) {
	is := assert.New(t)

	// []string => []int
	arr, err := arrutil.ConvType([]string{"1", "2"}, 1)
	is.Nil(err)
	is.Eq("[]int{1, 2}", fmt.Sprintf("%#v", arr))

	// []int => []string
	arr1, err := arrutil.ConvType([]int{1, 2}, "1")
	is.Nil(err)
	is.Eq(`[]string{"1", "2"}`, fmt.Sprintf("%#v", arr1))

	// not need conv
	arr2, err := arrutil.ConvType([]string{"1", "2"}, "1")
	is.Nil(err)
	is.Eq(`[]string{"1", "2"}`, fmt.Sprintf("%#v", arr2))

	// conv error
	arr3, err := arrutil.ConvType([]string{"ab", "cd"}, 1)
	is.Err(err)
	is.Nil(arr3)
}

func TestJoinSlice(t *testing.T) {
	assert.Eq(t, "", arrutil.JoinSlice(","))
	assert.Eq(t, "", arrutil.JoinSlice(",", nil))
	assert.Eq(t, "a,23,b", arrutil.JoinSlice(",", "a", 23, "b"))
}
