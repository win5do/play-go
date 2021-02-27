package main

func singleNumber(nums []int) int {
	var x int
	for _, v := range nums {
		x ^= v
	}
	return x
}
