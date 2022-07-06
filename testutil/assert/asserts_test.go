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
	assert.Err(t, err, "user custom message")
	assert.ErrMsg(t, err, "this is a error")
}

func TestContains(t *testing.T) {
	str := "abc+123"
	assert.StrContains(t, str, "123")
}

func TestEq(t *testing.T) {
	str := "abc"

	assert2.Equal(t, "abc", str)
	assert2.NotEmpty(t, str)
	assert2.Panics(t, func() {
		panic("hh")
	})
	assert2.Len(t, str, 3)
	assert2.NotEqual(t, "def", str)

	assert.Eq(t, "abc", str)
}
