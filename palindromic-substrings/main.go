package main

// https://leetcode.com/problems/palindromic-substrings/

func countSubstrings(s string) int {
	cnt := 0

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			cnt++
			r++
			l--
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			cnt++
			r++
			l--
		}
	}

	return cnt
}
