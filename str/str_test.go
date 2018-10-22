package str_test

import (
	"github.com/gookit/goutil/str"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimilarity(t *testing.T) {
	is := assert.New(t)
	_, ok := str.Similarity("hello", "he", 0.3)
	is.True(ok)
}