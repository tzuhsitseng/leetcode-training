package main

import (
	"container/heap"
	"math"
)

// https://leetcode.com/problems/last-stone-weight/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	size := len(*h)
	res := (*h)[size-1]
	*h = (*h)[:size-1]
	return res
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	h := IntHeap(stones)
	heap.Init(&h)

	for h.Len() >= 2 {
		s1 := heap.Pop(&h).(int)
		s2 := heap.Pop(&h).(int)
		if s1 != s2 {
			heap.Push(&h, int(math.Abs(float64(s1)-float64(s2))))
		}
	}

	if h.Len() > 0 {
		return h[0]
	}
	return 0
}
