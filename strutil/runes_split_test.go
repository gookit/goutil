package strutil_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/assert"
)

func TestUtf8Split_invalidWidth(t *testing.T) {
	// a width below 1 could not be split and used to panic (divide by zero / a negative
	// make capacity), the whole string is one item now.
	for _, w := range []int{0, -1, -10} {
		t.Run("w"+strconv.Itoa(w), func(t *testing.T) {
			for _, split := range []func(string, int) []string{strutil.Utf8Split, strutil.TextSplit} {
				assert.Eq(t, []string{"hello 你好"}, split("hello 你好", w))
				assert.Eq(t, []string{""}, split("", w))
			}
		})
	}
}

func TestUtf8SplitPreservesZeroWidthRunes(t *testing.T) {
	tests := []struct {
		name, input string
		width       int
		want        []string
	}{
		{"ASCII control", "abcd", 2, []string{"ab", "cd"}},
		{"combining marks", "a\u0301b\u0301", 1, []string{"a\u0301", "b\u0301"}},
		{"trailing combining mark", "ab\u0301", 1, []string{"a", "b\u0301"}},
		{"variation selectors", "a\ufe0fb\ufe0f", 1, []string{"a\ufe0f", "b\ufe0f"}},
		{"format rune", "ab\u200d", 1, []string{"a", "b\u200d"}},
		{"wide bases", "你\u0301好\u0301", 2, []string{"你\u0301", "好\u0301"}},
		{"oversized bases with marks", "你\u0301好\u0301", 1, []string{"你\u0301", "好\u0301"}},
		{"partial final chunk", "abc\u0301", 2, []string{"ab", "c\u0301"}},
		{"only zero width", "\u0301\u200d", 1, []string{"\u0301\u200d"}},
		{"one rune exceeds width", "你好", 1, []string{"你", "好"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit := tt.width
			for _, r := range tt.input {
				if width := strutil.RuneWidth(r); width > limit {
					limit = width
				}
			}
			for _, split := range []func(string, int) []string{strutil.Utf8Split, strutil.TextSplit} {
				got := split(tt.input, tt.width)
				assert.Eq(t, tt.want, got)
				assert.Eq(t, tt.input, strings.Join(got, ""))
				for _, part := range got {
					assert.NotEmpty(t, part)
					// A wide base may exceed the target, including its zero-width marks.
					assert.True(t, strutil.Utf8Width(part) <= limit)
				}
			}
		})
	}
}
