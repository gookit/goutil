package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/stretchr/testify/assert"
)

func TestToInt64s(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.ToInt64s([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = arrutil.MustToInt64s([]string{"1", "2"})
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = arrutil.MustToInt64s([]interface{}{"1", "2"})
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	ints = arrutil.SliceToInt64s([]interface{}{"1", "2"})
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.ToInt64s([]string{"a", "b"})
	is.Error(err)
}

func TestToStrings(t *testing.T) {
	is := assert.New(t)

	ss, err := arrutil.ToStrings([]int{1, 2})
	is.Nil(err)
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = arrutil.MustToStrings([]int{1, 2})
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = arrutil.MustToStrings([]interface{}{1, 2})
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	ss = arrutil.SliceToStrings([]interface{}{1, 2})
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))

	as := arrutil.StringsToSlice([]string{"1", "2"})
	is.Equal(`[]interface {}{"1", "2"}`, fmt.Sprintf("%#v", as))

	_, err = arrutil.ToStrings("b")
	is.Error(err)

	_, err = arrutil.ToStrings([]interface{}{[]int{1}, nil})
	is.Error(err)
}

func TestStringsToString(t *testing.T) {
	is := assert.New(t)

	is.Equal("a,b", arrutil.JoinStrings(",", []string{"a", "b"}...))
	is.Equal("a,b", arrutil.StringsJoin(",", []string{"a", "b"}...))
	is.Equal("a,b", arrutil.StringsJoin(",", "a", "b"))
}

func TestAnyToString(t *testing.T) {
	is := assert.New(t)
	arr := [2]string{"a", "b"}

	is.Equal("[]", arrutil.AnyToString(nil))
	is.Equal("[]", arrutil.AnyToString([]string{}))
	is.Equal("[a,b]", arrutil.AnyToString(arr))
	is.Equal("[a,b]", arrutil.AnyToString([]string{"a", "b"}))
	is.Equal("", arrutil.AnyToString("invalid"))
}

func TestSliceToString(t *testing.T) {
	is := assert.New(t)

	is.Equal("[]", arrutil.SliceToString(nil))
	is.Equal("[a,b]", arrutil.SliceToString("a", "b"))
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.StringsToInts([]string{"a", "b"})
	is.Error(err)
}

func TestJoinSlice(t *testing.T) {
	assert.Equal(t, "", arrutil.JoinSlice(","))
	assert.Equal(t, "", arrutil.JoinSlice(",", nil))
	assert.Equal(t, "a,23,b", arrutil.JoinSlice(",", "a", 23, "b"))
}
