package main

import (
	"fmt"
)

func main() {
	fmt.Println(largest1BorderedSquare([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
}

func largest1BorderedSquare1(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dx, dy := make([][]int, n+1), make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dx[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dy[i] = make([]int, m+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dy[i+1][j+1] = dy[i+1][j] + matrix[i][j] // 横着
			dx[j+1][i+1] = dx[j+1][i] + matrix[i][j] // 竖着
		}
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == 0 {
				continue
			}

			// 从后往前算，这样少于ans的就直接略过，减少计算
			for y1 := n; y1 >= j; y1-- {
				wid := y1 - j + 1
				if i+wid > m {
					break
				}

				if wid <= ans {
					break
				}
				// 横
				if dy[i][y1]-dy[i][j-1] != wid {
					continue
				}
				if dy[i+wid][y1]-dy[i+wid][j-1] != wid {
					continue
				}

				// 竖
				if dx[j][i+wid]-dx[j][i-1] != wid {
					continue
				}
				if dx[y1][wid+i]-dx[j+wid][i-1] != wid {
					continue
				}

				ans = max(ans, wid)
			}
		}
	}

	return ans * ans
}

func largest1BorderedSquare(matrix [][]int) int {
	sh, he := len(matrix), len(matrix[0])
	dsh, dhe := make([][]int, sh), make([][]int, he)
	for col := 0; col < sh; col++ {
		dsh[col] = make([]int, sh+1)
	}
	for row := 0; row < he; row++ {
		dhe[row] = make([]int, he+1)
	}

	for col := 1; col < sh; col++ {
		for row := 1; row < he; row++ {
			dsh[1]
		}
	}

	return ans * ans
}
