package main

// https://leetcode.com/problems/jump-game-ii/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	furthest := 0
	l, r := 0, 0

	for r < len(nums)-1 {
		for i := l; i <= r; i++ {
			furthest = max(furthest, i+nums[i])
		}
		l = r + 1
		r = furthest
		res++
	}

	return res
}
