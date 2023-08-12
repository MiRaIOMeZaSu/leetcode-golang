package main

import "fmt"

func minInt(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minFallingPathSum(grid [][]int) int {
	length := len(grid)
	if length == 1 {
		return grid[0][0]
	}
	eachMax := grid[1]
	for i := 1; i < length; i++ {
		eachMax = grid[i]
		for j := 0; j < length; j++ {
			if j == 0 {
				eachMax[j] += grid[i-1][j+1]
			} else if j == length-1 {
				eachMax[j] += grid[i-1][j-1]
			} else {
				eachMax[j] += minInt(grid[i-1][j+1], grid[i-1][j-1])
			}
			fmt.Println(eachMax[j])
		}
	}
	max := eachMax[0]
	for i := 1; i < length; i++ {
		if eachMax[i] < max {
			max = eachMax[i]
		}
	}
	return max
}

func main() {
	minFallingPathSum([][]int{
		[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9},
	})
}
