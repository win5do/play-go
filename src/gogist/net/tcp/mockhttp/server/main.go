package main

import (
	"bufio"
	"log"
	"net"
	"time"

	"playGo/src/gogist/net/tcp/mockhttp"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen to 8888")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("conn err:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Println("disconnect")
	}()

	log.Printf("new connect：%s", conn.RemoteAddr())
	done := make(chan struct{})

	go func() {
		tick := time.NewTicker(3 * time.Second)
		select {
		case <-tick.C:
			log.Println("timeout")
			conn.Close()
			return
		case <-done:
			tick.Stop()
			log.Println("done")
			return
		}
	}()

	_, err := mockhttp.ReadConn(conn)
	if err != nil {
		log.Println("read err:", err)
		return
	}

	done <- struct{}{}
	sendResp(conn)
}

func sendResp(conn net.Conn) {
	wt := bufio.NewWriter(conn)
	wt.WriteString("HTTP/1.1 200 OK")
	wt.Write(mockhttp.RN)
	wt.WriteString("Date: " + time.Now().String())
	wt.Write(mockhttp.RN)
	wt.WriteString("Content-Length: 5")
	wt.Write(mockhttp.RN)
	wt.WriteString("Content-Type: text/plain")
	wt.Write(append(mockhttp.RN, mockhttp.RN...))
	wt.WriteString("hello")
	err := wt.Flush()
	if err != nil {
		log.Println("Flush err: ", err)
		return
	}
	log.Printf("write done: %s", conn.RemoteAddr())
}
