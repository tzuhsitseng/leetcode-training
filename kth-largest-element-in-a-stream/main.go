package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	size := len(*h)
	res := (*h)[size-1]
	*h = (*h)[:size-1]
	return res
}

type KthLargest struct {
	k  int
	hp IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	hp := IntHeap(nums)
	heap.Init(&hp)

	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(&hp)
	}

	return KthLargest{k, hp}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.hp, val)

	if this.hp.Len() > this.k {
		heap.Pop(&this.hp)
	}

	return this.hp[0]
}

func main() {
	kth := Constructor(1, []int{})
	fmt.Println(kth.Add(-3))
	fmt.Println(kth.Add(-2))
	fmt.Println(kth.Add(-4))
	fmt.Println(kth.Add(0))
	fmt.Println(kth.Add(4))
}
