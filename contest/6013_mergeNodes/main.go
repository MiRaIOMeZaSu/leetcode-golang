package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	curr := head.Next
	ans := &ListNode{
		-1,
		nil,
	}
	curr_ans := ans
	for curr != nil && curr.Next != nil {
		temp := 0
		for curr.Val != 0 {
			temp += curr.Val
			curr = curr.Next
		}
		curr = curr.Next
		curr_ans.Next = &ListNode{
			temp, nil,
		}
		curr_ans = curr_ans.Next
	}
	return ans.Next
}
