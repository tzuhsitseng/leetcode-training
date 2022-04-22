package main

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	cnt := 0

	for num > 0 {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
	}

	return cnt
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
