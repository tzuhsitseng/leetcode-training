package main

func plusOne(digits []int) []int {
	size := len(digits)
	for i := size - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
