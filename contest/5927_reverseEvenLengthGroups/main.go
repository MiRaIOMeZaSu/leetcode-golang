package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func handle(head *ListNode, count int) *ListNode {
	curr := head
	for i := 0; i < count-1 && curr.Next != nil; i++ {
		curr = curr.Next
	}
	return curr
}

func handleEven(pres *ListNode, head *ListNode, count int) *ListNode {
	var list []*ListNode
	curr := head
	list = append(list, curr)
	for i := 0; i < count-1 && curr.Next != nil; i++ {
		curr = curr.Next
		list = append(list, curr)
	}
	next := curr.Next
	for i := len(list) - 1; i >= 0; i-- {
		pres.Next = list[i]
		pres = pres.Next
	}
	pres.Next = next
	return pres
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	// 直接使用数组?
	// 先计算长度
	rest := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		rest++
	}
	curr = head
	target := 1
	var pres *ListNode
	for curr != nil {
		if min(target, rest)%2 != 0 {
			tail := handle(curr, target)
			pres = tail
			curr = tail.Next
		} else {
			// 偶数
			tail := handleEven(pres, curr, target)
			pres = tail
			curr = tail.Next
		}
		rest -= target
		target++
	}
	return head
}

func main() {
	head := ListNode{5, nil}
	_1 := ListNode{2, nil}
	_2 := ListNode{6, nil}
	_3 := ListNode{3, nil}
	_4 := ListNode{9, nil}
	head.Next = &_1
	_1.Next = &_2
	_2.Next = &_3
	_3.Next = &_4
	reverseEvenLengthGroups(&head)
}
