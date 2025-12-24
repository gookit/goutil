package netutil_test

import (
	"net"
	"testing"

	"github.com/gookit/goutil/netutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIncrIP(t *testing.T) {
	ip := net.IP{192, 168, 1, 1}
	netutil.IncrIP(ip)
	assert.Eq(t, "192.168.1.2", ip.String())
}
