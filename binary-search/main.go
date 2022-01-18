package main

func search(nums []int, target int) int {
	min, max := 0, len(nums)-1

	for max >= min {
		mid := (max + min) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		}
	}

	return -1
}
