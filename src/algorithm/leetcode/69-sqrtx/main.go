package main

func mySqrt(x int) int {
	ans := -1
	start := 0
	end := x

	for start <= end {
		m := (start + end) / 2
		if m*m <= x {
			ans = m
			start = m + 1
		} else {
			end = m - 1
		}
	}

	return ans
}
