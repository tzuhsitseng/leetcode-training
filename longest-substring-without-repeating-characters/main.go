package main

import "strings"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	left, right := 0, 0
	sub := ""
	res := 0

	for right < len(s) {
		if strings.Contains(sub, string(s[right])) {
			res = max(res, len(sub))
			sub = ""
			left, right = left+1, left+1
		} else {
			sub = s[left : right+1]
			right++
		}
	}

	res = max(res, len(sub))
	return res
}
