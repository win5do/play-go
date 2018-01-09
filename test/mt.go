package main

import (
	"bufio"
	"bytes"
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
	//t := time.Now().Unix()
	for {
		err := readTcp(conn)
		if err != nil {
			break
			//if time.Now().Unix()-t >= 1 {
			//	break
			//}
		}
		writeTcp(conn)
		//t = time.Now().Unix()
	}
}

func readTcp(conn net.Conn) error {
	var err error
	var buf bytes.Buffer
	rd := bufio.NewReader(conn)
	for {
		var line []byte
		line, _, err = rd.ReadLine()
		fmt.Printf("Received data: %s\n", line)
		buf.Write(line)
		if err != nil {
			fmt.Println("line err", err, time.Now())
			break
		}
	}

	fmt.Printf("all data: %s\n", buf)
	return err
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

	//bfw := bufio.NewWriter(conn)
	//bfw.WriteString("HTTP/1.1 200 OK\r\n")
	//bfw.WriteString("Date:" + time.Now().String() + "\r\n")
	//bfw.WriteString("Content-Type:text;charset=utf8\r\n")
	//bfw.WriteString("Content-Length:" + string(len("hello")) + "\r\n")
	//bfw.WriteString("\r\n")
	//bfw.WriteString("hello")
	//err := bfw.Flush()
	//if err != nil {
	//	fmt.Println(err)
	//}
}
