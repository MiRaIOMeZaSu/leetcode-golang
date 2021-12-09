package main

func validTicTacToe(board []string) bool {
	// X的数量必须大于等于O的数量
	// 胜利的情况
	// 未胜利的情况
	countO := 0
	countX := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'O' {
				countO++
			} else {
				countX++
			}
		}
	}
	if countX < countO || countX-countO > 1 {
		return false
	}
	// 此时数量已经满足
	// 遍历所有成功的手段
	flag := ""
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][2] == board[i][1] {
			flag += string(board[i][0])
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			flag += string(board[0][i])
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		flag += string(board[0][0])
	}
	if board[2][0] == board[1][1] && board[1][1] == board[2][0] {
		flag += string(board[1][1])
	}
	if len(flag) > 1 {
		return false
	}
	if len(flag) == 1 {
		if flag == "X" {
			return countX-countO == 1
		} else {
			return countX == countO
		}
	}
	return true
}
