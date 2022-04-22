package main

import "fmt"

// https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	round := 31

	for num > 0 {
		b := num & 1
		res |= b << round
		round--
		num >>= 1
	}

	fmt.Println(round)
	return res
}
