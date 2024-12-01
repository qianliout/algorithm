package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(shortestPathAllKeys([]string{"@.a..", "###.#", "b.A.B"}))
	fmt.Println(shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
}

type node struct {
	x, y int
	stat int //  最多只有6个，所以可以使用 int 表示钥匙的状态
}

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	g := make([][]byte, m)
	for i := range g {
		g[i] = []byte(grid[i])
	}
	cnt := 0
	start := node{}
	for i, row := range grid {
		for j, ch := range row {
			if ch >= 'a' && ch <= 'f' {
				cnt++
			}
			if ch == '@' {
				start.x = i
				start.y = j
			}
		}
	}
	inf := math.MaxInt / 100

	visit := make(map[node]bool)
	visit[start] = true

	var dfs func(no node, dis int) int
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(no node, dis int) int {
		x, y, stat := no.x, no.y, no.stat
		if x < 0 || x >= m || y < 0 || y >= n {
			return inf
		}
		ans := inf
		// 全部找到了
		if no.stat == (1<<cnt)-1 {
			return dis
		}
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			c := g[nx][ny]
			if c == '#' || c >= 'A' && c <= 'F' && stat>>int(c-'A')&1 == 0 {
				continue
			}

			nod := node{x: nx, y: ny, stat: stat}

			if c >= 'a' && c <= 'f' {
				ns := stat | 1<<int(c-'a')
				nod.stat = ns
			}

			if visit[nod] {
				continue
			}

			visit[nod] = true

			ans = min(ans, dfs(nod, dis+1))
		}
		return ans
	}
	ans := dfs(start, 0)
	if ans == inf {
		return -1
	}
	return ans
}

// bfs 这样写是不对的
func shortestPathAllKeys1(grid []string) int {

	m, n := len(grid), len(grid[0])
	g := make([][]byte, m)
	for i := range g {
		g[i] = []byte(grid[i])
	}
	cnt := 0
	q := make([]node, 0)
	for i, row := range grid {
		for j, ch := range row {
			if ch >= 'a' && ch <= 'z' {
				cnt++
			}
			if ch == '@' {
				q = append(q, node{x: i, y: j})
			}
		}
	}

	exit := make(map[byte]bool)
	find := 0
	dis := 0
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		lev := make([]node, 0)
		for _, no := range q {
			for _, dir := range dirs {
				nx, ny := no.x+dir[0], no.y+dir[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}
				c := g[nx][ny]
				if c == '#' {
					continue
				}
				if c == '.' {
					lev = append(lev, node{x: nx, y: ny})
					g[nx][ny] = '#'
					continue
				}
				if c >= 'a' && c <= 'z' {
					exit[c] = true
					lev = append(lev, node{x: nx, y: ny})
					find++
					if find == cnt {
						return dis + 1
					}
					g[nx][ny] = '#'
					continue
				}
				if c >= 'A' && c <= 'Z' {
					if exit[byte(c+32)] {
						g[nx][ny] = '#'
						lev = append(lev, node{x: nx, y: ny})
					}
				}
			}
		}
		q = lev
		dis++
	}
	return -1
}
