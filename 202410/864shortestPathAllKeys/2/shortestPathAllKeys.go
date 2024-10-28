package main

import (
	"math/bits"
)

func main() {

}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	queue := make([]pair, 0)

	// 这里是否返问过的状态容易出错，一个格子，从不同的格子走过来，结果不是同的,所以还得加一维状态，表示上一次的钥匙情况
	// visit:=make([][]int,m) // 这种方式就不得行
	visit := make(map[pair]bool)
	for i, ch := range grid {
		for j, a := range ch {
			if a >= 'a' && a <= 'z' {
				cnt++
			}
			if a == '@' {
				no := pair{x: i, y: j}
				queue = append(queue, no)
				visit[no] = true
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ans := 0
	for len(queue) > 0 {
		lev := make([]pair, 0)
		for _, no := range queue {
			s := no.state
			if bits.OnesCount(uint(s)) == cnt {
				return ans
			}
			for _, dir := range dirs {
				x, y := no.x+dir[0], no.y+dir[1]
				if !in(m, n, x, y) {
					continue
				}
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
					state: s,
				}
				if c >= 'a' && c <= 'z' {
					d.state |= 1 << (c - 'a')
				}
				if !visit[d] {
					visit[d] = true
					lev = append(lev, d)
				}
			}
		}
		if len(lev) > 0 {
			ans++
		}
		queue = lev
	}

	return -1
}

type pair struct {
	x     int
	y     int
	state int
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
