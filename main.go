package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/rustymotors/lotus/internal/authlogin"
	"github.com/rustymotors/lotus/internal/shard"
)

func launchHTTPServer(host string, port int, handler http.HandlerFunc) {
	log.Printf("Starting HTTP server on %s:%d", host, port)
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func launchTCPServer(host string, port int, handler func(net.Conn)) {
	log.Printf("Starting TCP server on %s:%d", host, port)
	s, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn)
	}
}

var (
	tcpListenPorts = []int{8226, 8227, 8228, 7003, 43300}
)

func handleTCPConnection(conn net.Conn) {
	log.Println("TCP connection from", conn.RemoteAddr())
	body := make([]byte, 1024)

	_, err := conn. Read(body)

	if err != nil {
		log.Println("Error reading:", err.Error())
	}

	log.Println("Received:", fmt.Sprintf("%X", body))

}


func main() {
	fmt.Println("Hello, 世界, welcome to Go!")

	// launch HTTP server
	go launchHTTPServer("0.0.0.0", 3000, myHandler)

	// launch TCP servers
	for _, port := range tcpListenPorts {
		go launchTCPServer("0.0.0.0", port, handleTCPConnection)
	}
	log.Println("Servers started")

	// block forever
	select {}


}



func myHandler(w http.ResponseWriter, r *http.Request) {
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


