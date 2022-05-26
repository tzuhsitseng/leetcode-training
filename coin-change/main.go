package main

// https://leetcode.com/problems/coin-change/

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := map[int]int{}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			coin := coins[j]

			if i < coin {
				continue
			}

			sub, ok := dp[i-coin]
			if !ok {
				continue
			}

			if _, ok = dp[i]; !ok {
				dp[i] = 1 + sub
			} else {
				dp[i] = min(1+sub, dp[i])
			}
		}
	}

	if v, ok := dp[amount]; ok {
		return v
	}

	return -1
}

// --
// It's a neetcode solution
// --
//func coinChange(coins []int, amount int) int {
//	dp := make([]int, amount+1)
//	for i := range dp {
//		dp[i] = amount + 1
//	}
//	dp[0] = 0
//
//	for a := 1; a <= amount; a++ {
//		for _, c := range coins {
//			if a-c >= 0 {
//				dp[a] = min(dp[a], 1+dp[a-c])
//			}
//		}
//	}
//
//	if v := dp[amount]; v != amount+1 {
//		return v
//	}
//	return -1
//}
