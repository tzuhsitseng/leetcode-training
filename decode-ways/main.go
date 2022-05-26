package main

import (
	"strings"
)

// https://leetcode.com/problems/decode-ways/

func numDecodings(s string) int {
	dp := map[int]int{}

	var dfs func(idx int) int
	dfs = func(idx int) int {
		res := 0

		if v, ok := dp[idx]; ok {
			return v
		}

		if idx == len(s) {
			return 1
		}

		if s[idx] == '0' {
			return 0
		}

		res += dfs(idx + 1)

		if idx < len(s)-1 &&
			(s[idx] == '1' ||
				(s[idx] == '2' && strings.Contains("0123456", string(s[idx+1])))) {
			res += dfs(idx + 2)
		}

		dp[idx] = res
		return res
	}

	return dfs(0)
}
