package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewFormatter(t *testing.T) {
	arr := [2]string{"a", "b"}
	str := arrutil.FormatIndent(arr, "  ")
	assert.Contains(t, str, "\n  ")
	fmt.Println(str)

	str = arrutil.FormatIndent(arr, "")
	assert.NotContains(t, str, "\n  ")
	assert.Eq(t, "[a, b]", str)
	fmt.Println(str)

	assert.Eq(t, "", arrutil.FormatIndent("invalid", ""))
	assert.Eq(t, "[]", arrutil.FormatIndent([]string{}, ""))
}
