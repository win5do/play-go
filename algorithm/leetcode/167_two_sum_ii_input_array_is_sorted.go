package leetcode

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		add := numbers[l] + numbers[r]
		if add == target {
			return []int{l + 1, r + 1}
		}

		if add > target {
			r--
		} else {
			l++
		}
	}

	return nil
}
