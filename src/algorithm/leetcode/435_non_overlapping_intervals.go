package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		// 按照结尾的大小进行增序排序，选择的区间结尾越小，余留给其它区间的空间 就越大
		return intervals[i][1] <= intervals[j][1]
	})

	i, j := 0, 1
	count := 0

	leng := len(intervals)
	for i < leng && j < leng {
		left := intervals[i]
		right := intervals[j]

		if left[1] <= right[0] {
			i = j
			j++
			continue
		}

		j++
		count++
	}
	return count
}
