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
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	var curr *TreeNode

	for len(queue) > 0 {
		temp := make([]int, 0)
		for range queue {
			curr, queue = dequeue(queue)
			temp = append(temp, curr.Val)
			if curr.Left != nil {
				queue = enqueue(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = enqueue(queue, curr.Right)
			}
		}
		result = append(result, temp)
	}
	return result
}

func enqueue(queue []*TreeNode, elements ...*TreeNode) []*TreeNode {
	return append(queue, elements...)
}

func dequeue(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	return queue[0], queue[1:]
}
