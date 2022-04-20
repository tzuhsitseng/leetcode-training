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
	if root == nil {
		return true
	}

	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	if left == -1 || right == -1 || math.Abs(float64(left-right)) > 1 {
		return -1
	}

	return 1 + max(left, right)
}
