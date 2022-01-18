package main

import (
	"strings"
)

func addBinary(a string, b string) string {
	result := make([]string, 0)
	carry := "0"
	i, j := len(a)-1, len(b)-1

	for carry == "1" || i >= 0 || j >= 0 {
		var ai, bj string
		if i >= 0 {
			ai = string(a[i])
		} else {
			ai = "0"
		}
		if j >= 0 {
			bj = string(b[j])
		} else {
			bj = "0"
		}

		if ai == bj {
			result = append(result, carry)
			carry = ai
		} else {
			if carry == "0" {
				result = append(result, "1")
			} else {
				result = append(result, "0")
			}
		}

		i--
		j--
	}

	return reverse(strings.Join(result, ""))
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
