package main

// https://leetcode.com/problems/diameter-of-binary-tree/

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

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var dfs func(n *TreeNode) int
	dfs = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := dfs(n.Left), dfs(n.Right)
		res = max(res, l+r)
		return 1 + max(l, r)
	}

	dfs(root)
	return res
}
