package errorx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gookit/goutil/errorx"
)

func TestNew(t *testing.T) {
	err := errorx.New("error message")
	// dump.V(err)
	fmt.Println(err)

	// err = errox.With(err)
}

func TestWrap(t *testing.T) {
	err := errors.New("first error message")
	fmt.Println(err)
	fmt.Println("----------------------------------")

	err = errorx.Wrap(err, "second error message")
	fmt.Println(err)
	fmt.Println("----------------------------------")

	err = errorx.Wrap(err, "third error message")
	fmt.Println(err)
}
