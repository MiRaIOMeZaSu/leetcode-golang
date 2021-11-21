package main

type Node struct {
	Val      int
	Children []*Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for i := 0; i < len(root.Children); i++ {
		maxChildDepth = max(maxChildDepth, maxDepth(root.Children[i]))
	}
	return maxChildDepth + 1
}
