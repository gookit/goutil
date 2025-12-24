package netutil

import "net"

// IncrIP 将IP地址递增1 eg: 192.168.1.1 -> 192.168.1.2
func IncrIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

