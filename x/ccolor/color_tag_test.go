package ccolor_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/ccolor"
)

func TestCommon(t *testing.T) {
	assert.NotEmpty(t, ccolor.ColorTags())
	assert.True(t, ccolor.IsDefinedTag("info"))
}

func TestApplyTag(t *testing.T) {
	assert.Equal(t,"\x1b[0;32mMSG\x1b[0m", ccolor.ApplyTag("info", "MSG"))
}
