package main

import (
	"bufio"
	"log"
	"net"
	"time"

	"playGo/gogist/net/tcp/mockhttp"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("disconnect")
		conn.Close()
	}()

	sendReq(conn)

	buf, err := mockhttp.ReadConn(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("total: %d", buf.Len())
}

func sendReq(conn net.Conn) {
	wt := bufio.NewWriter(conn)
	wt.WriteString("GET / HTTP/1.1")
	wt.Write(mockhttp.RN)
	wt.WriteString("Date: " + time.Now().String())
	wt.Write(append(mockhttp.RN, mockhttp.RN...))
	err := wt.Flush()
	if err != nil {
		log.Printf("Flush err: %s", err)
		return
	}
	log.Printf("write done: %s", conn.RemoteAddr())
}
