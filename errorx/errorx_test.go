package errorx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := errorx.New("error message")
	// dump.V(err)
	fmt.Println(err)
	// err = errorx.With(err)
}

func TestWrap(t *testing.T) {
	err := errors.New("first error message")
	assert.Nil(t, errorx.Unwrap(err))
	assert.Equal(t, "first error message", errorx.Cause(err).Error())

	fmt.Println(err)
	fmt.Println("----------------------------------")

	err = errorx.Wrap(err, "second error message")
	fmt.Println(err)
	fmt.Println("----------------------------------")

	err = errorx.Wrap(err, "third error message")
	fmt.Println(err)

	assert.Equal(t, "first error message", errorx.Cause(err).Error())
	assert.Equal(t, "second error message", errorx.Unwrap(err).Error())
}
