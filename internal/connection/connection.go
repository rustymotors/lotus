package connection

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rustymotors/lotus/internal/authlogin"
	"github.com/rustymotors/lotus/internal/shard"
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
	return nil
}

func HandleTCPConnection(conn net.Conn) {
	log.Println("TCP connection from", conn.RemoteAddr())

	// read incoming data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err)
		return
	}

	// unmarshal the raw packet
	var packet RawPacket
	err = packet.UnmarshalBinary(buf[:n])
	if err != nil {
		log.Println("Error unmarshalling:", err)
		return
	}

	localPort := conn.LocalAddr().(*net.TCPAddr).Port

	// print the raw packet
	log.Println("Received packet from", localPort, ":", packet)
}


func HandleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// print the request
	fmt.Println("Request received: ", r.RemoteAddr, r.Method, r.URL.Path, r.URL.Query())

	// print the request body
	fmt.Println("Request body: ", r.Body)

	switch r.URL.Path {
	case "/AuthLogin":
		// handle AuthLogin
		authlogin.HandleAuthLogin(r, w)

	case "/ShardList/":
		// handle ShardList
		shard.HandleShardList(r, w)

	default:
		// handle all other requests
		fmt.Println("Other")
	}
}
