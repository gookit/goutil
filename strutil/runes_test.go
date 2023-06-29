package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestRuneWidth(t *testing.T) {
	assert.Eq(t, 3, len("你"))
	assert.Eq(t, 1, len("\n"))
	assert.Eq(t, 2, strutil.RuneWidth('你'))
	assert.Eq(t, 1, strutil.RuneWidth('a'))
	assert.Eq(t, 0, strutil.RuneWidth('\n'))
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
