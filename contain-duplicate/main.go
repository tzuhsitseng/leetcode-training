package main

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for i := range nums {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
