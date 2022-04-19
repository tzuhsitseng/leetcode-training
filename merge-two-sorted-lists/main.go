package main

// https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := &ListNode{}
	iter := dummy

	for l1 != nil && l2 != nil {
		if l2.Val > l1.Val {
			iter.Next = l1
			l1 = l1.Next
		} else {
			iter.Next = l2
			l2 = l2.Next
		}

		iter = iter.Next
	}

	if l1 != nil {
		iter.Next = l1
	} else if l2 != nil {
		iter.Next = l2
	}

	result := dummy.Next
	return result
}
