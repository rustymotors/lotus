package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/rustymotors/lotus/internal/authlogin"
)

func main() {
	fmt.Println("Hello, 世界, welcome to Go!")

	s := &http.Server{
		Addr:           ":3000",
		Handler:        http.HandlerFunc(myHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

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


	default:
		// handle all other requests
		fmt.Println("Other")
	}
}


