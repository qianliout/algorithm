package main

func main() {

}

func solveSudoku(board [][]byte) {
	n := len(board)

	var backtrack func(row, col int) bool
	// 有返回值的目的ch程在到正确解时不再 归
	backtrack = func(row, col int) bool {
		// 如果已经填完最后一行，返回成功
		if row == n {
			return true
		}
		// 如果当前列已经到达末尾，转到下一行
		if col == n {
			return backtrack(row+1, 0)
		}

		// 如果当前位置已有数字，直接处理下一个位置
		if board[row][col] != '.' {
			return backtrack(row, col+1)
		}

		// 尝试填入1-9
		for num := byte('1'); num <= '9'; num++ {
			// 检查是否可以填入
			if isValid(board, row, col, num) {
				// 填入数字
				board[row][col] = num
				// 继续填下一个位置
				if backtrack(row, col+1) {
					return true
				}
				// 回溯
				board[row][col] = '.'
			}
		}
		// 无法找到有效解
		return false
	}

	backtrack(0, 0)
}

// isValid 检查在(row, col)位置放置数字num是否有效
func isValid(board [][]byte, row, col int, num byte) bool {
	n := len(board)
	// 检查行
	for j := 0; j < n; j++ {
		if board[row][j] == num {
			return false
		}
	}
	// 检查列
	for i := 0; i < n; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// 检查3x3宫格
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

/*
编写一个程序，通过填充空格来解决数独问题。

数独的解法需 遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。


*/
