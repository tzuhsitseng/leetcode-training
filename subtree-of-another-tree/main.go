package main

// https://leetcode.com/problems/subtree-of-another-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(n *TreeNode) bool

	dfs = func(n *TreeNode) bool {
		if n == nil {
			return false
		}
		return (n.Val == subRoot.Val && isSameTree(n, subRoot)) || dfs(n.Left) || dfs(n.Right)
	}
	return dfs(root)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
