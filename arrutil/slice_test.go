package arrutil_test

import (
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
