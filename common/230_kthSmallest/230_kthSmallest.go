package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// 返回数字
	// 使用栈实现中序遍历
	index := 0
	// 使用栈模拟函数调用
	var stack []TreeNode
	curr := root
	for index != k {
		// 直到去到最右边,即区间最小值
		if curr.Left != nil {
			// 入栈,继续往左寻找
			stack = append(stack, *curr)
			curr = curr.Left
			continue
		} else {
			if index == 0 {
				index = 1
				if k == 1 {
					return curr.Val
				}
			}
			// 此时curr左结点为空,根据中序遍历往右走
			if curr.Right != nil {
				curr = curr.Right
			} else {
				// 可以出栈了
				curr = &stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				curr.Left = nil
			}
			index++
		}
	}
	return curr.Val
}

func main() {
	var root TreeNode
	// [3,1,4,null,2]
	root.Val = 5
	var _1 TreeNode
	var _2 TreeNode
	var _3 TreeNode
	var _4 TreeNode
	var _5 TreeNode
	_1.Val = 3
	_2.Val = 6
	_3.Val = 2
	_4.Val = 4
	_5.Val = 1
	root.Left = &_1
	root.Right = &_2
	_1.Left = &_3
	_3.Left = &_5
	_1.Right = &_4
	/*
		[3,1,4,null,2]
		root.Val = 3
		var _1 TreeNode
		var _2 TreeNode
		var _3 TreeNode
		_1.Val = 1
		_2.Val = 4
		_3.Val = 2
		root.Left = &_1
		root.Right = &_2
		_1.Right = &_3
	*/
	/*
		root.Val = 25
		var _1 TreeNode
		var _2 TreeNode
		var _3 TreeNode
		var _4 TreeNode
		_1.Val = 28
		_2.Val = 26
		_3.Val = 27
		_4.Val = 29
		root.Right = &_1
		_1.Left = &_2
		_1.Right = &_4
		_2.Right = &_3
	*/
	fmt.Print(kthSmallest(&root, 3))
}
