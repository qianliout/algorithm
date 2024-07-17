package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPossible(4, [][]int{{1, 2}, {1, 3}, {1, 4}}))
}

func isPossible(n int, edges [][]int) bool {
	g := make([]map[int]bool, n+1)
	for i := range g {
		g[i] = make(map[int]bool)
	}

	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x][y] = true
		g[y][x] = true
	}
	odd := make([]int, 0)
	for i, ch := range g {
		if len(ch)&1 == 1 {
			odd = append(odd, i)
		}
	}
	if len(odd) == 0 {
		return true
	}
	if len(odd) == 2 {
		x, y := odd[0], odd[1]
		if !g[y][x] {
			return true
		}
		for i := 1; i <= n; i++ {
			// 一个不是 x 也不是 y 的 i 值，可以和 x 联边，也可以和 y 联边
			if i != x && i != y && !g[i][x] && !g[i][y] {
				return true
			}
		}
		return false
	}
	if len(odd) == 4 {
		a, b, c, d := odd[0], odd[1], odd[2], odd[3]
		a1 := !g[a][b] && !g[c][d] // a和 b 联边，c和d 联边
		a2 := !g[a][c] && !g[b][d] // a和 c 联边，g 和 d联边
		a3 := !g[b][c] && !g[a][d] // b和c联边，a 和 d 联边
		return a1 || a2 || a3
	}
	return false
}
