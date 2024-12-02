package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// fmt.Println(shortestPathAllKeys([]string{"@.a..", "###.#", "b.A.B"}))
	fmt.Println(shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
}

type node struct {
	x, y int
	stat int //  最多只有6个，所以可以使用 int 表示钥匙的状态
}

// 这个写法是对的，因为 dfs 从不同的路走过去的距离是不一样的，如果前一次返问了一个节点，后面一也能范围到该节点
// 并且比前一次更短，但是因为前一次已经标计已访问，导致后一次是访问不到的
func shortestPathAllKeys2(grid []string) int {
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

				// 是锁，但是没有钥匙
				if c >= 'A' && c <= 'Z' && (no.state&(1<<(c-'A'))) == 0 {
					continue
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
		// 这里的 ans 判断时可以不用判断 lev 是否是空
		ans++
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
