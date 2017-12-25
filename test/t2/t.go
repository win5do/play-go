package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	go a()
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	m1()
}
func m1() {
	m2()
}
func m2() {
	m3()
}
func m3() {
	time.Sleep(5 * time.Second)
}
func a() {
	time.Sleep(5 * time.Second)
}
