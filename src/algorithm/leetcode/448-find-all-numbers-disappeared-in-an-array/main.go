package main

func findDisappearedNumbers(nums []int) (ans []int) {
	n := len(nums)

	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}

	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}

	return
}
