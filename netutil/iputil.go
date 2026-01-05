package netutil

import (
	"fmt"
	"net"
	"os"
)

// AllLocalIPv4 Get all non-loop IPv4 addresses. 获取所有非回环 IPv4 地址
func AllLocalIPv4() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	ips := filterIpv4(addrs)

	if len(ips) == 0 {
		return nil, fmt.Errorf("no IPv4 addresses found")
	}
	return ips, nil
}

func filterIpv4(addrs []net.Addr) []string {
	var ips []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				ips = append(ips, ip4.String())
			}
		}
	}
	return ips
}

// GetLocalIPs get local IPs(ipv4+ipv6), will panic on error.
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

// InternalIP get local first internal IP v4/v6 addr.
func InternalIP() string { return getLocalIP(0) }

// InternalIPv4 get internal IPv4 for host.
func InternalIPv4() string { return IPv4() }

// IPv4 get local first internal IPv4 addr.
func IPv4() string { return getLocalIP(4) }

// InternalIPv6 get first internal IPv6 addr
func InternalIPv6() string { return IPv6() }

// IPv6 get local first internal IPv6 addr
func IPv6() string { return getLocalIP(6) }

// MustIPv4 get first internal IP v4 addr, will panic on error.
func MustIPv4() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic("MustIPv4: " + err.Error())
	}
	return getFirstIP(4, addrs)
}

// get local first internal IP v4/v6 addr.
//
// ver: 4, 6 or not limit
func getLocalIP(ver uint8) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	return getFirstIP(ver, addrs)
}

func getFirstIP(ver uint8, addrs []net.Addr) string {
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// need ip v4
			if ver == 4 {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					return ip4.String()
				}
			} else {
				return addr.String()
			}
		}
	}
	return ""
}

// HostIP returns the IP addresses of the local hostname.
func HostIP() ([]string, error) {
	name, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return net.LookupHost(name)
}

// IncrIP 将IP地址递增1 eg: 192.168.1.1 -> 192.168.1.2
func IncrIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
