package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.com/problems/koko-eating-bananas/

func calculate(piles []int, k int) int {
	res := 0
	for _, p := range piles {
		res += int(math.Ceil(float64(p) / float64(k)))
	}
	return res
}

func minEatingSpeed(piles []int, h int) int {
	if h < len(piles) {
		return -1
	}

	sort.Ints(piles)
	max := piles[len(piles)-1]

	l, r := 1, max

	for l <= r {
		m := (l + r) / 2
		need := calculate(piles, m)

		if need <= h {
			if m == 1 {
				return m
			}
			need = calculate(piles, m-1)
			if need > h {
				return m
			}
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func main() {
	//fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
}
