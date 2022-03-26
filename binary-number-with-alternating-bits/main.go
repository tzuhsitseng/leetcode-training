package main

import (
	"fmt"
)

// https://leetcode.com/problems/binary-number-with-alternating-bits/

func hasAlternatingBits(n int) bool {
	m := n ^ n>>1       // m should be all of 1 bits if n has alternating bits (e.g. 101 ^ 010)
	return m&(m+1) == 0 // m+1 should be all of 0 bits except the first bit
}

//func hasAlternatingBits(n int) bool {
//	b := strconv.FormatInt(int64(n), 2)
//	for i := 0; i < len(b)-1; i++ {
//		if b[i] == b[i+1] {
//			return false
//		}
//	}®
//	return true
//}

func main() {
	fmt.Println(hasAlternatingBits(5))
}
