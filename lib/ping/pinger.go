package ping

import (
	"log"
	"net"
	"os"

	"github.com/hachi-n/pingo/lib/icmp/echo"
	"github.com/hachi-n/pingo/lib/resolv"
)

func Do(hostname string) {
	ip := resolv.GetIPv4ByName(hostname)

	conn, err := net.DialIP("ip4:icmp", nil, ip)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Println("PING", hostname, "("+ip.String()+")")

	sigc := make(chan os.Signal, 1)

	//Fix me...
	identifier := os.Getpid()

	go echo.SendRequest(conn, sigc, identifier)
	result := echo.GetReply(conn, sigc, ip.String(), identifier)

	result.Dump()
}
