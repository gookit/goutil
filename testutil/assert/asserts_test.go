package assert_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
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

	assert.Eq(t, "abc", str)
	assert.Panics(t, func() {
		panic("hh")
	})
}
