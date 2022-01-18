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

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	return dfs(root, result)
}

func dfs(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	result = dfs(root.Left, result)
	result = append(result, root.Val)
	result = dfs(root.Right, result)
	return result
}
