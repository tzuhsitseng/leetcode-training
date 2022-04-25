package main

import "fmt"

// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, str := range strs {
		counter := [26]int{}

		for _, c := range str {
			idx := int(c) - int('a')
			counter[idx]++
		}

		key := fmt.Sprint(counter)
		m[key] = append(m[key], str)
	}

	result := make([][]string, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
