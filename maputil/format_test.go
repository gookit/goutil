package maputil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNewFormatter(t *testing.T) {
	mp := map[string]interface{}{"a": "v0", "b": 23}

	mf := maputil.NewFormatter(mp)
	assert.Contains(t, mf.String(), "b:23")

	buf := testutil.NewTestWriter()
	mf = maputil.NewFormatter(mp).WithFn(func(f *maputil.MapFormatter) {
		f.Indent = "   "
	})
	mf.FormatTo(buf)
	assert.Contains(t, buf.String(), "\n   ")
	fmt.Println(buf.String())

	s := maputil.FormatIndent(mp, "  ")
	fmt.Println(s)
	assert.Contains(t, s, "\n  ")

	s = maputil.FormatIndent(mp, "")
	fmt.Println(s)
	assert.NotContains(t, s, "\n  ")
}

func TestFormatIndent_mlevel(t *testing.T) {
	mp := map[string]interface{}{"a": "v0", "b": 23}

	mp["subs"] = map[string]string{
		"sub_k1": "sub val1",
		"sub_k2": "sub val2",
	}

	s := maputil.FormatIndent(mp, "")
	fmt.Println(s)
	assert.NotContains(t, s, "\n  ")

	s = maputil.FormatIndent(mp, "  ")
	fmt.Println(s)
}
