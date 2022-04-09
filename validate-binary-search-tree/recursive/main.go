package main

import "math"

// https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validBST(root, math.Inf(-1), math.Inf(1))
}

func validBST(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}

	v := float64(root.Val)

	return v > min && v < max && validBST(root.Left, min, v) && validBST(root.Right, v, max)
}
