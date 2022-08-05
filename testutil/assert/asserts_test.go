package assert_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	assert2 "github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	err := errors.New("this is a error")
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
	assert2.NotPanics(t, func() {})
	assert2.Len(t, str, 3)
	assert2.NotEqual(t, "def", str)
	assert2.Contains(t, str, "a")

	assert.Eq(t, "abc", str)
}
