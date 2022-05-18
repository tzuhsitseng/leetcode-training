package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	l, r := 0, 0
	m := map[uint8]bool{}

	for r < len(s) {
		if m[s[r]] {
			res = max(res, r-l)
			m = map[uint8]bool{}
			l, r = l+1, l+1
		} else {
			m[s[r]] = true
			r++
		}
	}

	res = max(res, r-l)
	return res
}
