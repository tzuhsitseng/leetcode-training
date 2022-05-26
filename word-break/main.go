package main

// https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	dictMap := map[string]bool{}
	dp := map[string]bool{}

	for _, w := range wordDict {
		dictMap[w] = true
	}

	var dfs func(str string) bool
	dfs = func(str string) bool {
		if str == "" {
			return true
		}

		if v, ok := dp[str]; ok {
			return v
		}

		for i := range str {
			if _, ok := dictMap[str[:i+1]]; ok {
				if dfs(str[i+1:]) {
					dp[str] = true
					return true
				}
			}
		}

		dp[str] = false
		return false
	}
	return dfs(s)
}
