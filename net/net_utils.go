package net

import (
	"net"
	"os"
)

// InternalIP
func InternalIP() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error())
		os.Exit(1)
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

	// os.Exit(0)
	return
}
