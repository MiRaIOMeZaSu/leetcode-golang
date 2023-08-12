package main

func diagonalSum(mat [][]int) int {
	l := len(mat)
	sum := 0
	x := l - 1
	for i := 0; i < l; i++ {
		sum += mat[i][i]
		if i != x {
			sum += mat[x][i]
		}
		x -= 1
	}
	return sum
}
