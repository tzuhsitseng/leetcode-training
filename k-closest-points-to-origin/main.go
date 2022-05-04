package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/k-closest-points-to-origin/

type Point struct {
	p []int
	v int
}

type PointHeap []Point

func (h PointHeap) Len() int {
	return len(h)
}

func (h PointHeap) Less(i, j int) bool {
	return h[i].v < h[j].v
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) == 0 {
		return nil
	}

	h := PointHeap{}
	heap.Init(&h)

	for _, p := range points {
		p1 := p[0]
		p2 := p[1]
		heap.Push(&h, Point{
			p: p,
			v: p1*p1 + p2*p2,
		})
	}

	res := make([][]int, 0)
	for i := 0; i < k; i++ {
		p := heap.Pop(&h).(Point)
		res = append(res, p.p)
	}

	return res
}

func main() {
	fmt.Println(kClosest([][]int{{-2, -6}, {-7, -2}, {-9, 6}, {10, 3}, {-8, 1}, {2, 8}}, 5))
}
