package comfunc_test

import (
	"testing"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSplitLineToKv(t *testing.T) {
	tests := []struct {
		line string
		k, v string
	}{
		{"key=val", "key", "val"},
		{"key = val ", "key", "val"},
		{"key =val\n", "key", "val"},
		{"key= val\r\n", "key", "val"},
		{" key=val\r", "key", "val"},
		{"key=val\t ", "key", "val"},
		{" key=val\t\n", "key", "val"},
		{"key=val\t\r\n", "key", "val"},
		{"key = val\nue", "key", "val\nue"},
		{" key-one =val ", "key-one", "val"},
		{" key_one = val", "key_one", "val"},
		{" valid=", "valid", ""},
		// invalid input
		{"invalid", "", ""},
		{"=invalid", "", ""},
		{" = invalid", "", ""},
		{"  ", "", ""},
	}

	for _, tt := range tests {
		k, v := comfunc.SplitLineToKv(tt.line, "=")
		assert.Eq(t, tt.k, k)
		assert.Eq(t, tt.v, v)
	}
}