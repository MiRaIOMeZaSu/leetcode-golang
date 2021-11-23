package buildtree

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

func build(inorder []int, postorder []int, inL int, inR int, postL int, postR int) *TreeNode {
	return new(TreeNode)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 需要递归实现
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}
