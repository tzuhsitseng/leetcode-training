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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	q := NewTreeNodeQueue()
	q.enqueue(root)

	for !q.isEmpty() {
		items := make([]*TreeNode, 0)
		values := make([]int, 0)
		for !q.isEmpty() {
			item := q.dequeue()
			items = append(items, item)
			values = append(values, item.Val)
		}
		for _, v := range items {
			if v.Left != nil {
				q.enqueue(v.Left)
			}
			if v.Right != nil {
				q.enqueue(v.Right)
			}
		}
		res = append(res, values)
	}

	return res
}
