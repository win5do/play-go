package main

func canWinNim(n int) bool {
	return n%4 != 0
}

func canWinNim_bitwise(n int) bool {
	// n&3 == n%s
	return n&3 != 0
}

func canWinNim_dp(n int) bool {
	if n < 0 {
		return false
	}

	if n <= 3 {
		return true
	}

	dp := make([]bool, n+1)
	dp[0] = false
	dp[1] = true
	dp[2] = true
	dp[3] = true
	for i := 4; i <= n; i++ {
		dp[i] = !(dp[i-1] && dp[i-2] && dp[i-3])
	}

	return dp[n]
}
