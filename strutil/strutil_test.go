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
}

func TestUpperFirst(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{"a", "A"},
		{"", ""},
		{"ab", "Ab"},
		{"Ab", "Ab"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.UpperFirst(tt.give))
	}
}

func TestLowerFirst(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{"A", "a"},
		{"", ""},
		{"Ab", "ab"},
		{"ab", "ab"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.LowerFirst(tt.give))
	}
}

func TestUpperWord(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{"a", "A"},
		{"", ""},
		{"ab", "Ab"},
		{"Ab", "Ab"},
		{"hi lo", "Hi Lo"},
		{"hi lo wr", "Hi Lo Wr"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.UpperWord(tt.give))
	}
}
