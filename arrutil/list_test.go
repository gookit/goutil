package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestInts_methods(t *testing.T) {
	ints := arrutil.Ints[int]{23, 10, 12}
	assert.False(t, ints.Has(999))
	assert.Eq(t, "[23,10,12]", ints.String())

	ints.Sort()
	assert.Eq(t, "[10,12,23]", ints.String())
	assert.Eq(t, 10, ints.First())
	assert.Eq(t, 23, ints.Last())

	ints = arrutil.Ints[int]{}
	assert.Eq(t, 1, ints.First(1))
	assert.Eq(t, 2, ints.Last(2))

	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() {
			ints.First()
		})
		assert.Panics(t, func() {
			ints.Last()
		})
	})
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
	assert.Eq(t, "b", ss.Last())

	ss = arrutil.Strings{"b", "a"}
	ss.Sort()
	assert.Eq(t, "a", ss.First())

	ss = arrutil.Strings{}
	assert.Eq(t, "default1", ss.First("default1"))
	assert.Eq(t, "default2", ss.Last("default2"))

	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() {
			ss.First()
		})
		assert.Panics(t, func() {
			ss.Last()
		})
	})
}

func TestSortedList_methods(t *testing.T) {
	ls := arrutil.SortedList[string]{"a", "", "b"}
	assert.Eq(t, "a", ls.First())
	assert.Eq(t, "b", ls.Last())
	assert.True(t, ls.Has("a"))
	assert.False(t, ls.Has("e"))
	assert.False(t, ls.IsEmpty())
	assert.Eq(t, "[a,b]", ls.Filter().String())

	ls = ls.Remove("")
	assert.Eq(t, "[a,b]", ls.String())

	ls1 := arrutil.SortedList[int]{4, 3}
	ls1.Sort()
	assert.Eq(t, "[3,4]", ls1.String())

	ls = arrutil.SortedList[string]{}
	assert.Eq(t, "default1", ls.First("default1"))
	assert.Eq(t, "default2", ls.Last("default2"))

	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() {
			ls.First()
		})
		assert.Panics(t, func() {
			ls.Last()
		})
	})
}
