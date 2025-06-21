package termenv_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/termenv"
)

func TestCommon(t *testing.T) {
	termenv.ForceEnableColor()
	defer termenv.RevertColorSupport()

	assert.True(t, termenv.IsSupportColor())
	assert.True(t, termenv.IsSupport256Color())
	assert.False(t, termenv.NoColor())
	assert.False(t, termenv.IsSupportTrueColor())
}
