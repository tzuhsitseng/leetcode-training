package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/happy-number/

func sqrt(n int) int {
	return n * n
}

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	visited := map[int]bool{}

	for {
		s := strconv.Itoa(n)
		sum := 0

		for _, c := range s {
			i, _ := strconv.Atoi(string(c))
			sum += sqrt(i)
		}

		if sum == 1 {
			return true
		}

		if visited[sum] {
			return false
		}

		visited[sum] = true
		n = sum
	}
}

func main() {
	fmt.Println(isHappy(19))
}
