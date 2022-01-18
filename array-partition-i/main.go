package main

func arrayPairSum(nums []int) int {
	result := 0
	//sort.Ints(nums)
	nums = bubbleSort(nums)

	for i := 0; i < len(nums); i += 2 {
		result += nums[i]
	}

	return result
}

func bubbleSort(nums []int) []int {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
