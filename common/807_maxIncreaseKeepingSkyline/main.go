package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	// 只要一栋建筑四个方向都有比自己高的即可
	// 先寻找每列和每行的最大值
	m, n := len(grid), len(grid[0])
	// 列
	var col []int
	for i := 0; i < n; i++ {
		col = append(col, grid[0][i])
		for j := 1; j < m; j++ {
			col[i] = max(col[i], grid[j][i])
		}
	}
	// 行
	var row []int
	for i := 0; i < m; i++ {
		row = append(row, grid[i][0])
		for j := 1; j < n; j++ {
			row[i] = max(row[i], grid[i][j])
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp := min(row[i], col[j])
			if temp > grid[i][j] {
				ans += temp - grid[i][j]
			}
		}
	}
	return ans
}
