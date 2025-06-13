package main

func main() {

}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	used := make([][]bool, m+1)
	for i := range used {
		used[i] = make([]bool, n+1)
	}

	var dfs func(start int, i, j int) bool
	find := false
	dfs = func(start int, i, j int) bool {
		if find {
			return true
		}
		if start >= len(word) {
			find = true
			return true
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}
		if word[start] != board[i][j] {
			return false
		}
		if used[i][j] {
			return false
		}
		used[i][j] = true

		ans := false
		ans = ans || dfs(start+1, i+1, j)
		ans = ans || dfs(start+1, i-1, j)
		ans = ans || dfs(start+1, i, j+1)
		ans = ans || dfs(start+1, i, j-1)
		used[i][j] = false
		return ans
	}
	//  题目中说了，可以ww任务位开始
	// 尝试从网格中的每个位置开始搜索
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}

/*
题目：给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/
