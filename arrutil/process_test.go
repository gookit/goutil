package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestReverse(t *testing.T) {
	ss := []string{"a", "b", "c"}
	arrutil.Reverse(ss)
	assert.Eq(t, []string{"c", "b", "a"}, ss)

	ints := []int{1, 2, 3}
	arrutil.Reverse(ints)
	assert.Eq(t, []int{3, 2, 1}, ints)
}

func TestRemove(t *testing.T) {
	ss := []string{"a", "b", "c"}
	ns := arrutil.Remove(ss, "b")
	assert.Eq(t, []string{"a", "c"}, ns)

	ints := []int{1, 2, 3}
	ni := arrutil.Remove(ints, 2)
	assert.Eq(t, []int{1, 3}, ni)
}

func TestFilter(t *testing.T) {
	is := assert.New(t)
	ss := arrutil.Filter([]string{"a", "", "b", ""})
	is.Eq([]string{"a", "b"}, ss)
}
