package termenv_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/termenv"
)

func TestLastErr(t *testing.T) {
	assert.Nil(t, termenv.LastErr())
}
