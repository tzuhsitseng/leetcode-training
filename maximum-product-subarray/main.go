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
	res := -10 // constraints of question
	curMin, curMax := 1, 1

	for _, n := range nums {
		tmpMin := curMin
		tmpMax := curMax
		curMin = min(n, n*tmpMin)
		curMin = min(curMin, n*tmpMax)
		curMax = max(n, n*tmpMin)
		curMax = max(curMax, n*tmpMax)
		res = max(res, curMax)
	}

	return res
}
