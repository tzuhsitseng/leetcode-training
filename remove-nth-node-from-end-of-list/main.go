package main

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l, r := head, head
	for i := 0; i < n; i++ {
		r = r.Next
	}

	var prev *ListNode
	for r != nil {
		prev = l
		l = l.Next
		r = r.Next
	}

	if prev == nil { // means the first element should be removed
		head = head.Next
	} else {
		prev.Next = l.Next
	}

	return head
}

// --
// The following solution is in one pass
// --

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//
//	left, right := head, head
//	for n > 0 && right != nil {
//		right = right.Next
//		n--
//	}
//
//	if right == nil {
//		return left.Next
//	}
//
//	var prev *ListNode
//	for right != nil {
//		prev = left
//		left = left.Next
//		right = right.Next
//	}
//	prev.Next = left.Next
//	return head
//}
