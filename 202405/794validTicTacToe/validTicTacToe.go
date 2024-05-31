package main

import (
	"fmt"
)

func main() {
	fmt.Println(validTicTacToe([]string{"XXX", "XOO", "OO "}))
}

func validTicTacToe(board []string) bool {
	xc, oc := 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				xc++
			} else if board[i][j] == 'O' {
				oc++
			}
		}
	}
	if oc > xc || xc-oc > 1 {
		return false
	}
	a := check(board, 'X')
	b := check(board, 'O')

	if a && !b && xc-oc == 1 {
		return true
	}
	if b && !a && xc == oc {
		return true
	}
	if !b && !a {
		return true
	}

	return false
}

func check(board []string, by byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][2] == by {
			return true
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[2][i] == by {
			return true
		}
	}
	// 斜线
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[1][1] == by {
		return true
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[1][1] == by {
		return true
	}

	return false
}
