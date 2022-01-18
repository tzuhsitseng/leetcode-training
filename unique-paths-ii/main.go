package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || obstacleGrid[0][0] == 1 {
		return 0
	}

	// 宣告 m * n 的二維陣列
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 唯一只知道起點方法數為 1
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 判斷障礙物
			if obstacleGrid[i][j] == 1 {
				continue
			}

			// 確認邊界條件
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
