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

type PointPQ []Point

func (h PointPQ) Len() int { return len(h) }

func (h PointPQ) Less(i, j int) bool { return h[i].v < h[j].v }

func (h PointPQ) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *PointPQ) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointPQ) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) == 1 {
		return points
	}

	pq := PointPQ{}
	heap.Init(&pq)

	for _, p := range points {
		heap.Push(&pq, Point{
			p: p,
			v: p[0]*p[0] + p[1]*p[1],
		})
	}

	res := make([][]int, 0, k)

	for k > 0 {
		p := heap.Pop(&pq).(Point)
		res = append(res, p.p)
		k--
	}

	return res
}

func main() {
	fmt.Println(kClosest([][]int{{-2, -6}, {-7, -2}, {-9, 6}, {10, 3}, {-8, 1}, {2, 8}}, 5))
}
