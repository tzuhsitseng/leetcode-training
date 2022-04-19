package main

// https://leetcode.com/problems/valid-parentheses/

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

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	brackets := map[string]uint8{
		"(": ')',
		"{": '}',
		"[": ']',
	}

	stack := NewStringStack()

	for i := 0; i < len(s); i++ {
		if _, ok := brackets[string(s[i])]; ok {
			stack.push(string(s[i]))
		} else {
			left := stack.pop()
			right := s[i]
			if left == nil || brackets[*left] != right {
				return false
			}
		}
	}

	return stack.isEmpty()
}
