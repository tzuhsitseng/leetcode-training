package main

// https://leetcode.com/problems/maximum-subarray/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := nums[0]
	sum := 0

	for _, n := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += n
		res = max(res, sum)
	}

	return res
}
