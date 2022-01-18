package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	i, j, currentVal, nextVal := 0, 0, nums[0], nums[0]

	for i < len(nums) && j < len(nums) {
		for j < len(nums) {
			if nums[j] != currentVal {
				nextVal = nums[j]
				break
			}
			j++
		}

		nums[i] = currentVal
		currentVal = nextVal
		i++
	}

	return i
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(n, nums)
}
