package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"time"
)

const rn = "\r\n"

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 8888")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")
	fmt.Println("新连接：", conn.RemoteAddr())
	t := time.Now().Unix()

	// 超时
	go func(t *int64) {
		for {
			if time.Now().Unix() - *t >= 5 {
				fmt.Println("超时")
				conn.Close()
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(&t)

	for {
		data, err := readTcp(conn)
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		}
		if (data > 0) {
			writeTcp(conn)
			t = time.Now().Unix()
		} else {
			break
		}
	}
}

func readTcp(conn net.Conn) (int, error) {
	var buf bytes.Buffer
	var err error
	rd := bufio.NewScanner(conn)
	total := 0

	for rd.Scan() {
		var n int
		n, err = buf.Write(rd.Bytes())
		if err != nil {
			panic(err)
		}
		buf.Write([]byte(rn))
		total += n
		fmt.Println("读到字节：", n)
		if n == 0 {
			break
		}
	}

	err = rd.Err()

	fmt.Println("总字节数：", total)
	fmt.Println("内容：", rn, buf.String())

	return total, err
}

func writeTcp(conn net.Conn) {
	wt := bufio.NewWriter(conn)
	wt.WriteString("HTTP/1.1 200 OK" + rn)
	wt.WriteString("Date: " + time.Now().String() + rn)
	wt.WriteString("Content-Length: 5" + rn)
	wt.WriteString("Content-Type: text/plain" + rn)
	wt.WriteString(rn)
	wt.WriteString("hello")
	err := wt.Flush()
	if err != nil {
		fmt.Println("Flush err: ", err)
	}
	fmt.Println("写入完毕", conn.RemoteAddr())
}
