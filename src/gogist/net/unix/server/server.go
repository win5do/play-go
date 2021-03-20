package main

// https://stackoverflow.com/a/4759231/7986742

import (
	"fmt"
	"net"
	"os"
	"os/signal"
)

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			panic("Write: " + err.Error())
		}
	}
}

func main() {
	l, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		println("listen error", err.Error())
		return
	}
	defer func() {
		// close实际调用unlink /tmp/echo.sock
		// 如果没有close，则下次会报错 bind: address already in use
		println("close sock")
		err := l.Close()
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		println("start listening")
		for {
			fd, err := l.Accept()
			if err != nil {
				println("accept error", err.Error())
				return
			}

			go echoServer(fd)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
	return
	// 如果使用exit而不是return则不会执行defer
	// os.Exit(0)
}
