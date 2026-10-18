package strutil_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestToTimeRFC3339Fractions(t *testing.T) {
	for _, input := range []string{"2026-01-02T03:04:05.1Z", "2026-01-02T03:04:05.123Z", "2026-01-02T03:04:05.123456789+01:00"} {
		t.Run(input, func(t *testing.T) {
			want, err := time.Parse(time.RFC3339Nano, input)
			assert.NoErr(t, err)
			got, err := strutil.ToTimeIn(input, time.UTC)
			assert.NoErr(t, err)
			assert.True(t, got.Equal(want))
		})
	}
	loc := time.FixedZone("local", 3*60*60)
	got, err := strutil.ToTimeIn("2026-01-02T03:04:05Z", loc)
	assert.NoErr(t, err)
	assert.True(t, got.Equal(time.Date(2026, 1, 2, 3, 4, 5, 0, loc)))
}
