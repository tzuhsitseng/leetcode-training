package main

import "fmt"

// https://leetcode.com/problems/interleaving-string/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := map[string]bool{}

	var dfs func(i1, i2 int) bool
	dfs = func(i1, i2 int) bool {
		if i1 == len(s1) && i2 == len(s2) {
			return true
		}

		key := fmt.Sprint(i1, i2)
		if v, ok := dp[key]; ok {
			return v
		}

		if i1 < len(s1) && s1[i1] == s3[i1+i2] {
			if dfs(i1+1, i2) {
				dp[key] = true
				return true
			}
		}
		if i2 < len(s2) && s2[i2] == s3[i1+i2] {
			if dfs(i1, i2+1) {
				dp[key] = true
				return true
			}
		}

		dp[key] = false
		return dp[key]
	}
	return dfs(0, 0)
}
