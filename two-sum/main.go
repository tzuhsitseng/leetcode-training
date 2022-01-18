package main

// Keyword: hash table

func twoSum(nums []int, target int) []int {
	records := map[int]int{}
	for idx, val := range nums {
		// 若有查到 target - nums[idx] 的 hash table, 代表當前的 idx 跟查找到的資料匹配
		if v, ok := records[target-val]; ok {
			return []int{v, idx}
		}

		// 有走訪過就紀錄起來, 往後走會用到
		records[val] = idx
	}
	return []int{-1, -1}
}
