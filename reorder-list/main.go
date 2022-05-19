package main

// https://leetcode.com/problems/reorder-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}

	var prev *ListNode
	s, f := head, head
	for f != nil && f.Next != nil {
		prev = s
		s = s.Next
		f = f.Next.Next
	}

	// cut input in half
	prev.Next = nil
	l1 := head

	// reverse l2
	var l2 *ListNode
	for s != nil {
		tmp := s.Next
		s.Next = l2
		l2 = s
		s = tmp
	}

	for l1 != nil && l2 != nil {
		prev = l2
		nxt1, nxt2 := l1.Next, l2.Next
		l1.Next = l2
		l2.Next = nxt1
		l1 = nxt1
		l2 = nxt2
	}
	if l2 != nil {
		prev.Next = l2
	}
}

func main() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	//n4 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	reorderList(n1)
}
