package arrutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/stretchr/testify/assert"
)

func TestNewFormatter(t *testing.T) {
	arr := [2]string{"a", "b"}
	str := arrutil.FormatIndent(arr, "  ")
	assert.Contains(t, str, "\n  ")
	fmt.Println(str)

	str = arrutil.FormatIndent(arr, "")
	assert.NotContains(t, str, "\n  ")
	assert.Equal(t, "[a, b]", str)
	fmt.Println(str)

	assert.Equal(t, "", arrutil.FormatIndent("invalid", ""))
	assert.Equal(t, "[]", arrutil.FormatIndent([]string{}, ""))
}
