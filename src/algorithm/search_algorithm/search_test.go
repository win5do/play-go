package search_algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{}
	i := BinarySearch(arr, 6)
	fmt.Println(i)
}
