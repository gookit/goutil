package panics_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/panics"
)

func TestIsTrue(t *testing.T) {
	assert.Panics(t, func() {
		panics.IsTrue(false)
	})
}
