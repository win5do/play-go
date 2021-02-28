package main

func removeDuplicates(nums []int) int {
	l := 0

	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[l] {
			swap(nums, l+1, r)
			l++
		}
	}

	return l + 1
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
