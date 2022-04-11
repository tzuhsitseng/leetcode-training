package main

// https://leetcode.com/problems/longest-palindrome/

func longestPalindrome(s string) int {
	result := 0
	odd := false
	counter := map[rune]int{}

	for _, c := range s {
		counter[c]++
	}

	for _, v := range counter {
		result += v / 2 * 2

		if !odd && v%2 != 0 {
			odd = true
		}
	}

	if odd {
		result += 1
	}

	return result
}
