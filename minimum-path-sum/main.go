package main

import "math"

// https://leetcode.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	r, c := len(grid), len(grid[0])
	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	result[0][0] = grid[0][0]
	for i := 1; i < c; i++ {
		result[0][i] = result[0][i-1] + grid[0][i]
	}
	for i := 1; i < r; i++ {
		result[i][0] = result[i-1][0] + grid[i][0]
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			up := result[i-1][j]
			left := result[i][j-1]
			min := int(math.Min(float64(up), float64(left)))
			result[i][j] = min + grid[i][j]
		}
	}

	return result[r-1][c-1]
}
