package main

import (
	"fmt"
)

func main() {
	fmt.Println(updateMatrix([][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}))
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	q := make([]node, 0)
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				q = append(q, node{x: i, y: j, dis: 0})
			}
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		lev := make([]node, 0)
		for _, fir := range q {
			for _, dir := range dirs {
				nx, ny := fir.x+dir[0], fir.y+dir[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}
				if mat[nx][ny] == 0 {
					continue
				}

				ans[nx][ny] = fir.dis + 1
				// ans[nx][ny] = min(ans[nx][ny], fir.dis+1)
				lev = append(lev, node{x: nx, y: ny, dis: fir.dis + 1})
				mat[nx][ny] = 0
			}
		}
		q = lev
	}

	return ans
}

type node struct {
	x, y int
	dis  int
}
