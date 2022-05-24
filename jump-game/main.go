package main

// https://leetcode.com/problems/jump-game/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	furthest := nums[0]

	for i := 1; i < len(nums); i++ {
		if furthest < i {
			return false
		}
		furthest = max(furthest, i+nums[i])
	}

	return true
}
