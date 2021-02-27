package main

// 贪心
func maxProfit(prices []int) (ans int) {
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
