package main

import (
	"fmt"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	dp := map[string]int{}

	var dfs func(i int, buying bool) int
	dfs = func(i int, buying bool) int {
		if i >= len(prices) {
			return 0
		}

		mKey := fmt.Sprintf("%d%t", i, buying)
		if v, ok := dp[mKey]; ok {
			return v
		}

		cooldown := dfs(i+1, buying)
		if buying {
			buy := dfs(i+1, !buying) - prices[i]
			dp[mKey] = max(buy, cooldown)
		} else {
			sell := dfs(i+2, !buying) + prices[i]
			dp[mKey] = max(sell, cooldown)
		}

		return dp[mKey]
	}

	return dfs(0, true)
}
