package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(highestRankedKItems([][]int{{1, 2, 0, 1}, {1, 3, 0, 1}, {0, 2, 5, 1}}, []int{2, 5}, []int{0, 0}, 3))
	fmt.Println(highestRankedKItems([][]int{{2, 0, 2}, {4, 5, 3}, {2, 0, 2}}, []int{2, 5}, []int{1, 1}, 9))
}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	ans := make([][]int, 0)
	if len(grid) == 0 || len(grid[0]) == 0 {
		return ans
	}
	m, n := len(grid), len(grid[0])
	low, height := pricing[0], pricing[1]
	c, r := start[0], start[1]
	q := []pair{pair{col: c, row: r, pri: grid[c][r]}}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	visit[c][r] = true

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			if q[i].pri != q[j].pri {
				return q[i].pri < q[j].pri
			}
			if q[i].col != q[j].col {
				return q[i].col < q[j].col
			}
			return q[i].row < q[j].row
		})

		lev := make([]pair, 0)
		for _, ch := range q {
			if ch.pri > 1 && ch.pri >= low && ch.pri <= height {
				ans = append(ans, []int{ch.col, ch.row})
				if len(ans) >= k {
					return ans
				}
			}

			for _, d := range dirs {
				x, y := ch.col+d[0], ch.row+d[1]
				if x >= 0 && y >= 0 && x < m && y < n && !visit[x][y] && grid[x][y] > 0 {
					p := pair{
						col: x,
						row: y,
						pri: grid[x][y],
					}
					visit[x][y] = true
					lev = append(lev, p)
				}
			}
		}

		q = lev
	}

	return ans
}

type pair struct {
	row, col int
	pri      int
}
