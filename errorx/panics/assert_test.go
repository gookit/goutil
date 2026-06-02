package panics_test

import (
	"testing"

	"github.com/gookit/goutil/errorx/panics"
	"github.com/gookit/goutil/x/assert"
)

func TestIsTrue(t *testing.T) {
	assert.Panics(t, func() {
		panics.IsTrue(false)
	})
}
