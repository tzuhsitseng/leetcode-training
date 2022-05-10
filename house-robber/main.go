package main

// https://leetcode.com/problems/house-robber/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}

	dp := make([]int, size)
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])

	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[size-1]
}
