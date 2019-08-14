package main

import (
	"fmt"
	"testing"
)

func TestPrint1ToN_string(t *testing.T) {
	print1ToN_string(3)
	fmt.Println("---")
	print1ToN_recurse(3)
}
