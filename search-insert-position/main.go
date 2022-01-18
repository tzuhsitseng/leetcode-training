package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	// 黃金交叉後代表查找結束, 迴圈退出
	for low <= high {
		mid := (high + low) / 2

		// 如果直接找 mid 就找到，就直接回傳
		if target == nums[mid] {
			return mid
		}

		// 如果發現 target 比 nums[mid] 小，繼續往前找，所以把 high 下修
		// 如果發現 target 比 nums[mid] 大，繼續往後找，所以把 low 上修
		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// 黃金交叉後若都沒找到，low 即為 insert position
	return low
}
