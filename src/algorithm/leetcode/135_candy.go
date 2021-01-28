package leetcode

func candy(ratings []int) int {
	leng := len(ratings)
	if leng < 2 {
		return leng
	}

	candys := make([]int, leng)
	candys[0] = 1
	for i := 1; i < leng; i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		} else {
			candys[i] = 1
		}
	}
	for i := leng - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i], candys[i+1]+1)
		}
	}

	r := 0
	for _, v := range candys {
		r += v
	}
	return r
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
