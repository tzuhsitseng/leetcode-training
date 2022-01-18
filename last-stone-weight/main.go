package main

import "container/heap"

func main() {
	lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, stone := range stones {
		heap.Push(h, -stone)
	}

	for h.Len() > 0 {
		if h.Len() == 1 {
			return -(heap.Pop(h).(int))
		}
		p1 := heap.Pop(h).(int)
		p2 := heap.Pop(h).(int)
		if p2-p1 > 0 {
			heap.Push(h, -(p2 - p1))
		}
	}
	return 0
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
