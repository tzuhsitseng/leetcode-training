package main

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	iSize := len(prices)

	res := 0
	l, r := 0, 1

	for r < iSize {
		if prices[l] > prices[r] {
			l, r = r, r+1
			continue
		}

		res = max(res, prices[r]-prices[l])
		r++
	}

	return res
}
