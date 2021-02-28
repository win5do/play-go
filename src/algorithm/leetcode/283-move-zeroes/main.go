package main

func moveZeroes(nums []int) {
	l := 0
	r := 0
	leng := len(nums)

	for r < leng {
		if nums[r] != 0 {
			if l != r {
				swap(nums, l, r)
			}
			l++
		}

		r++
	}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
