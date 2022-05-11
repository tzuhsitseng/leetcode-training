package main

// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	res := ""

	for i := range s {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			size := r - l + 1
			if size > len(res) {
				res = s[l : r+1]
			}
			r++
			l--
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			size := r - l + 1
			if size > len(res) {
				res = s[l : r+1]
			}
			r++
			l--
		}
	}

	return res
}
