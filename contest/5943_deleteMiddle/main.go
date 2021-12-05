package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	// 开始正式的双下标
	slow := head
	fast := head
	var last *ListNode
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			last.Next = slow.Next
			return head
		}
		last = slow
		slow = slow.Next
	}
	// 若此时index为0
	last.Next = slow.Next
	return head
}
