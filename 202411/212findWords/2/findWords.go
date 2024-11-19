package main

import (
	"sort"
)

func main() {

}

// 会超时
func findWords(board [][]byte, words []string) []string {
	tr := make(map[string]bool)
	for _, ch := range words {
		tr[ch] = true
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int)
	var word []byte
	ans := make([]string, 0)
	m, n := len(board), len(board[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	dfs = func(i, j int) {
		s := string(word)
		if len(s) > 10 {
			return
		}
		if tr[s] {
			delete(tr, s)
			ans = append(ans, s)
			// 这里不能 return ,因为abc abcd 如果直接返回了，那么就找不到 abcd
		}

		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || visit[x][y] {
				continue
			}
			visit[x][y] = true
			word = append(word, board[x][y])
			dfs(x, y)
			visit[x][y] = false
			word = word[:len(word)-1]
		}
	}
	for i := range board {
		for j := range board[i] {
			word = append(word, board[i][j])
			visit[i][j] = true
			dfs(i, j)
			visit[i][j] = false
			word = word[:len(word)-1]
		}
	}
	sort.Strings(ans)
	return ans
}

func exist(board [][]byte, word string) bool {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int) bool
	m, n := len(board), len(board[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	path := make([]byte, 0)
	dfs = func(i, j int) bool {
		s := string(path)
		if len(s) > len(word) {
			return false
		}
		if s == word {
			return true
		}

		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || visit[x][y] {
				continue
			}
			visit[x][y] = true
			path = append(path, board[x][y])
			if dfs(x, y) {
				return true
			}
			visit[x][y] = false
			path = path[:len(path)-1]
		}
		return false
	}
	for i := range board {
		for j := range board[i] {
			path = append(path, board[i][j])
			visit[i][j] = true
			if dfs(i, j) {
				return true
			}
			visit[i][j] = false
			path = path[:len(path)-1]
		}
	}
	return false
}
