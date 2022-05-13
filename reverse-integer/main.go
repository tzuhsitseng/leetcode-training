package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	res := 0
	max := math.MaxInt32
	min := math.MinInt32

	for x != 0 {
		d := x % 10
		x /= 10

		if res > max/10 || (res == max/10 && d >= max%10) {
			return 0
		}

		if res < min/10 || (res == min/10 && d <= min%10) {
			return 0
		}

		res = res*10 + d
	}

	return res
}

func main() {
	fmt.Println(reverse(-123))
}
