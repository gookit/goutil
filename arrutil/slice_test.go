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
