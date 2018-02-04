package main

import "fmt"

func main() {
	//var a interface{}
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	ref(a)
	fmt.Println(a)
}

func ref(a []int) {
	b := append(a[:1], a[2:]...)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
