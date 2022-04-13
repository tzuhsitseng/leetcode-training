package main

import "fmt"

// https://leetcode.com/problems/sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	left, right := head, slow
	left, right = sortList(left), sortList(right)
	result := &ListNode{}
	iter := result

	for left != nil && right != nil {
		if left.Val < right.Val {
			iter.Next = left
			left = left.Next
		} else {
			iter.Next = right
			right = right.Next
		}
		iter = iter.Next
	}

	for left != nil {
		iter.Next = left
		left = left.Next
		iter = iter.Next
	}

	for right != nil {
		iter.Next = right
		right = right.Next
		iter = iter.Next
	}

	return result.Next
}

func main() {
	n1 := &ListNode{Val: 3, Next: nil}
	n2 := &ListNode{Val: 1, Next: n1}
	n3 := &ListNode{Val: 2, Next: n2}
	n4 := &ListNode{Val: 4, Next: n3}
	result := sortList(n4)
	fmt.Println(result, result.Next, result.Next.Next, result.Next.Next.Next)
}
