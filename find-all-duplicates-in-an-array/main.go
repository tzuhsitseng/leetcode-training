package main

import "math"

func findDuplicates(nums []int) []int {
	ans := make([]int, 0)

	for _, num := range nums {
		n := int(math.Abs(float64(num))) - 1

		if nums[n] < 0 {
			ans = append(ans, n+1)
		} else {
			nums[n] *= -1
		}

	}

	return ans
}
