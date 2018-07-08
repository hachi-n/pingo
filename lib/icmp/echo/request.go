package echo

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/hachi-n/pingo/lib/icmp"
)

func SendRequest(conn *net.IPConn, sigc chan os.Signal, identifier int) {
	ticker := time.NewTicker(1 * time.Second)
	seq := 0

	for {
		select {
		case <-sigc:
			break

		case <-ticker.C:
			data, err := time.Now().MarshalBinary()
			if err != nil {
				log.Fatal(err)
			}

			m := &icmp.Message{
				Type:       icmp.EchoRequest,
				Code:       0,
				Checksum:   0,
				Identifier: uint16(identifier),
				Sequences:  uint16(seq),
				Data:       data,
			}

			seq++
			mByte := m.Marshal()

			if _, err := conn.Write(mByte); err != nil {
				log.Fatal(err)
			}
		}
	}
	ticker.Stop()
}
