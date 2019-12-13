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

	ss = strutil.Split(" ", ",")
	assert.Nil(t, ss)
}

func TestSubstr(t *testing.T) {
	assert.Equal(t, "abc", strutil.Substr("abcDef", 0, 3))
	assert.Equal(t, "cD", strutil.Substr("abcDef", 2, 2))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, "aaa", strutil.Repeat("a", 3))
	assert.Equal(t, "D", strutil.Repeat("D", 1))
	assert.Equal(t, "D", strutil.Repeat("D", 0))
	assert.Equal(t, "D", strutil.Repeat("D", -3))
}

func TestPadding(t *testing.T) {
	tests := []struct{
		want, give, pad string
		len int
		pos uint8
	} {
		{"ab000", "ab", "0", 5, strutil.PosRight},
		{"000ab", "ab", "0", 5, strutil.PosLeft},
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
