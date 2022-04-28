package main

import "fmt"

// https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	top, bot := 0, len(matrix)-1

	for top <= bot {
		mid := (bot + top) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			top = mid + 1
		} else {
			bot = mid - 1
		}
	}

	row := bot
	if row < 0 {
		return false
	}

	left, right := 0, len(matrix[0])-1

	for left <= right {
		mid := (left + right) / 2

		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	//fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	//fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
	fmt.Println(searchMatrix([][]int{{1}}, 0))
}
