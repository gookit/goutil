package assert_test

import (
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/testutil/assert"
	assert2 "github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	err := errorx.Raw("this is a error")
	// assert2.EqualError(t, err, "user custom message")
	assert.NoErr(t, err, "user custom message")
	assert.ErrMsg(t, err, "")
}

func TestEq(t *testing.T) {
	str := "abc"
	assert2.Equal(t, "def", str)
	assert2.Empty(t, str)
	assert2.Len(t, str, 2)
	assert2.NotEqual(t, "def", str)
	assert.Eq(t, "def", str)
}
