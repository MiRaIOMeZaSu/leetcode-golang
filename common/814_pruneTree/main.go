package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	// 深度优先法
	if dfs(root) {
		return root
	}
	return nil
}

func dfs(node *TreeNode) bool {
	if node == nil {
		return false
	}
	left := node.Left
	right := node.Right
	if !dfs(left) {
		node.Left = nil
	}
	if !dfs(right) {
		node.Right = nil
	}
	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return false
	}
	return true
}
