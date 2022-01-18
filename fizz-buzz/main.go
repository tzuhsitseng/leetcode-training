package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	objs := []struct {
		Key int
		Val string
	}{
		{Key: 3, Val: "Fizz"},
		{Key: 5, Val: "Buzz"},
	}
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		str := ""
		for _, obj := range objs {
			if i%obj.Key == 0 {
				str += obj.Val
			}
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		result = append(result, str)
	}
	return result
}
