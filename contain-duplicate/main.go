package main

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	counter := map[int]bool{}

	for _, v := range nums {
		if ok := counter[v]; ok {
			return true
		}
		counter[v] = true
	}

	return false
}
