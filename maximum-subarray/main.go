package main

// https://leetcode.com/problems/maximum-subarray/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, curSum := nums[0], 0

	for _, v := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += v
		maxSum = max(curSum, maxSum)
	}

	return maxSum
}
