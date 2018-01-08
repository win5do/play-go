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
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", buf[:n])
		fmt.Println("r")
		writeTcp(conn)
		if time.Now().Second()-t >= 3 {
			break
		}
	}
}

func readTcp(conn net.Conn) {
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", buf[:n])
	fmt.Println("r")

	//rd := bufio.NewReader(conn)
	//for {
	//	s, err := rd.ReadString('\n')
	//	fmt.Println([]byte(s))
	//	fmt.Printf("Received data: %s\n", s)
	//	if err != nil {
	//		fmt.Println("Error reading", err.Error())
	//		break
	//	}
	//
	//	if s == "\r\n" {
	//		fmt.Println("什么鬼")
	//	}
	//}
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
