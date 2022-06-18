package main

import "fmt"

// https://leetcode.com/problems/target-sum/

func findTargetSumWays(nums []int, target int) int {
	dp := map[string]int{}

	var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}

		key := fmt.Sprint(i, sum)
		if v, ok := dp[key]; ok {
			return v
		}

		dp[key] = dfs(i+1, sum-nums[i]) + dfs(i+1, sum+nums[i])
		return dp[key]
	}

	return dfs(0, 0)
}
