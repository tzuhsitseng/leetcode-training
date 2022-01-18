package main

func twoSum(numbers []int, target int) []int {
	temp := map[int]int{}
	for numIdx, v := range numbers {
		if tempIdx, ok := temp[target-v]; ok {
			return []int{tempIdx + 1, numIdx + 1}
		}
		temp[v] = numIdx
	}
	return nil
}
