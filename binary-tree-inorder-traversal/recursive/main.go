package main

// https://leetcode.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return travelByDFS(root, make([]int, 0))
}

func travelByDFS(n *TreeNode, result []int) []int {
	if n == nil {
		return result
	}

	// left
	if n.Left != nil {
		result = travelByDFS(n.Left, result)
	}

	// middle
	result = append(result, n.Val)

	// right
	if n.Right != nil {
		result = travelByDFS(n.Right, result)
	}

	return result
}
