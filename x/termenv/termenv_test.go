package termenv_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/termenv"
)

func TestBasicFunc(t *testing.T) {
	assert.Nil(t, termenv.LastErr())

	if termenv.NoColor() {
		assert.True(t, termenv.NoColor())
	} else {
		assert.False(t, termenv.NoColor())
	}

	termenv.SetDebugMode(true)
	termenv.SetDebugMode(false)
}
