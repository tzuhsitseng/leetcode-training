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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 != nil && root2 == nil {
		return root1
	}

	if root1 == nil && root2 != nil {
		return root2
	}

	stack := push(make([]*TreeNode, 0), root1, root2)
	var t1, t2 *TreeNode

	for len(stack) > 0 {
		t2, stack = pop(stack)
		t1, stack = pop(stack)

		if t2 == nil {
			continue
		}

		// 這時候代表 t2 也不為 nil
		if t1 != nil {
			t1.Val += t2.Val

			if t1.Left == nil {
				t1.Left = t2.Left
			} else {
				stack = push(stack, t1.Left, t2.Left)
			}

			if t1.Right == nil {
				t1.Right = t2.Right
			} else {
				stack = push(stack, t1.Right, t2.Right)
			}
		}
	}
	return root1
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
