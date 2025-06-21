package ccolor_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/ccolor"
)

func TestStyleString(t *testing.T) {
	assert.Eq(t, "32", ccolor.Info.String())
}
