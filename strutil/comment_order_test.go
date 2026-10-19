package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSplitInlineCommentFirstMarker(t *testing.T) {
	for _, strict := range []bool{true, false} {
		t.Run(map[bool]string{true: "strict", false: "non-strict"}[strict], func(t *testing.T) {
			value, comment := strutil.SplitInlineComment("value // first # second", strict)
			assert.Eq(t, "value", value)
			assert.Eq(t, "// first # second", comment)
		})
	}
}
