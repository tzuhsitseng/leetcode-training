package main

// https://leetcode.com/problems/longest-increasing-subsequence/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 1
	}

	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	res := 0
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

// --
// It's the v1 thought
// --

//func lengthOfLIS(nums []int) int {
//	size := len(nums)
//	if size == 1 {
//		return 1
//	}
//
//	dp := make([]int, size)
//	for i := range dp {
//		dp[i] = 1
//	}
//
//	res := 0
//	for i := 0; i < size; i++ {
//		for j := i - 1; j >= 0; j-- {
//			if nums[i] > nums[j] {
//				dp[i] = max(dp[i], dp[j]+1)
//			}
//		}
//		res = max(res, dp[i])
//	}
//
//	return res
//}
