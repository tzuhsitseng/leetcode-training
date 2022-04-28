package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.com/problems/koko-eating-bananas/

func minEatingSpeed(piles []int, h int) int {
	if h < len(piles) {
		return -1
	}

	sort.Ints(piles)

	l, r := 1, piles[len(piles)-1]

	for l <= r {
		mid := (l + r) / 2
		hr := getHours(piles, mid)

		if hr <= h {
			if mid == 1 || getHours(piles, mid-1) > h {
				return mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func getHours(piles []int, k int) int {
	sum := 0
	for _, v := range piles {
		sum += int(math.Ceil(float64(v) / float64(k)))
	}
	return sum
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
