package main

import (
	"fmt"
)

func main() {
	fmt.Println(constructProductMatrix([][]int{{1, 2}, {3, 4}}))
	fmt.Println(constructProductMatrix([][]int{{12345}, {2}, {1}}))
}

func constructProductMatrix(grid [][]int) [][]int {
	base := 12345
	if len(grid) == 0 || len(grid[0]) == 0 {
		return grid
	}
	n, m := len(grid), len(grid[0])
	pp := make([][]int, n)
	for i := range pp {
		pp[i] = make([]int, m)
	}
	suf := 1
	// 在pp中先填充后缀所有的积
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			pp[i][j] = suf
			suf = (suf * grid[i][j]) % base
		}
	}
	pre := 1
	// 再填充前缀
	for i := range grid {
		for j, ch := range grid[i] {
			// 这里的 pre 值前面所有元素的积
			// 两种写法是一样的
			// pp[i][j] = (pp[i][j] * pre) % base
			pp[i][j] = pp[i][j] * pre % base

			// 一定是后更新
			pre = (pre * ch) % base
		}
	}
	return pp
}
