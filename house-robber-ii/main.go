package main

import "math"

//func rob(nums []int) int {
//	l := len(nums)
//
//	// 判斷邊界條件
//	if l == 0 {
//		return 0
//	} else if l == 1 {
//		return nums[0]
//	}
//
//	// 先看偷第 0 棟的狀況
//	tmp := 0
//	prev, curr := nums[0], nums[0]
//	for i := 2; i < l-1; i++ {
//		tmp, prev = prev, curr
//		curr = int(math.Max(float64(nums[i]+tmp), float64(prev)))
//	}
//	result := curr
//
//	// 不偷第 0 棟的狀況
//	prev, curr = 0, nums[1]
//	for i := 2; i < l; i++ {
//		tmp, prev = prev, curr
//		curr = int(math.Max(float64(nums[i]+tmp), float64(prev)))
//	}
//
//	// 確認情境 1 還是情境 2 偷到的比較多
//	return int(math.Max(float64(result), float64(curr)))
//}

func rob(nums []int) int {
	size := len(nums)

	if size == 0 {
		return 0
	}

	if size == 1 {
		return nums[0]
	}

	// case 1: steal from the 0
	robbedFromFirst := make([]int, size)
	robbedFromFirst[0] = nums[0]
	robbedFromFirst[1] = nums[0]
	for i := 2; i < size-1; i++ {
		robbedFromFirst[i] = int(math.Max(float64(robbedFromFirst[i-1]), float64(robbedFromFirst[i-2]+nums[i])))
	}

	// case 2: steal from the 1
	robbedFromSecond := make([]int, size)
	robbedFromSecond[0] = 0
	robbedFromSecond[1] = nums[1]
	for i := 2; i < size; i++ {
		robbedFromSecond[i] = int(math.Max(float64(robbedFromSecond[i-1]), float64(robbedFromSecond[i-2]+nums[i])))
	}

	return int(math.Max(float64(robbedFromFirst[size-2]), float64(robbedFromSecond[size-1])))
}
