package main

import "math"

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

func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := depth(node.Left)
	r := depth(node.Right)

	// 若發現深度已經不匹配或是深度差異大於 1 了，就回傳 -1 代表 unbalanced
	if l == -1 || r == -1 || math.Abs(float64(l-r)) > 1 {
		return -1
	}

	// 否則深度就是自己 + 左右子樹 max depth
	return 1 + int(math.Max(float64(l), float64(r)))
}
