package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
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

type KthLargest struct {
	kth  int
	data *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	input := IntHeap(nums)
	heap.Init(&input)

	for i := 0; i < len(nums)-k; i++ {
		_ = heap.Pop(&input)
	}

	return KthLargest{
		kth:  k,
		data: &input,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.data, val)

	if this.data.Len() > this.kth {
		_ = heap.Pop(this.data)
	}

	return (*this.data)[0]
}

func main() {
	kth := Constructor(1, []int{})
	fmt.Println(kth.Add(-3))
	fmt.Println(kth.Add(-2))
	fmt.Println(kth.Add(-4))
	fmt.Println(kth.Add(0))
	fmt.Println(kth.Add(4))
}
