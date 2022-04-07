package main

// https://leetcode.com/problems/merge-two-binary-trees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeQueue struct {
	queue []*TreeNode
}

func NewTreeNodeQueue() *TreeNodeQueue {
	return &TreeNodeQueue{
		queue: make([]*TreeNode, 0),
	}
}

func (q *TreeNodeQueue) enqueue(element *TreeNode) {
	if q.queue == nil {
		q.queue = make([]*TreeNode, 0)
	}
	q.queue = append(q.queue, element)
}

func (q *TreeNodeQueue) dequeue() *TreeNode {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *TreeNodeQueue) peek() *TreeNode {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	return result
}

func (q *TreeNodeQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil && root2 != nil {
		return root2
	} else if root1 != nil && root2 == nil {
		return root1
	}

	q := NewTreeNodeQueue()
	q.enqueue(root1)
	q.enqueue(root2)

	for !q.isEmpty() {
		n1 := q.dequeue()
		n2 := q.dequeue()

		if n2 == nil {
			continue
		}

		if n1 != nil {
			n1.Val += n2.Val

			if n1.Left == nil {
				n1.Left = n2.Left
			} else {
				q.enqueue(n1.Left)
				q.enqueue(n2.Left)
			}

			if n1.Right == nil {
				n1.Right = n2.Right
			} else {
				q.enqueue(n1.Right)
				q.enqueue(n2.Right)
			}
		}
	}

	return root1
}
