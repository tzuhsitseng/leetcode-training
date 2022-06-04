package main

// https://www.lintcode.com/problem/394/

func FirstWillWin(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 || n == 2 {
		return true
	}

	dp := make([]bool, n+1)
	dp[0] = false
	dp[1] = true
	dp[2] = true

	for i := 3; i <= n; i++ {
		dp[i] = !dp[i-1] || !dp[i-2]
	}

	return dp[n]
}
