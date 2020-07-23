package netutil_test

import (
	"testing"

	"github.com/gookit/goutil/netutil"
	"github.com/stretchr/testify/assert"
)

func TestInternalIP(t *testing.T) {
	assert.NotEmpty(t, netutil.InternalIP())
}
