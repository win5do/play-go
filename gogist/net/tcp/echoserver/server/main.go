package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Println("Listening on localhost:3333")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err)
			os.Exit(1)
		}
		//logs an incoming message
		log.Printf("Received message %s -> %s", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
