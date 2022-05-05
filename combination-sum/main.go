package main

// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	res := make([][]int, 0)
	cur := make([]int, 0)

	var dfs func(int, int)
	dfs = func(idx, sum int) {
		if sum > target {
			return
		}

		if sum == target {
			data := make([]int, len(cur))
			copy(data, cur)
			res = append(res, data)
			return
		}

		for i := idx; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			dfs(i, sum+candidates[i])
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, 0)
	return res
}
