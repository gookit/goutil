package ccolor_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/ccolor"
)

func TestCommon(t *testing.T) {
	assert.Nil(t, ccolor.LastErr())

	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()
	assert.Gt(t, ccolor.Level(), 0)
}
