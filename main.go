package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/rustymotors/lotus/internal/connection"
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




func main() {
	fmt.Println("Hello, 世界, welcome to Go!")

	// launch HTTP server
	go launchHTTPServer("0.0.0.0", 3000, connection.HandleHTTPRequest)

	// launch TCP servers
	for _, port := range tcpListenPorts {
		go launchTCPServer("0.0.0.0", port, connection.HandleTCPConnection)
	}
	log.Println("Servers started")

	// block forever
	select {}


}




