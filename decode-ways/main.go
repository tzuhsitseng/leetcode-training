package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/decode-ways/

func numDecodings(s string) int {
	dp := map[int]int{
		len(s): 1,
	}
	//result := findByRecursive(s, 0, dp)
	result := findByIterative(s, dp)
	return result
}

func findByIterative(s string, dp map[int]int) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
		}

		if i < len(s)-1 &&
			(s[i] == '1' ||
				s[i] == '2' && strings.Index("0123456", string(s[i+1])) != -1) {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

func findByRecursive(s string, i int, dp map[int]int) int {
	if v, ok := dp[i]; ok {
		return v
	}

	if s[i] == '0' {
		return 0
	}

	result := findByRecursive(s, i+1, dp)

	if i < len(s)-1 &&
		(s[i] == '1' ||
			s[i] == '2' && strings.Index("0123456", string(s[i+1])) != -1) {
		result += findByRecursive(s, i+2, dp)
	}

	dp[i] = result

	return result
}

func main() {
	fmt.Println(numDecodings("226"))
}
