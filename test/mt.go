package main

import "fmt"

func fib(n int) {
	a, b := -1, 1
	for n > 0 {
		a, b = b, a+b
		fmt.Println(b)
		n--
	}
}

func main() {
	fib(9)
}
