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

func TestIsAlphaChar(t *testing.T) {
	tests := []struct {
		name     string
		input    byte
		expected bool
	}{
		{"Lowercase letter", 'a', true},
		{"Uppercase letter", 'Z', true},
		{"Digit", '5', false},
		{"Special character", '@', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := byteutil.IsAlphaChar(tt.input)
			assert.Eq(t, tt.expected, result)
		})
	}
}