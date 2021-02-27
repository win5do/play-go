package main

import "math"

func maxProfit(prices []int) int {
	r := 0
	min := math.MaxInt32

	for _, v := range prices {
		if v > min {
			x := v - min
			if x > r {
				r = x
			}
		}

		if v < min {
			min = v
		}
	}

	return r
}
