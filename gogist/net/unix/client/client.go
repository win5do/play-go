package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got:", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	go reader(c)
	for {
		fmt.Print("Enter text: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			panic(err)
		}
		_, err = c.Write([]byte(input))
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
