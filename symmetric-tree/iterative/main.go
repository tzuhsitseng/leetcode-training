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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := enqueue(make([]*TreeNode, 0), root.Left, root.Right)
	var l1, l2 *TreeNode

	for len(queue) > 0 {
		l1, queue = dequeue(queue)
		l2, queue = dequeue(queue)

		if l1 == nil && l2 == nil {
			continue
		}

		if l1 != nil && l2 != nil {
			if l1.Val != l2.Val {
				return false
			}
			queue = enqueue(queue, l1.Left, l2.Right, l1.Right, l2.Left)
			continue
		}
		return false
	}
	return true
}

func enqueue(queue []*TreeNode, elements ...*TreeNode) []*TreeNode {
	return append(queue, elements...)
}

func dequeue(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	return queue[0], queue[1:]
}
