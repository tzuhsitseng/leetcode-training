package main

import "fmt"

// https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	row := -1

	for l <= r {
		m := (l + r) / 2
		if target >= matrix[m][0] && target <= matrix[m][len(matrix[m])-1] {
			row = m
			break
		} else if target < matrix[m][0] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	if row == -1 {
		return false
	}

	l, r = 0, len(matrix[row])-1

	for l <= r {
		m := (l + r) / 2
		if target == matrix[row][m] {
			return true
		} else if target < matrix[row][m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	//fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
	//fmt.Println(searchMatrix([][]int{{1}}, 0))
}
