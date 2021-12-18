package main

import "fmt"

func countBattleships(board [][]byte) int {
	// 不查看左边或上面的
	ans := 0
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				flag := 0
				if i-1 < 0 || board[i-1][j] != 'X' {
					flag++
				}
				if j-1 < 0 || board[i][j-1] != 'X' {
					flag++
				}
				if flag == 2 {
					ans++
				}
			}
		}
	}
	return ans
}

func main() {
	ans := countBattleships([][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}})
	fmt.Print(ans)
}
