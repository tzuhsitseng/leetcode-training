package main

// https://leetcode.com/problems/min-stack/

type IntegerStack struct {
	stack []int
}

func NewIntegerStack() *IntegerStack {
	return &IntegerStack{
		stack: make([]int, 0),
	}
}

func (s *IntegerStack) push(element int) {
	if s.stack == nil {
		s.stack = make([]int, 0)
	}
	s.stack = append(s.stack, element)
}

func (s *IntegerStack) pop() *int {
	if s.isEmpty() {
		return nil
	}

	l := len(s.stack)
	result := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return &result
}

func (s *IntegerStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *IntegerStack) peek() *int {
	if s.isEmpty() {
		return nil
	}

	return &s.stack[len(s.stack)-1]
}

type MinStack struct {
	IntegerStack
	min IntegerStack
}

func Constructor() MinStack {
	return MinStack{
		IntegerStack: *NewIntegerStack(),
		min:          *NewIntegerStack(),
	}
}

func (this *MinStack) Push(val int) {
	this.push(val)

	if this.min.isEmpty() || val <= *this.min.peek() {
		this.min.push(val)
	}
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		return
	}

	l := len(this.stack)
	result := this.stack[l-1]
	this.stack = this.stack[:l-1]

	if min := this.min.peek(); min != nil && result == *min {
		this.min.pop()
	}
}

func (this *MinStack) Top() int {
	return *this.peek()
}

func (this *MinStack) GetMin() int {
	if this.min.isEmpty() {
		return 0
	}
	return *this.min.peek()
}
