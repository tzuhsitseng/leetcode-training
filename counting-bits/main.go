package main

// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 0; i <= n; i++ {
		cnt := 0
		tmp := i
		for tmp > 0 {
			cnt += tmp & 1
			tmp >>= 1
		}
		res[i] = cnt
	}

	return res
}

// --
// dp solution
// --

//func countBits(n int) []int {
//	dp := make([]int, n+1)
//	dp[0] = 0
//	offset := 1
//
//	for i := 1; i <= n; i++ {
//		if offset*2 == i {
//			offset = i
//		}
//		dp[i] = 1 + dp[i-offset]
//	}
//
//	return dp
//}
