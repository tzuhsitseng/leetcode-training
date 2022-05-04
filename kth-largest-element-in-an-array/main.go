package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/kth-largest-element-in-an-array/

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	l := h.Len()
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)

	res := 0
	for i := 0; i < k; i++ {
		res = heap.Pop(&h).(int)
	}
	return res
}

// --
// The following solution is quick sort
// --

//func findKthLargest(nums []int, k int) int {
//	if len(nums) == 0 {
//		return -1
//	}
//
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	if k < 1 || k > len(nums) {
//		return -1
//	}
//
//	idx := len(nums) - k
//	return find(nums, idx)
//}
//
//func find(nums []int, idx int) int {
//	l := make([]int, 0)
//	r := make([]int, 0)
//	m := len(nums) / 2
//	pivot := nums[m]
//
//	for i, v := range nums {
//		if v <= pivot && i != m {
//			l = append(l, v)
//		} else if v > pivot {
//			r = append(r, v)
//		}
//	}
//
//	if idx == len(l) {
//		return pivot
//	} else if idx < len(l) {
//		return find(l, idx)
//	} else {
//		return find(r, idx-1-len(l))
//	}
//}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
