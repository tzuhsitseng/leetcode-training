package main

func repeatedNTimes(nums []int) int {
	m := map[int]bool{}

	for _, n := range nums {
		if exists := m[n]; exists {
			return n
		}
		m[n] = true
	}

	return -1
}
