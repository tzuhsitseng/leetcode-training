package main

// https://leetcode.com/problems/subsets/

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	cur := make([]int, 0)
	res := make([][]int, 0)

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			subset := make([]int, len(cur))
			copy(subset, cur)
			res = append(res, subset)
			return
		}

		cur = append(cur, nums[idx])
		dfs(idx + 1)

		cur = cur[:len(cur)-1]
		dfs(idx + 1)
	}

	dfs(0)
	return res
}
