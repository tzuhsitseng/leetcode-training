package main

// https://leetcode.com/problems/binary-tree-level-order-traversal/

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
	q.queue = append(q.queue, element)
}

func (q *TreeNodeQueue) dequeue() *TreeNode {
	if q.isEmpty() {
		return nil
	}
	res := q.queue[0]
	q.queue = q.queue[1:]
	return res
}

func (q *TreeNodeQueue) peek() *TreeNode {
	if q.isEmpty() {
		return nil
	}
	return q.queue[0]
}

func (q *TreeNodeQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	q := NewTreeNodeQueue()
	q.enqueue(root)

	for !q.isEmpty() {
		nodes := make([]*TreeNode, 0)
		sub := make([]int, 0)

		for !q.isEmpty() {
			nodes = append(nodes, q.dequeue())
		}
		for _, n := range nodes {
			sub = append(sub, n.Val)
			if n.Left != nil {
				q.enqueue(n.Left)
			}
			if n.Right != nil {
				q.enqueue(n.Right)
			}
		}

		res = append(res, sub)
	}

	return res
}
