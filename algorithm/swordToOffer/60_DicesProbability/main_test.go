package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	n := 6
	printProbability_recurse(n)
	fmt.Println("---")
	printProbability_loop(n)
}
