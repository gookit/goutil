package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestInts_methods(t *testing.T) {
	tests := []struct {
		is    arrutil.Ints
		val   int
		want  bool
		want2 string
	}{
		{
			arrutil.Ints{12, 23},
			12,
			true,
			"12,23",
		},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, tt.is.Has(tt.val))
		assert.False(t, tt.is.Has(999))
		assert.Eq(t, tt.want2, tt.is.String())
	}

	ints := arrutil.Ints{23, 10, 12}
	ints.Sort()
	assert.Eq(t, "10,12,23", ints.String())
	assert.Eq(t, 10, ints.First())
	assert.Eq(t, 23, ints.Last())
}

func TestStrings_methods(t *testing.T) {
	tests := []struct {
		ss    arrutil.Strings
		val   string
		want  bool
		want2 string
	}{
		{
			arrutil.Strings{"a", "b"},
			"a",
			true,
			"a,b",
		},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, tt.ss.Has(tt.val))
		assert.False(t, tt.ss.Has("not-exists"))
		assert.Eq(t, tt.want2, tt.ss.String())
	}

	ss := arrutil.Strings{"a", "b"}
	assert.Eq(t, "a b", ss.Join(" "))
	assert.Eq(t, "a", ss.First())
}

func TestScalarList_methods(t *testing.T) {
	ls := arrutil.ScalarList[string]{"a", "", "b"}
	assert.Eq(t, "a", ls.First())
	assert.Eq(t, "b", ls.Last())
	assert.True(t, ls.Has("a"))
	assert.False(t, ls.IsEmpty())
	assert.Eq(t, "[a,b]", ls.Filter().String())
	assert.Eq(t, "[a,b]", ls.Remove("").String())

	t.Run("panic", func(t *testing.T) {
		ls = arrutil.ScalarList[string]{}
		assert.Panics(t, func() {
			ls.First()
		})
		assert.Panics(t, func() {
			ls.Last()
		})
	})
}
