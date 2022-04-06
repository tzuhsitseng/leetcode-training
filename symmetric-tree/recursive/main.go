package recursive

// https://leetcode.com/problems/symmetric-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	} else if n1 != nil && n2 == nil {
		return false
	} else if n1 == nil && n2 != nil {
		return false
	}

	if n1.Val != n2.Val {
		return false
	}

	return checkSymmetric(n1.Left, n2.Right) && checkSymmetric(n1.Right, n2.Left)
}
