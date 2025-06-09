package termenv_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/termenv"
)

func TestCurrentShell(t *testing.T) {
	path := termenv.CurrentShell(true)

	if path != "" {
		assert.NotEmpty(t, path)
		assert.True(t, termenv.HasShellEnv(path))

		path = termenv.CurrentShell(false)
		assert.NotEmpty(t, path)
	}

	assert.True(t, termenv.IsShellSpecialVar('$'))
	assert.True(t, termenv.IsShellSpecialVar('@'))
	assert.False(t, termenv.IsShellSpecialVar('a'))
}
