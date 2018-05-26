package sort_algorithm

func quickSort(arr []int) []int {
	return quickRecurse(arr, 0, len(arr)-1)
}

func quickRecurse(arr []int, start, end int) [] int {
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

func quickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	lp, rp := 0, len(arr)-1
	m := arr[0]

	for lp < rp {
		// arr[0]为基准数 arr[lp]一直等于arr[0]
		if arr[lp+1] > m {
			// 移到后面
			swap(arr, lp+1, rp)
			rp--
		} else {
			// 向前移
			swap(arr, lp+1, lp)
			lp++
		}
	}

	quickSort2(arr[:lp])
	quickSort2(arr[lp+1:])

	return arr
}
