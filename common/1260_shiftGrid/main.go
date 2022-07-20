package main

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	// 去掉重复
	count := n * m
	rest := k % count
	// 使用全局下标计算
	row := rest / n
	col := rest % n
	result := [][]int{}
	for i := 0; i < m; i++ {
		temp := []int{}
		for j := 0; j < n; j++ {
			next_i := i - row
			next_j := j - col
			if next_j < 0 {
				next_j += n
				next_i--
			}
			if next_i < 0 {
				next_i += m
			}
			temp = append(temp, grid[next_i][next_j])
		}
		result = append(result, temp)
	}
	return result
}

func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	shiftGrid(grid, 1)
}
