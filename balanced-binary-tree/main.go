package main

import "math"

// https://leetcode.com/problems/balanced-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalanced(root *TreeNode) bool {
	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := dfs(n.Left), dfs(n.Right)
		if l == -1 || r == -1 || math.Abs(float64(l)-float64(r)) > 1 {
			return -1
		}
		return 1 + max(l, r)
	}
	return dfs(root) != -1
}
