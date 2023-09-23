// Package netutil provide some network util functions.
package netutil

import (
	"net"
	"net/netip"
	"os"
)

// InternalIPv1 get internal IP buy old logic
func InternalIPv1() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("Oops: " + err.Error())
	}

	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				return
			}
		}
	}
	return
}

// GetLocalIPs get local IPs, will panic if error.
func GetLocalIPs() (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("get local IPs error: " + err.Error())
	}

	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			ips = append(ips, ipNet.IP.String())
		}
	}
	return
}

// InternalIP get internal IP for host.
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

// InternalIPv4 get internal IPv4 for host.
func InternalIPv4() string { return IPv4() }

// IPv4 get internal IPv4 for host.
func IPv4() (ip string) {
	addr := netip.IPv4Unspecified()
	if addr.IsValid() {
		return addr.String()
	}
	return ""
}

// InternalIPv6 get internal IPv6
func InternalIPv6() string { return IPv6() }

// IPv6 get internal IPv6
func IPv6() (ip string) {
	addr := netip.IPv6Unspecified()
	if addr.IsValid() {
		return addr.String()
	}
	return ""
}

// HostIP returns the IP addresses of the localhost.
func HostIP() ([]string, error) {
	name, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return net.LookupHost(name)
}

// FreePort returns a free port.
func FreePort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", addr); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}
