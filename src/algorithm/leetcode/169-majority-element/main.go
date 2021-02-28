package main

// 必须要有众数，count > len/2，否则结果错误
func majorityElement(nums []int) int {
	r := -1
	count := 0

	for _, v := range nums {
		if count == 0 {
			r = v
		}

		if v == r {
			count++
		} else {
			count--
		}
	}

	return r
}
