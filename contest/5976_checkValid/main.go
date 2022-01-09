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

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	for i := 0; i < n; i++ {
		visit := make(map[int]bool)
		minN := 101
		maxN := 0
		for j := 0; j < n; j++ {
			num := matrix[i][j]
			minN = min(minN, num)
			maxN = max(maxN, num)
			if _, ok := visit[num]; !ok {
				visit[num] = true
			} else {
				return false
			}
		}
		if minN != 1 || maxN != n {
			return false
		}
	}
	for i := 0; i < n; i++ {
		visit := make(map[int]bool)
		minN := 101
		maxN := 0
		for j := 0; j < n; j++ {
			num := matrix[j][i]
			minN = min(minN, num)
			maxN = max(maxN, num)
			if _, ok := visit[num]; !ok {
				visit[num] = true
			} else {
				return false
			}
		}
		if minN != 1 || maxN != n {
			return false
		}
	}
	return true
}
