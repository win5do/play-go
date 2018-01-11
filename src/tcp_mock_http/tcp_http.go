package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 8888")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	t := time.Now().Second()
	for {
		err := readTcp(conn)
		if err != nil {
			continue
		}

		writeTcp(conn)
		if time.Now().Second()-t >= 3 {
			break
		}
	}
}

func readTcp(conn net.Conn) error {
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", buf[:n])
	fmt.Println("r")
	return nil
}

func writeTcp(conn net.Conn) {
	rn := "\r\n"
	s := "HTTP/1.1 200 OK" + rn +
		"Date: " + time.Now().String() + rn +
		"Content-Length: 5" + rn +
		"Content-Type: text/plain" + rn + rn +
		"hello"

	fmt.Println(s)

	conn.Write([]byte(s))
	fmt.Println("w")
}
