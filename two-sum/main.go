package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	records := map[int]int{}

	for i, v := range nums {
		wanted := target - v
		if r, ok := records[wanted]; ok {
			return []int{r, i}
		}
		records[v] = i
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
