package main

func uniquePaths(m int, n int) int {
	// 宣告 m * n 的二維陣列
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 最上面一排的方法數都是 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// 最左邊一排的方法數都是 1
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	// 在某節點的方法數會等於左節點跟上節點方法數的總和
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
