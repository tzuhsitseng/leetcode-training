package main

import (
	"fmt"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if k < 1 || k > len(nums) {
		return -1
	}

	idx := len(nums) - k
	return find(nums, idx)
}

func find(nums []int, idx int) int {
	l := make([]int, 0)
	r := make([]int, 0)
	m := len(nums) / 2
	pivot := nums[m]

	for i, v := range nums {
		if v <= pivot && i != m {
			l = append(l, v)
		} else if v > pivot {
			r = append(r, v)
		}
	}

	if idx == len(l) {
		return pivot
	} else if idx < len(l) {
		return find(l, idx)
	} else {
		return find(r, idx-1-len(l))
	}
}

//func quickSort(nums []int, l, r int) {
//	if r < l {
//		return
//	}
//
//	left, right := l, r
//	m := (r-l)/2 + l
//	pivot := nums[m]
//
//	for right >= left {
//		for right >= left && nums[left] < pivot {
//			left++
//		}
//		for right >= left && nums[right] > pivot {
//			right--
//		}
//		if right >= left {
//			nums[left], nums[right] = nums[right], nums[left]
//			left++
//			right--
//		}
//	}
//
//	quickSort(nums, l, right)
//	quickSort(nums, left, r)
//}
