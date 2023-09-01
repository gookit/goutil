package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBaseConv(t *testing.T) {
	tests := []struct {
		give string
		from int
		to   int
		want string
	}{
		// 10 -> 16
		{"123", 10, 16, "7b"},
		{"7b", 16, 10, "123"},
		// 10 -> 32
		{"123", 10, 32, "3r"},
		{"3r", 32, 10, "123"},
		// 10 -> 62
		{"123", 10, 62, "1Z"},
		{"1Z", 62, 10, "123"},
		// 10 -> 64
		{"123", 10, 64, "1X"},
		{"1X", 64, 10, "123"},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := strutil.BaseConv(tt.give, tt.from, tt.to)
			if got != tt.want {
				t.Errorf("BaseConv() got = %v, want %v; give=%s", got, tt.want, tt.give)
			}
		})
	}
}

func TestBase10Conv(t *testing.T) {
	// fmt.Println(time.Now().Format("20060102150405"))
	date := "20230829194900" // seconds
	t.Run("date sec", func(t *testing.T) {
		base16 := strutil.Base10Conv(date, 16)
		assert.Eq(t, "12665b633e94", base16)

		base32 := strutil.Base10Conv(date, 32)
		assert.Eq(t, "icpdm6fkk", base32)

		base62 := strutil.Base10Conv(date, 62)
		assert.Eq(t, "5KaR3zRW", base62)

		base64 := strutil.Base10Conv(date, 64)
		assert.Eq(t, "4CproPWk", base64)
	})

	// fmt.Println(time.Now().Format("20060102150405.000"))
	msDate := "20230829195105843"
	t.Run("date ms", func(t *testing.T) {
		base16 := strutil.Base10Conv(msDate, 16)
		assert.Eq(t, "47dfd4fbaf9633", base16)

		base32 := strutil.Base10Conv(msDate, 32)
		assert.Eq(t, "huvqjtqv5hj", base32)

		base62 := strutil.Base10Conv(msDate, 62)
		assert.Eq(t, "1uEL5LJptx", base62)

		base64 := strutil.Base10Conv(msDate, 64)
		assert.Eq(t, "17TZjXHVoP", base64)
	})

	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() {
			strutil.Base10Conv("invalid", 16)
		})
	})
}
