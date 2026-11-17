package strutil_test

import (
	"testing"
	"time"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/assert"
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
}

// A "Z" is the UTC designator, with or without fractional seconds, no matter which
// location is given. See https://github.com/gookit/goutil/issues/321
func TestToTimeRFC3339ZIsUTC(t *testing.T) {
	loc := time.FixedZone("local", 3*60*60)

	want := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	cases := map[string]time.Time{
		"2026-01-02T03:04:05Z":   want,
		"2026-01-02T03:04:05.5Z": want.Add(500 * time.Millisecond),
		"2026-01-02 03:04:05Z":   want,
	}
	for input, want := range cases {
		got, err := strutil.ToTimeIn(input, loc)
		assert.NoErr(t, err, "input %s", input)
		assert.True(t, got.Equal(want), "input %s => got %s want %s", input, got, want)
	}

	// without a zone the given location is still used
	got, err := strutil.ToTimeIn("2026-01-02T03:04:05", loc)
	assert.NoErr(t, err)
	assert.True(t, got.Equal(time.Date(2026, 1, 2, 3, 4, 5, 0, loc)))
}
