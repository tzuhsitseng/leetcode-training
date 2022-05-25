package main

// https://leetcode.com/problems/house-robber-ii/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	inputSize := len(nums)

	if inputSize == 1 {
		return nums[0]
	}

	dp := make([]int, inputSize)
	dp[0] = nums[0]
	dp[1] = nums[0]

	for i := 2; i < inputSize-1; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	res1 := dp[inputSize-2]

	dp = make([]int, inputSize)
	dp[0] = 0
	dp[1] = nums[1]

	for i := 2; i < inputSize; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	res2 := dp[inputSize-1]

	return max(res1, res2)
}
