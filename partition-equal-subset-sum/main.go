package main

import "fmt"

// https://leetcode.com/problems/partition-equal-subset-sum/

func canPartition(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 != 0 {
		return false
	}

	avg := sum / 2
	dp := map[int]bool{}

	for _, n := range nums {
		if n == avg {
			return true
		}

		candidates := map[int]bool{}
		for m := range dp {
			if n+m == avg {
				return true
			}

			candidates[n+m] = true
		}

		dp[n] = true
		for c := range candidates {
			dp[c] = true
		}
	}

	return false
}

func main() {
	fmt.Println(canPartition([]int{2, 2, 3, 5}))
}
