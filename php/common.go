package php

import (
	"net"
	"strings"
)

type RemoteAddr struct {
	Ip   string
	Port string
}

func ResolveRemoteAddr(addr string) *RemoteAddr {
	ipPlusPort := strings.Split(addr, ":")
	port := ipPlusPort[len(ipPlusPort)-1:]
	portStr := strings.Join(port, "")

	// "[::1]:55083"
	if strings.Count(addr, ":") > 2 {
		// [::1]
		ipBlock := addr[0:len(portStr)]
		ipV6 := ipBlock[1 : len(ipBlock)-1]
		ip := net.ParseIP(ipV6)
		return &RemoteAddr{
			Ip:   ip.String(),
			Port: portStr,
		}
	}

	ip := net.ParseIP(ipPlusPort[0])
	return &RemoteAddr{
		Ip:   ip.String(),
		Port: portStr,
	}
}

// THis method was copied from Go std library !
// isIPv4 reports whether addr contains an IPv4 address.
func isIPv4(ip net.IP) bool {
	return ip.To4() != nil

}
