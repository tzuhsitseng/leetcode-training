package recursive

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
	return check(root.Left, root.Right)
}

func check(l1, l2 *TreeNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}

	if l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		return check(l1.Left, l2.Right) && check(l1.Right, l2.Left)
	}
	return false
}
