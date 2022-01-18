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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var pre *TreeNode
	curr := root
	stack := make([]*TreeNode, 0)

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = push(stack, curr)
			curr = curr.Left
		}
		curr, stack = pop(stack)
		if pre != nil && pre.Val >= curr.Val {
			return false
		}
		curr, pre = curr.Right, curr
	}
	return true
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
