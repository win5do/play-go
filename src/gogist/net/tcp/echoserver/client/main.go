package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3333")
	if err != nil {
		log.Fatalf("Error connecting:", err)
	}
	defer conn.Close()
	log.Println("Connecting to localhost:3333")
	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}
func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\n"))
		if e != nil {
			log.Println("Error to send message because of ", e.Error())
			break
		}
	}
}
func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for i := 1; i <= 10; i++ {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			log.Println("Error to read message because of ", err)
			return
		}
		log.Println(line)
	}
}
