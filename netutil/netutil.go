// Package netutil provide some network util functions.
package netutil

import (
	"net"
)

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

// AllMacAddrs get all mac addresses
func AllMacAddrs() ([]string, error) {
	var macAddrs []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return macAddrs, err
	}

	for _, iface := range interfaces {
		// 跳过回环接口（如 lo）和未启用的接口
		if iface.Flags&net.FlagLoopback != 0 || iface.HardwareAddr == nil {
			continue
		}

		mac := iface.HardwareAddr.String()
		if mac != "" {
			macAddrs = append(macAddrs, mac)
		}
	}
	return macAddrs, nil
}