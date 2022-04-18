package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	l, r := 0, 1
	result := 0

	for r < len(prices) {
		if prices[r] < prices[l] {
			l, r = r, r+1
			continue
		}

		if p := prices[r] - prices[l]; p > result {
			result = p
		}
		r++
	}

	return result
}
