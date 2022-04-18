package main

import "fmt"

// https://leetcode.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counterS := map[uint8]int{}
	counterT := map[uint8]int{}

	for i := range s {
		counterS[s[i]] += 1
		counterT[t[i]] += 1
	}

	for k := range counterS {
		if counterS[k] != counterT[k] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
