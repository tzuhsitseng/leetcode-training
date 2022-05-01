package main

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation/

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

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	operators := "+-*/"
	operands := NewStringStack()

	for _, v := range tokens {
		if strings.Contains(operators, v) {
			right, _ := strconv.Atoi(*operands.pop())
			left, _ := strconv.Atoi(*operands.pop())
			value := 0

			switch v {
			case "+":
				value = left + right
			case "-":
				value = left - right
			case "*":
				value = left * right
			case "/":
				value = left / right
			}

			operands.push(strconv.Itoa(value))
		} else {
			operands.push(v)
		}
	}

	res, _ := strconv.Atoi(*operands.pop())
	return res
}
