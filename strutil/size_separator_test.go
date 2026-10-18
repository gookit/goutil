package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestParseSizeRangeCustomSeparator(t *testing.T) {
	for _, tt := range []struct {
		expr     string
		min, max uint64
	}{{"1KB:2KB", 1024, 2048}, {":2KB", 0, 2048}, {"1KB:", 1024, 0}} {
		t.Run(tt.expr, func(t *testing.T) {
			min, max, err := strutil.ParseSizeRange(tt.expr, &strutil.ParseSizeOpt{SepChar: ':'})
			assert.NoErr(t, err)
			assert.Eq(t, tt.min, min)
			assert.Eq(t, tt.max, max)
		})
	}
}
