package main

import "fmt"

func main() {
	//var a interface{}
	a := []int{1, 2, 3, 4}
	fmt.Println(a[4+1:])
}
