package main

func majorityElement(nums []int) int {
	result := nums[0]
	cnt := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == result {
			cnt++
		} else {
			if cnt > 0 {
				cnt--
			} else {
				result = nums[i]
				cnt = 1
			}

		}
	}

	return result
}
