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

	ss, err = arrutil.ToStrings("b")
	is.Nil(err)
	is.Eq(`[]string{"b"}`, fmt.Sprintf("%#v", ss))

	is.Empty(arrutil.AnyToStrings(234))
	is.Panics(func() {
		arrutil.MustToStrings(234)
	})

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

	is.Eq("[]", arrutil.ToString[any](nil))
	is.Eq("[a,b]", arrutil.ToString([]any{"a", "b"}))

	is.Eq("[a,b]", arrutil.SliceToString("a", "b"))
}

func TestAnyToSlice(t *testing.T) {
	is := assert.New(t)

	sl, err := arrutil.AnyToSlice([]int{1, 2})
	is.NoErr(err)
	is.Eq("[]interface {}{1, 2}", fmt.Sprintf("%#v", sl))

	_, err = arrutil.AnyToSlice(123)
	is.Err(err)
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

func TestJoinTyped(t *testing.T) {
	assert.Eq(t, "", arrutil.JoinTyped[any](","))
	assert.Eq(t, "", arrutil.JoinTyped[any](",", nil))
	assert.Eq(t, "1,2", arrutil.JoinTyped(",", 1, 2))
	assert.Eq(t, "a,b", arrutil.JoinTyped(",", "a", "b"))
	assert.Eq(t, "1,a", arrutil.JoinTyped[any](",", 1, "a"))
}

func TestJoinSlice(t *testing.T) {
	assert.Eq(t, "", arrutil.JoinSlice(","))
	assert.Eq(t, "", arrutil.JoinSlice(",", nil))
	assert.Eq(t, "a,23,b", arrutil.JoinSlice(",", "a", 23, "b"))
}

func TestIntsToString(t *testing.T) {
	assert.Eq(t, "[]", arrutil.IntsToString([]int{}))
	assert.Eq(t, "[1,2,3]", arrutil.IntsToString([]int{1, 2, 3}))
}

func TestCombineToMap(t *testing.T) {
	keys := []string{"key0", "key1"}

	mp := arrutil.CombineToMap(keys, []int{1, 2})
	assert.Len(t, mp, 2)
	assert.Eq(t, 1, mp["key0"])
	assert.Eq(t, 2, mp["key1"])

	mp = arrutil.CombineToMap(keys, []int{1})
	assert.Len(t, mp, 1)
	assert.Eq(t, 1, mp["key0"])
}

func TestCombineToSMap(t *testing.T) {
	keys := []string{"key0", "key1"}

	mp := arrutil.CombineToSMap(keys, []string{"val0", "val1"})
	assert.Len(t, mp, 2)
	assert.Eq(t, "val0", mp["key0"])

	mp = arrutil.CombineToSMap(keys, []string{"val0"})
	assert.Len(t, mp, 2)
	assert.Eq(t, "val0", mp["key0"])
	assert.Eq(t, "", mp["key1"])
}
