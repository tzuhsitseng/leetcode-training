package main

import "sort"

// https://leetcode.com/problems/top-k-frequent-elements/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	for _, v := range nums {
		counter[v]++
	}

	elementsByCnt := map[int][]int{}
	for k, v := range counter {
		if _, ok := elementsByCnt[v]; !ok {
			elementsByCnt[v] = []int{k}
		} else {
			elementsByCnt[v] = append(elementsByCnt[v], k)
		}
	}

	sorted := make([]int, 0, len(elementsByCnt))
	for k := range elementsByCnt {
		sorted = append(sorted, k)
	}

	sort.Ints(sorted)

	result := make([]int, 0)
	total := 0

	for i := len(sorted) - 1; i >= max(0, len(sorted)-k) && total < k; i-- {
		elements := elementsByCnt[sorted[i]]
		total += len(elements)
		result = append(result, elements...)
	}

	return result
}

// --
// 可用 bubble sort 達到 O(n) 的效果
// --
