package main

import "container/heap"

// https://leetcode.com/problems/last-stone-weight/

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	size := len(*h)
	res := (*h)[size-1]
	*h = (*h)[:size-1]
	return res
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}

	hp := IntHeap(stones)
	heap.Init(&hp)

	for hp.Len() > 1 {
		p1 := heap.Pop(&hp).(int)
		p2 := heap.Pop(&hp).(int)
		if p1-p2 > 0 {
			heap.Push(&hp, p1-p2)
		}
	}

	if hp.Len() == 1 {
		return hp[0]
	}

	return 0
}
