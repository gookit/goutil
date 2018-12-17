package strUtil_test

import (
	"fmt"
	"github.com/gookit/goutil/strUtil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := strUtil.Similarity("hello", "he", 0.3)
	is.True(ok)
}

func TestSplit(t *testing.T) {
	ss := strUtil.Split("a, , b,c", ",")
	assert.Equal(t, `[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))
}
