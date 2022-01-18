package main

func hasAlternatingBits(n int) bool {
	n = (n >> 1) ^ n
	return n&(n+1) == 0
}
