package main

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	size := len(numbers)

	l, r := 0, size-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return nil
}
