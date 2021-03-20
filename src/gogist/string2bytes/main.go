package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	s := strings.Repeat("abc", 3)
	b := str2bytes(s)
	s2 := bytes2str(b)
	fmt.Println(b, s2)
}
