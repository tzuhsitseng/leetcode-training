package main

// https://leetcode.com/problems/clone-graph/

type Node struct {
	Val       int
	Neighbors []*Node
}

type NodeQueue struct {
	queue []*Node
}

func NewNodeQueue() *NodeQueue {
	return &NodeQueue{
		queue: make([]*Node, 0),
	}
}

func (q *NodeQueue) enqueue(element *Node) {
	if q.queue == nil {
		q.queue = make([]*Node, 0)
	}
	q.queue = append(q.queue, element)
}

func (q *NodeQueue) dequeue() *Node {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	q.queue = q.queue[1:]
	return result
}

func (q *NodeQueue) peek() *Node {
	if len(q.queue) == 0 {
		return nil
	}
	result := q.queue[0]
	return result
}

func (q *NodeQueue) isEmpty() bool {
	return len(q.queue) == 0
}

//func cloneGraph(node *Node) *Node {
//	if node == nil {
//		return nil
//	}
//
//	result := map[int]*Node{}
//	result[node.Val] = &Node{Val: node.Val}
//	q := NewNodeQueue()
//	q.enqueue(node)
//
//	for !q.isEmpty() {
//		n := q.dequeue()
//		newNei := make([]*Node, 0, len(n.Neighbors))
//
//		for _, v := range n.Neighbors {
//			if _, ok := result[v.Val]; !ok {
//				q.enqueue(v)
//				result[v.Val] = &Node{Val: v.Val}
//			}
//			newNei = append(newNei, result[v.Val])
//		}
//
//		result[n.Val].Neighbors = newNei
//	}
//
//	return result[1]
//}

// --
// The following solution is a dfs version

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := map[int]*Node{}
	var clone func(*Node) *Node

	clone = func(n *Node) *Node {
		if v, ok := visited[n.Val]; ok {
			return v
		}

		newNode := &Node{Val: n.Val}

		for _, v := range n.Neighbors {
			newNei := clone(v)
			newNode.Neighbors = append(newNode.Neighbors, newNei)
		}

		visited[n.Val] = newNode
		return newNode
	}

	return clone(node)
}
