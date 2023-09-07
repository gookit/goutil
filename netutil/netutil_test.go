package netutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/netutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestInternalIP(t *testing.T) {
	assert.NotEmpty(t, netutil.InternalIPv1())
	assert.NotEmpty(t, netutil.InternalIP())
	assert.NotEmpty(t, netutil.InternalIPv4())
	assert.NotEmpty(t, netutil.InternalIPv6())
	assert.NotEmpty(t, netutil.GetLocalIPs())
}

func TestHostIP(t *testing.T) {
	addrs, err := netutil.HostIP()
	if err != nil {
		t.Skip("skip test for error: " + err.Error())
		return
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, addrs)
	fmt.Println(addrs)
}

func TestFreePort(t *testing.T) {
	port, err := netutil.FreePort()
	assert.NoError(t, err)
	assert.Gt(t, port, 0)
	assert.Lt(t, port, 65536)
}
