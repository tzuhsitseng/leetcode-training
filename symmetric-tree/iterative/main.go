package main

// https://leetcode.com/problems/symmetric-tree

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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := NewTreeNodeQueue()
	q.enqueue(root.Left)
	q.enqueue(root.Right)

	for !q.isEmpty() {
		n1 := q.dequeue()
		n2 := q.dequeue()

		if n1 == nil && n2 == nil {
			continue
		} else if n1 == nil && n2 != nil {
			return false
		} else if n1 != nil && n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		q.enqueue(n1.Left)
		q.enqueue(n2.Right)
		q.enqueue(n1.Right)
		q.enqueue(n2.Left)
	}

	return true
}
