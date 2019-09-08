package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	rd := NewStreamReader()

	for i := 100; i > 0; i-- {
		rd.insert(i)
	}

	fmt.Println(rd.getMedian())
}
