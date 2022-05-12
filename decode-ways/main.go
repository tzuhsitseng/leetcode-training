package main

import "strings"

// https://leetcode.com/problems/decode-ways/

func numDecodings(s string) int {
	dp := map[int]int{
		len(s): 1,
	}

	var dfs func(i int) int
	dfs = func(i int) int {
		res := 0

		if v, ok := dp[i]; ok {
			return v
		}

		if s[i] == '0' {
			return 0
		}

		res = dfs(i + 1)
		if i < len(s)-1 &&
			(s[i] == '1' || s[i] == '2' &&
				strings.Index("0123456", string(s[i+1])) != -1) {
			res += dfs(i + 2)
		}
		dp[i] = res
		return res
	}

	return dfs(0)
}
