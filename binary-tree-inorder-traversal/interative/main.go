package main

import "fmt"

// https://leetcode.com/problems/binary-tree-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeStack struct {
	stack []*TreeNode
}

func NewTreeNodeStack() *TreeNodeStack {
	return &TreeNodeStack{
		stack: make([]*TreeNode, 0),
	}
}

func (s *TreeNodeStack) push(element *TreeNode) {
	if s.stack == nil {
		s.stack = make([]*TreeNode, 0)
	}
	s.stack = append(s.stack, element)
}

func (s *TreeNodeStack) pop() *TreeNode {
	if s.isEmpty() {
		return nil
	}

	l := len(s.stack)
	result := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return result
}

func (s *TreeNodeStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *TreeNodeStack) peek() *TreeNode {
	if s.isEmpty() {
		return nil
	}

	return s.stack[len(s.stack)-1]
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := NewTreeNodeStack()
	cur := root

	for cur != nil || !stack.isEmpty() {
		for cur != nil {
			stack.push(cur)
			cur = cur.Left
		}

		cur = stack.pop()

		if cur != nil {
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}

	return result
}

func main() {
	t := &TreeNode{Val: 1}
	fmt.Println(inorderTraversal(t))
}
