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
	stack := make([]*TreeNode, 0)
	iter := root

	for len(stack) > 0 || iter != nil {
		for iter != nil {
			stack = push(stack, iter)
			iter = iter.Left
		}
		iter, stack = pop(stack)
		result = append(result, iter.Val)
		iter = iter.Right
	}
	return result
}

func push(stack []*TreeNode, elements ...*TreeNode) []*TreeNode {
	return append(stack, elements...)
}

func pop(stack []*TreeNode) (*TreeNode, []*TreeNode) {
	l := len(stack)
	if l > 0 {
		return stack[l-1], stack[:l-1]
	}
	return nil, nil
}
