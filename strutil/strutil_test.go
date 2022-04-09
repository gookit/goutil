package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := strutil.Similarity("hello", "he", 0.3)
	is.True(ok)
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
