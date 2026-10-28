package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/assert"
)

func TestRuneWidth(t *testing.T) {
	assert.Eq(t, 3, len("你"))
	assert.Eq(t, 1, len("\n"))
	assert.Eq(t, 2, strutil.RuneWidth('你'))
	assert.Eq(t, 1, strutil.RuneWidth('a'))
	assert.Eq(t, 0, strutil.RuneWidth('\n'))

	// fullwidth punctuation and Hangul are wide
	assert.Eq(t, 2, strutil.RuneWidth('，'))
	assert.Eq(t, 2, strutil.RuneWidth('한'))
	// combining marks and format characters take no column
	assert.Eq(t, 0, strutil.RuneWidth('\u0301'))
	assert.Eq(t, 0, strutil.RuneWidth('\u200d'))
	// ambiguous runes render as one column in common terminals
	assert.Eq(t, 1, strutil.RuneWidth('é'))
	assert.Eq(t, 1, strutil.RuneWidth('Ж'))
	assert.Eq(t, 1, strutil.RuneWidth('─'))
	// neutral runes (e.g. Arabic) take one column
	assert.Eq(t, 1, strutil.RuneWidth('ع'))
	assert.Eq(t, 0, strutil.RuneWidth('\t'))
	// "é" written as e + combining acute, then a wide rune
	assert.Eq(t, 3, strutil.TextWidth("e\u0301你"))
}

func TestUtf8Len(t *testing.T) {
	str := "Hello, 世界"

	assert.Eq(t, 7, len("Hello, "))
	assert.Eq(t, 13, len(str))
	assert.Eq(t, 9, strutil.RuneCount(str))
	assert.Eq(t, 9, strutil.Utf8len(str))
	assert.Eq(t, 9, strutil.Utf8Len(str))
	assert.Eq(t, 11, strutil.TextWidth(str))
	assert.Eq(t, 11, strutil.Utf8Width(str))
	assert.True(t, strutil.IsValidUtf8(str))
}

func TestUtf8Width(t *testing.T) {
	assert.Eq(t, 0, strutil.TextWidth(""))

	// case from bun.js
	// assert.Eq(t, 2, strutil.TextWidth("🇺🇸"))     // Now: 2 (was: 1) - flag emoji
	// assert.Eq(t, 2, strutil.TextWidth("👋🏽"))     // Now: 2 (was: 4) - emoji + skin tone
	// assert.Eq(t, 2, strutil.TextWidth("👨‍👩‍👧"))  // Now: 2 (was: 8) - ZWJ family sequence
	// assert.Eq(t, 0, strutil.TextWidth("\u2060")) // Now: 0 (was: 1) - word joiner
}

func TestUtf8Truncate(t *testing.T) {
	s := "hello 你好, world 世界"
	assert.Eq(t, "hello 你好", strutil.Truncate(s, 10, ""))
	assert.Eq(t, "hello ...", strutil.TextTruncate(s, 10, "..."))
	assert.Eq(t, "hello 你好", strutil.TextTruncate("hello 你好", 20, "..."))
}

func TestUtf8Split(t *testing.T) {
	s := "hello 你好, world 世界"
	assert.Eq(t, []string{"hello ", "你好, ", "world ", "世界"}, strutil.TextSplit(s, 6))
	assert.Eq(t, []string{"hello 你好"}, strutil.TextSplit("hello 你好", 10))
}

func TestWidthWrap(t *testing.T) {
	s := "hello 你好, world 世界"
	assert.Eq(t, "hello \n你好, \nworld \n世界", strutil.TextWrap(s, 6))
	s = "hello, world"
	assert.Eq(t, "hel\nlo,\n wo\nrld", strutil.TextWrap(s, 3))
}

func TestWordWrap(t *testing.T) {
	s := "hello 你好, world 世界"
	assert.Eq(t, "hello\n你好,\nworld\n世界", strutil.WordWrap(s, 6))
	s = "hello, world"
	assert.Eq(t, "hello,\nworld", strutil.WordWrap(s, 3))
}

func TestWordWrap_preservesWordSeparators(t *testing.T) {
	tests := []struct {
		name  string
		input string
		width int
		want  string
	}{
		{"unwrapped words", "hello world", 80, "hello world"},
		{"exact width", "alpha beta", 10, "alpha beta"},
		{"separator exceeds width", "alpha beta", 9, "alpha\nbeta"},
		{"multiple lines", "one two three", 6, "one\ntwo\nthree"},
		{"wide runes", "你好 world", 10, "你好 world"},
		{"combining rune", "a\u0301 b", 3, "a\u0301 b"},
		{"zero width combining token", "\u0301 a", 1, "\u0301\na"},
		{"zero width format token", "\u200b a", 1, "\u200b\na"},
		{"zero width control token", "\t a", 1, "\t\na"},
		{"existing newline without extra blank line", "\n a", 1, "\na"},
		{"multiple spaces", "one  two", 20, "one  two"},
		{"surrounding spaces", " one two ", 20, " one two "},
		{"only spaces", "  ", 5, "  "},
		{"existing newline without wrap", "one\ntwo three", 80, "one\ntwo three"},
		{"single word", "hello", 80, "hello"},
		{"long words", "hello, world", 3, "hello,\nworld"},
		{"existing wide wrap", "hello 你好, world 世界", 6, "hello\n你好,\nworld\n世界"},
		{"empty", "", 5, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Eq(t, tt.want, strutil.WordWrap(tt.input, tt.width))
		})
	}
}
