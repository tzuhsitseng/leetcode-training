package main

import "fmt"

// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := map[rune]int{}

	for _, c := range s {
		if _, ok := counter[c]; ok {
			counter[c]++
		} else {
			counter[c] = 1
		}
	}

	for _, c := range t {
		if _, ok := counter[c]; !ok {
			return false
		}
		counter[c]--
		if counter[c] == 0 {
			delete(counter, c)
		}
	}

	return len(counter) == 0
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
