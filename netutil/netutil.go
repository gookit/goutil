package netutil

import (
	"net"
)

// InternalIP get internal IP
func InternalIP() (ip string) {
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
