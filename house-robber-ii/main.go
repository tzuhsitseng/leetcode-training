package main

import "math"

// https://leetcode.com/problems/house-robber-ii/

func rob(nums []int) int {
	size := len(nums)

	if size == 0 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	// case 1: steal from the 0
	robbedFromFirst := make([]int, size)
	robbedFromFirst[0] = nums[0]
	robbedFromFirst[1] = nums[0]
	for i := 2; i < size-1; i++ {
		robbedFromFirst[i] = int(math.Max(float64(robbedFromFirst[i-1]), float64(robbedFromFirst[i-2]+nums[i])))
	}

	// case 2: steal from the 1
	robbedFromSecond := make([]int, size)
	robbedFromSecond[0] = 0
	robbedFromSecond[1] = nums[1]
	for i := 2; i < size; i++ {
		robbedFromSecond[i] = int(math.Max(float64(robbedFromSecond[i-1]), float64(robbedFromSecond[i-2]+nums[i])))
	}

	return int(math.Max(float64(robbedFromFirst[size-2]), float64(robbedFromSecond[size-1])))
}
