package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minFlips([][]int{{0, 0}, {0, 1}}))
}

func minFlips(mat [][]int) int {
	dirs := [][]int{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(mat), len(mat[0])
	reverse := func(i, j int) {
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if in(m, n, x, y) {
				if mat[x][y] == 1 {
					mat[x][y] = 0
				} else {
					mat[x][y] = 1
				}
			}
		}
	}

	var dfs func(i int) int
	dfs = func(state int) int {
		if state >= m*n || state < 0 {
			if sum(mat) == 0 {
				return 0
			}
			return math.MaxInt64 / 10
		}
		x, y := state/n, state%n       // 状态位转横纵坐标
		res := dfs(state + 1)          // 不翻转
		reverse(x, y)                  // 翻转
		res = min(res, dfs(state+1)+1) // 求翻转后的值
		reverse(x, y)                  // 恢复
		return res
	}
	res := dfs(0)
	if res == math.MaxInt64/10 {
		return -1
	}
	return res
}

func in(m, n, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= m || y >= n {
		return false
	}
	return true
}

func sum(mat [][]int) int {
	ans := 0
	for i := 0; i < len(mat); i++ {
		for _, ch := range mat[i] {
			ans += ch
		}
	}
	return ans
}
