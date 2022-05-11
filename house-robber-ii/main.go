package main

// https://leetcode.com/problems/house-robber-ii/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	dp1 := make([]int, length)
	dp2 := make([]int, length)

	dp1[0] = nums[0]
	dp1[1] = dp1[0]
	for i := 2; i < length-1; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
	}

	dp2[0] = 0
	dp2[1] = nums[1]
	for i := 2; i < length; i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}

	return max(dp1[length-2], dp2[length-1])
}
