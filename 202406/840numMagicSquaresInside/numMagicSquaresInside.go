package main

import (
	"fmt"
)

func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
}

func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}
	ans := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i <= n-3; i++ {
		for j := 0; j <= m-3; j++ {
			if check(grid, i, j) {
				ans++
			}
		}
	}
	return ans
}

func check(grid [][]int, co, row int) bool {
	a := 0
	// 行
	for i := row; i < row+3; i++ {
		va := grid[co][i]
		if va < 1 || va > 9 {
			return false
		}
		a += va
	}
	// 列
	b := 0
	for i := co; i < co+3; i++ {
		b += grid[i][row]
	}
	if a != b {
		return false
	}
	// 135 xie
	c := 0
	for i, j := co, row; i < co+3 && j < row+3; i, j = i+1, j+1 {
		c += grid[i][j]
	}
	if c != a {
		return false
	}

	d := 0
	for i, j := co, row+2; i < co+3 && j >= row; i, j = i+1, j-1 {
		d += grid[i][j]
	}
	if d != a {
		return false
	}
	return true
}
