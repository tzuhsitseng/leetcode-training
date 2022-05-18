package main

import (
	"regexp"
	"strings"
)

// https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("^[a-z0-9]+$")
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	if s == "" {
		return true
	}

	l, r := 0, len(s)-1
	for l <= r {
		if !reg.Match([]byte{s[l]}) {
			l++
			continue
		}
		if !reg.Match([]byte{s[r]}) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
