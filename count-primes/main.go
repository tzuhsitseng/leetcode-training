package main

import "math"

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	primes := make([]bool, n)

	for i := range primes {
		primes[i] = true
	}

	primes[0] = false
	primes[1] = false

	for i := 2; i < int(math.Sqrt(float64(n))+1); i++ {
		if primes[i] == true {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	cnt := 0
	for _, v := range primes {
		if v {
			cnt++
		}
	}

	return cnt
}
