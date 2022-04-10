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
	assert.Error(t, err)

	fmt.Println(err)
	fmt.Printf("%v\n", err)
	fmt.Printf("%#v\n", err)
	fmt.Printf("%+v\n", err)

	err = errorx.Newf("error %s", "message")
	assert.Error(t, err)
}

func TestWithPrev(t *testing.T) {
	err1 := errorx.New("first error message")
	assert.Error(t, err1)

	err2 := errorx.WithPrev(err1, "second error message")
	assert.Error(t, err2)
	assert.True(t, errorx.Has(err2, err1))

	// fmt.Println(err2)
	// fmt.Printf("%#v\n", err2)
	// fmt.Printf("%+v\n", err2)
	fmt.Printf("%v\n", err2)
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
