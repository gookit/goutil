package capp_test

import (
	"testing"

	"github.com/gookit/goutil/cflag/capp"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewCmd(t *testing.T) {
	c1 := capp.NewCmd("demo", "this is a demo command", capp.WrapRunFunc(func() error {
		return nil
	}))

	assert.NotNil(t, c1.Func)
}
