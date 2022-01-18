package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	iter := root

	for iter != nil {
		if val == iter.Val {
			return iter
		}
		if val < iter.Val {
			iter = iter.Left
		} else {
			iter = iter.Right
		}
	}

	return nil
}
