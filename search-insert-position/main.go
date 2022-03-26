package main

import "fmt"

// https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for high >= low {
		mid := (low + high) / 2
		midVal := nums[mid]
		if target == midVal {
			return mid
		} else if target > midVal {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
