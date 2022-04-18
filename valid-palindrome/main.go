package main

import (
	"regexp"
	"strings"
)

// https://leetcode.com/problems/valid-palindrome/

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("^[a-z0-9]+$")

	if len(s) == 0 {
		return true
	}

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	if len(s) == 0 || len(s) == 1 {
		return true
	}

	l := 0
	r := len(s) - 1

	for l < r {
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

func main() {

}
