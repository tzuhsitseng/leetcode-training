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

func deleteDuplicates(head *ListNode) *ListNode {
	ite := head

	for ite != nil {
		temp := ite.Next

		// 將 temp 走到不相等為止
		for temp != nil && temp.Val == ite.Val {
			temp = temp.Next
		}

		// 將當前 node 的 next 指到最新的 temp node，且將 iterator 移到最新 node 繼續往下走
		ite.Next = temp
		ite = temp
	}
	return head
}
