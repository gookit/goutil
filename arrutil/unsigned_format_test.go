package arrutil_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/x/assert"
)

type unsignedFormatID uint64
type signedFormatID int64

func TestIntsToStringUnsignedRange(t *testing.T) {
	cases := []struct{ name, got, want string }{
		{"uint64 maximum", arrutil.IntsToString([]uint64{0, math.MaxInt64 + 1, math.MaxUint64}), "[0,9223372036854775808,18446744073709551615]"},
		{"defined unsigned", arrutil.IntsToString([]unsignedFormatID{math.MaxUint64}), "[18446744073709551615]"},
		{"uint maximum", arrutil.IntsToString([]uint{^uint(0)}), "[" + strconv.FormatUint(uint64(^uint(0)), 10) + "]"},
		{"uintptr maximum", arrutil.IntsToString([]uintptr{^uintptr(0)}), "[" + strconv.FormatUint(uint64(^uintptr(0)), 10) + "]"},
		{"signed bounds", arrutil.IntsToString([]int64{math.MinInt64, -1, 0, math.MaxInt64}), "[-9223372036854775808,-1,0,9223372036854775807]"},
		{"defined signed", arrutil.IntsToString([]signedFormatID{-10, 20}), "[-10,20]"},
		{"small unsigned", arrutil.IntsToString([]uint8{0, 255}), "[0,255]"},
		{"nil unsigned", arrutil.IntsToString([]uint64(nil)), "[]"},
		{"empty signed", arrutil.IntsToString([]int{}), "[]"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) { assert.Eq(t, tt.want, tt.got) })
	}
}
