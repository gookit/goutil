package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestPadding(t *testing.T) {
	tests := []struct {
		want, give, pad string
		len             int
		pos             strutil.PosFlag
	}{
		{"ab000", "ab", "0", 5, strutil.PosRight},
		{"000ab", "ab", "0", 5, strutil.PosLeft},
		{"ab012", "ab012", "0", 4, strutil.PosLeft},
		{"ab   ", "ab", "", 5, strutil.PosRight},
		{"   ab", "ab", "", 5, strutil.PosLeft},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, strutil.Padding(tt.give, tt.pad, tt.len, tt.pos))

		if tt.pos == strutil.PosRight {
			assert.Eq(t, tt.want, strutil.PadRight(tt.give, tt.pad, tt.len))
		} else {
			assert.Eq(t, tt.want, strutil.PadLeft(tt.give, tt.pad, tt.len))
		}
	}
}

func TestResize(t *testing.T) {
	tests := []struct {
		want, give string
		len        int
		pos        strutil.PosFlag
	}{
		{"ab   ", "ab", 5, strutil.PosRight},
		{"   ab", "ab", 5, strutil.PosLeft},
		{"ab012", "ab012", 5, strutil.PosLeft},
		{"ab", "ab", 2, strutil.PosLeft},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, strutil.Resize(tt.give, tt.len, tt.pos))
	}

	want := "       title        "
	assert.Eq(t, want, strutil.Resize("title", 20, strutil.PosMiddle))
}

func TestRepeat(t *testing.T) {
	assert.Eq(t, "aaa", strutil.Repeat("a", 3))
	assert.Eq(t, "DD", strutil.Repeat("D", 2))
	assert.Eq(t, "D", strutil.Repeat("D", 1))
	assert.Eq(t, "", strutil.Repeat("A", 0))
	assert.Eq(t, "", strutil.Repeat("D", -3))
}

func TestRepeatRune(t *testing.T) {
	tests := []struct {
		want  []rune
		give  rune
		times int
	}{
		{[]rune("bbb"), 'b', 3},
		{[]rune("..."), '.', 3},
		{[]rune("  "), ' ', 2},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.want, strutil.RepeatRune(tt.give, tt.times))
	}
}

func TestRepeatBytes(t *testing.T) {
	assert.Eq(t, []byte("aaa"), strutil.RepeatBytes('a', 3))
	assert.Eq(t, []byte{}, strutil.RepeatBytes('a', 0))
	assert.Eq(t, []byte{}, strutil.RepeatBytes(' ', 0))
	assert.Eq(t, []byte{}, strutil.RepeatBytes(' ', -2))
	assert.Eq(t, []byte{' '}, strutil.RepeatBytes(' ', 1))
	assert.Len(t, strutil.RepeatBytes(' ', 20), 20)
}

func TestPadChars(t *testing.T) {
	tests := []struct {
		wt  []byte
		ls  []byte
		pad byte
		pln int
	}{
		{
			[]byte("aaaabc"), []byte("abc"), 'a', 6,
		},
		{
			[]byte("abc"), []byte("abcd"), 'a', 3,
		},
	}
	for _, item := range tests {
		assert.Eq(t, item.wt, strutil.PadChars(item.ls, item.pad, item.pln, strutil.PosLeft))
		assert.Eq(t, item.wt, strutil.PadBytes(item.ls, item.pad, item.pln, strutil.PosLeft))
		assert.Eq(t, item.wt, strutil.PadBytesLeft(item.ls, item.pad, item.pln))
	}

	tests2 := []struct {
		wt  []byte
		ls  []byte
		pad byte
		pln int
	}{
		{
			[]byte("abcaaa"), []byte("abc"), 'a', 6,
		},
		{
			[]byte("abc"), []byte("abcd"), 'a', 3,
		},
	}
	for _, item := range tests2 {
		assert.Eq(t, item.wt, strutil.PadChars(item.ls, item.pad, item.pln, strutil.PosRight))
		assert.Eq(t, item.wt, strutil.PadBytes(item.ls, item.pad, item.pln, strutil.PosRight))
		assert.Eq(t, item.wt, strutil.PadBytesRight(item.ls, item.pad, item.pln))
	}
}

func TestPadRunes(t *testing.T) {
	assert.Eq(t, []rune("hi123aaa"), strutil.PadRunesRight([]rune("hi123"), 'a', 8))
	assert.Eq(t, []rune("aaahi123"), strutil.PadRunesLeft([]rune("hi123"), 'a', 8))
	assert.Eq(t, []rune("hi123aaa"), strutil.PadRunes([]rune("hi123"), 'a', 8, strutil.PosRight))
}

// runtime error: make slice: cap out of range #76
// https://github.com/gookit/goutil/issues/76
func TestIssues_76(t *testing.T) {
	// TODO
}
