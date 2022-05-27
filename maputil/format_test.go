package maputil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/stretchr/testify/assert"
)

func TestNewFormatter(t *testing.T) {
	mp := map[string]interface{}{"a": "v0", "b": 23}

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
