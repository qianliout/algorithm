package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestPathAllKeys([]string{"@.a..", "###.#", "b.A.B"}))
	fmt.Println(shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	queue := make([]pair, 0)

	for i, ch := range grid {
		for j, a := range ch {
			if a >= 'a' && a <= 'z' {
				cnt++
			}
			if a == '@' {
				queue = append(queue, pair{x: i, y: j, state: 0})
			}
		}
	}
	visit := make([][][]int, m)
	for i := range visit {
		visit[i] = make([][]int, n)
		for j := range visit[i] {
			visit[i][j] = make([]int, 1<<cnt)
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		no := queue[0]
		queue = queue[1:]
		step := visit[no.x][no.y][no.state]
		for _, dir := range dirs {
			x, y, s := no.x+dir[0], no.y+dir[1], no.state
			if !in(m, n, x, y) {
				continue
			}
			c := grid[x][y]
			if c == '#' {
				continue
			}

			if c >= 'A' && c <= 'Z' {
				if (s & (1 << (c - 'A'))) == 0 {
					continue
				}
			}
			if c >= 'a' && c <= 'z' {
				s |= 1 << (c - 'a')
			}
			if s == 1<<cnt-1 {
				return step + 1
			}

			if visit[x][y][s] == 0 || step+1 < visit[x][y][s] {
				visit[x][y][s] = step + 1

				d := pair{
					x:     x,
					y:     y,
					state: s,
				}
				queue = append(queue, d)
			}
		}
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
