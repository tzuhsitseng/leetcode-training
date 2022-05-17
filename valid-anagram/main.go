package main

import "fmt"

// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	sc, tc := map[rune]int{}, map[rune]int{}

	for _, c := range s {
		sc[c]++
	}
	for _, c := range t {
		tc[c]++
	}

	for k := range sc {
		if sc[k] != tc[k] {
			return false
		}
		delete(tc, k)
	}
	for k := range tc {
		if tc[k] != sc[k] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
