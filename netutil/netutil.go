// Package netutil provide some network util functions.
package netutil

import (
	"net"
	"net/netip"
)

// InternalIPOld get internal IP buy old logic
func InternalIPOld() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("Oops: " + err.Error())
	}

	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				// os.Stdout.WriteString(ipNet.IP.String() + "\n")
				ip = ipNet.IP.String()
				return
			}
		}
	}
	return
}

// InternalIP get internal IP
func InternalIP() (ip string) {
	addr := netip.IPv4Unspecified()
	if addr.IsValid() {
		return addr.String()
	}

	addr = netip.IPv6Unspecified()
	if addr.IsValid() {
		return addr.String()
	}
	return ""
}

// InternalIPv4 get internal IPv4
func InternalIPv4() (ip string) {
	addr := netip.IPv4Unspecified()
	if addr.IsValid() {
		return addr.String()
	}
	return ""
}

// InternalIPv6 get internal IPv6
func InternalIPv6() (ip string) {
	addr := netip.IPv6Unspecified()
	if addr.IsValid() {
		return addr.String()
	}
	return ""
}
