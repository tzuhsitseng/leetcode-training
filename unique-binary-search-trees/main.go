package main

func numTrees(n int) int {
	numsTree := map[int]int{}
	numsTree[0] = 1
	numsTree[1] = 1

	for nodes := 2; nodes <= n; nodes++ {
		result := 0

		for root := 1; root <= nodes; root++ {
			left := root - 1
			right := nodes - root
			result += numsTree[left] * numsTree[right]
		}

		numsTree[nodes] = result
	}
	return numsTree[n]
}
