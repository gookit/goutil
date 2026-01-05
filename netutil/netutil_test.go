package netutil_test

import (
	"testing"

	"github.com/gookit/goutil/netutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFreePort(t *testing.T) {
	port, err := netutil.FreePort()
	assert.NoError(t, err)
	assert.Gt(t, port, 0)
	assert.Lt(t, port, 65536)
}
