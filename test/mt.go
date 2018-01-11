package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

	go func(t *int64) {
		for {
			if time.Now().Unix() - *t >= 5 {
				conn.Close()
			}

			time.Sleep(100 * time.Millisecond)
		}
	}(&t)

	for {
		fmt.Println("这是连接：", conn.RemoteAddr())
		data, err := readTcp(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println(err, conn.RemoteAddr()) // todo del
				continue
			} else {
				fmt.Println(err)
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
		total += n
		fmt.Println("读到字节：", n)
		if n == 0 {
			break
		}
	}

	if err = rd.Err(); err != nil {
		fmt.Println("read err：", err)
	}

	fmt.Println("总字节数：", total, " 连接：", conn.RemoteAddr())

	return total, err
}

func writeTcp(conn net.Conn) {
	rn := "\r\n"
	wt := bufio.NewWriter(conn)
	wt.WriteString("HTTP/1.1 200 OK" + rn)
	wt.WriteString("Date: " + time.Now().String() + "\r\n")
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
