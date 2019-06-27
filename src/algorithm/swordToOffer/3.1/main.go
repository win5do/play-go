package main

import (
	"fmt"
)

func duplicate(arr []int) (int, bool) {
	if len(arr) <= 1 {
		return -1, false
	}

	for _, v := range arr {
		if v < 0 || v > len(arr)-1 {
			// invalid arr
			return -1, false
		}
	}

	for i := 0; i < len(arr); i++ {
		for i != arr[i] {
			if arr[i] == arr[arr[i]] {
				return arr[i], true
			}

			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		}
	}

	return -1, false
}

func main() {
	fmt.Println(duplicate([]int{1}))
	fmt.Println(duplicate([]int{1, 5, 4, 6, 5, 4, 8}))
	fmt.Println(duplicate([]int{1, 5, 4, 6, 5, 4, 3}))
	fmt.Println(duplicate([]int{0, 0, 4, 6, 5, 4, 3}))
	fmt.Println(duplicate([]int{0, 1, 2, 3, 4, 5, 6}))
	fmt.Println(duplicate([]int{0, 1, 2, 3, 4, 5, 5}))
}
