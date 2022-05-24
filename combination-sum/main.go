package main

// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)

	var dfs func(int, int)
	dfs = func(idx, sum int) {
		if sum == target {
			com := make([]int, len(cur))
			copy(com, cur)
			res = append(res, com)
			return
		}

		if sum > target || idx == len(candidates) {
			return
		}

		cur = append(cur, candidates[idx])
		dfs(idx, sum+candidates[idx])

		cur = cur[:len(cur)-1]
		dfs(idx+1, sum)
	}

	dfs(0, 0)
	return res
}

// --
// The following solution is v1 version
// --

//func combinationSum(candidates []int, target int) [][]int {
//	if len(candidates) == 0 {
//		return nil
//	}
//
//	res := make([][]int, 0)
//	cur := make([]int, 0)
//
//	var dfs func(int, int)
//	dfs = func(idx, sum int) {
//		if idx == len(candidates) || sum > target {
//			return
//		}
//
//		if sum == target {
//			data := make([]int, len(cur))
//			copy(data, cur)
//			res = append(res, data)
//			return
//		}
//
//		for i := idx; i < len(candidates); i++ {
//			cur = append(cur, candidates[i])
//			dfs(i, sum+candidates[i])
//			cur = cur[:len(cur)-1]
//		}
//	}
//
//	dfs(0, 0)
//	return res
//}
