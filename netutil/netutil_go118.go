//go:build go1.18
// +build go1.18

package netutil

import "net/netip"

// InternalIP get internal IP
// func InternalIP() (ip string) {
// 	addr := netip.IPv4Unspecified()
// 	if addr.IsValid() {
// 		return addr.String()
// 	}
//
// 	addr = netip.IPv6Unspecified()
// 	if addr.IsValid() {
// 		return addr.String()
// 	}
//
// 	return ""
// }

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
