package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	ite := head

	for ite != nil {
		tmp := ite.Next

		for tmp != nil && tmp.Val == ite.Val {
			tmp = tmp.Next
		}

		ite.Next = tmp
		ite = ite.Next
	}

	return head
}

//func deleteDuplicates(head *ListNode) *ListNode {
//	ite := head
//
//	for ite != nil {
//		tmp := ite.Val
//		subIte := ite.Next
//
//		for {
//			if subIte == nil || tmp != subIte.Val {
//				ite.Next = subIte
//				break
//			}
//			subIte = subIte.Next
//		}
//
//		ite = ite.Next
//	}
//
//	return head
//}
