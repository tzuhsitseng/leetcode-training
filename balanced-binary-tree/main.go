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
	return getDepth(root) != -1
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getDepth(node.Left)
	right := getDepth(node.Right)

	if left == -1 || right == -1 || math.Abs(float64(left-right)) > 1 {
		return -1
	}

	return 1 + int(math.Max(float64(left), float64(right)))
}
