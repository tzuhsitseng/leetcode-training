package main

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	iter := root
	for iter != nil {
		if p.Val > iter.Val && q.Val > iter.Val {
			iter = iter.Right
		} else if p.Val < iter.Val && q.Val < iter.Val {
			iter = iter.Left
		} else {
			return iter
		}

	}
	return nil
}
