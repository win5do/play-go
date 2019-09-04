package sortAlgorithm

func quickSort(arr []int) []int {
	return quickRecurse(arr, 0, len(arr)-1)
}

func quickRecurse(arr []int, start, end int) []int {
	if start > end-1 {
		return arr
	}

	m := arr[start]
	lp := start
	rp := end

	for lp < rp {
		// 从右往左找到第一个比m小的下标
		for arr[rp] >= m && lp < rp {
			rp--
		}

		// 从左往右找到第一个比m大的下标
		for arr[lp] <= m && lp < rp {
			lp++
		}

		swap(arr, lp, rp)
	}

	// lp == rp
	swap(arr, lp, start)

	quickRecurse(arr, start, lp-1)
	quickRecurse(arr, lp+1, end)

	return arr
}
