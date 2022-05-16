package main

// https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	dic := map[string]bool{}
	for _, word := range wordDict {
		dic[word] = true
	}

	mem := map[string]bool{}
	var dfs func(sub string) bool
	dfs = func(sub string) bool {
		if sub == "" {
			return true
		}

		if v, ok := mem[sub]; ok {
			return v
		}

		for i := range sub {
			if ok := dic[sub[:i+1]]; ok {
				if dfs(sub[i+1:]) {
					mem[sub[i+1:]] = true
					return true
				}
			}
		}

		mem[sub] = false
		return false
	}

	return dfs(s)
}
