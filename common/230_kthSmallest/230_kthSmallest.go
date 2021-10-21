package main

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
	index := 0
	// 使用栈模拟函数调用
	var stack []TreeNode
	curr := root
	for {
		// 直到去到最右边,即区间最小值
		if curr.Left != nil {
			stack = append(stack, *curr)
			curr = curr.Left
			continue
		} else if index == 0 {
			index = 1
			if k == 1 {
				return curr.Val
			}
		} else {
			for {
				curr = &stack[len(stack)-1]
				if curr.Left != nil {
					stack = stack[:len(stack)-1]
					break
				}
				stack = stack[:len(stack)-1]
				index++
				if k == index {
					return curr.Val
				}
			}
		}
	}
}
