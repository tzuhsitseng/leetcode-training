package main

// https://leetcode.com/problems/reorder-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head

	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	var secondHalf *ListNode
	iter := slow
	for iter != nil {
		tmp := iter.Next
		iter.Next = secondHalf
		secondHalf = iter
		iter = tmp
	}

	iter = head
	sIter := secondHalf
	sPrev := sIter

	for iter != nil && sIter != nil {
		tmp := iter.Next
		rTmp := sIter.Next
		iter.Next = sIter
		sIter.Next = tmp
		sPrev = sIter
		iter = tmp
		sIter = rTmp
	}

	// 在 nodes 為單數的情境下，在 second half 的部分會剩下一個 e.g. [1, 2, 3, 4, 5]
	if sIter != nil {
		sPrev.Next = sIter
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
