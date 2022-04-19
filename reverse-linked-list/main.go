package main

// https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var result *ListNode
	iter := head

	for iter != nil {
		next := iter.Next
		iter.Next = result
		result = iter
		iter = next
	}

	return result
}

// the following is recursive version

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	newHead := head
//
//	if head.Next != nil {
//		newHead = reverseList(head.Next)
//		head.Next.Next = head
//	}
//
//	head.Next = nil
//
//	return newHead
//}
