package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := []*Node{root}

	for len(queue) > 0 {
		cnt := len(queue)
		sameLevel := make([]int, 0)

		for cnt > 0 {
			var p *Node
			p, queue = dequeue(queue)
			if len(p.Children) > 0 {
				for _, v := range p.Children {
					queue = enqueue(queue, v)
				}
			}
			sameLevel = append(sameLevel, p.Val)
			cnt--
		}

		result = append(result, sameLevel)
	}

	return result
}

func enqueue(queue []*Node, elements ...*Node) []*Node {
	return append(queue, elements...)
}

func dequeue(queue []*Node) (*Node, []*Node) {
	return queue[0], queue[1:]
}
