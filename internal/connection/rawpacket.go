package connection

import (
	"fmt"
)

type RawPacket struct {
	Header uint16
	Length uint16
	Data   []byte
}

func (p RawPacket) String() string {
	return fmt.Sprintf("Header: %X, Length: %v, Data: %X", p.Header, p.Length, p.Data)
}

func (p RawPacket) MarshalBinary() ([]byte, error) {
	return []byte{byte(p.Header >> 8), byte(p.Header), byte(p.Length >> 8), byte(p.Length)}, nil
}

func (p *RawPacket) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("data too short")
	}
	p.Header = uint16(data[0])<<8 | uint16(data[1])
	p.Length = uint16(data[2])<<8 | uint16(data[3])
	p.Data = data[4:]

	if len(p.Data) != int(p.Length) {
		return fmt.Errorf("data length mismatch: expected %d, got %d", p.Length, len(p.Data))
	}

	return nil
}
