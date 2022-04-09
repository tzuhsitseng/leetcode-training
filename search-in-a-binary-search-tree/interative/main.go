package main

// https://leetcode.com/problems/search-in-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	curr := root

	for curr != nil {
		if curr.Val == val {
			return curr
		} else if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return nil
}
