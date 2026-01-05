package netutil_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/netutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestInternalIP(t *testing.T) {
	t.Run("InternalIP", func(t *testing.T) {
		ip := netutil.InternalIP()
		assert.NotEmpty(t, ip)
		dump.P(ip)
	})

	t.Run("InternalIPv4", func(t *testing.T) {
		assert.NotEmpty(t, netutil.MustIPv4())

		ipv4 := netutil.InternalIPv4()
		assert.NotEmpty(t, ipv4)
		dump.P(ipv4)
	})

	t.Run("InternalIPv6", func(t *testing.T) {
		ipv6 := netutil.InternalIPv6()
		assert.NotEmpty(t, ipv6)
		dump.P(ipv6)
	})

	t.Run("AllLocalIPv4", func(t *testing.T) {
		ips, err := netutil.AllLocalIPv4()
		assert.NoError(t, err)
		assert.NotEmpty(t, ips)
		dump.P(ips)
	})

	t.Run("GetLocalIPs", func(t *testing.T) {
		ips := netutil.GetLocalIPs()
		assert.NotEmpty(t, ips)
		dump.P(ips)
	})
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

func TestIncrIP(t *testing.T) {
	ip := net.IP{192, 168, 1, 1}
	netutil.IncrIP(ip)
	assert.Eq(t, "192.168.1.2", ip.String())
}
