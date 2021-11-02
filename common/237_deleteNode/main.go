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

func deleteNode(node *ListNode) {
	var lastNode *ListNode
	lastNode = node
	node.Val = node.Next.Val
	node = node.Next
	for node.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
		lastNode = lastNode.Next
	}
	lastNode.Next = nil
}
