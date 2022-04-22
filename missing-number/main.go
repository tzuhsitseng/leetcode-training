package main

// https://leetcode.com/problems/missing-number/

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	expected := len(nums)
	actually := 0

	for i := 0; i < len(nums); i++ {
		expected ^= i
		actually ^= nums[i]
	}

	return expected ^ actually
}
