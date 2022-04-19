package main

// https://leetcode.com/problems/binary-search/

func search(nums []int, target int) int {
	size := len(nums)

	if size == 0 {
		return -1
	}

	if size == 1 && nums[0] == target {
		return 0
	}

	l, r := 0, size-1

	for r >= l {
		m := (r-l)/2 + l
		if target == nums[m] {
			return m
		} else if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
