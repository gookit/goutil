package strUtil_test

import (
	"github.com/gookit/goutil/strUtil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := strUtil.Similarity("hello", "he", 0.3)
	is.True(ok)
}
