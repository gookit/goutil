package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestStringsToSlice(t *testing.T) {
	is := assert.New(t)

	as := arrutil.StringsToSlice([]string{"1", "2"})
	is.Eq(`[]interface {}{"1", "2"}`, fmt.Sprintf("%#v", as))
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

func TestStringsUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"EmptyInput", []string{}, []string{}},
		{"NoDuplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"WithDuplicates", []string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
		{"AllDuplicates", []string{"a", "a", "a"}, []string{"a"}},
		{"Mixed", []string{"a", "b", "c", "a", "d", "b"}, []string{"a", "b", "c", "d"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := arrutil.StringsUnique(tt.input)
			assert.Eq(t, tt.expected, result)
		})
	}
}

func TestStringsRemove(t *testing.T) {
	ss := []string{"a", "b", "c"}
	ns := arrutil.StringsRemove(ss, "b")

	assert.Contains(t, ns, "a")
	assert.NotContains(t, ns, "b")
	assert.Len(t, ns, 2)
}

func TestStringsFilter(t *testing.T) {
	is := assert.New(t)

	ss := arrutil.StringsFilter([]string{"a", "", "b", ""})
	is.Eq([]string{"a", "b"}, ss)
}

func TestStringsMap(t *testing.T) {
	is := assert.New(t)

	ss := arrutil.StringsMap([]string{"a", "b", "c"}, func(s string) string {
		return s + "1"
	})
	is.Eq([]string{"a1", "b1", "c1"}, ss)
}

func TestTrimStrings(t *testing.T) {
	is := assert.New(t)

	// TrimStrings
	ss := arrutil.TrimStrings([]string{" a", "b ", " c "})
	is.Eq("[a b c]", fmt.Sprint(ss))
	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
	is.Eq("[a b c]", fmt.Sprint(ss))
	ss = arrutil.TrimStrings([]string{",a", "b.", ",.c,"}, ",", ".")
	is.Eq("[a b c]", fmt.Sprint(ss))
}
