package icmp

import "encoding/binary"

const (
	EchoReply   = 0
	EchoRequest = 8
)

type Message struct {
	Type       uint8
	Code       uint8
	Checksum   uint16
	Identifier uint16
	Sequences  uint16
	Data       []byte
}

func MessageUnmashal(b []byte) (*Message, error) {
	hlen := int(b[0]&0x0f) << 2
	b = b[hlen:]

	m := &Message{
		Type:       b[0],
		Code:       uint8(b[1]),
		Checksum:   uint16(binary.BigEndian.Uint16(b[2:4])),
		Identifier: uint16(binary.BigEndian.Uint16(b[4:6])),
		Sequences:  uint16(binary.BigEndian.Uint16(b[6:8])),
	}

	m.Data = make([]byte, len(b)-8)
	copy(m.Data, b[8:])

	return m, nil
}

func (m *Message) Marshal() []byte {
	// Fix max len
	b := make([]byte, 8+len(m.Data))

	// Type
	b[0] = byte(m.Type)
	// Code
	b[1] = byte(m.Code)

	//Checksum
	b[2] = 0
	b[3] = 0

	//Identifier
	binary.BigEndian.PutUint16(b[4:6], m.Identifier)

	//Sequenses
	binary.BigEndian.PutUint16(b[6:8], m.Sequences)

	//Data
	copy(b[8:], m.Data)

	//Checksum
	cs := checksum(b)
	b[2] = byte(cs >> 8)
	b[3] = byte(cs)

	return b
}

func checksum(b []byte) uint16 {
	count := len(b)
	sum := uint32(0)
	for i := 0; i < count-1; i += 2 {
		sum += uint32(b[i])<<8 | uint32(b[i+1])
	}
	if count&1 != 0 {
		sum += uint32(b[count-1]) << 8
	}
	for (sum >> 16) > 0 {
		sum = (sum & 0xffff) + (sum >> 16)
	}
	return ^(uint16(sum))
}
