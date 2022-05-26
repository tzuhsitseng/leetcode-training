package main

// https://leetcode.com/problems/maximum-product-subarray/

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxVal, minVal := 1, 1
	res := 0

	for _, n := range nums {
		tmpMax, tmpMin := maxVal, minVal
		maxVal = max(max(n*tmpMax, n*tmpMin), n)
		minVal = min(min(n*tmpMin, n*tmpMax), n)
		res = max(res, maxVal)
	}

	return res
}
