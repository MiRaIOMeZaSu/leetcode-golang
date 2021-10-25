package distancek

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

var ans []int
var path []bool
var flag = -1

func search(root *TreeNode, index int, isLeft bool, target int) {
	if len(path) <= index {
		path = append(path, isLeft)
	} else {
		path[index] = isLeft
	}

	if root.Val == target {
		flag = index
		return
	}

	if root.Left != nil && flag == -1 {
		search(root.Left, index+1, true, target)
	}
	if root.Right != nil && flag == -1 {
		search(root.Right, index+1, false, target)
	}
}

func solve(root *TreeNode, index int, k int) {
	// 使用中序遍历
	curr := len(path) - index
	if curr == 0 {
		return
	}
	if curr == k {
		ans = append(ans, root.Val)
	}
	next := curr + 1
	fmt.Println(next)
	if path[index] {
		// 此时往左走是会找到target
		if next <= k && root.Right != nil {
			add(root.Right, next, k)
		}
		solve(root.Left, index+1, k)
	} else {
		// 此时往右走会找到target
		if next <= k && root.Left != nil {
			add(root.Left, next, k)
		}
		solve(root.Right, index+1, k)
	}

}

func add(root *TreeNode, val int, k int) {
	if val == k {
		ans = append(ans, root.Val)
	}
	if root.Left != nil {
		add(root.Left, val+1, k)
	}
	if root.Right != nil {
		add(root.Right, val+1, k)
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// 结点的值是唯一的
	// 一边是前序遍历,一边是中序遍历
	// 还是先找到目的结点
	ans = []int{}
	path = []bool{}
	flag = -1
	if k == 0 {
		ans = append(ans, target.Val)
		return ans
	}
	search(root, 0, false, target.Val)
	// 此时已成功获得应该走的路径
	path = path[:flag+1]
	solve(root, 1, k)
	if target.Left != nil {
		add(target.Left, 1, k)
	}
	if target.Right != nil {
		add(target.Right, 1, k)
	}
	return ans
}
