package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
