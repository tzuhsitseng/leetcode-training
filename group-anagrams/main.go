package main

import "fmt"

// https://leetcode.com/problems/group-anagrams/

func groupAnagrams(strs []string) [][]string {
	gm := map[string][]string{}

	for _, str := range strs {
		m := map[rune]int{}

		for _, c := range str {
			m[c]++
		}

		key := fmt.Sprint(m)
		if _, ok := gm[key]; ok {
			gm[key] = append(gm[key], str)
		} else {
			gm[key] = []string{str}
		}
	}

	res := make([][]string, 0, len(gm))
	for _, s := range gm {
		res = append(res, s)
	}

	return res
}
