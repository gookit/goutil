package arrutil_test

import (
	"bytes"
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

	sl := []string{"c", "d"}
	f := arrutil.NewFormatter(sl)
	f.WithFn(func(f *arrutil.ArrFormatter) {
		f.Indent = "  "
		f.ClosePrefix = "-"
	})
	str = f.String()
	assert.Eq(t, "[\n  c,\n  d\n-]", str)

	f = arrutil.NewFormatter(sl)
	w := &bytes.Buffer{}
	f.FormatTo(w)
	assert.Eq(t, "[c, d]", w.String())
}
