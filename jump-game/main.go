package main

// https://leetcode.com/problems/jump-game/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	furthest := 0

	for i := 0; i < len(nums); i++ {
		if furthest < i {
			return false
		}
		furthest = max(furthest, i+nums[i])
	}
	return true
}
