package assert_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestAssertions_Chain(t *testing.T) {
	// err := "error message"
	err := errors.New("error message")

	as := assert.New(t).
		NotEmpty(err).
		NotNil(err).
		Err(err).
		ErrMsg(err, "error message").
		Eq("error message", err.Error()).
		Neq("message", err.Error()).
		Equal("error message", err.Error()).
		Contains(err.Error(), "message").
		StrContains(err.Error(), "message").
		NotContains(err.Error(), "success").
		Gt(4, 3).
		Lt(2, 3)

	assert.True(t, as.IsOk())
	assert.False(t, as.IsFail())

	iv := 23
	as = assert.New(t).
		IsType(1, iv).
		NotEq(22, iv).
		NotEqual(22, iv).
		Lte(iv, 23).
		Gte(iv, 23).
		Empty(0).
		True(true).
		False(false).
		Nil(nil)

	assert.True(t, as.IsOk())
}
