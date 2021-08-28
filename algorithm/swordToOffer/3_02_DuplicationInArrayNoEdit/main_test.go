package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(duplicate([]int{1}))
	fmt.Println(duplicate([]int{1, 5, 4, 6, 5, 4, 8}))
	fmt.Println(duplicate([]int{1, 5, 4, 6, 5, 4, 3}))
	fmt.Println(duplicate([]int{0, 0, 4, 6, 5, 4, 3}))
	fmt.Println(duplicate([]int{0, 1, 2, 3, 4, 5, 6}))
	fmt.Println(duplicate([]int{0, 1, 2, 3, 4, 5, 5}))
}
