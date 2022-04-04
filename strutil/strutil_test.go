package strutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := strutil.Similarity("hello", "he", 0.3)
	is.True(ok)
}

func TestSplit(t *testing.T) {
	ss := strutil.Split("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitValid("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitN("a, , b,c", ",", 3)
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitN("a, , b,c", ",", 2)
	assert.Equal(t, `[]string{"a", "b,c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.Split(" ", ",")
	assert.Nil(t, ss)
}

func TestSplitTrimmed(t *testing.T) {
	ss := strutil.SplitTrimmed("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "", "b", "c"}`, fmt.Sprintf("%#v", ss))

	ss = strutil.SplitNTrimmed("a, , b,c", ",", 2)
	assert.Equal(t, `[]string{"a", ", b,c"}`, fmt.Sprintf("%#v", ss))
}

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", strutil.Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", strutil.Substr("abcDef", 2, 2))
	assert.Equal(t, "cDef", strutil.Substr("abcDef", 2, 0))
	assert.Equal(t, "", strutil.Substr("abcDEF", 23, 5))
	assert.Equal(t, "cDEF12", strutil.Substr("abcDEF123", 2, -1))
	assert.Equal(t, "cDEF", strutil.Substr("abcDEF123", 2, -3))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, "aaa", strutil.Repeat("a", 3))
	assert.Equal(t, "D", strutil.Repeat("D", 1))
	assert.Equal(t, "D", strutil.Repeat("D", 0))
	assert.Equal(t, "D", strutil.Repeat("D", -3))
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
		assert.Equal(t, tt.want, strutil.RepeatRune(tt.give, tt.times))
	}
}

func TestReplaces(t *testing.T) {
	assert.Equal(t, "tom age is 20", strutil.Replaces(
		"{name} age is {age}",
		map[string]string{
			"{name}": "tom",
			"{age}":  "20",
		}))
}

func TestPadding(t *testing.T) {
	tests := []struct {
		want, give, pad string
		len             int
		pos             uint8
	}{
		{"ab000", "ab", "0", 5, strutil.PosRight},
		{"000ab", "ab", "0", 5, strutil.PosLeft},
		{"ab012", "ab012", "0", 4, strutil.PosLeft},
		{"ab   ", "ab", "", 5, strutil.PosRight},
		{"   ab", "ab", "", 5, strutil.PosLeft},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.Padding(tt.give, tt.pad, tt.len, tt.pos))

		if tt.pos == strutil.PosRight {
			assert.Equal(t, tt.want, strutil.PadRight(tt.give, tt.pad, tt.len))
		} else {
			assert.Equal(t, tt.want, strutil.PadLeft(tt.give, tt.pad, tt.len))
		}
	}
}

func TestPrettyJSON(t *testing.T) {
	tests := []interface{}{
		map[string]int{"a": 1},
		struct {
			A int `json:"a"`
		}{1},
	}
	want := `{
    "a": 1
}`
	for _, sample := range tests {
		got, err := strutil.PrettyJSON(sample)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	}
}
