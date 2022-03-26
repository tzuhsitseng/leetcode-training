package main

import "fmt"

// https://leetcode.com/problems/first-bad-version/

func isBadVersion(version int) bool {
	// ... (provided by leetcode)
	return false
}

func firstBadVersion(n int) int {
	lo, hi := 1, n

	if isBadVersion(1) {
		return 1
	}

	for hi >= lo {
		mi := (lo + hi) / 2

		if isBadVersion(mi) {
			if !isBadVersion(mi - 1) {
				return mi
			}

			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(firstBadVersion(5))
}
