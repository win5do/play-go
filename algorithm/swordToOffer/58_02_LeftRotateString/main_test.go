package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(leftRotateString("abcd", 2))
	fmt.Println(leftRotateString("abcd", 1))
	fmt.Println(leftRotateString("abcd", 6))
	fmt.Println(leftRotateString("你好明天", 2))
}
