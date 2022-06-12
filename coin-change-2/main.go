package main

import "fmt"

// https://leetcode.com/problems/coin-change-2/

func change(amount int, coins []int) int {
	dp := map[string]int{}

	var dfs func(i, a int) int
	dfs = func(i, a int) int {
		if i == len(coins) {
			return 0
		}
		if a == amount {
			return 1
		}
		if a > amount {
			return 0
		}

		key := fmt.Sprintf("%d:%d", i, a)
		if v, ok := dp[key]; ok {
			return v
		}

		dp[key] = dfs(i, a+coins[i]) + dfs(i+1, a)
		return dp[key]
	}

	return dfs(0, 0)
}
