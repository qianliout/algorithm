package main

func main() {

}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if board[i][j] != 'O' {
			return
		}
		board[i][j] = 'Y'
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}
			dfs(x, y)
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
