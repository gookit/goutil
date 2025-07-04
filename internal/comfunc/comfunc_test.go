package comfunc_test

import (
	"testing"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/testutil/assert"
)

func TestEnviron(t *testing.T) {
	assert.NotEmpty(t, comfunc.Environ())
}
