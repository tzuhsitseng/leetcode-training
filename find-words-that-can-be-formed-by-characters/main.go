package main

func countCharacters(words []string, chars string) int {
	charsHashes := map[rune]int{}

	// 先算出魔王單字所有 char 的 count
	for _, char := range chars {
		charsHashes[char]++
	}

	result := 0

	for _, word := range words {
		wordHashes := map[rune]int{}
		good := true

		for _, char := range word {
			// 如果出現了魔王單字中沒有的 char，則直接當成 not good
			if _, ok := charsHashes[char]; !ok {
				good = false
				break
			}

			wordHashes[char]++
		}

		if good {
			for char, cnt := range wordHashes {
				// 如果 char 的 count 大於原本魔王單字中該 char 的 count，則直接當成 not good
				if cnt > charsHashes[char] {
					good = false
					break
				}
			}
		}

		// 闖關成功，將該單字的 length 加到 result 中
		if good {
			result += len(word)
		}
	}

	return result
}
