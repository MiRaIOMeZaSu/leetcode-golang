package searchmatrix

func searchMatrix(matrix [][]int, target int) bool {
	// 直接搜索
	m := len(matrix)
	n := len(matrix[0])
	x := 0
	y := 0
	if matrix[x][y] > target || matrix[m-1][n-1] < target {
		return false
	}
	for x >= 0 && x < m && y >= 0 && y < n {
		if matrix[x][y] < target {
			if x == m-1 {
				y++
			} else {
				x++
			}
		} else if matrix[x][y] > target {
			for x >= 0 && matrix[x][y] > target {
				x--
			}
			if x < 0 || matrix[x][y] > target {
				return false
			} else if matrix[x][y] == target {
				return true
			}
			y++
		} else {
			return true
		}
	}
	return false
}
