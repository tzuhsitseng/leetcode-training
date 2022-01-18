package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	iter := temp

	// 若 l1 跟 l2 都還有節點就繼續作比較
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			iter.Next = l1
			l1 = l1.Next
		} else {
			iter.Next = l2
			l2 = l2.Next
		}
		iter = iter.Next
	}

	// 若 l1 還有節點，但 l2 已經空了，就直接串上 l1
	// 反之，若 l2 還有節點，但 l1 已經空了，就直接串上 l2
	if l1 != nil {
		iter.Next = l1
	} else if l2 != nil {
		iter.Next = l2
	}
	return temp.Next
}
