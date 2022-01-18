package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CustomTreeNode struct {
	*TreeNode
	Depth int
}

func minDepth(root *TreeNode) int {
	queue := make([]*CustomTreeNode, 0)
	var n *CustomTreeNode

	if root == nil {
		return 0
	}

	queue = enqueue(queue, &CustomTreeNode{root, 1})

	for len(queue) > 0 {
		n, queue = dequeue(queue)

		if n.Left == nil && n.Right == nil {
			return n.Depth
		}

		if n.Left != nil {
			queue = enqueue(queue, &CustomTreeNode{n.Left, n.Depth + 1})
		}

		if n.Right != nil {
			queue = enqueue(queue, &CustomTreeNode{n.Right, n.Depth + 1})
		}
	}

	return 0
}

func enqueue(queue []*CustomTreeNode, elements ...*CustomTreeNode) []*CustomTreeNode {
	return append(queue, elements...)
}

func dequeue(queue []*CustomTreeNode) (*CustomTreeNode, []*CustomTreeNode) {
	return queue[0], queue[1:]
}
