package main

import (
	"fmt"
	"testing"
)

func TestMovingCount(t *testing.T) {
	r := movingCount(18, 50, 50)
	fmt.Println(r)
}
