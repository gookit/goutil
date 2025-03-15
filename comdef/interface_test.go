package comdef_test

import (
	"testing"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMatchFunc_Match(t *testing.T) {
	fn := comdef.MatchFunc[int](func(v int) bool {
		return v > 5
	})

	input := 10
	result := fn.Match(input)
	assert.True(t, result, "Expected match to succeed for non-zero value")

	input = 0
	result = fn.Match(input)
	assert.False(t, result, "Expected match to fail for zero value")
}

func TestStringMatchFunc_Match(t *testing.T) {
	tests := []struct {
		name     string
		fn       comdef.StringMatchFunc
		input    string
		expected bool
	}{
		{
			name: "Match_ReturnsTrue",
			fn: func(s string) bool {
				return true
			},
			input:    "test",
			expected: true,
		},
		{
			name: "Match_ReturnsFalse",
			fn: func(s string) bool {
				return false
			},
			input:    "test",
			expected: false,
		},
		{
			name: "Match_WithPanic",
			fn: func(s string) bool {
				panic("test panic")
			},
			input:    "test",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					assert.Eq(t, tt.name, "Match_WithPanic", "Expected panic for test case: %s", tt.name)
				}
			}()

			result := tt.fn.Match(tt.input)
			assert.Equal(t, tt.expected, result, "Expected %v, got %v", tt.expected, result)
		})
	}
}

func TestStringHandleFunc_Handle(t *testing.T) {
	tests := []struct {
		name     string
		fn       comdef.StringHandleFunc
		input    string
		expected string
	}{
		{
			name: "Handle_ReturnsInput",
			fn: func(s string) string {
				return s + ".log"
			},
			input:    "test",
			expected: "test.log",
		},
		{
			name: "Handle_WithPanic",
			fn: func(s string) string {
				panic("test panic")
			},
			input: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					assert.Eq(t, tt.name, "Handle_WithPanic", "Expected panic for test case: %s", tt.name)
				}
			}()

			result := tt.fn.Handle(tt.input)
			assert.Equal(t, tt.expected, result, "Expected %v, got %v", tt.expected, result)
		})
	}
}
