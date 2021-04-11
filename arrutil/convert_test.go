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

	_, err = arrutil.ToStrings("b")
	is.Error(err)

	_, err = arrutil.ToStrings([]interface{}{[]int{1}, nil})
	is.Error(err)
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := arrutil.StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = arrutil.StringsToInts([]string{"a", "b"})
	is.Error(err)
}
