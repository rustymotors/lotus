package connection

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rustymotors/lotus/internal/authlogin"
	"github.com/rustymotors/lotus/internal/shard"
)

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
