package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/assert"
)

func TestBaseConvZero(t *testing.T) {
	for _, tt := range []struct{ name, got, want string }{
		{"base conversion", strutil.BaseConv("0", 10, 16), "0"},
		{"decimal conversion", strutil.Base10Conv("0", 64), "0"},
		{"integer conversion", strutil.BaseConvInt(0, 64), "0"},
		{"custom alphabet", strutil.BaseConvIntByTpl(0, "ab"), "a"},
		{"empty alphabet", strutil.BaseConvIntByTpl(0, ""), ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Eq(t, tt.want, tt.got)
		})
	}
}
