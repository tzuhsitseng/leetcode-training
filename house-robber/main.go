package main

import "math"

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}

	robbed := make([]int, size)
	robbed[0] = nums[0]
	robbed[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < size; i++ {
		robbed[i] = int(math.Max(float64(robbed[i-1]), float64(robbed[i-2]+nums[i])))
	}

	return robbed[size-1]
}
