package main

func maxCount(m int, n int, ops [][]int) int {
	// 要求的是所有矩形相交的地方
	minI, minJ := m, n
	for _, op := range ops {
		if minI > op[0] {
			minI = op[0]
		}
		if minJ > op[1] {
			minJ = op[1]
		}
	}
	return minI * minJ
}
