package main

import (
	"math/bits"
)

func main() {

}

// 为啥是错的
func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	queue := make([]pair, 0)
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	for i, ch := range grid {
		for j, a := range ch {
			if a >= 'a' && a <= 'z' {
				cnt++
			}
			if a == '@' {
				queue = append(queue, pair{x: i, y: j, state: 0, cnt: 0})
				visit[i][j] = true
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ans := 0
	state := 0
	for len(queue) > 0 {
		lev := make([]pair, 0)
		for _, no := range queue {
			for _, dir := range dirs {
				x, y := no.x+dir[0], no.y+dir[1]
				if !in(m, n, x, y) {
					continue
				}
				if visit[x][y] {
					continue
				}
				visit[x][y] = true
				c := grid[x][y]
				if c == '#' {
					continue
				}

				if c >= 'A' && c <= 'Z' {
					if (no.state & (1 << (c - 'A'))) == 0 {
						continue
					}
				}
				d := pair{
					x:     x,
					y:     y,
					state: no.state,
					cnt:   no.cnt,
				}
				if c >= 'a' && c <= 'z' {
					state |= 1 << (c - 'a')
					d.state |= 1 << (c - 'a')
					d.cnt++
				}
				lev = append(lev, d)
			}
		}
		if len(lev) > 0 {
			ans++
		}
		if bits.OnesCount(uint(state)) == cnt {
			return ans
		}

		queue = lev
	}

	return -1
}

type pair struct {
	x     int
	y     int
	state int
	cnt   int
}

func in(m, n int, x, y int) bool {
	if x >= m || y >= n {
		return false
	}
	if x < 0 || y < 0 {
		return false
	}
	return true
}
