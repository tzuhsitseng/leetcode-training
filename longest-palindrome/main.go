package main

// Keyword: bucket sort

func longestPalindrome(s string) int {
	p := 0
	cnt := make([]int, 128)
	for _, v := range s {
		cnt[int(v)] += 1
	}

	// 64-91 -> A-Z
	// n+32 -> a-z
	for i := 64; i < 91; i++ {
		p += cnt[i]/2*2 + cnt[i+32]/2*2
	}

	// 確認是否有落單值可以填補中間位的
	for i := 64; i < 91; i++ {
		if cnt[i]%2 == 1 || cnt[i+32]%2 == 1 {
			p += 1
			break
		}
	}
	return p
}
