package main

import (
	"fmt"
)

func duplicate(arr []int) (int, bool) {
	if len(arr) <= 1 {
		return -1, false
	}

	start := 1
	end := len(arr)

	for end >= start {
		// equal (start + end) / 2
		mid := (start + end) >> 1
		count := countRange(arr, start, mid)
		if start == end {
			if count > 1 {
				return start, true
			} else {
				break
			}
		}

		if count > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return -1, false
}

func countRange(arr []int, start, end int) int {
	count := 0
	for _, v := range arr {
		if v >= start && v <= end {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(duplicate([]int{1}))
	fmt.Println(duplicate([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(duplicate([]int{1, 7, 6, 5, 4, 3, 2, 1}))
}
