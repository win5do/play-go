package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"time"
)

const rn = "\r\n"

func main() {
	conn, err := net.Dial("tcp", ":8888")
	defer conn.Close()
	defer fmt.Println("断开")
	if err != nil {
		panic(err)
	}
	sendReq(conn)
	for {
		total, err := readResp(conn)
		if err != nil {
			panic(err)
		}
		if total > 0 {
			break
		}
	}
}

func sendReq(conn net.Conn) {
	wt := bufio.NewWriter(conn)
	wt.WriteString("GET / HTTP/1.1" + rn)
	wt.WriteString("Date: " + time.Now().String() + rn)
	wt.WriteString(rn)
	err := wt.Flush()
	if err != nil {
		fmt.Println("Flush err: ", err)
	}
	fmt.Println("写入完毕", conn.RemoteAddr())
}

func readResp(conn net.Conn) (int, error) {
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

	if (total > 0) {
		fmt.Println("resp:", rn, buf.String())
	}

	return total, err
}
