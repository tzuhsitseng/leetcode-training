package main

import "math"

func minFlipsMonoIncr(S string) int {
	f0, f1 := 0, 0

	// case 0: 全部都是 0 的狀況
	// case 1: 前面 0，但後面都變成 1

	for _, v := range S {
		if string(v) == "0" {
			f1 += 1
		} else {
			f1 = int(math.Min(float64(f0), float64(f1)))
			f0 += 1
		}
	}

	return int(math.Min(float64(f0), float64(f1)))
}
