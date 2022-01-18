package main

// Keyword: hash table

func isAnagram(s string, t string) bool {
	counter := map[string]int{}

	// 若長度就不一樣了就肯定不是 anagram
	if len(s) != len(t) {
		return false
	}

	for _, val := range s {
		// 將 rune 先轉成 string
		character := string(val)

		// 若有找到就 count ++, 若沒有就直接從頭開始 = 1
		if _, ok := counter[character]; !ok {
			counter[character] = 1
		} else {
			counter[character] += 1
		}
	}

	for _, val := range t {
		// 將 rune 先轉成 string
		character := string(val)

		v, ok := counter[character]
		// 如果 counter 找不到或是已經 count = 0 了，就代表不是 anagram
		if !ok || v == 0 {
			return false
		}

		// 若 count 已經歸零, 就直接 delete, 以利最終可以用長度判斷是否為 anagram
		// 否則就 count --
		if v-1 == 0 {
			delete(counter, character)
		} else {
			counter[character] = v - 1
		}
	}
	return len(counter) == 0
}
