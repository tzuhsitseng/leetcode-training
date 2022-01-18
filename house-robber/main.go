package main

import "math"

func rob(nums []int) int {
	l := len(nums)
	rob := make([]int, l)

	// 判斷邊界條件
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}

	// 先算好節點 1 跟 2
	rob[0] = nums[0]
	rob[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	// 開始計算怎麼偷會比較多
	for i := 2; i < l; i++ {
		rob[i] = int(math.Max(float64(nums[i]+rob[i-2]), float64(rob[i-1])))
	}
	return rob[l-1]
}
