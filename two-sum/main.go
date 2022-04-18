package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nil
	}

	m := map[int]int{}

	for i, v := range nums {
		if idx, ok := m[v]; ok {
			return []int{idx, i}
		}

		wanted := target - v
		m[wanted] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
