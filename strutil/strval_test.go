package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestStrVal(t *testing.T) {
	s := strutil.StrVal("abc-123")
	assert.True(t, s.IsStartWith("abc"))
	assert.True(t, s.HasPrefix("abc"))
	assert.True(t, s.IsEndWith("123"))
	assert.True(t, s.HasSuffix("123"))

	assert.Equal(t, "abc-123", s.Val())
	assert.Equal(t, "abc-123", s.String())

}
