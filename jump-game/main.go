package main

import (
	"math"
)

func canJump(nums []int) bool {
	farest := 0
	for i, v := range nums {
		if i > farest {
			return false
		}
		farest = int(math.Max(float64(farest), float64(i+v)))
	}
	return true
}
