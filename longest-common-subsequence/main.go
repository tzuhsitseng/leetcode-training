package main

// https://leetcode.com/problems/longest-common-subsequence/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[0][0]
}

// --
// The following solution will be "Time Limit Exceeded", although it use cache
// --

//func longestCommonSubsequence(text1 string, text2 string) int {
//	var dfs func(sub1, sub2 string) int
//	dfs = func(sub1, sub2 string) int {
//		if len(sub1) == 0 || len(sub2) == 0 {
//			return 0
//		}
//
//		if sub1[0] == sub2[0] {
//			d := dfs(sub1[1:], sub2[1:])
//			return 1 + d
//		}
//
//		d1, d2 := dfs(sub1, sub2[1:]), dfs(sub1[1:], sub2)
//		return max(d1, d2)
//	}
//
//	return dfs(text1, text2)
//}
