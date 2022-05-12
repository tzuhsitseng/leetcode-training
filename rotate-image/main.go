package main

// https://leetcode.com/problems/rotate-image/

func rotate(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		first := i
		last := length - 1 - i

		for j := first; j < last; j++ {
			offset := j - first
			tmp := matrix[first][j]
			matrix[first][j] = matrix[last-offset][first]
			matrix[last-offset][first] = matrix[last][last-offset]
			matrix[last][last-offset] = matrix[j][last]
			matrix[j][last] = tmp
		}
	}
}
