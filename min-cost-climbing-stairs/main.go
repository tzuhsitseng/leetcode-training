package main

// https://leetcode.com/problems/min-cost-climbing-stairs/

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	l := len(cost)

	if l == 0 {
		return 0
	} else if l == 1 {
		return cost[0]
	} else if l == 2 {
		return min(cost[0], cost[1])
	}

	dp := make([]int, l+1)
	dp[l] = 0
	dp[l-1] = cost[l-1]

	for i := l - 2; i >= 0; i-- {
		dp[i] = min(dp[i+1], dp[i+2]) + cost[i]
	}

	return min(dp[0], dp[1])
}

// --
// v1 thought
// --

//func minCostClimbingStairs(cost []int) int {
//	if len(cost) == 0 {
//		return 0
//	} else if len(cost) == 1 {
//		return cost[0]
//	} else if len(cost) == 2 {
//		return min(cost[0], cost[1])
//	}
//
//	dp0, dp1 := make([]int, len(cost)), make([]int, len(cost))
//
//	dp0[0] = cost[0]
//	dp0[1] = cost[1]
//	for i := 2; i <= len(cost)-1; i++ {
//		dp0[i] = min(dp0[i-1], dp0[i-2]) + cost[i]
//	}
//	result0 := min(dp0[len(cost)-1], dp0[len(cost)-2])
//
//	dp1[0] = 1000
//	dp1[1] = cost[1]
//	for i := 2; i <= len(cost)-1; i++ {
//		dp1[i] = min(dp1[i-1], dp1[i-2]) + cost[i]
//	}
//	result1 := min(dp1[len(cost)-1], dp1[len(cost)-2])
//
//	return min(result0, result1)
//}
