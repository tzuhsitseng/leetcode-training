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

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	return bfs(root, result, 0)
}

func bfs(root *TreeNode, result [][]int, level int) [][]int {
	if root == nil {
		return result
	}

	if level > len(result)-1 {
		result = append(result, []int{root.Val})
	} else {
		result[level] = append(result[level], root.Val)
	}

	result = bfs(root.Left, result, level+1)
	result = bfs(root.Right, result, level+1)
	return result
}
