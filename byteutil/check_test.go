package byteutil_test

import (
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsNumChar(t *testing.T) {
	tests := []struct {
		args byte
		want bool
	}{
		{'2', true},
		{'a', false},
		{'+', false},
	}
	for _, tt := range tests {
		assert.Eq(t, tt.want, byteutil.IsNumChar(tt.args))
	}
}
