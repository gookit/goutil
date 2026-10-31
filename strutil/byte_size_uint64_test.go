package strutil_test

import (
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/assert"
	"testing"
)

func TestToByteSizeUint64(t *testing.T) {
	for _, tc := range []struct {
		input     string
		want      uint64
		wantError bool
	}{
		{"0", 0, false},
		{"4294967295", 4294967295, false},
		{"4294967296", 4294967296, false},
		{"4294967297", 4294967297, false},
		{"18446744073709551615", 18446744073709551615, false},
		{"18446744073709551616", 0, true},
		{"-1", 0, true},
		{"invalid", 0, true},
	} {
		t.Run(tc.input, func(t *testing.T) {
			got, err := strutil.ToByteSize(tc.input)
			if tc.wantError {
				assert.Err(t, err)
				return
			}
			assert.NoErr(t, err)
			assert.Eq(t, tc.want, got)
		})
	}
	t.Run("safe byte size", func(t *testing.T) { assert.Eq(t, uint64(4294967296), strutil.SafeByteSize("4294967296")) })
	t.Run("size range", func(t *testing.T) {
		min, max, err := strutil.ParseSizeRange("4294967296~4294967297", nil)
		assert.NoErr(t, err)
		assert.Eq(t, uint64(4294967296), min)
		assert.Eq(t, uint64(4294967297), max)
	})
}
