package resolv

import (
	"log"
	"net"
)

func GetIPv4ByName(domain string) *net.IPAddr {
	ipaddr, err := net.ResolveIPAddr("ip4", domain)
	if err != nil {
		log.Fatal(err)
	}

	return ipaddr
}
