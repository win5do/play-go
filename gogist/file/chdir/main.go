package main

import (
	"fmt"
	"os"
)

// os.Chdir 改变工作目录

func main() {
	beforeDir, _ := os.Getwd()
	fmt.Println("before working dir：", beforeDir)
	err := os.Chdir("/")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	afterDir, _ := os.Getwd()
	fmt.Println("after working dir：", afterDir)
}
