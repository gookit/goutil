package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestValue_usage(t *testing.T) {
	s := strutil.Value("abc-123")
	assert.True(t, s.IsStartWith("abc"))
	assert.True(t, s.HasPrefix("abc"))
	assert.True(t, s.IsEndWith("123"))
	assert.True(t, s.HasSuffix("123"))

	assert.Equal(t, "abc-123", s.Val())
	assert.Equal(t, "abc-123", s.String())

	s1 := strutil.StrVal("abc-123")
	assert.NotEmpty(t, s1.Val())
	assert.Len(t, s1.Split("-"), 2)
	assert.Len(t, s1.SplitN("-", 2), 2)
}
