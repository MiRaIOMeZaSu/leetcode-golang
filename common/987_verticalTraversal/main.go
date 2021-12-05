package main

import (
	"fmt"
	"sort"
)

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

const (
	INT_MAX = int(^uint32(0) >> 1)
)

var arr [2001][1000]sort.IntSlice
var borderY [2001][2]int
var borderX [2]int

func verticalTraversal(root *TreeNode) [][]int {
	// 使用数组存储两边下标,使用三维数组?
	// 若如此,遍历时将耗费巨大?
	arr = [2001][1000]sort.IntSlice{}
	borderY = [2001][2]int{}
	borderX[0] = INT_MAX
	borderX[1] = 1000
	solve(root, 0, 0)
	var ans [][]int
	for i := borderX[0]; i <= borderX[1]; i++ {
		var temp []int
		for j := borderY[i][0]; j <= borderY[i][1]; j++ {
			sort.Sort(arr[i][j])
			temp = append(temp, arr[i][j]...)
		}
		ans = append(ans, temp)
	}
	return ans
}

func solve(node *TreeNode, x int, y int) {
	targetX := x + 1000
	if y != 0 && x != 0 && borderY[targetX][0] == 0 {
		borderY[targetX][0] = INT_MAX
	}
	if borderX[0] > targetX {
		borderX[0] = targetX
	}
	if borderX[1] < targetX {
		borderX[1] = targetX
	}
	if borderY[targetX][0] > y {
		borderY[targetX][0] = y
	}
	if borderY[targetX][1] < y {
		borderY[targetX][1] = y
	}
	arr[targetX][y] = append(arr[targetX][y], node.Val)
	if node.Left != nil {
		solve(node.Left, x-1, y+1)
	}
	if node.Right != nil {
		solve(node.Right, x+1, y+1)
	}
}

func main() {
	fmt.Println(INT_MAX)
}
