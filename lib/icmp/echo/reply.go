package echo

import (
	"log"
	"net"
	"os"

	"strconv"

	"time"

	"github.com/hachi-n/pingo/lib/icmp"
	"github.com/hachi-n/pingo/lib/statistic"
)

func GetReply(conn *net.IPConn, sigc chan os.Signal, ipAddr string, identifier int) *statistic.Statistics {
	mtu := 1500
	mByte := make([]byte, mtu)
	stat := new(statistic.Statistics)

	for {
		select {

		case <-sigc:
			break

		default:
			n, err := conn.Read(mByte)
			if err != nil {
				log.Println("packet read error..")
				continue
			}

			receivedTime := time.Now()

			m, err := icmp.MessageUnmashal(mByte[:n])
			if err != nil {
				log.Fatal(err)
			}

			if m.Type == icmp.EchoReply && m.Identifier == uint16(identifier) {
				sendTime := time.Time{}
				if err := sendTime.UnmarshalBinary(m.Data); err != nil {
					log.Fatal(err)
				}
				rtt := receivedTime.Sub(sendTime)
				stat.RTTS = append(stat.RTTS, rtt)
				log.Println(ipAddr+":", "icmp_seq =", strconv.Itoa(int(m.Sequences)), "time =", rtt)
			}
		}
	}
	return stat
}
