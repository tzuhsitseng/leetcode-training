package main

import "strings"

// https://leetcode.com/problems/generate-parentheses/

type StringStack struct {
	stack []string
}

func NewStringStack() *StringStack {
	return &StringStack{
		stack: make([]string, 0),
	}
}

func (s *StringStack) push(element string) {
	if s.stack == nil {
		s.stack = make([]string, 0)
	}
	s.stack = append(s.stack, element)
}

func (s *StringStack) pop() *string {
	if s.isEmpty() {
		return nil
	}

	l := len(s.stack)
	result := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return &result
}

func (s *StringStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *StringStack) peek() *string {
	if s.isEmpty() {
		return nil
	}

	return &s.stack[len(s.stack)-1]
}

func (s *StringStack) getSlice() []string {
	return s.stack
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	res := make([]string, 0)
	stack := NewStringStack()

	var backtrack func(openN, closedN int)
	backtrack = func(openN, closedN int) {
		if openN == closedN && n == openN {
			res = append(res, strings.Join(stack.getSlice(), ""))
			return
		}

		if openN < n {
			stack.push("(")
			backtrack(openN+1, closedN)
			stack.pop()
		}

		if closedN < openN {
			stack.push(")")
			backtrack(openN, closedN+1)
			stack.pop()
		}
	}

	backtrack(0, 0)
	return res
}
