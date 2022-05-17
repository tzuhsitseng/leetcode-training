package main

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-elements/

type Element struct {
	p int
	v int
}

type PriorityQueue []*Element

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].p > p[j].p
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(*Element))
}

func (p *PriorityQueue) Pop() interface{} {
	l := len(*p)
	res := (*p)[l-1]
	*p = (*p)[:l-1]
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	s := make([]*Element, 0, len(m))
	for val, counter := range m {
		s = append(s, &Element{
			p: counter,
			v: val,
		})
	}

	p := PriorityQueue(s)
	heap.Init(&p)

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&p).(*Element).v)
	}

	return res
}
