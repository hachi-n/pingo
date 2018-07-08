package statistic

import (
	"log"
	"time"
)

type Statistics struct {
	Transmit int
	RTTS     []time.Duration
}

func (s *Statistics) Dump() {
	packetTransmited := s.Transmit
	packetReceived := len(s.RTTS)

	max, avg, min, _ := dump(s)

	log.Println(packetTransmited, "packets transmitted")
	log.Println(packetReceived, "packets received")

	log.Println("max/avg/min", max, "/", avg, "avg", "/", min, "min")

}

func dump(s *Statistics) (max, avg, min, sum time.Duration) {
	max = 0
	min = 0
	sum = 0

	for _, rtt := range s.RTTS {
		sum += rtt

		if max < rtt {
			max = rtt
		}

		if min > rtt {
			min = rtt
		}
	}

	avg = time.Duration(int(sum.Nanoseconds()) / len(s.RTTS))

	return
}
