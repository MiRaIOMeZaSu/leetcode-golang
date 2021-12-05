package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func findTilt(root *TreeNode) int {
	// 求取坡度
	// 深度优先?!
	ans = 0
	solve(root)
	return ans
}
func solve(root *TreeNode) int {
	// 只需返回和
	if root == nil {
		return 0
	}
	sumL := solve(root.Left)
	sumR := solve(root.Right)
	ans += abs(sumL - sumR)
	return sumL + sumR + root.Val
}

func main() {
	var root TreeNode
	var _1 TreeNode
	var _2 TreeNode
	root.Val = 1
	_1.Val = 2
	_2.Val = 3
	root.Left = &_1
	root.Right = &_2
	findTilt(&root)
	fmt.Print(ans)
}
