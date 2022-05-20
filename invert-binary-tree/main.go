package main

// https://leetcode.com/problems/invert-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}

		n.Left, n.Right = n.Right, n.Left

		if n.Left != nil {
			dfs(n.Left)
		}
		if n.Right != nil {
			dfs(n.Right)
		}
	}

	dfs(root)
	return root
}
