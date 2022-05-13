package main

// https://leetcode.com/problems/sum-of-two-integers/

func getSum(a int, b int) int {
	for b != 0 {
		tmp := (a & b) << 1
		a = a ^ b // 不包括進位的部分
		b = tmp   // 包括進位的部分
	}
	return a
}
