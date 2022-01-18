package main

func validPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return checkPalindrome(s, i+1, j) || checkPalindrome(s, i, j-1)
		}
		i, j = i+1, j-1
	}
	return true
}

func checkPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}
