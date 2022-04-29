package main

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	size := 0
	iter := head
	for iter != nil {
		size++
		iter = iter.Next
	}

	target := size - n + 1
	dummy := &ListNode{}
	dummy.Next = head
	iter = dummy
	prev := iter
	i := 0

	for iter != nil {
		if i == target {
			prev.Next = iter.Next
			break
		}

		prev = iter
		iter = iter.Next
		i++
	}

	return dummy.Next
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
